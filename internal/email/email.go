package email

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base32"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"os"
	"text/template"
	"time"

	pb "github.com/mxc-foundation/lpwan-app-server/api"
	"github.com/mxc-foundation/lpwan-app-server/internal/config"
	"github.com/mxc-foundation/lpwan-app-server/internal/static"
	log "github.com/sirupsen/logrus"
)

var (
	senderID   string
	password   string
	host       string
	smtpServer string
	smtpPort   string
	disable    bool
)

// English sets the Language to en
const (
	English            = pb.Language_en
	Korean             = pb.Language_ko
	SimplifiedChinese  = pb.Language_zhcn
	TraditionalChinese = pb.Language_zhtw
	Japanese           = pb.Language_ja
)

// Setup configures the package.
func Setup(c config.Config) error {
	senderID = c.SMTP.Email
	password = c.SMTP.Password
	smtpServer = c.SMTP.Host
	smtpPort = c.SMTP.Port
	host = os.Getenv("APPSERVER")
	disable = false

	base32endocoding = base32.StdEncoding.WithPadding(base32.NoPadding)
	mailTemplates = make([]*template.Template, len(mailTemplateNames))
	for k, v := range mailTemplateNames {
		mailTemplates[k] = template.Must(
			template.New(v.templatePath).Parse(string(static.MustAsset(v.templatePath))))
	}

	return nil
}

// Disable stops emails from being sent.
func Disable() {
	disable = true
}

var (
	mailTemplates    []*template.Template
	base32endocoding *base32.Encoding

	mailTemplateNames = []struct {
		templatePath string
		url          string
	}{
		English: {
			templatePath: "templates/registration-confirm-en",
			url:          "/#/registration-confirm/",
		},
		Korean: {
			templatePath: "templates/registration-confirm-ko",
			url:          "/#/registration-confirm/",
		},
		SimplifiedChinese: {
			templatePath: "templates/registration-confirm-zhcn",
			url:          "/#/registration-confirm/",
		},
		TraditionalChinese: {
			templatePath: "templates/registration-confirm-zhtw",
			url:          "/#/registration-confirm/",
		},
		Japanese: {
			templatePath: "templates/registration-confirm-ja",
			url:          "/#/registration-confirm/",
		},
	}
)

// SendInvite ...
func SendInvite(user, token string, language int32) error {
	var err error

	if disable == true {
		return nil
	}

	if smtpServer == "" {
		log.Error("Tried to send registration email, but SMTP is not configured")
		return errors.New("Unable to send confirmation email")
	}

	link := host + mailTemplateNames[language].url + token

	logo := host + "/branding.png"

	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	messageID := time.Now().Format("20060102150405.") + base32endocoding.EncodeToString(b)

	var msg bytes.Buffer
	if err := mailTemplates[language].Execute(&msg, struct {
		From, To, Host, MsgID, Boundary, Link, Logo string
	}{
		From:     senderID,
		To:       user,
		Host:     host,
		MsgID:    messageID + "@" + host,
		Boundary: "----=_Part_" + messageID,
		Link:     link,
		Logo:     logo,
	}); err != nil {
		log.Error(err)
		return err
	}

	auth := smtp.PlainAuth(
		"",
		senderID,
		password,
		smtpServer,
	)

	err = SendMailUsingTLS(
		fmt.Sprintf("%s:%d", smtpServer, 465),
		auth,
		senderID,
		[]string{user},
		msg.Bytes(),
	)
	if err != nil {
		return err
	}

	return nil
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//split host address and port
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
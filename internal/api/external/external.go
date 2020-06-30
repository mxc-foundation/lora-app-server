package external

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/brocaar/chirpstack-api/go/v3/as/external/api"

	"github.com/mxc-foundation/lpwan-app-server/internal/api/external/auth"
	"github.com/mxc-foundation/lpwan-app-server/internal/api/external/oidc"
	"github.com/mxc-foundation/lpwan-app-server/internal/api/externalcus"
	"github.com/mxc-foundation/lpwan-app-server/internal/api/helpers"
	"github.com/mxc-foundation/lpwan-app-server/internal/config"
	"github.com/mxc-foundation/lpwan-app-server/internal/static"
	"github.com/mxc-foundation/lpwan-app-server/internal/storage"
)

var externalCtx struct {
	brandingRegistration    string
	brandingFooter          string
	openIDLoginLabel        string
	openIDConnectEnabled    bool
	registrationEnabled     bool
	registrationCallbackURL string

	bind            string
	tlsCert         string
	tlsKey          string
	jwtSecret       string
	corsAllowOrigin string

	applicationServerID uuid.UUID
}

// Setup configures the API package.
func Setup(conf config.Config) error {
	if conf.ApplicationServer.ExternalAPI.JWTSecret == "" {
		return errors.New("jwt_secret must be set")
	}

	externalCtx.brandingRegistration = conf.ApplicationServer.Branding.Registration
	externalCtx.brandingFooter = conf.ApplicationServer.Branding.Footer
	externalCtx.registrationEnabled = conf.ApplicationServer.UserAuthentication.OpenIDConnect.RegistrationEnabled
	externalCtx.registrationCallbackURL = conf.ApplicationServer.UserAuthentication.OpenIDConnect.RegistrationCallbackURL
	externalCtx.openIDConnectEnabled = conf.ApplicationServer.UserAuthentication.OpenIDConnect.Enabled
	externalCtx.openIDLoginLabel = conf.ApplicationServer.UserAuthentication.OpenIDConnect.LoginLabel

	externalCtx.bind = conf.ApplicationServer.ExternalAPI.Bind
	externalCtx.tlsCert = conf.ApplicationServer.ExternalAPI.TLSCert
	externalCtx.tlsKey = conf.ApplicationServer.ExternalAPI.TLSKey
	externalCtx.jwtSecret = conf.ApplicationServer.ExternalAPI.JWTSecret
	externalCtx.corsAllowOrigin = conf.ApplicationServer.ExternalAPI.CORSAllowOrigin

	if err := externalCtx.applicationServerID.UnmarshalText([]byte(conf.ApplicationServer.ID)); err != nil {
		return errors.Wrap(err, "decode application_server.id error")
	}

	return setupAPI(conf)
}

func setupAPI(conf config.Config) error {
	validator := auth.NewJWTValidator(storage.DB(), "HS256", externalCtx.jwtSecret)
	rpID, err := uuid.FromString(conf.ApplicationServer.ID)
	if err != nil {
		return errors.Wrap(err, "application-server id to uuid error")
	}

	grpcOpts := helpers.GetgRPCServerOptions()
	grpcServer := grpc.NewServer(grpcOpts...)
	pb.RegisterApplicationServiceServer(grpcServer, NewApplicationAPI(validator))
	pb.RegisterDeviceQueueServiceServer(grpcServer, NewDeviceQueueAPI(validator))
	pb.RegisterDeviceServiceServer(grpcServer, NewDeviceAPI(validator))
	pb.RegisterUserServiceServer(grpcServer, NewUserAPI(validator))
	pb.RegisterInternalServiceServer(grpcServer, NewInternalAPI(validator))
	pb.RegisterGatewayServiceServer(grpcServer, NewGatewayAPI(validator))
	pb.RegisterGatewayProfileServiceServer(grpcServer, NewGatewayProfileAPI(validator))
	pb.RegisterOrganizationServiceServer(grpcServer, NewOrganizationAPI(validator))
	pb.RegisterNetworkServerServiceServer(grpcServer, NewNetworkServerAPI(validator))
	pb.RegisterServiceProfileServiceServer(grpcServer, NewServiceProfileServiceAPI(validator))
	pb.RegisterDeviceProfileServiceServer(grpcServer, NewDeviceProfileServiceAPI(validator))
	pb.RegisterMulticastGroupServiceServer(grpcServer, NewMulticastGroupAPI(validator, rpID))
	pb.RegisterFUOTADeploymentServiceServer(grpcServer, NewFUOTADeploymentAPI(validator))

	err = externalcus.SetupAPI(grpcServer, externalCtx.applicationServerID)
	if err != nil {
		return errors.Wrap(err, "failed to register customized external APIs ")
	}

	// setup the client http interface variable
	// we need to start the gRPC service first, as it is used by the
	// grpc-gateway
	var clientHTTPHandler http.Handler

	// switch between gRPC and "plain" http handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			if clientHTTPHandler == nil {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}

			if externalCtx.corsAllowOrigin != "" {
				w.Header().Set("Access-Control-Allow-Origin", externalCtx.corsAllowOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Grpc-Metadata-Authorization")

				if r.Method == "OPTIONS" {
					return
				}
			}

			clientHTTPHandler.ServeHTTP(w, r)
		}
	})

	// start the API server
	go func() {
		log.WithFields(log.Fields{
			"bind":     externalCtx.bind,
			"tls-cert": externalCtx.tlsCert,
			"tls-key":  externalCtx.tlsKey,
		}).Info("api/external: starting api server")

		if externalCtx.tlsCert == "" || externalCtx.tlsKey == "" {
			log.Fatal(http.ListenAndServe(externalCtx.bind, h2c.NewHandler(handler, &http2.Server{})))
		} else {
			log.Fatal(http.ListenAndServeTLS(
				externalCtx.bind,
				externalCtx.tlsCert,
				externalCtx.tlsKey,
				h2c.NewHandler(handler, &http2.Server{}),
			))
		}
	}()

	// give the http server some time to start
	time.Sleep(time.Millisecond * 100)

	// setup the HTTP handler
	clientHTTPHandler, err = setupHTTPAPI(conf)
	if err != nil {
		return err
	}

	return nil
}

func setupHTTPAPI(conf config.Config) (http.Handler, error) {
	r := mux.NewRouter()

	// setup json api handler
	jsonHandler, err := getJSONGateway(context.Background())
	if err != nil {
		return nil, err
	}

	log.WithField("path", "/api").Info("api/external: registering rest api handler and documentation endpoint")
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		data, err := static.Asset("swagger/index.html")
		if err != nil {
			log.WithError(err).Error("get swagger template error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}).Methods("get")
	r.PathPrefix("/api").Handler(jsonHandler)

	if err := oidc.Setup(conf, r); err != nil {
		return nil, errors.Wrap(err, "setup openid connect error")
	}

	// setup static file server
	r.PathPrefix("/").Handler(http.FileServer(&assetfs.AssetFS{
		Asset:     static.Asset,
		AssetDir:  static.AssetDir,
		AssetInfo: static.AssetInfo,
		Prefix:    "",
	}))

	return wsproxy.WebsocketProxy(r), nil
}

func getJSONGateway(ctx context.Context) (http.Handler, error) {
	// dial options for the grpc-gateway
	var grpcDialOpts []grpc.DialOption

	if externalCtx.tlsCert == "" || externalCtx.tlsKey == "" {
		grpcDialOpts = append(grpcDialOpts, grpc.WithInsecure())
	} else {
		b, err := ioutil.ReadFile(externalCtx.tlsCert)
		if err != nil {
			return nil, errors.Wrap(err, "read external api tls cert error")
		}
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			return nil, errors.Wrap(err, "failed to append certificate")
		}
		grpcDialOpts = append(grpcDialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			// given the grpc-gateway is always connecting to localhost, does
			// InsecureSkipVerify=true cause any security issues?
			InsecureSkipVerify: true,
			RootCAs:            cp,
		})))
	}

	bindParts := strings.SplitN(externalCtx.bind, ":", 2)
	if len(bindParts) != 2 {
		log.Fatal("get port from bind failed")
	}
	apiEndpoint := fmt.Sprintf("localhost:%s", bindParts[1])

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			EnumsAsInts:  false,
			EmitDefaults: true,
		},
	))

	err := externalcus.CusGetJSONGateway(ctx, mux, apiEndpoint, grpcDialOpts)
	if err != nil {
		return nil, errors.Wrap(err, "register customized external service handler error")
	}

	if err := pb.RegisterApplicationServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register application handler error")
	}
	if err := pb.RegisterDeviceQueueServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register downlink queue handler error")
	}
	if err := pb.RegisterDeviceServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register node handler error")
	}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register user handler error")
	}
	if err := pb.RegisterInternalServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register internal handler error")
	}
	if err := pb.RegisterGatewayServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register gateway handler error")
	}
	if err := pb.RegisterGatewayProfileServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register gateway-profile handler error")
	}
	if err := pb.RegisterOrganizationServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register organization handler error")
	}
	if err := pb.RegisterNetworkServerServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register network-server handler error")
	}
	if err := pb.RegisterServiceProfileServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register service-profile handler error")
	}
	if err := pb.RegisterDeviceProfileServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register device-profile handler error")
	}
	if err := pb.RegisterMulticastGroupServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register multicast-group handler error")
	}
	if err := pb.RegisterFUOTADeploymentServiceHandlerFromEndpoint(ctx, mux, apiEndpoint, grpcDialOpts); err != nil {
		return nil, errors.Wrap(err, "register fuota deployment handler error")
	}

	return mux, nil
}

package auth

import (
	"context"

	"github.com/gofrs/uuid"
)

// Authenticator authenticates the user and returns Credentials
type Authenticator interface {
	GetCredentials(ctx context.Context, opts *Options) (*Credentials, error)
}

type Options struct {
	Audience         string
	RequireOTP       bool
	AllowNonExisting bool
	OrgID            int64
	ApplicationID    int64
	DeviceProfileID  string

	// when GetOrgIDFromToken is true, extract organization id from jwt then assign it to user's credential
	GetOrgIDFromToken bool
	// external auth
	ExternalLimited bool
}

func NewOptions() *Options {
	return &Options{
		Audience: "lora-app-server",
	}
}

// WithOrgIDFromToken sets GetOrgIDFromToken = true
func (o *Options) WithOrgIDFromToken() *Options {
	o.GetOrgIDFromToken = true
	return o
}

func (o *Options) WithAudience(audience string) *Options {
	o.Audience = audience
	return o
}

func (o *Options) WithRequireOTP() *Options {
	o.RequireOTP = true
	return o
}

func (o *Options) WithAllowNonExisting() *Options {
	o.AllowNonExisting = true
	return o
}

func (o *Options) WithOrgID(orgID int64) *Options {
	o.OrgID = orgID
	return o
}

// WithApplicationID will take applicationID into consideration
func (o *Options) WithApplicationID(applicationID int64) *Options {
	o.ApplicationID = applicationID
	return o
}

// WithDeviceProfileID will take DeviceProfileID into consideration
func (o *Options) WithDeviceProfileID(deviceProfileID string) *Options {
	o.DeviceProfileID = deviceProfileID
	return o
}

// WithExternalLimited restricts checking external credentials only
func (o *Options) WithExternalLimited() *Options {
	o.ExternalLimited = true
	return o
}

// ExternalServiceName defines const type: name of external services
const (
	EMAIL   string = "email"
	WECHAT  string = "wechat"
	TG      string = "telegram"
	SHOPIFY string = "shopify"
)

// Credentials provides information about user's credentials
type Credentials struct {
	// UserID is the id of the user
	UserID int64
	// Username is the username
	Username string
	// IsExisting is true if the user exist
	IsExisting bool
	// IsGlobalAdmin is true if user is a global admin
	IsGlobalAdmin bool
	// OrgID is the ID of the organization for which org credentials were checked
	OrgID int64
	// IsOrgUser is true if the user belongs to the organization
	IsOrgUser bool
	// IsOrgAdmin is true if the user is org admin
	IsOrgAdmin bool
	// IsDeviceAdmin is true if the user is device admin for the org
	IsDeviceAdmin bool
	// IsGatewayAdmin is true if the user is device admin for the org
	IsGatewayAdmin bool
	// IsApplicationUser is true if the user's organization owns the given application
	IsApplicationUser bool
	// DeviceProfileIsValid is true if the user's organization owns the given device profile
	DeviceProfileIsValid bool
	// ExternalUserID is the id of external user
	ExternalUserID string
	// Service is the name of external user's service
	Service string
	// ExternalUsername is the nickname of the external user
	ExternalUsername string
}

// User contains information about the user
type User struct {
	ID            int64
	Email         string
	IsGlobalAdmin bool
}

// OrgUser contains information about the role of the user in organisation
type OrgUser struct {
	IsOrgUser      bool
	IsOrgAdmin     bool
	IsDeviceAdmin  bool
	IsGatewayAdmin bool
}

// Store provides access to information about users and their roles
type Store interface {
	// AuthGetUser returns user's information given that there is an active user
	// with the given username
	AuthGetUser(ctx context.Context, username string) (User, error)
	// AuthGetOrgUser returns user's role in the listed organization
	AuthGetOrgUser(ctx context.Context, userID int64, orgID int64) (OrgUser, error)
	// ApplicationOwnedByOrganization returns true when given orgID owns given applicationID
	ApplicationOwnedByOrganization(ctx context.Context, orgID, applicationID int64) (bool, error)
	// DeviceProfileOwnedByOrganization returns true when given orgID owns given device profile
	DeviceProfileOwnedByOrganization(ctx context.Context, orgID int64, deviceProfile uuid.UUID) (bool, error)
}

// NewCredentials returns credential set of an user
func NewCredentials(ctx context.Context, st Store, username string, orgID int64, service string,
	applicationID int64, deviceProfileID string) (*Credentials, error) {
	user, err := st.AuthGetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	c := &Credentials{
		UserID:        user.ID,
		Username:      user.Email,
		IsGlobalAdmin: user.IsGlobalAdmin,
		IsExisting:    true,
		Service:       service,
	}
	if orgID > 0 {
		orgUser, err := st.AuthGetOrgUser(ctx, user.ID, orgID)
		if err != nil {
			return nil, err
		}
		c.OrgID = orgID
		c.IsOrgUser = orgUser.IsOrgUser || c.IsGlobalAdmin
		c.IsOrgAdmin = orgUser.IsOrgAdmin || c.IsGlobalAdmin
		c.IsDeviceAdmin = orgUser.IsDeviceAdmin || c.IsOrgAdmin
		c.IsGatewayAdmin = orgUser.IsGatewayAdmin || c.IsOrgAdmin

		if applicationID > 0 {
			appOwnedByOrg, err := st.ApplicationOwnedByOrganization(ctx, orgID, applicationID)
			if err != nil {
				return nil, err
			}
			c.IsApplicationUser = appOwnedByOrg
		}

		if deviceProfileID != "" {
			dpID, err := uuid.FromString(deviceProfileID)
			if err != nil {
				return nil, err
			}
			dpOwnedByOrg, err := st.DeviceProfileOwnedByOrganization(ctx, orgID, dpID)
			if err != nil {
				return nil, err
			}
			c.DeviceProfileIsValid = dpOwnedByOrg
		}
	}
	return c, nil
}

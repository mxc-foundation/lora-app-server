package m2m_ui

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	m2m_api "github.com/mxc-foundation/lpwan-app-server/api/m2m_server"
	api "github.com/mxc-foundation/lpwan-app-server/api/m2m_ui"
	"github.com/mxc-foundation/lpwan-app-server/internal/api/external/auth"
	"github.com/mxc-foundation/lpwan-app-server/internal/api/helpers"
	"github.com/mxc-foundation/lpwan-app-server/internal/backend/m2m_client"
	"github.com/mxc-foundation/lpwan-app-server/internal/config"
	"github.com/mxc-foundation/lpwan-app-server/internal/storage"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ExtAccountServerAPI struct {
	validator auth.Validator
}

func NewMoneyServerAPI(validator auth.Validator) *ExtAccountServerAPI {
	return &ExtAccountServerAPI{
		validator: validator,
	}
}

func (s *ExtAccountServerAPI) ModifyMoneyAccount(ctx context.Context, req *api.ModifyMoneyAccountRequest) (*api.ModifyMoneyAccountResponse, error) {
	log.WithField("orgId", req.OrgId).Info("grpc_api/ModifyMoneyAccount")

	m2mClient, err := m2m_client.GetPool().Get(config.C.M2MServer.M2MServer, []byte(config.C.M2MServer.CACert),
		[]byte(config.C.M2MServer.TLSCert), []byte(config.C.M2MServer.TLSKey))
	if err != nil {
		return &api.ModifyMoneyAccountResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	moneyAbbr := m2m_api.Money(req.MoneyAbbr)

	resp, err := m2mClient.ModifyMoneyAccount(ctx, &m2m_api.ModifyMoneyAccountRequest{
		OrgId:          req.OrgId,
		MoneyAbbr:      moneyAbbr,
		CurrentAccount: req.CurrentAccount,
	})
	if err != nil {
		return &api.ModifyMoneyAccountResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	username, err := auth.JWTValidator{}.GetUsername(ctx)
	if nil != err {
		return nil, helpers.ErrToRPCError(err)
	}

	// Get the user id based on the username.
	user, err := storage.GetUserByUsername(ctx, storage.DB(), username)
	if nil != err {
		return nil, helpers.ErrToRPCError(err)
	}

	prof, err := storage.GetProfile(ctx, storage.DB(), user.ID)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	//userProfile := api.GetDeviceListResponse.GetUserProfile(prof)
	userProfile := api.ProfileResponse{
		User: &api.User{
			Id:         string(prof.User.ID),
			Username:   prof.User.Username,
			SessionTtl: prof.User.SessionTTL,
			IsAdmin:    prof.User.IsAdmin,
			IsActive:   prof.User.IsActive,
		},
		Settings: &api.ProfileSettings{
			DisableAssignExistingUsers: auth.DisableAssignExistingUsers,
		},
	}

	for _, org := range prof.Organizations {
		row := api.OrganizationLink{
			OrganizationId:   org.ID,
			OrganizationName: org.Name,
			IsAdmin:          org.IsAdmin,
		}

		row.CreatedAt, err = ptypes.TimestampProto(org.CreatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		row.UpdatedAt, err = ptypes.TimestampProto(org.UpdatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}

		userProfile.Organizations = append(userProfile.Organizations, &row)
	}

	return &api.ModifyMoneyAccountResponse{
		Status:      resp.Status,
		UserProfile: &userProfile,
	}, nil
}

func (s *ExtAccountServerAPI) GetChangeMoneyAccountHistory(ctx context.Context, req *api.GetMoneyAccountChangeHistoryRequest) (*api.GetMoneyAccountChangeHistoryResponse, error) {
	log.WithField("orgId", req.OrgId).Info("grpc_api/GetChangeMoneyAccountHistory")

	m2mClient, err := m2m_client.GetPool().Get(config.C.M2MServer.M2MServer, []byte(config.C.M2MServer.CACert),
		[]byte(config.C.M2MServer.TLSCert), []byte(config.C.M2MServer.TLSKey))
	if err != nil {
		return &api.GetMoneyAccountChangeHistoryResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	moneyAbbr := m2m_api.Money(req.MoneyAbbr)

	resp, err := m2mClient.GetChangeMoneyAccountHistory(ctx, &m2m_api.GetMoneyAccountChangeHistoryRequest{
		OrgId:     req.OrgId,
		Offset:    req.Offset,
		Limit:     req.Limit,
		MoneyAbbr: moneyAbbr,
	})
	if err != nil {
		return &api.GetMoneyAccountChangeHistoryResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	//changeHist := api.GetMoneyAccountChangeHistoryResponse.GetChangeHistory(&resp.ChangeHistory)
	hists := api.GetMoneyAccountChangeHistoryResponse{}.ChangeHistory

	for _, v := range resp.ChangeHistory {
		hist := api.MoneyAccountChangeHistory{}
		hist.CreatedAt = v.CreatedAt
		hist.Addr = v.Addr
		hist.Status = v.Status

		hists = append(hists, &hist)
	}

	username, err := auth.JWTValidator{}.GetUsername(ctx)
	if nil != err {
		return nil, helpers.ErrToRPCError(err)
	}

	// Get the user id based on the username.
	user, err := storage.GetUserByUsername(ctx, storage.DB(), username)
	if nil != err {
		return nil, helpers.ErrToRPCError(err)
	}

	prof, err := storage.GetProfile(ctx, storage.DB(), user.ID)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	//userProfile := api.GetDeviceListResponse.GetUserProfile(prof)
	userProfile := api.ProfileResponse{
		User: &api.User{
			Id:         string(prof.User.ID),
			Username:   prof.User.Username,
			SessionTtl: prof.User.SessionTTL,
			IsAdmin:    prof.User.IsAdmin,
			IsActive:   prof.User.IsActive,
		},
		Settings: &api.ProfileSettings{
			DisableAssignExistingUsers: auth.DisableAssignExistingUsers,
		},
	}

	for _, org := range prof.Organizations {
		row := api.OrganizationLink{
			OrganizationId:   org.ID,
			OrganizationName: org.Name,
			IsAdmin:          org.IsAdmin,
		}

		row.CreatedAt, err = ptypes.TimestampProto(org.CreatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		row.UpdatedAt, err = ptypes.TimestampProto(org.UpdatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}

		userProfile.Organizations = append(userProfile.Organizations, &row)
	}

	return &api.GetMoneyAccountChangeHistoryResponse{
		Count:         resp.Count,
		ChangeHistory: hists,
		UserProfile:   &userProfile,
	}, nil
}

func (s *ExtAccountServerAPI) GetActiveMoneyAccount(ctx context.Context, req *api.GetActiveMoneyAccountRequest) (*api.GetActiveMoneyAccountResponse, error) {
	log.WithField("orgId", req.OrgId).Info("grpc_api/GetActiveMoneyAccount")

	m2mClient, err := m2m_client.GetPool().Get(config.C.M2MServer.M2MServer, []byte(config.C.M2MServer.CACert),
		[]byte(config.C.M2MServer.TLSCert), []byte(config.C.M2MServer.TLSKey))
	if err != nil {
		return &api.GetActiveMoneyAccountResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	moneyAbbr := m2m_api.Money(req.MoneyAbbr)

	resp, err := m2mClient.GetActiveMoneyAccount(ctx, &m2m_api.GetActiveMoneyAccountRequest{
		OrgId:     req.OrgId,
		MoneyAbbr: moneyAbbr,
	})
	if err != nil {
		return &api.GetActiveMoneyAccountResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	username, err := auth.JWTValidator{}.GetUsername(ctx)
	if nil != err {
		return nil, helpers.ErrToRPCError(err)
	}

	// Get the user id based on the username.
	user, err := storage.GetUserByUsername(ctx, storage.DB(), username)
	if nil != err {
		return nil, helpers.ErrToRPCError(err)
	}

	prof, err := storage.GetProfile(ctx, storage.DB(), user.ID)
	if err != nil {
		return nil, helpers.ErrToRPCError(err)
	}

	//userProfile := api.GetDeviceListResponse.GetUserProfile(prof)
	userProfile := api.ProfileResponse{
		User: &api.User{
			Id:         string(prof.User.ID),
			Username:   prof.User.Username,
			SessionTtl: prof.User.SessionTTL,
			IsAdmin:    prof.User.IsAdmin,
			IsActive:   prof.User.IsActive,
		},
		Settings: &api.ProfileSettings{
			DisableAssignExistingUsers: auth.DisableAssignExistingUsers,
		},
	}

	for _, org := range prof.Organizations {
		row := api.OrganizationLink{
			OrganizationId:   org.ID,
			OrganizationName: org.Name,
			IsAdmin:          org.IsAdmin,
		}

		row.CreatedAt, err = ptypes.TimestampProto(org.CreatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}
		row.UpdatedAt, err = ptypes.TimestampProto(org.UpdatedAt)
		if err != nil {
			return nil, helpers.ErrToRPCError(err)
		}

		userProfile.Organizations = append(userProfile.Organizations, &row)
	}

	return &api.GetActiveMoneyAccountResponse{
		ActiveAccount: resp.ActiveAccount,
		UserProfile:   &userProfile,
	}, nil
}
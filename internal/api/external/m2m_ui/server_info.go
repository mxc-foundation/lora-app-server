package m2m_ui

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/mxc-foundation/lpwan-app-server/api/m2m_ui"
	"github.com/mxc-foundation/lpwan-app-server/internal/api/external/auth"
	"github.com/mxc-foundation/lpwan-app-server/internal/backend/m2m_client"
	"github.com/mxc-foundation/lpwan-app-server/internal/config"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type M2MServerAPI struct {
	validator auth.Validator
}

func NewM2MServerAPI(validator auth.Validator) *M2MServerAPI {
	return &M2MServerAPI{
		validator: validator,
	}
}

func (s *M2MServerAPI) GetVersion(ctx context.Context, req *empty.Empty) (*api.GetVersionResponse, error) {
	log.WithField("", "").Info("grpc_api/GetVersion")

	/*prof, err := getUserProfileByJwt(s.validator, ctx, 0)
	if err != nil{
		return &api.GetVersionResponse{}, status.Errorf(codes.Unauthenticated, err.Error())
	}*/

	m2mClient, err := m2m_client.GetPool().Get(config.C.M2MServer.M2MServer, []byte(config.C.M2MServer.CACert),
		[]byte(config.C.M2MServer.TLSCert), []byte(config.C.M2MServer.TLSKey))
	if err != nil {
		return &api.GetVersionResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	verClient := api.NewServerInfoServiceClient(m2mClient)

	resp, err := verClient.GetVersion(ctx, nil)
	if err != nil {
		return &api.GetVersionResponse{}, status.Errorf(codes.Unavailable, err.Error())
	}

	return &api.GetVersionResponse{Version: resp.Version}, nil
}
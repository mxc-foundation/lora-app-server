package download

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"path/filepath"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/mxc-foundation/lpwan-app-server/api/appserver-serves-ui"
	pb "github.com/mxc-foundation/lpwan-app-server/api/m2m-serves-appserver"
	"github.com/mxc-foundation/lpwan-app-server/internal/auth"
)

type testMXPCli struct {
	pb.FinanceReportServiceClient
}

func (tc *testMXPCli) GetMXCMiningReportByDate(ctx context.Context, in *pb.GetMXCMiningReportByDateRequest,
	opts ...grpc.CallOption) (*pb.GetMXCMiningReportByDateResponse, error) {
	return &pb.GetMXCMiningReportByDateResponse{}, nil
}

type testAuth struct {
	auth.Authenticator
	validator func(opts *auth.Options) (*auth.Credentials, error)
}

func (ta *testAuth) GetCredentials(ctx context.Context, opts *auth.Options) (*auth.Credentials, error) {
	if ta.validator != nil {
		return ta.validator(opts)
	}
	return nil, fmt.Errorf("validator is not defined")
}

func TestGetMiningReportCSVFileURI(t *testing.T) {
	ctx := context.Background()
	mxpCli := &testMXPCli{}
	ta := &testAuth{
		validator: func(opts *auth.Options) (*auth.Credentials, error) {
			if opts.OrgID != 1 {
				return &auth.Credentials{
					UserID:     2,
					Username:   "user2",
					IsExisting: true,
					OrgID:      2,
				}, nil
			}
			return &auth.Credentials{
				UserID:     1,
				Username:   "user1",
				IsExisting: true,
				OrgID:      1,
				IsOrgAdmin: true,
			}, nil
		},
	}
	server := NewServer(mxpCli, ta, "local_test_server")
	request := api.MiningReportRequest{
		OrganizationId: 1,
		Currency:       []string{"ETH_MXC"},
		FiatCurrency:   "usd",
		Start:          timestamppb.New(time.Now().AddDate(-1, 0, 0)),
		End:            timestamppb.New(time.Now()),
		Decimals:       6,
	}
	resp, err := server.MiningReportCSV(ctx, &request)
	if err != nil {
		t.Fatal(err)
	}

	sy, sm, sd := request.Start.AsTime().Date()
	ey, em, ed := request.End.AsTime().Date()
	filename := fmt.Sprintf("mining_report_%s_org_%d_%s_%s_%s.csv", "local_test_server", 1, "usd",
		fmt.Sprintf("%04d-%02d-%02d", sy, sm, sd), fmt.Sprintf("%04d-%02d-%02d", ey, em, ed))
	// drawGrid(pdf, format)
	filePath := filepath.Join("/tmp/mining-report", filename)
	if resp.ReportUri != filePath {
		t.Fatalf("expected %s, got %s", filePath, resp.ReportUri)
	}
}

func TestGetMiningReportPDFFileURI(t *testing.T) {
	ctx := context.Background()
	mxpCli := &testMXPCli{}
	ta := &testAuth{
		validator: func(opts *auth.Options) (*auth.Credentials, error) {
			if opts.OrgID != 1 {
				return &auth.Credentials{
					UserID:     2,
					Username:   "user2",
					IsExisting: true,
					OrgID:      2,
				}, nil
			}
			return &auth.Credentials{
				UserID:     1,
				Username:   "user1",
				IsExisting: true,
				OrgID:      1,
				IsOrgAdmin: true,
			}, nil
		},
	}
	server := NewServer(mxpCli, ta, "local_test_server")
	request := api.MiningReportRequest{
		OrganizationId: 1,
		Currency:       []string{"ETH_MXC"},
		FiatCurrency:   "usd",
		Start:          timestamppb.New(time.Now().AddDate(-1, 0, 0)),
		End:            timestamppb.New(time.Now()),
		Decimals:       6,
	}
	resp, err := server.MiningReportPDF(ctx, &request)
	if err != nil {
		t.Fatal(err)
	}

	sy, sm, sd := request.Start.AsTime().Date()
	ey, em, ed := request.End.AsTime().Date()
	filename := fmt.Sprintf("mining_report_%s_org_%d_%s_%s_%s.pdf", "local_test_server", 1, "usd",
		fmt.Sprintf("%04d-%02d-%02d", sy, sm, sd), fmt.Sprintf("%04d-%02d-%02d", ey, em, ed))
	// drawGrid(pdf, format)
	filePath := filepath.Join("/tmp/mining-report", filename)
	if resp.ReportUri != filePath {
		t.Fatalf("expected %s, got %s", filePath, resp.ReportUri)
	}
}

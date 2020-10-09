package devprofile

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/brocaar/chirpstack-api/go/v3/ns"

	nscli "github.com/mxc-foundation/lpwan-app-server/internal/networkserver_portal"
	"github.com/mxc-foundation/lpwan-app-server/internal/storage/store"
	mgr "github.com/mxc-foundation/lpwan-app-server/internal/system_manager"

	. "github.com/mxc-foundation/lpwan-app-server/internal/modules/device-profile/data"
)

func init() {
	mgr.RegisterModuleSetup(moduleName, Setup)
}

const moduleName = "device_profile"

type controller struct {
	st *store.Handler
}

var ctrl *controller

func Setup(name string, h *store.Handler) error {
	if name != moduleName {
		return errors.New(fmt.Sprintf("Calling SettingsSetup for %s, but %s is called", name, moduleName))
	}
	ctrl = &controller{
		st: h,
	}

	return nil
}

// CreateDeviceProfile creates the given device-profile.
// This will create the device-profile at the network-server side and will
// create a local reference record.
func CreateDeviceProfile(ctx context.Context, dp *DeviceProfile) error {
	if err := ctrl.st.Tx(ctx, func(ctx context.Context, handler *store.Handler) error {
		if err := handler.CreateDeviceProfile(ctx, dp); err != nil {
			return err
		}

		n, err := handler.GetNetworkServer(ctx, dp.NetworkServerID)
		if err != nil {
			return errors.Wrap(err, "get network-server error")
		}

		nsStruct := nscli.NSStruct{
			Server:  n.Server,
			CACert:  n.CACert,
			TLSCert: n.TLSCert,
			TLSKey:  n.TLSKey,
		}
		nsClient, err := nsStruct.GetNetworkServiceClient()
		if err != nil {
			return errors.Wrap(err, "get network-server client error")
		}

		_, err = nsClient.CreateDeviceProfile(ctx, &ns.CreateDeviceProfileRequest{
			DeviceProfile: &dp.DeviceProfile,
		})
		if err != nil {
			return errors.Wrap(err, "create device-profile errror")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// UpdateDeviceProfile updates the given device-profile.
func UpdateDeviceProfile(ctx context.Context, dp *DeviceProfile) error {
	if err := ctrl.st.Tx(ctx, func(ctx context.Context, handler *store.Handler) error {
		if err := handler.UpdateDeviceProfile(ctx, dp); err != nil {
			return err
		}

		n, err := handler.GetNetworkServer(ctx, dp.NetworkServerID)
		if err != nil {
			return errors.Wrap(err, "get network-server error")
		}

		nstruct := nscli.NSStruct{
			Server:  n.Server,
			CACert:  n.CACert,
			TLSCert: n.TLSCert,
			TLSKey:  n.TLSKey,
		}

		nsClient, err := nstruct.GetNetworkServiceClient()
		if err != nil {
			return errors.Wrap(err, "get network-server client error")
		}
		_, err = nsClient.UpdateDeviceProfile(ctx, &ns.UpdateDeviceProfileRequest{
			DeviceProfile: &dp.DeviceProfile,
		})
		if err != nil {
			return errors.Wrap(err, "update device-profile error")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// DeleteDeviceProfile deletes the device-profile matching the given id.
func DeleteDeviceProfile(ctx context.Context, id uuid.UUID) error {
	if err := ctrl.st.Tx(ctx, func(ctx context.Context, handler *store.Handler) error {
		if err := handler.DeleteDeviceProfile(ctx, id); err != nil {
			return err
		}

		n, err := handler.GetNetworkServerForDeviceProfileID(ctx, id)
		if err != nil {
			return errors.Wrap(err, "get network-server error")
		}

		nsStruct := nscli.NSStruct{
			Server:  n.Server,
			CACert:  n.CACert,
			TLSCert: n.TLSCert,
			TLSKey:  n.TLSKey,
		}
		nsClient, err := nsStruct.GetNetworkServiceClient()
		if err != nil {
			return errors.Wrap(err, "get network-server client error")
		}

		_, err = nsClient.DeleteDeviceProfile(ctx, &ns.DeleteDeviceProfileRequest{
			Id: id.Bytes(),
		})
		if err != nil && status.Code(err) != codes.NotFound {
			return errors.Wrap(err, "delete device-profile error")
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

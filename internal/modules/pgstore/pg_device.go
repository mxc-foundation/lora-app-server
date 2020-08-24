package pgstore

import (
	"context"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/brocaar/chirpstack-api/go/v3/ns"
	"github.com/brocaar/lorawan"

	m2m_api "github.com/mxc-foundation/lpwan-app-server/api/m2m-serves-appserver"

	m2mcli "github.com/mxc-foundation/lpwan-app-server/internal/clients/mxprotocol-server"
	nscli "github.com/mxc-foundation/lpwan-app-server/internal/clients/networkserver"
	"github.com/mxc-foundation/lpwan-app-server/internal/logging"
	"github.com/mxc-foundation/lpwan-app-server/internal/modules/store"

	"github.com/mxc-foundation/lpwan-app-server/internal/modules/networkserver"
)

func (ps *pgstore) CheckCreateNodeAccess(ctx context.Context, username string, applicationID int64, userID int64) (bool, error) {
	userQuery := `
		select
			1
		from
			"user" u
		left join organization_user ou
			on u.id = ou.user_id
		left join organization o
			on o.id = ou.organization_id
		left join application a
			on a.organization_id = o.id
	`

	// global admin
	// organization admin
	// organization device admin
	var userWhere = [][]string{
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "u.is_admin = true"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "ou.is_admin = true", "a.id = $2"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "ou.is_device_admin = true", "a.id = $2"},
	}

	var ors []string
	for _, ands := range userWhere {
		ors = append(ors, "(("+strings.Join(ands, ") and (")+"))")
	}
	whereStr := strings.Join(ors, " or ")
	userQuery = "select count(*) from (" + userQuery + " where " + whereStr + " limit 1) count_only"

	var count int64
	if err := sqlx.GetContext(ctx, ps.db, &count, userQuery, username, applicationID, userID); err != nil {
		return false, errors.Wrap(err, "select error")
	}
	return count > 0, nil
}

func (ps *pgstore) CheckListNodeAccess(ctx context.Context, username string, applicationID int64, userID int64) (bool, error) {
	userQuery := `
		select
			1
		from
			"user" u
		left join organization_user ou
			on u.id = ou.user_id
		left join organization o
			on o.id = ou.organization_id
		left join application a
			on a.organization_id = o.id
	`
	// global admin
	// organization user
	var userWhere = [][]string{
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "u.is_admin = true"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "a.id = $2"},
	}

	var ors []string
	for _, ands := range userWhere {
		ors = append(ors, "(("+strings.Join(ands, ") and (")+"))")
	}
	whereStr := strings.Join(ors, " or ")
	userQuery = "select count(*) from (" + userQuery + " where " + whereStr + " limit 1) count_only"

	var count int64
	if err := sqlx.GetContext(ctx, ps.db, &count, userQuery, username, applicationID, userID); err != nil {
		return false, errors.Wrap(err, "select error")
	}
	return count > 0, nil
}

func (ps *pgstore) CheckReadNodeAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error) {
	userQuery := `
		select
			1
		from
			"user" u
		left join organization_user ou
			on u.id = ou.user_id
		left join organization o
			on o.id = ou.organization_id
		left join application a
			on a.organization_id = o.id
		left join device d
			on a.id = d.application_id
	`
	// global admin
	// organization user
	var userWhere = [][]string{
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "u.is_admin = true"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "d.dev_eui = $2"},
	}

	var ors []string
	for _, ands := range userWhere {
		ors = append(ors, "(("+strings.Join(ands, ") and (")+"))")
	}
	whereStr := strings.Join(ors, " or ")
	userQuery = "select count(*) from (" + userQuery + " where " + whereStr + " limit 1) count_only"

	var count int64
	if err := sqlx.GetContext(ctx, ps.db, &count, userQuery, username, devEUI[:], userID); err != nil {
		return false, errors.Wrap(err, "select error")
	}
	return count > 0, nil
}

func (ps *pgstore) CheckUpdateNodeAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error) {
	userQuery := `
		select
			1
		from
			"user" u
		left join organization_user ou
			on u.id = ou.user_id
		left join organization o
			on o.id = ou.organization_id
		left join application a
			on a.organization_id = o.id
		left join device d
			on a.id = d.application_id
	`
	// global admin
	// organization admin
	// organization device admin
	var userWhere = [][]string{
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "u.is_admin = true"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "ou.is_admin = true", "d.dev_eui = $2"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "ou.is_device_admin = true", "d.dev_eui = $2"},
	}

	var ors []string
	for _, ands := range userWhere {
		ors = append(ors, "(("+strings.Join(ands, ") and (")+"))")
	}
	whereStr := strings.Join(ors, " or ")
	userQuery = "select count(*) from (" + userQuery + " where " + whereStr + " limit 1) count_only"

	var count int64
	if err := sqlx.GetContext(ctx, ps.db, &count, userQuery, username, devEUI[:], userID); err != nil {
		return false, errors.Wrap(err, "select error")
	}
	return count > 0, nil
}

func (ps *pgstore) CheckDeleteNodeAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error) {
	userQuery := `
		select
			1
		from
			"user" u
		left join organization_user ou
			on u.id = ou.user_id
		left join organization o
			on o.id = ou.organization_id
		left join application a
			on a.organization_id = o.id
		left join device d
			on a.id = d.application_id
	`
	// global admin
	// organization admin
	// organization device admin
	var userWhere = [][]string{
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "u.is_admin = true"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "ou.is_admin = true", "d.dev_eui = $2"},
		{"(u.email = $1 or u.id = $3)", "u.is_active = true", "ou.is_device_admin = true", "d.dev_eui = $2"},
	}

	var ors []string
	for _, ands := range userWhere {
		ors = append(ors, "(("+strings.Join(ands, ") and (")+"))")
	}
	whereStr := strings.Join(ors, " or ")
	userQuery = "select count(*) from (" + userQuery + " where " + whereStr + " limit 1) count_only"

	var count int64
	if err := sqlx.GetContext(ctx, ps.db, &count, userQuery, username, devEUI[:], userID); err != nil {
		return false, errors.Wrap(err, "select error")
	}
	return count > 0, nil
}

// UpdateDeviceActivation updates the device address and the AppSKey.
func (ps *pgstore) UpdateDeviceActivation(ctx context.Context, devEUI lorawan.EUI64, devAddr lorawan.DevAddr, appSKey lorawan.AES128Key) error {
	res, err := ps.db.ExecContext(ctx, `
		update device
		set
			dev_addr = $2,
			app_s_key = $3
		where
			dev_eui = $1`,
		devEUI[:],
		devAddr[:],
		appSKey[:],
	)
	if err != nil {
		return errors.Wrap(err, "update last-seen and dr error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errors.New("ErrDoesNotExist")
	}

	log.WithFields(log.Fields{
		"dev_eui":  devEUI,
		"dev_addr": devAddr,
		"ctx_id":   ctx.Value(logging.ContextIDKey),
	}).Info("device activation updated")

	return nil
}

// CreateDevice creates the given device.
func (ps *pgstore) CreateDevice(ctx context.Context, d *store.Device, applicationServerID uuid.UUID) error {
	if err := d.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	_, err := ps.db.ExecContext(ctx, `
		insert into device (
			dev_eui,
			created_at,
			updated_at,
			application_id,
			device_profile_id,
			name,
			description,
			device_status_battery,
			device_status_margin,
			device_status_external_power_source,
			last_seen_at,
			latitude,
			longitude,
			altitude,
			dr,
			variables,
			tags,
			dev_addr,
			app_s_key
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`,
		d.DevEUI[:],
		d.CreatedAt,
		d.UpdatedAt,
		d.ApplicationID,
		d.DeviceProfileID,
		d.Name,
		d.Description,
		d.DeviceStatusBattery,
		d.DeviceStatusMargin,
		d.DeviceStatusExternalPower,
		d.LastSeenAt,
		d.Latitude,
		d.Longitude,
		d.Altitude,
		d.DR,
		d.Variables,
		d.Tags,
		d.DevAddr[:],
		d.AppSKey,
	)
	if err != nil {
		return errors.Wrap(err, "insert error")
	}

	log.WithFields(log.Fields{
		"dev_eui": d.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device created")

	return nil
}

// GetDevice returns the device matching the given DevEUI.
// When forUpdate is set to true, then tx must be a tx transaction.
// When localOnly is set to true, no call to the network-server is made to
// retrieve additional device data.
func (ps *pgstore) GetDevice(ctx context.Context, devEUI lorawan.EUI64, forUpdate bool) (store.Device, error) {
	var fu string
	if forUpdate {
		fu = " for update"
	}

	var d store.Device
	err := sqlx.GetContext(ctx, ps.db, &d, "select * from device where dev_eui = $1"+fu, devEUI[:])
	if err != nil {
		return d, errors.Wrap(err, "select error")
	}

	return d, nil
}

// GetDeviceCount returns the number of devices.
func (ps *pgstore) GetDeviceCount(ctx context.Context, filters store.DeviceFilters) (int, error) {
	if filters.Search != "" {
		filters.Search = "%" + filters.Search + "%"
	}

	query, args, err := sqlx.BindNamed(sqlx.DOLLAR, `
		select
			count(distinct d.*)
		from device d
		inner join application a
			on d.application_id = a.id
		left join device_multicast_group dmg
			on d.dev_eui = dmg.dev_eui
	`+filters.SQL(), filters)
	if err != nil {
		return 0, errors.Wrap(err, "named query error")
	}

	var count int
	err = sqlx.GetContext(ctx, ps.db, &count, query, args...)
	if err != nil {
		return 0, errors.Wrap(err, "select query error")
	}

	return count, nil
}

// GetAllDeviceEuis returns a slice of devices.
func (ps *pgstore) GetAllDeviceEuis(ctx context.Context) ([]string, error) {
	var devEuiList []string
	var list []lorawan.EUI64
	err := sqlx.SelectContext(ctx, ps.db, &list, "select dev_eui from device ORDER BY created_at DESC")
	if err != nil {
		return nil, errors.Wrap(err, "select error")
	}

	for _, devEui := range list {
		devEuiList = append(devEuiList, devEui.String())
	}

	return devEuiList, nil
}

// GetDevices returns a slice of devices.
func (ps *pgstore) GetDevices(ctx context.Context, filters store.DeviceFilters) ([]store.DeviceListItem, error) {
	if filters.Search != "" {
		filters.Search = "%" + filters.Search + "%"
	}

	query, args, err := sqlx.BindNamed(sqlx.DOLLAR, `
		select
			distinct d.*,
			dp.name as device_profile_name
		from
			device d
		inner join device_profile dp
			on dp.device_profile_id = d.device_profile_id
		inner join application a
			on d.application_id = a.id
		left join device_multicast_group dmg
			on d.dev_eui = dmg.dev_eui
		`+filters.SQL()+`
		order by
			d.name
		limit :limit
		offset :offset
	`, filters)
	if err != nil {
		return nil, errors.Wrap(err, "named query error")
	}

	var devices []store.DeviceListItem
	err = sqlx.SelectContext(ctx, ps.db, &devices, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "select error")
	}

	return devices, nil
}

// UpdateDevice updates the given device.
// When localOnly is set, it will not update the device on the network-server.
func (ps *pgstore) UpdateDevice(ctx context.Context, d *store.Device) error {
	if err := d.Validate(); err != nil {
		return errors.Wrap(err, "validate error")
	}

	d.UpdatedAt = time.Now()

	res, err := ps.db.ExecContext(ctx, `
        update device
        set
            updated_at = $2,
            application_id = $3,
            device_profile_id = $4,
            name = $5,
			description = $6,
			device_status_battery = $7,
			device_status_margin = $8,
			last_seen_at = $9,
			latitude = $10,
			longitude = $11,
			altitude = $12,
			device_status_external_power_source = $13,
			dr = $14,
			variables = $15,
			tags = $16
        where
            dev_eui = $1`,
		d.DevEUI[:],
		d.UpdatedAt,
		d.ApplicationID,
		d.DeviceProfileID,
		d.Name,
		d.Description,
		d.DeviceStatusBattery,
		d.DeviceStatusMargin,
		d.LastSeenAt,
		d.Latitude,
		d.Longitude,
		d.Altitude,
		d.DeviceStatusExternalPower,
		d.DR,
		d.Variables,
		d.Tags,
	)
	if err != nil {
		return errors.Wrap(err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errors.New("not exist")
	}

	log.WithFields(log.Fields{
		"dev_eui": d.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device updated")

	return nil
}

// DeleteDevice deletes the device matching the given DevEUI.
func (ps *pgstore) DeleteDevice(ctx context.Context, devEUI lorawan.EUI64) error {
	n, err := networkserver.Service.St.GetNetworkServerForDevEUI(ctx, devEUI)
	if err != nil {
		return errors.Wrap(err, "get network-server error")
	}

	res, err := ps.db.ExecContext(ctx, "delete from device where dev_eui = $1", devEUI[:])
	if err != nil {
		return errors.Wrap(err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errors.New("not exist")
	}

	// delete device from networkserver
	nsStruct := nscli.NSStruct{
		Server:  n.Server,
		CACert:  n.CACert,
		TLSCert: n.TLSCert,
		TLSKey:  n.TLSKey,
	}
	client, err := nsStruct.GetNetworkServiceClient()
	if err != nil {
		return errors.Wrap(err, "get network-server client error")
	}

	_, err = client.DeleteDevice(ctx, &ns.DeleteDeviceRequest{
		DevEui: devEUI[:],
	})
	if err != nil && status.Code(err) != codes.NotFound {
		return errors.Wrap(err, "delete device error")
	}

	// delete device from m2m server, this procedure should not block delete device from appserver once it's deleted from
	// network server successfully
	dvClient, err := m2mcli.GetM2MDeviceServiceClient()
	if err == nil {
		_, err = dvClient.DeleteDeviceInM2MServer(context.Background(), &m2m_api.DeleteDeviceInM2MServerRequest{
			DevEui: devEUI.String(),
		})
		if err != nil && status.Code(err) != codes.NotFound {
			log.WithError(err).Error("m2m-server delete device api error")
		}
	} else {
		log.WithError(err).Error("get m2m-server client error")
	}

	log.WithFields(log.Fields{
		"dev_eui": devEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device deleted")

	return nil
}

// CreateDeviceKeys creates the keys for the given device.
func (ps *pgstore) CreateDeviceKeys(ctx context.Context, dc *store.DeviceKeys) error {
	now := time.Now()
	dc.CreatedAt = now
	dc.UpdatedAt = now

	_, err := ps.db.ExecContext(ctx, `
        insert into device_keys (
            created_at,
            updated_at,
            dev_eui,
			nwk_key,
			app_key,
			join_nonce,
			gen_app_key
        ) values ($1, $2, $3, $4, $5, $6, $7)`,
		dc.CreatedAt,
		dc.UpdatedAt,
		dc.DevEUI[:],
		dc.NwkKey[:],
		dc.AppKey[:],
		dc.JoinNonce,
		dc.GenAppKey[:],
	)
	if err != nil {
		return errors.Wrap(err, "insert error")
	}

	log.WithFields(log.Fields{
		"dev_eui": dc.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device-keys created")

	return nil
}

// GetDeviceKeys returns the device-keys for the given DevEUI.
func (ps *pgstore) GetDeviceKeys(ctx context.Context, devEUI lorawan.EUI64) (store.DeviceKeys, error) {
	var dc store.DeviceKeys

	err := sqlx.GetContext(ctx, ps.db, &dc, "select * from device_keys where dev_eui = $1", devEUI[:])
	if err != nil {
		return dc, errors.Wrap(err, "select error")
	}

	return dc, nil
}

// UpdateDeviceKeys updates the given device-keys.
func (ps *pgstore) UpdateDeviceKeys(ctx context.Context, dc *store.DeviceKeys) error {
	dc.UpdatedAt = time.Now()

	res, err := ps.db.ExecContext(ctx, `
        update device_keys
        set
            updated_at = $2,
			nwk_key = $3,
			app_key = $4,
			join_nonce = $5,
			gen_app_key = $6
        where
            dev_eui = $1`,
		dc.DevEUI[:],
		dc.UpdatedAt,
		dc.NwkKey[:],
		dc.AppKey[:],
		dc.JoinNonce,
		dc.GenAppKey[:],
	)
	if err != nil {
		return errors.Wrap(err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errors.New("not exist")
	}

	log.WithFields(log.Fields{
		"dev_eui": dc.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device-keys updated")

	return nil
}

// DeleteDeviceKeys deletes the device-keys for the given DevEUI.
func (ps *pgstore) DeleteDeviceKeys(ctx context.Context, devEUI lorawan.EUI64) error {
	res, err := ps.db.ExecContext(ctx, "delete from device_keys where dev_eui = $1", devEUI[:])
	if err != nil {
		return errors.Wrap(err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected errro")
	}
	if ra == 0 {
		return errors.New("not exist")
	}

	log.WithFields(log.Fields{
		"dev_eui": devEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device-keys deleted")

	return nil
}

// CreateDeviceActivation creates the given device-activation.
func (ps *pgstore) CreateDeviceActivation(ctx context.Context, da *store.DeviceActivation) error {
	da.CreatedAt = time.Now()

	err := sqlx.GetContext(ctx, ps.db, &da.ID, `
        insert into device_activation (
            created_at,
            dev_eui,
            dev_addr,
			app_s_key
        ) values ($1, $2, $3, $4)
        returning id`,
		da.CreatedAt,
		da.DevEUI[:],
		da.DevAddr[:],
		da.AppSKey[:],
	)
	if err != nil {
		return errors.Wrap(err, "insert error")
	}

	log.WithFields(log.Fields{
		"id":      da.ID,
		"dev_eui": da.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device-activation created")

	return nil
}

// GetLastDeviceActivationForDevEUI returns the most recent device-activation for the given DevEUI.
func (ps *pgstore) GetLastDeviceActivationForDevEUI(ctx context.Context, devEUI lorawan.EUI64) (store.DeviceActivation, error) {
	var da store.DeviceActivation

	err := sqlx.GetContext(ctx, ps.db, &da, `
        select *
        from device_activation
        where
            dev_eui = $1
        order by
            created_at desc
        limit 1`,
		devEUI[:],
	)
	if err != nil {
		return da, errors.Wrap(err, "select error")
	}

	return da, nil
}

// DeleteAllDevicesForApplicationID deletes all devices given an application id.
func (ps *pgstore) DeleteAllDevicesForApplicationID(ctx context.Context, applicationID int64) error {
	var devs []store.Device
	err := sqlx.SelectContext(ctx, ps.db, &devs, "select * from device where application_id = $1", applicationID)
	if err != nil {
		return errors.Wrap(err, "select error")
	}

	for _, dev := range devs {
		err = ps.DeleteDevice(ctx, dev.DevEUI)
		if err != nil {
			return errors.Wrap(err, "delete device error")
		}
	}

	return nil
}

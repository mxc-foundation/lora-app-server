package pgstore

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/brocaar/lorawan"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	errHandler "github.com/mxc-foundation/lpwan-app-server/internal/errors"
	"github.com/mxc-foundation/lpwan-app-server/internal/logging"
	. "github.com/mxc-foundation/lpwan-app-server/internal/modules/device/data"
)

type DevicePgstore interface {
	CheckCreateNodeAccess(ctx context.Context, username string, applicationID int64, userID int64) (bool, error)
	CheckListNodeAccess(ctx context.Context, username string, applicationID int64, userID int64) (bool, error)
	CheckReadNodeAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error)
	CheckUpdateNodeAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error)
	CheckDeleteNodeAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error)
	CheckCreateListDeleteDeviceQueueAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error)
	UpdateDeviceActivation(ctx context.Context, devEUI lorawan.EUI64, devAddr lorawan.DevAddr, appSKey lorawan.AES128Key) error
	CreateDevice(ctx context.Context, d *Device) error
	GetDevice(ctx context.Context, devEUI lorawan.EUI64, forUpdate bool) (Device, error)
	GetDeviceCount(ctx context.Context, filters DeviceFilters) (int, error)
	GetAllDeviceEuis(ctx context.Context) ([]string, error)
	GetDevices(ctx context.Context, filters DeviceFilters) ([]DeviceListItem, error)
	UpdateDevice(ctx context.Context, d *Device) error
	UpdateDeviceLastSeenAndDR(ctx context.Context, devEUI lorawan.EUI64, ts time.Time, dr int) error
	DeleteDevice(ctx context.Context, devEUI lorawan.EUI64) error
	CreateDeviceKeys(ctx context.Context, dc *DeviceKeys) error
	GetDeviceKeys(ctx context.Context, devEUI lorawan.EUI64) (DeviceKeys, error)
	UpdateDeviceKeys(ctx context.Context, dc *DeviceKeys) error
	DeleteDeviceKeys(ctx context.Context, devEUI lorawan.EUI64) error
	CreateDeviceActivation(ctx context.Context, da *DeviceActivation) error
	GetLastDeviceActivationForDevEUI(ctx context.Context, devEUI lorawan.EUI64) (DeviceActivation, error)
	DeleteAllDevicesForApplicationID(ctx context.Context, applicationID int64) error
	//EnqueueDownlinkPayload(ctx context.Context, devEUI lorawan.EUI64, confirmed bool, fPort uint8, data []byte) (uint32, error)
	GetDevicesActiveInactive(ctx context.Context, organizationID int64) (DevicesActiveInactive, error)
	GetDevicesDataRates(ctx context.Context, organizationID int64) (DevicesDataRates, error)
}

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

func (ps *pgstore) CheckCreateListDeleteDeviceQueueAccess(ctx context.Context, username string, devEUI lorawan.EUI64, userID int64) (bool, error) {
	userQuery := `
		select
			1
		from
			"user" u
		left join organization_user ou
			on u.id = ou.user_id
		left join application a
			on a.organization_id = ou.organization_id
		left join device d
			on a.id = d.application_id
	`
	// global admin
	// organization user
	userWhere := [][]string{
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
		return handlePSQLError(Update, err, "update last-seen and dr error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errHandler.ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"dev_eui":  devEUI,
		"dev_addr": devAddr,
		"ctx_id":   ctx.Value(logging.ContextIDKey),
	}).Info("device activation updated")

	return nil
}

// CreateDevice creates the given device.
func (ps *pgstore) CreateDevice(ctx context.Context, d *Device) error {
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
		return handlePSQLError(Insert, err, "insert error")
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
func (ps *pgstore) GetDevice(ctx context.Context, devEUI lorawan.EUI64, forUpdate bool) (Device, error) {
	var fu string
	if forUpdate {
		fu = " for update"
	}

	var d Device
	err := sqlx.GetContext(ctx, ps.db, &d, "select * from device where dev_eui = $1"+fu, devEUI[:])
	if err != nil {
		if err == sql.ErrNoRows {
			return d, errHandler.ErrDoesNotExist
		}
		return d, errors.Wrap(err, "select error")
	}

	return d, nil
}

// GetDeviceCount returns the number of devices.
func (ps *pgstore) GetDeviceCount(ctx context.Context, filters DeviceFilters) (int, error) {
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
		return 0, handlePSQLError(Select, err, "select error")
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
func (ps *pgstore) GetDevices(ctx context.Context, filters DeviceFilters) ([]DeviceListItem, error) {
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

	var devices []DeviceListItem
	err = sqlx.SelectContext(ctx, ps.db, &devices, query, args...)
	if err != nil {
		return nil, handlePSQLError(Select, err, "select error")
	}

	return devices, nil
}

// UpdateDevice updates the given device.
// When localOnly is set, it will not update the device on the network-server.
func (ps *pgstore) UpdateDevice(ctx context.Context, d *Device) error {
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
			tags = $16,
			dev_addr = $17,
			app_s_key = $18
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
		d.DevAddr,
		d.AppSKey,
	)
	if err != nil {
		return handlePSQLError(Update, err, "update error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errHandler.ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"dev_eui": d.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device updated")

	return nil
}

// UpdateDeviceLastSeenAndDR updates the device last-seen timestamp and data-rate.
func (ps *pgstore) UpdateDeviceLastSeenAndDR(ctx context.Context, devEUI lorawan.EUI64, ts time.Time, dr int) error {
	res, err := ps.db.ExecContext(ctx, `
		update device
		set
			last_seen_at = $2,
			dr = $3
		where
			dev_eui = $1`,
		devEUI[:],
		ts,
		dr,
	)
	if err != nil {
		return handlePSQLError(Update, err, "update last-seen and dr error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errHandler.ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"dev_eui": devEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device last-seen and dr updated")

	return nil
}

// DeleteDevice deletes the device matching the given DevEUI.
func (ps *pgstore) DeleteDevice(ctx context.Context, devEUI lorawan.EUI64) error {
	res, err := ps.db.ExecContext(ctx, "delete from device where dev_eui = $1", devEUI[:])
	if err != nil {
		return handlePSQLError(Delete, err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected error")
	}
	if ra == 0 {
		return errors.New("not exist")
	}

	log.WithFields(log.Fields{
		"dev_eui": devEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device deleted")

	return nil
}

// CreateDeviceKeys creates the keys for the given device.
func (ps *pgstore) CreateDeviceKeys(ctx context.Context, dc *DeviceKeys) error {
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
		return handlePSQLError(Insert, err, "insert error")
	}

	log.WithFields(log.Fields{
		"dev_eui": dc.DevEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device-keys created")

	return nil
}

// GetDeviceKeys returns the device-keys for the given DevEUI.
func (ps *pgstore) GetDeviceKeys(ctx context.Context, devEUI lorawan.EUI64) (DeviceKeys, error) {
	var dc DeviceKeys

	err := sqlx.GetContext(ctx, ps.db, &dc, "select * from device_keys where dev_eui = $1", devEUI[:])
	if err != nil {
		return dc, handlePSQLError(Select, err, "select error")
	}

	return dc, nil
}

// UpdateDeviceKeys updates the given device-keys.
func (ps *pgstore) UpdateDeviceKeys(ctx context.Context, dc *DeviceKeys) error {
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
		return handlePSQLError(Update, err, "update error")
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
		return handlePSQLError(Delete, err, "delete error")
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "get rows affected errro")
	}
	if ra == 0 {
		return errHandler.ErrDoesNotExist
	}

	log.WithFields(log.Fields{
		"dev_eui": devEUI,
		"ctx_id":  ctx.Value(logging.ContextIDKey),
	}).Info("device-keys deleted")

	return nil
}

// CreateDeviceActivation creates the given device-activation.
func (ps *pgstore) CreateDeviceActivation(ctx context.Context, da *DeviceActivation) error {
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
func (ps *pgstore) GetLastDeviceActivationForDevEUI(ctx context.Context, devEUI lorawan.EUI64) (DeviceActivation, error) {
	var da DeviceActivation

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
	var devs []Device
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

// GetDevicesActiveInactive returns the active / inactive devices.
func (ps *pgstore) GetDevicesActiveInactive(ctx context.Context, organizationID int64) (DevicesActiveInactive, error) {
	var out DevicesActiveInactive
	err := sqlx.GetContext(ctx, ps.db, &out, `
		with device_active_inactive as (
			select
				make_interval(secs => dp.uplink_interval / 1000000000) * 1.5 as uplink_interval,
				d.last_seen_at as last_seen_at
			from
				device d
			inner join device_profile dp
				on d.device_profile_id = dp.device_profile_id
			inner join application a
				on d.application_id = a.id
			where
				$1 = 0 or a.organization_id = $1
		)
		select
			coalesce(sum(case when last_seen_at is null then 1 end), 0) as never_seen_count,
			coalesce(sum(case when (now() - uplink_interval) > last_seen_at then 1 end), 0) as inactive_count,
			coalesce(sum(case when (now() - uplink_interval) <= last_seen_at then 1 end), 0) as active_count
		from
			device_active_inactive
	`, organizationID)
	if err != nil {
		return out, errors.Wrap(err, "get device active/inactive count error")
	}

	return out, nil
}

// GetDevicesDataRates returns the device counts by data-rate.
func (ps *pgstore) GetDevicesDataRates(ctx context.Context, organizationID int64) (DevicesDataRates, error) {
	out := make(DevicesDataRates)

	rows, err := ps.db.QueryxContext(ctx, `
		select
			d.dr,
			count(1)
		from
			device d
		inner join application a
			on d.application_id = a.id
		where
			($1 = 0 or a.organization_id = $1)
			and d.dr is not null
		group by d.dr
	`, organizationID)
	if err != nil {
		return out, errors.Wrap(err, "get device count per data-rate error")
	}

	for rows.Next() {
		var dr, count uint32

		if err := rows.Scan(&dr, &count); err != nil {
			return out, errors.Wrap(err, "scan row error")
		}

		out[dr] = count
	}

	return out, nil
}
package abi

// Operation identifies a request the connector can handle.
type Operation string

const (
	// Connection.
	Describe        Operation = "describe"
	PrepareSettings Operation = "prepare_settings"
	Probe           Operation = "probe"
	Capabilities    Operation = "capabilities"

	// Zones.
	ListZones       Operation = "list_zones"
	ListZonesPage   Operation = "list_zones_page"
	GetZone         Operation = "get_zone"
	GetZoneMetadata Operation = "get_zone_metadata"
	CreateZone      Operation = "create_zone"
	UpdateZone      Operation = "update_zone"
	DeleteZone      Operation = "delete_zone"

	// Zone settings.
	GetZoneSettings   Operation = "get_zone_settings"
	PatchZoneSettings Operation = "patch_zone_settings"

	// RRsets.
	ListRRsetsPage   Operation = "list_rrsets_page"
	GetRRset         Operation = "get_rrset"
	ReplaceRRset     Operation = "replace_rrset"
	DeleteRRset      Operation = "delete_rrset"
	GetRRsetByID     Operation = "get_rrset_by_id"
	ReplaceRRsetByID Operation = "replace_rrset_by_id"
	DeleteRRsetByID  Operation = "delete_rrset_by_id"

	// Individual records.
	ListRecords     Operation = "list_records"
	ListRecordsPage Operation = "list_records_page"
	GetRecord       Operation = "get_record"
	CreateRecord    Operation = "create_record"
	UpdateRecord    Operation = "update_record"
	DeleteRecord    Operation = "delete_record"

	// Record batches.
	SubmitRecordBatch    Operation = "submit_record_batch"
	GetRecordBatchStatus Operation = "get_record_batch_status"

	// DNSSEC.
	GetDNSSEC Operation = "get_dnssec"
	SetDNSSEC Operation = "set_dnssec"

	// Zone actions.
	RectifyZone Operation = "rectify_zone"
	NotifyZone  Operation = "notify_zone"
	AxfrZone    Operation = "axfr_zone"
)

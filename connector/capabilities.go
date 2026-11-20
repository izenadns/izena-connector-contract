package connector

// ZoneAction identifies an operation handled by the connector.
type ZoneAction string

const (
	ActionRectify ZoneAction = "rectify"
	ActionNotify  ZoneAction = "notify"
	ActionAXFR    ZoneAction = "axfr"
)

type SigningCapability string

const (
	SigningSupported   SigningCapability = "supported"
	SigningUnsupported SigningCapability = "unsupported"
	SigningUnknown     SigningCapability = "unknown"
)

type DNSSECCapabilities struct {
	Inspect     bool
	KeyMetadata bool
	Signing     SigningCapability
}

// RecordModel identifies how a connection addresses and modifies DNS records.
type RecordModel string

const (
	RecordModelRRsets     RecordModel = "rrsets"
	RecordModelIndividual RecordModel = "records"
)

type ServerCapabilities struct {
	Provider  string
	ZoneKinds []ZoneKindOption

	RecordModel       RecordModel
	RecordTypes       []string
	CustomRecordTypes bool

	Pagination PaginationCapabilities
	DNSSEC     DNSSECCapabilities
	Actions    []ZoneAction
}

// ZoneKindOption describes a provider-defined zone kind and its creation inputs.
type ZoneKindOption struct {
	ID     string
	Label  string
	Fields []ZoneCreationField
}

type ZoneCreationField struct {
	Key         string
	Label       string
	Description string
	Placeholder string
	Required    bool
}

type ZoneCapabilities struct {
	Settings        []SettingField `json:",omitempty"`
	UpdateZoneKinds []string
	EditSOA         bool

	RecordTypes        []string
	ManagedRecordTypes []string
	CustomRecordTypes  bool

	WriteRRsets     bool
	WriteRecords    bool
	RecordComments  bool
	RecordDisabling bool
	RRsetIDs        bool                     `json:",omitempty"`
	AutomaticTTL    bool                     `json:",omitempty"`
	RecordOptions   []SettingField           `json:",omitempty"`
	RecordBatch     *RecordBatchCapabilities `json:",omitempty"`

	DNSSEC  DNSSECCapabilities
	Actions []ZoneAction
}

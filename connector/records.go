package connector

import "context"

// IndividualRecordOperations is implemented by connectors that advertise
// RecordModelIndividual. Unsupported writes return ErrUnsupported.
type IndividualRecordOperations interface {
	ListRecords(ctx context.Context, zoneID string) ([]IndividualRecord, error)
	GetRecord(ctx context.Context, zoneID, recordID string) (IndividualRecord, error)
	CreateRecord(ctx context.Context, zoneID string, input IndividualRecordInput) (IndividualRecord, error)
	UpdateRecord(ctx context.Context, zoneID, recordID string, input IndividualRecordInput) (IndividualRecord, error)
	DeleteRecord(ctx context.Context, zoneID, recordID string) error
}

// IndividualRecord identifies a record independently of its name or type.
// IDs are opaque and scoped to the zone in which the record was read.
type IndividualRecord struct {
	ID string
	IndividualRecordInput
}

// IndividualRecordInput describes a complete record to create or update.
type IndividualRecordInput struct {
	Name     string
	Type     string
	TTL      int64
	Content  string
	Disabled bool
	Comments []Comment
	// AutomaticTTL requests a provider-selected TTL. Leave TTL at zero on writes in this mode.
	AutomaticTTL bool                    `json:",omitempty"`
	Options      map[string]SettingValue `json:",omitempty"`
}

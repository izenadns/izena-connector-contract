package connector

import "context"

type RecordBatchOperations interface {
	SubmitRecordBatch(ctx context.Context, zoneID string, input RecordBatchInput) (RecordBatchReceipt, error)
	GetRecordBatchStatus(ctx context.Context, zoneID, batchID string) (RecordBatchStatus, error)
}

type RecordBatchCapabilities struct {
	// Atomic applies to acceptance and upstream application, not DNS propagation.
	Atomic           bool
	TrackPropagation bool
}

type RecordChangeAction string

const (
	RecordChangeCreate  RecordChangeAction = "create"
	RecordChangeReplace RecordChangeAction = "replace"
	RecordChangeDelete  RecordChangeAction = "delete"
)

// RecordChange selects exactly one record model. Create supplies input,
// replace supplies the target and input, and delete supplies only the target.
type RecordChange struct {
	Action RecordChangeAction
	RRset  *RRsetChange            `json:",omitempty"`
	Record *IndividualRecordChange `json:",omitempty"`
}

// RRsetChange identifies an existing RRset by ID or by name and type.
type RRsetChange struct {
	ID    string             `json:",omitempty"`
	Name  string             `json:",omitempty"`
	Type  string             `json:",omitempty"`
	Input *RRsetReplaceInput `json:",omitempty"`
}

type IndividualRecordChange struct {
	ID    string                 `json:",omitempty"`
	Input *IndividualRecordInput `json:",omitempty"`
}

// RecordBatchInput uses the selected record model for every change.
type RecordBatchInput struct {
	Changes []RecordChange
}

// RecordBatchReceipt indicates acceptance. An ID is required for a batch with
// trackable propagation and is scoped to the connection and zone.
type RecordBatchReceipt struct {
	ID string `json:",omitempty"`
}

type PropagationState string

const (
	PropagationUnknown  PropagationState = "unknown"
	PropagationPending  PropagationState = "pending"
	PropagationComplete PropagationState = "complete"
	PropagationFailed   PropagationState = "failed"
)

type RecordBatchStatus struct {
	Propagation PropagationState
}

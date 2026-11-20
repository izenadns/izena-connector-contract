// Package abi defines the versioned messages exchanged by a connector
// reactor and Izena. The host owns credentials and network access.
package abi

import (
	"bytes"
	"encoding/json"

	"github.com/izenadns/izena-connector-contract/connector"
)

const (
	Version         = 1
	ImportModule    = "izena_connector_v1"
	MaxMessageBytes = 64 << 20
)

// Individual-record operations use ZoneID. GetRecord, UpdateRecord, and
// DeleteRecord also use RecordID. CreateRecord and UpdateRecord take an
// IndividualRecordInput in Input and return an IndividualRecord in Result.Value.
// GetRecord returns an IndividualRecord. ListRecords returns a slice of
// IndividualRecord. DeleteRecord has no result value.
//
// Zone settings use ZoneID and a ZoneSettingsPatch for writes. The by-ID RRset
// operations use ZoneID and RRsetID. Paginated operations take PageInput and
// return Page with the corresponding item type. Batch submission takes
// RecordBatchInput and returns a RecordBatchReceipt. Status uses ZoneID and
// BatchID and returns RecordBatchStatus.
type Request struct {
	Operation Operation
	ZoneID    string            `json:",omitempty"`
	RRsetID   string            `json:",omitempty"`
	BatchID   string            `json:",omitempty"`
	RecordID  string            `json:",omitempty"`
	Name      string            `json:",omitempty"`
	Type      string            `json:",omitempty"`
	Settings  map[string]string `json:",omitempty"`
	Input     json.RawMessage   `json:",omitempty"`
}

// Result carries an optional failure and any operation result. A successful
// operation with a return value must supply a non-null Value. An operation
// without a return value can omit Value. A failure can omit Value or include
// a partial result when the operation defines one.
type Result struct {
	Value json.RawMessage `json:",omitempty"`
	Error *Failure        `json:",omitempty"`
}

// Manifest declares how the host binds HTTP to persisted connection fields.
// Settings are passed to the guest separately from secret values.
type Manifest struct {
	Description connector.Description
	HTTP        HTTPBinding
}

type HTTPBinding struct {
	URLSetting  string
	CASetting   string
	Credentials []HTTPHeaderCredential
}

type HTTPHeaderCredential struct {
	Header      string
	SecretField string
	Prefix      string
}

// Marshal preserves presentation strings without HTML expansion on the ABI.
func Marshal(value any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(buffer.Bytes(), []byte("\n")), nil
}

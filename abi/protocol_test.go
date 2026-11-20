package abi

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/izenadns/izena-connector-contract/connector"
)

func TestIndividualRecordWireFormat(t *testing.T) {
	input := connector.IndividualRecordInput{
		Name: "www.example.test.", Type: "A", TTL: 300, Content: "192.0.2.1",
	}
	payload, err := Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	request := Request{Operation: UpdateRecord, ZoneID: "zone/42", RecordID: "record/007?active", Input: payload}
	wire, err := Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	const wantRequest = `{"Operation":"update_record","ZoneID":"zone/42","RecordID":"record/007?active","Input":{"Name":"www.example.test.","Type":"A","TTL":300,"Content":"192.0.2.1","Disabled":false,"Comments":null}}`
	if string(wire) != wantRequest {
		t.Fatalf("request = %s, want %s", wire, wantRequest)
	}

	record := connector.IndividualRecord{ID: request.RecordID, IndividualRecordInput: input}
	value, err := Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	const wantRecord = `{"ID":"record/007?active","Name":"www.example.test.","Type":"A","TTL":300,"Content":"192.0.2.1","Disabled":false,"Comments":null}`
	if string(value) != wantRecord {
		t.Fatalf("record = %s, want %s", value, wantRecord)
	}
	var decoded connector.IndividualRecord
	if err := json.Unmarshal(value, &decoded); err != nil || !reflect.DeepEqual(decoded, record) {
		t.Fatalf("decoded record = %+v, %v", decoded, err)
	}
}

func TestRecordModelCapabilitiesRoundTrip(t *testing.T) {
	for _, model := range []connector.RecordModel{connector.RecordModelRRsets, connector.RecordModelIndividual} {
		t.Run(string(model), func(t *testing.T) {
			capabilities := connector.ServerCapabilities{Provider: "example-dns", RecordModel: model}
			raw, err := Marshal(capabilities)
			if err != nil {
				t.Fatal(err)
			}
			var decoded connector.ServerCapabilities
			if err := json.Unmarshal(raw, &decoded); err != nil || decoded.RecordModel != model {
				t.Fatalf("decoded record model = %q, %v, want %q", decoded.RecordModel, err, model)
			}
		})
	}
}

func TestZoneSettingsPatchWireFormat(t *testing.T) {
	policy, disabled := "DEFAULT", false
	input := connector.ZoneSettingsPatch{
		Set: map[string]connector.SettingValue{
			"api_rectify":  {Boolean: &disabled},
			"soa_edit_api": {String: &policy},
		},
		Reset: []string{"soa_edit"},
	}
	wire, err := Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"Set":{"api_rectify":{"Boolean":false},"soa_edit_api":{"String":"DEFAULT"}},"Reset":["soa_edit"]}`
	if string(wire) != want {
		t.Fatalf("patch = %s, want %s", wire, want)
	}
}

func TestPaginatedRRsetsPreserveIdentityAndOptions(t *testing.T) {
	enabled := true
	page := connector.Page[connector.RRset]{
		Items: []connector.RRset{
			{ID: "primary:blue", Name: "www.example.test.", Type: "A", AutomaticTTL: true,
				Options: map[string]connector.SettingValue{"proxied": {Boolean: &enabled}}},
			{ID: "primary:green", Name: "www.example.test.", Type: "A", TTL: 300},
		},
		NextCursor: "opaque/+/next?",
	}
	value, err := Marshal(page)
	if err != nil {
		t.Fatal(err)
	}
	var decoded connector.Page[connector.RRset]
	if err := json.Unmarshal(value, &decoded); err != nil || !reflect.DeepEqual(decoded, page) {
		t.Fatalf("page round trip = %+v, %v", decoded, err)
	}
}

func TestRecordBatchWireFormat(t *testing.T) {
	input := connector.RecordBatchInput{Changes: []connector.RecordChange{
		{Action: connector.RecordChangeDelete, RRset: &connector.RRsetChange{ID: "opaque/rrset"}},
		{Action: connector.RecordChangeCreate, RRset: &connector.RRsetChange{
			Input: &connector.RRsetReplaceInput{
				Name: "www.example.test.", Type: "A", TTL: 300,
				Records: []connector.Record{{Content: "192.0.2.1"}},
			},
		}},
	}}
	raw, err := Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	var decoded connector.RecordBatchInput
	if err := json.Unmarshal(raw, &decoded); err != nil || !reflect.DeepEqual(decoded, input) {
		t.Fatalf("batch round trip = %+v, %v", decoded, err)
	}
	request, err := Marshal(Request{Operation: GetRecordBatchStatus, ZoneID: "zone", BatchID: "opaque/batch+1"})
	if err != nil {
		t.Fatal(err)
	}
	const want = `{"Operation":"get_record_batch_status","ZoneID":"zone","BatchID":"opaque/batch+1"}`
	if string(request) != want {
		t.Fatalf("status request = %s, want %s", request, want)
	}
}

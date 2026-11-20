package connector

import (
	"context"
	"time"
)

// RRsetIDOperations addresses provider-specific RRsets with opaque zone-scoped
// identifiers, including RRsets that share an owner name and type.
type RRsetIDOperations interface {
	GetRRsetByID(ctx context.Context, zoneID, rrsetID string) (RRset, error)
	ReplaceRRsetByID(ctx context.Context, zoneID, rrsetID string, input RRsetReplaceInput) error
	DeleteRRsetByID(ctx context.Context, zoneID, rrsetID string) error
}

type RRset struct {
	// ID is opaque and scoped to this zone. It is empty for name-and-type RRsets.
	ID       string `json:",omitempty"`
	Name     string
	Type     string
	TTL      int64
	Records  []Record
	Comments []Comment
	// AutomaticTTL reports a provider-selected TTL. TTL is the effective value
	// when the provider supplies it.
	AutomaticTTL bool                    `json:",omitempty"`
	Options      map[string]SettingValue `json:",omitempty"`
}

type RRsetReplaceInput struct {
	Name     string
	Type     string
	TTL      int64
	Records  []Record
	Comments []Comment
	// AutomaticTTL requests a provider-selected TTL. Leave TTL at zero in this mode.
	AutomaticTTL bool                    `json:",omitempty"`
	Options      map[string]SettingValue `json:",omitempty"`
}

type Record struct {
	Content  string
	Disabled bool
}

type Comment struct {
	Content    string
	Account    string
	ModifiedAt time.Time
}

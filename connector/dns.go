package connector

import "context"

// ZoneReader reads live DNS state from the provider. All zone IDs are opaque.
type ZoneReader interface {
	ListZones(context.Context) ([]Zone, error)
	GetZone(context.Context, string) (Zone, error)
	GetRRset(context.Context, string, string, string) (RRset, error)
}

// ZoneMutator supports the DNS workflows Izena currently exposes. A connector
// returns an Error with ErrUnsupported for an operation it cannot perform.
type ZoneMutator interface {
	ZoneReader
	DNSSECOperations
	ZoneActions

	GetZoneMetadata(context.Context, string) (ZoneMetadata, error)
	CreateZone(context.Context, ZoneCreateInput) (Zone, error)
	// UpdateZone reports whether an upstream mutation was attempted, even on failure.
	UpdateZone(context.Context, string, ZoneUpdateInput) (ZoneUpdateResult, error)
	DeleteZone(context.Context, string) error

	ReplaceRRset(context.Context, string, RRsetReplaceInput) error
	DeleteRRset(context.Context, string, string, string) error
}

type ZoneActions interface {
	RectifyZone(context.Context, string) error
	NotifyZone(context.Context, string) error
	AxfrZone(context.Context, string) error
}

type ZoneMetadata struct {
	ID      string
	Name    string
	Kind    string
	Serial  *int64
	Actions []ZoneAction
}

type Zone struct {
	ID     string
	Name   string
	Kind   string
	Serial *int64
	RRsets []RRset
	// DNSSEC reports signing configuration when available.
	DNSSEC       *bool
	Capabilities *ZoneCapabilities
}

type ZoneCreateInput struct {
	Name     string
	Kind     string
	Settings map[string][]string
}

type ZoneUpdateInput struct {
	Kind string
}

type ZoneUpdateResult struct {
	Attempted bool
	Previous  *ZoneMetadata
}

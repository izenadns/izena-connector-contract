// Package connector defines the in-process contract between Izena and an
// authoritative DNS provider. Izena owns access control, persistence, and audit.
// Providers translate live upstream operations and report classified failures.
package connector

import "context"

// Provider describes a connector and creates clients for stored connections.
// Normalization and validation are local operations and must perform no network
// I/O. A provider must support concurrent calls for different connections.
type Provider interface {
	Describe() Description
	PrepareSettings(context.Context, map[string]string) (PreparedSettings, error)
	Connect(ConnectionConfig) (Client, error)
}

// PreparedSettings contains canonical settings and the stable identity of the
// upstream inventory. The identity excludes credentials and mutable settings.
type PreparedSettings struct {
	Settings  map[string]string
	TargetKey string
}

// Description contains public metadata supplied by the provider.
type Description struct {
	ID               string
	Name             string
	ConnectionFields []ConnectionField
}

// ConnectionField describes a string setting rendered by Izena. Secret fields
// are write-only and their values are supplied separately to Connect.
type ConnectionField struct {
	Key              string
	Label            string
	Help             string
	Default          string
	Placeholder      string
	MaxLength        int
	URL              bool
	Required         bool
	Secret           bool
	WriteOnly        bool // implied for secret fields
	Multiline        bool
	Editable         bool
	ClearLabel       string
	ClearDescription string
}

// ConnectionConfig contains persisted settings and the secrets for one
// authorized connection. Connect must treat both maps as read-only. It must
// construct a client without making network requests or logging configuration.
type ConnectionConfig struct {
	// ConnectionID is an opaque Izena identifier, never sent upstream.
	ConnectionID string
	Settings     map[string]string
	Secrets      map[string]string
}

// Client is bound to one upstream account or server. Clients that advertise
// RecordModelIndividual also implement IndividualRecordOperations and return
// ErrUnsupported for any RRset operations they do not support.
type Client interface {
	ZoneMutator
	Capabilities(context.Context) (ServerCapabilities, error)
	Probe(context.Context) (Profile, error)
}

// Profile identifies the service discovered at a connection.
type Profile struct {
	DaemonType string
	Version    string
}

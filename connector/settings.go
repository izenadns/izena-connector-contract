package connector

import "context"

type SettingKind string

const (
	SettingString     SettingKind = "string"
	SettingBoolean    SettingKind = "boolean"
	SettingInteger    SettingKind = "integer"
	SettingStringList SettingKind = "string_list"
	SettingEnum       SettingKind = "enum"
)

type SettingOption struct {
	Value string
	Label string
}

// SettingField describes a provider-owned setting. Option values are opaque
// identifiers, validated against the choices advertised by the connector.
type SettingField struct {
	Key         string
	Label       string
	Description string
	Kind        SettingKind
	ReadOnly    bool
	Resettable  bool
	Options     []SettingOption `json:",omitempty"`
}

// SettingValue carries exactly one value. Enums use String. Pointers preserve
// the difference between an absent value and zero, false, or an empty string.
type SettingValue struct {
	String     *string   `json:",omitempty"`
	Boolean    *bool     `json:",omitempty"`
	Integer    *int64    `json:",omitempty"`
	StringList *[]string `json:",omitempty"`
}

// ZoneSettings reports live values. A missing key means the setting is unset.
type ZoneSettings struct {
	Values map[string]SettingValue
}

// ZoneSettingsPatch sets or resets named fields. Unlisted fields are preserved.
// A key may appear in Set or Reset, never both.
type ZoneSettingsPatch struct {
	Set   map[string]SettingValue `json:",omitempty"`
	Reset []string                `json:",omitempty"`
}

type ZoneSettingsOperations interface {
	GetZoneSettings(ctx context.Context, zoneID string) (ZoneSettings, error)
	PatchZoneSettings(ctx context.Context, zoneID string, patch ZoneSettingsPatch) error
}

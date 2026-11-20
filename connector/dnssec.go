package connector

import "context"

type DNSSECReader interface {
	GetDNSSEC(context.Context, string) (DNSSEC, error)
}

type DNSSECOperations interface {
	DNSSECReader
	SetDNSSEC(context.Context, string, bool) error
}

// DNSSEC contains public status and key metadata. Unknown active status is nil.
type DNSSEC struct {
	Enabled bool
	Active  *bool
	Keys    []DNSSECKey
}

type DNSSECKey struct {
	ID        int64
	KeyType   string
	Active    bool
	Published bool
	DNSKEY    string
	KeyTag    int32
	Algorithm *string
	Bits      *int32
	DS        []string
}

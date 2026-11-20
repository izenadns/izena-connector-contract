package connector

import "context"

// PageInput uses a positive requested limit and an opaque continuation cursor.
// The first request omits the cursor. Pass NextCursor back unchanged.
type PageInput struct {
	Limit  int
	Cursor string `json:",omitempty"`
}

// Page has no continuation when NextCursor is empty.
type Page[T any] struct {
	Items      []T
	NextCursor string `json:",omitempty"`
}

type PaginationCapabilities struct {
	Zones   bool
	RRsets  bool
	Records bool
}

type PaginatedReadOperations interface {
	ListZonesPage(ctx context.Context, input PageInput) (Page[ZoneMetadata], error)
	ListRRsetsPage(ctx context.Context, zoneID string, input PageInput) (Page[RRset], error)
	ListRecordsPage(ctx context.Context, zoneID string, input PageInput) (Page[IndividualRecord], error)
}

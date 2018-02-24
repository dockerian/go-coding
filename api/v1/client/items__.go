// Package client :: items__.go - Query result items
package client

import "encoding/json"

// QueryResults struct defines a query result with count, and a collection of items
type QueryResults struct {

	// Count defines number of items in items
	Count int `json:"count,omitempty"`

	// Count defines total number of items in all pages, if page and count queries is provided
	CountAll int `json:"countAll,omitempty"`

	// PageOffset defines zero-based page offset
	PageOffset int `json:"pageOffset,omitempty"`

	// PageSize defines page size of current selection
	PageSize int `json:"pageSize,omitempty"`

	// NextURL specifies the URL for fetching next page, if available
	NextURL string `json:"nextUrl,omitempty"`

	// Items represents a collection of query results
	Items json.RawMessage `json:"items,omitempty"`
}

package index

import "github.com/google/uuid"

type Indexer interface {
	Index(doc *Document) error

	FindByID(linkID uuid.UUID) (*Document, error)

	// You almost never need a pointer to an interface, since interfaces are just pointers themselves (basically).
	// https://www.reddit.com/r/golang/comments/2s2jv8/interfaces_and_pointers_to_interfaces/
	Search(query Query) (Iterator, error)

	UpdateScore(linkID uuid.UUID, score float64) error
}

// Iterator is implemented by objects that can paginate search results.
type Iterator interface {
	// Close the iterator and release any allocated resources.
	Close() error

	// Next loads the next document matching the search query.
	// It returns false if no more documents are available.
	Next() bool

	// Error returns the last error encountered by the iterator.
	Error() error

	// Document returns the current document from the result set.
	Document() *Document

	// TotalCount returns the approximate number of search results.
	TotalCount() uint64
}

// QueryType describes the types of queries supported by the indexer
// implementations.
type QueryType uint8

// Declaring multiple const together with same value and type
// const (
//   a string = "circle"
//   b
// )
const (
	// QueryTypeMatch requests the indexer to match each expression term.
	// https://github.com/golang/go/wiki/Iota
	// Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers.
	QueryTypeMatch QueryType = iota

	// QueryTypePhrase searches for an exact phrase match.
	QueryTypePhrase
)

// Query encapsulates a set of parameters to use when searching indexed
// documents.
type Query struct {
	// The way that the indexer should interpret the search expression.
	Type QueryType

	// The search expression.
	Expression string

	// The number of search results to skip.
	Offset uint64
}

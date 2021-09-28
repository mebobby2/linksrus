package graph

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID         uuid.UUID
	URL        string
	RetrieveAt time.Time
}

type Edge struct {
	ID        uuid.UUID
	Src       uuid.UUID
	Dst       uuid.UUID
	UpdatedAt time.Time
}

type Iterator interface {
	Next() bool
	Error() error
	Close() error
}

type LinkIterator interface {
	Iterator

	Link() *Link
}

type EdgeIterator interface {
	Iterator

	Edge() *Edge
}

type Graph interface {
	UpsertLink(link *Link) error
	FindLink(id uuid.UUID) (*Link, error)

	UpsertEdge(edge *Edge) error
	RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error

	Links(fromID, toID uuid.UUID, retrieveBefore time.Time) (LinkIterator, error)
	Edges(fromID, toID uuid.UUID, retrieveBefore time.Time) (EdgeIterator, error)
}

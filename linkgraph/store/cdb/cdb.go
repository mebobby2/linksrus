package cdb

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mebobby2/linksrus/linkgraph/graph"
	"golang.org/x/xerrors"
)

var (
	upsertLinkQuery = `
INSERT INTO links (url, retrieved_at) VALUES ($1, $2)
ON CONFLICT (url) DO UPDATE SET retrieved_at=GREATEST(links.retrieved_at, $2)
RETURNING id, retrieved_at
	`

	findLinkQuery = "SELECT url, retrieved_at FROM links WHERE id=$1"

	upsertEdgeQuery = `
INSERT INTO edges (src, dst, updated_at) VALUES ($1, $2, NOW())
ON CONFLICT (src,dst) DO UPDATE SET updated_at=NOW()
RETURNING id, updated_at
	`

	// Compile-time check for ensuring CockroachDbGraph implements Graph.
	_ graph.Graph = (*CockroachDBGraph)(nil)
)

type CockroachDBGraph struct {
	db *sql.DB
}

func NewCockroachDbGraph(dsn string) (*CockroachDBGraph, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &CockroachDBGraph{db: db}, nil
}

func (c *CockroachDBGraph) UpsertLink(link *graph.Link) error {
	row := c.db.QueryRow(upsertLinkQuery, link.URL, link.RetrieveAt.UTC())
	if err := row.Scan(&link.ID, &link.RetrieveAt); err != nil {
		return xerrors.Errorf("upsert link: %w", err)
	}

	link.RetrieveAt = link.RetrieveAt.UTC()
	return nil
}

func (c *CockroachDBGraph) UpsertEdge(edge *graph.Edge) error {
	row := c.db.QueryRow(upsertEdgeQuery, edge.Src, edge.Dst)
	if err := row.Scan(&edge.ID, &edge.UpdatedAt); err != nil {
		if isForeignViolationError(err) {
			err = graph.ErrUnknownEdgeLinks
		}
		return xerrors.Errorf("upsert edge: %w", err)
	}

	edge.UpdatedAt = edge.UpdatedAt.UTC()
	return nil
}

func (c *CockroachDBGraph) FindLink(id uuid.UUID) (*graph.Link, error) {
	row := c.db.QueryRow(findLinkQuery, id)
	link := &graph.Link{ID: id}
	if err := row.Scan(&link.URL, &link.RetrieveAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, xerrors.Errorf("find link: %w", graph.ErrNotFound)
		}

		return nil, xerrors.Errorf("find link: %w", err)
	}

	link.RetrieveAt = link.RetrieveAt.UTC()
	return link, nil
}

func isForeignViolationError(err error) bool {
	pqErr, valid := err.(*pq.Error)
	// https://go.dev/tour/methods/15
	// A type assertion provides access to an interface value's underlying concrete value.
	// This statement asserts that the interface value err holds the concrete type *pq.Error and assigns the underlying *pq.Error value to the variable pqError.
	// If i does not hold a T, the statement will trigger a panic.

	if !valid {
		return false
	}

	return pqErr.Code.Name() == "foreign_key_violation"
}

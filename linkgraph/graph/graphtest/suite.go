package graphtest

import (
	"time"

	"github.com/google/uuid"
	"github.com/mebobby2/linksrus/linkgraph/graph"
	gc "gopkg.in/check.v1"
)

type SuiteBase struct {
	g graph.Graph
}

func (s *SuiteBase) SetGraph(g graph.Graph) {
	s.g = g
}

func (s *SuiteBase) TestUpsertLink(c *gc.C) {
	original := &graph.Link{
		URL:        "https://example.com",
		RetrieveAt: time.Now().Add(-10 * time.Hour),
	}

	err := s.g.UpsertLink(original)
	c.Assert(err, gc.IsNil)
	c.Assert(original.ID, gc.Not(gc.Equals), uuid.Nil, gc.Commentf("expected a linkID to be assigned to the new link"))
}
package memory

import (
	"testing"

	"github.com/mebobby2/linksrus/linkgraph/graph/graphtest"
	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(new(InMemoryGraphTestSuite))

type InMemoryGraphTestSuite struct {
	graphtest.SuiteBase // struct embedding https://medium.com/@simplyianm/why-gos-structs-are-superior-to-class-based-inheritance-b661ba897c67
}

func (s *InMemoryGraphTestSuite) SetUpTest(c *gc.C) {
	s.SetGraph(NewInMemoryGraph())
}

// Register our test-suite with go test
func Test(t *testing.T) { gc.TestingT(t) }

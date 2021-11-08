package memory

import "github.com/mebobby2/linksrus/linkgraph/graph"

type linkIterator struct {
	s *InMemoryGraph

	links    []*graph.Link
	curIndex int
}

func (i *linkIterator) Next() bool {
	if i.curIndex >= len(i.links) {
		return false
	}
	i.curIndex++
	return true
}

func (i *linkIterator) Link() *graph.Link {
	i.s.mu.RLock()
	link := new(graph.Link)
	*link = *i.links[i.curIndex-1]
	i.s.mu.RUnlock()
	return link
}

func (i *linkIterator) Error() error {
	return nil
}

func (i *linkIterator) Close() error {
	return nil
}

type edgeIterator struct {
	s *InMemoryGraph

	edges    []*graph.Edge
	curIndex int
}

func (i *edgeIterator) Next() bool {
	if i.curIndex >= len(i.edges) {
		return false
	}
	i.curIndex++
	return true
}

func (i *edgeIterator) Edge() *graph.Edge {
	i.s.mu.RLock()
	edge := new(graph.Edge)
	*edge = *i.edges[i.curIndex-1]
	i.s.mu.RUnlock()
	return edge
}

func (i *edgeIterator) Error() error {
	return nil
}

func (i *edgeIterator) Close() error {
	return nil
}

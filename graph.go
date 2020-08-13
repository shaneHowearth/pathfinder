// Package pathfinder -
package pathfinder

import "fmt"

// Node -
type Node struct {
	Parent   *Node
	Children []*Node
	Seen     bool
	Data     interface{}
}

// Graph -
type Graph struct {
	Agent   Searcher
	Goal    *Node
	Success func(goal, current *Node) (bool, error)
}

// AddChild -
func (n *Node) AddChild(child *Node) error {
	if child == nil {
		return fmt.Errorf("no child node supplied")
	}
	// Set the child's parent
	child.Parent = n

	n.Children = append(n.Children, child)
	return nil
}

// Searcher -
type Searcher interface {
	Search(start *Node, graph *Graph) (path []*Node, err error)
}

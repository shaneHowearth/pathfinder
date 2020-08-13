// Package dfs - Depth First Search
// Starting at the root node this algorithm searches the right most edge leading
// out of the node, that hasn't yet been searched, until it reachs the bottom of
// the search tree, or the goal. It then goes back up one node and searches the
// right most edge, and so on until the tree has been completely searched, or
// the goal has been found.
package dfs

import (
	"github.com/shanehowearth/pathfinder"
)

// DFS -
type DFS struct{}

// Search - search for the goal node and return the optimal path to that node
// from the root of the graph.
func (d *DFS) Search(start *pathfinder.Node, graph *pathfinder.Graph) ([]*pathfinder.Node, error) {

	// queue of Nodes to inspect
	var todo []*pathfinder.Node

	todo = append(todo, start)

	for {
	SearchLoop:
		// Take the first node from the queue for inspection
		var current *pathfinder.Node
		current, todo = todo[0], todo[1:]
		if current == nil {
			// no node to return, no error
			return nil, nil
		}

		// Check if this node has been processed already
		if current.Seen {
			continue
		}
		current.Seen = true

		// Check if this node is the goal
		win, err := graph.Success(graph.Goal, current)
		if err != nil {
			return nil, err
		}
		if win {
			return d.getPath(current, start)
		}

		// Add the next unseen RHS child to the todo queue.
		n := current
		for {
			// check this node's children
			for idx := range n.Children {
				if !n.Children[idx].Seen {
					todo = append(todo, n.Children[idx])
					goto SearchLoop
				}
			}

			if n.Parent == nil {
				// Search completed, nothing found, no error
				return nil, nil
			}

			// check the parent's next children
			n = n.Parent
		}
	}
}

// Getpath returns the path from the root node to the start node
// It assumes that the goal is always below the start, that is, it does not
// provide the path to the parent(s) of the start
func (d *DFS) getPath(goal, start *pathfinder.Node) ([]*pathfinder.Node, error) {
	var path []*pathfinder.Node

	node := goal
	for {
		node.Seen = !node.Seen // invert as we return to the start, so we can detect a loop
		if node.Seen {
			return path, nil
		}
		path = append(path, node)
		if node.Parent == nil {
			// reverse the path - root is first step
			for i, j := 0, len(path)-1; i <= j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path, nil
		}
		if node.Parent == start {

		}
		node = node.Parent
	}
}

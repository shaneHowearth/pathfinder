// Package bfs - Breadth First Search
// Starting at the root node this algorithm searches each edge that leads to a
// child node, in search of the goal node. Once all the edges leading out of the
// node have been searched, the algorithm moves to search each of the next level
// down (the children nodes found in the previous search) in the same way,
// continuing until the goal is found, or there are no more nodes to search.
package bfs

import (
	"github.com/shanehowearth/pathfinder"
)

// BFS -
type BFS struct{}

// Search - search for the goal node and return the optimal path to that node
// from the root of the graph.
func (b *BFS) Search(start *pathfinder.Node, graph *pathfinder.Graph) ([]*pathfinder.Node, error) {
	// queue of Nodes to inspect
	var todo []*pathfinder.Node

	todo = append(todo, start)

	for {
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
			return b.getPath(current)
		}

		// Add the node's children to the todo queue.
		todo = append(todo, current.Children...)
	}
}

// Getpath returns the path from the root node to the start node
// It assumes that the goal is always below the start, that is, it does not
// provide the path to the parent(s) of the start
func (b *BFS) getPath(node *pathfinder.Node) ([]*pathfinder.Node, error) {
	var path []*pathfinder.Node

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
		node = node.Parent
	}
}

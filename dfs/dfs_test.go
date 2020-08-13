package dfs_test

import (
	"testing"

	"github.com/shanehowearth/pathfinder"
	"github.com/shanehowearth/pathfinder/dfs"
	"github.com/stretchr/testify/assert"
)

var goalGood *pathfinder.Node
var startGood *pathfinder.Node
var node01, node02, node03, node04, node05 *pathfinder.Node
var node06, node07, node08, node09, node10 *pathfinder.Node
var node11, node12, node13, node14, node15 *pathfinder.Node
var node16, node17, node18, node19, node20 *pathfinder.Node

func buildGraphOne() {
	startGood.AddChild(node01)
	startGood.AddChild(node02)
	node01.AddChild(node03)
	node01.AddChild(node04)
	node01.AddChild(node05)
	node02.AddChild(node06)
	node02.AddChild(node07)
	node03.AddChild(node08)
	node03.AddChild(node09)
	node04.AddChild(node10)
	node05.AddChild(node11)
	node05.AddChild(node12)
	node06.AddChild(node13)
	node06.AddChild(goalGood)
	node07.AddChild(node14)
	node07.AddChild(node15)
	node08.AddChild(node16)
	node08.AddChild(node17)
	node10.AddChild(node18)
	node11.AddChild(node19)
	node13.AddChild(node20)
}
func Test(t *testing.T) {
	d := dfs.DFS{}
	startGood = &pathfinder.Node{Data: "start"}
	goalGood = &pathfinder.Node{Data: "Good"}
	node01 = &pathfinder.Node{}
	node02 = &pathfinder.Node{}
	node03 = &pathfinder.Node{}
	node04 = &pathfinder.Node{}
	node05 = &pathfinder.Node{}
	node06 = &pathfinder.Node{}
	node07 = &pathfinder.Node{}
	node08 = &pathfinder.Node{}
	node09 = &pathfinder.Node{}
	node10 = &pathfinder.Node{}
	node11 = &pathfinder.Node{}
	node12 = &pathfinder.Node{}
	node13 = &pathfinder.Node{}
	node14 = &pathfinder.Node{}
	node15 = &pathfinder.Node{}
	node16 = &pathfinder.Node{}
	node17 = &pathfinder.Node{}
	node18 = &pathfinder.Node{}
	node19 = &pathfinder.Node{}
	node20 = &pathfinder.Node{}

	testcases := map[string]struct {
		start         *pathfinder.Node
		graph         *pathfinder.Graph
		pathway       []*pathfinder.Node
		expectedError error
		graphBuild    func()
	}{
		"Happy Path": {
			start: startGood,
			graph: &pathfinder.Graph{Goal: goalGood,
				Success: func(goal, current *pathfinder.Node) (bool, error) {
					if current.Data != nil {
						return goal.Data.(string) == current.Data.(string), nil
					}
					return false, nil
				},
			},
			pathway:    []*pathfinder.Node{startGood, node02, node06, goalGood},
			graphBuild: buildGraphOne,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			tc.graphBuild()
			// call the search function
			path, err := d.Search(tc.start, tc.graph)
			// check that the search function behaved correctly.
			if tc.expectedError == nil {
				assert.Nil(t, err)
				assert.Equal(t, len(tc.pathway), len(path))
				for idx := range tc.pathway {
					assert.Equal(t, tc.pathway[idx], path[idx])
				}
			}

		})
	}
}

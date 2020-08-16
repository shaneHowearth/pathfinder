package pathfinder_test

import (
	"fmt"
	"testing"

	"github.com/shanehowearth/pathfinder"
	"github.com/stretchr/testify/assert"
)

func TestAddChild(t *testing.T) {
	testcases := map[string]struct {
		parent      *pathfinder.Node
		child       *pathfinder.Node
		expectedErr error
	}{
		"Happy Path": {
			parent: &pathfinder.Node{},
			child:  &pathfinder.Node{},
		},
		"Nil Child": {
			parent:      &pathfinder.Node{},
			expectedErr: fmt.Errorf("no child node supplied"),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			err := tc.parent.AddChild(tc.child)

			if tc.expectedErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, tc.child.Parent, tc.parent)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tc.expectedErr)
			}
		})
	}
}

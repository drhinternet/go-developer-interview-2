package linkedlist_test

import (
	"testing"

	"linkedlist"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		inputKeys        []int64
		expectLinkedList string
	}{
		{
			inputKeys:        nil,
			expectLinkedList: "",
		},
		{
			inputKeys:        []int64{50},
			expectLinkedList: "50",
		},
		{
			inputKeys:        []int64{11, 5, 2, 12, 14, 19, 4, 17, 16, 3, 9, 7, 15, 6, 13, 1, 8, 18, 20},
			expectLinkedList: "1:2:3:4:5:6:7:8:9:11:12:13:14:15:16:17:18:19:20",
		},
	}

	for _, test := range tests {
		var tree *linkedlist.Node

		if len(test.inputKeys) >= 1 {
			tree = &linkedlist.Node{Key: test.inputKeys[0]}
			for _, key := range test.inputKeys[1:] {
				tree.AddToTree(key)
			}
		}

		list := tree.TransformTreeToLinkedList()

		output, err := list.LinkedListString()

		assert.Nil(t, err)
		assert.Equal(t, test.expectLinkedList, output)
	}
}

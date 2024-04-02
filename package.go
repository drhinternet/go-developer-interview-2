package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Key   int64
	Left  *Node
	Right *Node
}

func (in *Node) TransformTreeToLinkedList() *Node {
	// This function will return a sorted linked list where the return value is
	// the first element of the linked list.

	return nil
}

func (in *Node) AddToTree(key int64) {
	var direction **Node

	if key < in.Key {
		direction = &in.Left
	} else {
		direction = &in.Right
	}

	if *direction != nil {
		(*direction).AddToTree(key)
	} else {
		*direction = &Node{Key: key}
	}
}

func (start *Node) LinkedListString() (string, error) {
	if start == nil {
		return "", nil
	}

	if start.Left != nil {
		return "", errors.New("list error: first node left is not nil")
	}

	var buf strings.Builder

	node := start

	for node != nil {
		if node != start {
			fmt.Fprintf(&buf, ":")
		}

		fmt.Fprintf(&buf, "%d", node.Key)

		if node.Right != nil {
			if node.Right.Left != node {
				return "", errors.New("list error: reverse linking is not correct")
			}
		}

		node = node.Right
	}

	return buf.String(), nil
}

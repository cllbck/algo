package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryTreeInsert(t *testing.T) {
	tree := &BinaryTree{
		value: 2,
		left: &BinaryTree{
			value: 1,
		},
		right: &BinaryTree{
			value: 3,
		},
	}
	tree.Insert(4)

	expectedTree := &BinaryTree{
		value: 2,
		left: &BinaryTree{
			value: 1,
		},
		right: &BinaryTree{
			value: 3,
			right: &BinaryTree{
				value: 4,
			},
		},
	}
	assert.Equal(t, expectedTree, tree)
}

func TestBinaryTreeSearch(t *testing.T) {
	tree := &BinaryTree{
		value: 2,
		left: &BinaryTree{
			value: 1,
		},
		right: &BinaryTree{
			value: 3,
		},
	}
	assert.True(t, tree.Search(2))
	assert.False(t, tree.Search(4))
}

func TestBinaryTreeRemove(t *testing.T) {
	tree := &BinaryTree{
		value: 5,
		left: &BinaryTree{
			value: 3,
			left: &BinaryTree{
				value: 1,
			},
			right: &BinaryTree{
				value: 4,
			},
		},
		right: &BinaryTree{
			value: 7,
			left: &BinaryTree{
				value: 6,
			},
			right: &BinaryTree{
				value: 8,
			},
		},
	}
	tree.Remove(5)
	assert.False(t, tree.Search(5))

	assert.True(t, tree.Search(3))
	tree.Remove(3)
	assert.False(t, tree.Search(3))
}

func TestBinaryTreeMax(t *testing.T) {
	tree := &BinaryTree{
		value: 5,
		left: &BinaryTree{
			value: 3,
			left: &BinaryTree{
				value: 1,
			},
			right: &BinaryTree{
				value: 4,
			},
		},
		right: &BinaryTree{
			value: 7,
			left: &BinaryTree{
				value: 6,
			},
			right: &BinaryTree{
				value: 8,
			},
		},
	}
	assert.Equal(t, 8, tree.Max())
}

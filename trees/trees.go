package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func (t *BinaryTree) Insert(value int) {
	if value < t.value {
		if t.left == nil {
			t.left = &BinaryTree{value, nil, nil}
		} else {
			t.left.Insert(value)
		}
	} else {
		if t.right == nil {
			t.right = &BinaryTree{value, nil, nil}
		} else {
			t.right.Insert(value)
		}
	}
}

func (t *BinaryTree) Search(value int) bool {
	if t.value == value {
		return true
	}
	if value < t.value {
		if t.left == nil {
			return false
		} else {
			return t.left.Search(value)
		}
	} else {
		if t.right == nil {
			return false
		} else {
			return t.right.Search(value)
		}
	}
}

func (t *BinaryTree) Remove(value int) {
	if t.value == value {
		if t.left == nil && t.right == nil {
			t = nil
		} else if t.left == nil {
			t = t.right
		} else if t.right == nil {
			t = t.left
		} else {
			max := t.left.Max()
			t.value = max
			t.left.Remove(max)
		}
	} else if value < t.value {
		if t.left == nil {
			return
		} else {
			t.left.Remove(value)
		}
	} else {
		if t.right == nil {
			return
		} else {
			t.right.Remove(value)
		}
	}
}

func (t *BinaryTree) Max() int {
	if t.right == nil {
		return t.value
	} else {
		return t.right.Max()
	}
}

func main() {
	bt1 := &BinaryTree{500, nil, nil}
	bt2 := &BinaryTree{500, nil, nil}
	N := 100000
	now := time.Now()
	for i := 0; i < N; i++ {
		bt1.Insert(i)
	}
	fmt.Println("Insert time bt1 ordered number: ", time.Since(now))
	now = time.Now()
	for i := 0; i < N; i++ {
		bt2.Insert(rand.Intn(N))
	}
	fmt.Println("Insert time bt2 random number: ", time.Since(now))

	now = time.Now()
	for i := 0; i < N; i += 10 {
		bt1.Search(rand.Intn(N))
	}
	fmt.Println("Search time bt1 random number: ", time.Since(now))

	now = time.Now()
	for i := 0; i < N; i += 10 {
		bt2.Search(rand.Intn(N))
	}
	fmt.Println("Search time bt2 random number: ", time.Since(now))

	now = time.Now()
	for i := 0; i < N; i += 10 {
		bt1.Remove(rand.Intn(N))
	}
	fmt.Println("Remove time bt1 random number: ", time.Since(now))

	now = time.Now()
	for i := 0; i < N; i += 10 {
		bt2.Remove(rand.Intn(N))
	}
	fmt.Println("Remove time bt2 random number: ", time.Since(now))
}

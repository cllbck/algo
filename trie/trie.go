package main

import "fmt"

type Trie struct {
	root *Node
}

func NewTrie() Trie {
	return Trie{
		root: NewNode(),
	}
}

func (t *Trie) Insert(word, value string) {
	node := t.root
	for _, r := range word {
		node = node.Add(string(r))
	}
	node.value = &value
}

func (t *Trie) Search(word string) *string {
	node := t.root
	for _, r := range word {
		if !node.Exist(string(r)) {
			return nil
		}
		node = node.Add(string(r))
	}
	return node.value
}

func (t *Trie) Delete(word string) {
	node := t.root
	for _, r := range word {
		if !node.Exist(string(r)) {
			return
		}
		node = node.Add(string(r))
	}
	node.value = nil
}

type Node struct {
	childNodes map[string]*Node
	value      *string
}

func (n *Node) Exist(c string) bool {
	_, ok := n.childNodes[c]
	return ok
}

func (n *Node) Add(s string) *Node {
	if n.Exist(s) {
		node := n.childNodes[s]
		return node
	}
	newNode := NewNode()
	n.childNodes[s] = newNode
	return newNode
}

func NewNode() *Node {
	return &Node{
		childNodes: map[string]*Node{},
	}
}

func main() {
	trie := NewTrie()
	trie.Insert("собака", "1")
	trie.Insert("соб", "2")
	trie.Insert("pong", "3")

	for _, v := range []string{"собака", "соб", "со", "pong1", "pong"} {
		value := trie.Search(v)
		if value != nil {
			fmt.Printf("%s -> %s\n", v, *value)
		} else {
			fmt.Printf("key %s doesn`t exist \n", v)
		}
	}

	fmt.Println("++++++++++++++++++++++++++++++++")

	trie.Delete("соб")

	for _, v := range []string{"собака", "соб", "со", "pong1", "pong"} {
		value := trie.Search(v)
		if value != nil {
			fmt.Printf("%s -> %s\n", v, *value)
		} else {
			fmt.Printf("key %s doesn`t exist \n", v)
		}
	}
}

package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method of LinkedList class
func (LinkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = LinkedList.headNode
	}
	LinkedList.headNode = &node
}

// main method
func main() {

	instancia := LinkedList{}

	instancia.AddToHead(1)
	instancia.AddToHead(3)
	fmt.Println(instancia.headNode.property)
}

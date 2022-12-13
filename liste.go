package liste

import (
	"fmt"
)

type Node struct {
	Value string 
	Next *Node //This is a recursive datatype
}

//Given the pointer to the first element node of a list, returns the number of elements
func (first *Node) Length() (c int) {
	var curs *Node
	for curs = first; curs != nil; curs = curs.Next {
		c++
	}
	return
}

//Given the pointer to the first element node of a list, prints the elements of the list
func (first *Node) Print() {
	var curs *Node
	for curs = first; curs != nil; curs = curs.Next {
		if curs.Next == nil {
			fmt.Printf("%s\n", curs.Value)
		} else {
			fmt.Printf("%s -> ", curs.Value)
		}
	}
}

//Given the pointer to the first element node of a list and a new string
//Adds the latter as the FIRST element of the list.
//Returns the pointer to the new First element
func (first *Node) AddFront(x string) (newFirst *Node){
	newFirst = new(Node)
	newFirst.Value = x
	newFirst.Next = first
	return
}

//Given the pointer to the first element node of a list and a new string
//Adds the latter as the LAST element of the list.
//Returns the pointer to the new First element (It can be useful is first is nil)
func AddLast(first *Node, x string) (newFirst *Node) {
	var curs *Node
	newNode := new(Node)
	newNode.Value = x

	if first == nil {
		newFirst = newNode
		return
	}
	//List is non-empty
	newFirst = first
	for curs = first; curs.Next != nil; curs = curs.Next {}
	//Curs is now pointing to the last node of the list
	curs.Next = newNode
	return
}

//Given the pointer to the first element node of a list and a new string
//Adds the latter in the correct position to maintain the list in (lexicographic) order.
//Returns the pointer to the new First element
func AddInOrder(first *Node, x string) (newFirst *Node){
	var curs, prev *Node
	newNode := new(Node)
	newNode.Value = x

	if first == nil {
		newFirst = newNode
		return
	}
	//List is not empty
	prev = nil
	for curs = first; curs != nil && curs.Value < newNode.Value; curs = curs.Next {
		prev = curs
	}
	//newNode comes before curs
	newNode.Next = curs
	if prev == nil { //Inserting as first element
		newFirst = newNode //newNode comes after nothing 
		return
	}
	//Inserting the newNode between prev and curs
	prev.Next = newNode //newNode comes after prev
	newFirst = first
	return
}

//Given the pointers to the first element of two linked lists,
//appends all the elements of the second list to the tail of the first list.
//Returns the pointer to the new first element of the first list 
func Concat(first1, first2 *Node) (newFirst *Node) {
	var curs1, curs2 *Node
	//If the first list is nil, the new list is the second list
	if first1 == nil {
		newFirst = first2
		return
	}
	for curs1 = first1; curs1.Next != nil; curs1 = curs1.Next {} 
	//curs1 is now pointing to the last element of the first list
	for curs2 = first2; curs2 != nil; curs2 = curs2.Next {
		curs1.Next = curs2
		curs1 = curs2
	}
	newFirst = first1
	return
}
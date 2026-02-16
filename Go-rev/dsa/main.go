package main


type Node struct {
	Value int
	Next  *Node
}

func main() {

	n1 := &Node{Value: 10}
	n2 := &Node{Value: 20}
    n3 := &Node{Value: 30}

	n1.Next = n2

	curr := n1

	for curr.Next != nil {
		curr = curr.Next
	}

 curr.Next = n3

 

}


func InsertHead(head *Node,value int) *Node	{
 
 newNode := &Node{Value: 30}
 newNode.Next = head

 head = newNode

 return head

}

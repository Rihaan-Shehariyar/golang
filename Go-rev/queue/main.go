package main

// type Queue struct {
// 	items []int
// }

// func (q *Queue) Push(val int) {
// 	q.items = append(q.items, val)
// }

// func (q *Queue)Pop(){
//  if len(q.items) == 0{
//    return
// }

//  q.items = q.items[1:]

// }

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Front *Node
	Rear  *Node
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{Value: 10}

	if q.Rear == nil {
		q.Rear = newNode
		q.Front = newNode
		return
	}

	q.Rear.Next = newNode
	q.Rear = newNode

}

func (q *Queue) Dequeue() {
	if q.Front == nil {
		return
	}

	q.Front = q.Front.Next

	if q.Front == nil {
		q.Rear = nil
	}

}



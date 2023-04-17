package utils

type Queue struct {
	front  *Node
	rear   *Node
	length float64
}

func (queue *Queue) Len() float64 {
	return queue.length
}

func (queue *Queue) FirstNode() *Node {
	return queue.rear
}

func (queue *Queue) isEmpty() bool {
	return queue.length == 0
}

func (queue *Queue) Enqueue(v float64) {
	newNode := &Node{
		value: v,
		next:  nil,
	}
	if queue.isEmpty() {
		queue.front = newNode
		queue.rear = newNode
	} else {
		if queue.Len() == 1 {
			queue.rear.next = newNode
		}
		queue.front.next = newNode
		queue.front = newNode
	}
	queue.length += 1
}

func (queue *Queue) Dequeue() {
	if queue.isEmpty() {
		return
	}
	queue.front = queue.front.next
	queue.length -= 1
}

func (queue *Queue) Sum() float64 {
	node := queue.rear
	sum := node.value

	for {
		if !node.HasNext() {
			return sum
		}
		node = node.next
		sum += node.value
	}
}

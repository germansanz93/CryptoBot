package utils

type Queue struct {
	front  *Node
	rear   *Node
	length int
}

func (queue *Queue) Len() int {
	return queue.length
}

func (queue *Queue) isEmpty() bool {
	return queue.length == 0
}

func (queue *Queue) Enqueue(v uint64) {
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

func (queue *Queue) Sum() uint64 {
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

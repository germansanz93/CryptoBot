package main

import (
	q "germansanz93/cryptobot/utils"
	"log"
)

func main() {
	log.Println("Hello, World!")
	queue := q.Queue{}
	queue.Enqueue(6)
	queue.Enqueue(5)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	log.Println(queue.Sum())
}

package main

import (
	ma "germansanz93/cryptobot/averages"
	q "germansanz93/cryptobot/utils"
	"log"
)

func main() {
	log.Println("Hello, World!")
	queue := q.Queue{}
	queue.Enqueue(6.2)
	queue.Enqueue(5.4)
	queue.Enqueue(13)
	queue.Enqueue(8)
	log.Println(queue.Sum())
	log.Println(ma.SimpleMovingAverage(queue))
}

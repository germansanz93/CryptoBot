package averages

import q "germansanz93/cryptobot/utils"

const EMA_SMOOTHING_FACTOR float64 = 2 //Higher smoothing factor will increase weight of more recent data

func SimpleMovingAverage(ds q.Queue) float64 {
	sum := ds.Sum()
	n := ds.Len()

	return sum / n
}

func exponentialMovingAverage(ds q.Queue, mqty float64) (q.Queue, float64) { //TODO change approach to use the recursive form of the formula, with memo maybe?

	fn := ds.FirstNode()
	node := fn

	for {
		if !node.HasNext() {

		}
	}
}

func emaToday(mqty float64, value float64, emay float64) float64 {
	alfaFactor := EMA_SMOOTHING_FACTOR / (1 + float64(mqty))
	return value*(alfaFactor) + emay*(1-(alfaFactor))
}

func smoothedMovingAverage() {

}

func linearMovingAverage() {

}

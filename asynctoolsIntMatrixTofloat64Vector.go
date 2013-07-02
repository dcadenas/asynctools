package asynctools

import "runtime"

func workerIntMatrixToFloat64Vector(mappingFunc func([]int) float64, inSlice [][]int, outSlice []float64, doneChan chan struct{}) {
	for i, valForI := range inSlice {
		outSlice[i] = mappingFunc(valForI)
	}

	doneChan <- struct{}{}
}

//TODO: DRY and test
func MapIntMatrixToFloat64Vector(mappable [][]int, mappingFunc func([]int) float64) []float64 {
  size := len(mappable)
	resultSlice := make([]float64, size)

	if size == 0 {
		return resultSlice
	}

	cpus := runtime.NumCPU()
	chunkSize := size / cpus
	remainder := size % cpus

	doneChan := make(chan struct{})
	head, goRoutinesCount := 0, 0
	for tail := chunkSize + remainder; head < size; tail += chunkSize {
		inSlice := mappable[head:tail]
		goRoutinesCount++
		go workerIntMatrixToFloat64Vector(mappingFunc, inSlice, resultSlice[head:tail], doneChan)
		head = tail
	}

	for i := 0; i < goRoutinesCount; i++ {
		<-doneChan
	}

	return resultSlice
}

package asynctools

import "runtime"

type Mappable interface {
	At(int) interface{}
	Len() int
	Slice(int, int) Mappable
}

type mappingFuncType func(interface{}) interface{}

func worker(mappingFunc mappingFuncType, inSlice Mappable, outSlice []interface{}, doneChan chan struct{}) {
	for i := 0; i < inSlice.Len(); i++ {
		outSlice[i] = mappingFunc(inSlice.At(i))
	}

	doneChan <- struct{}{}
}

func Map(mappable Mappable, mappingFunc mappingFuncType) []interface{} {
	resultSlice := make([]interface{}, mappable.Len())

	if mappable.Len() == 0 {
		return resultSlice
	}

	cpus := runtime.NumCPU()
	chunkSize := mappable.Len() / cpus
	remainder := mappable.Len() % cpus

	doneChan := make(chan struct{})
	head, goRoutinesCount := 0, 0
	for tail := chunkSize + remainder; head < mappable.Len(); tail += chunkSize {
		inSlice := mappable.Slice(head, tail)
		goRoutinesCount++
		go worker(mappingFunc, inSlice, resultSlice[head:tail], doneChan)
		head = tail
	}

	for i := 0; i < goRoutinesCount; i++ {
		<-doneChan
	}

	return resultSlice
}

package asynctools

import (
  "runtime"
  "sync"
)

type Mappable interface {
	At(int) interface{}
	Len() int
	Slice(int, int) Mappable
}

type mappingFuncType func(interface{}) interface{}

func worker(mappingFunc mappingFuncType, inSlice Mappable, outSlice []interface{}, wg *sync.WaitGroup) {
  defer wg.Done()

	for i := 0; i < inSlice.Len(); i++ {
		outSlice[i] = mappingFunc(inSlice.At(i))
	}
}

var cpus int

func init() {
  cpus = runtime.NumCPU()
}

func Map(mappable Mappable, mappingFunc mappingFuncType) []interface{} {
	resultSlice := make([]interface{}, mappable.Len())

	if mappable.Len() == 0 {
		return resultSlice
	}

	chunkSize := mappable.Len() / cpus
	remainder := mappable.Len() % cpus

  var wg sync.WaitGroup
	head := 0

  //for each chunk
	for tail := chunkSize + remainder; head < mappable.Len(); tail += chunkSize {
		inSlice := mappable.Slice(head, tail)

    wg.Add(1)

    go worker(mappingFunc, inSlice, resultSlice[head:tail], &wg)

		head = tail
	}

  wg.Wait()

	return resultSlice
}

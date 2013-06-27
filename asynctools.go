package asynctools

import "runtime"

type Mappable interface {
  At(int) interface{}
  Len()int
}

type sliceElement struct {
  index int
  value interface{}
}

type mappingFuncType func(interface{})interface{}

func worker(mappingFunc mappingFuncType, inChannel, outChannel chan sliceElement) {
  for elem := range inChannel {
    outChannel <- sliceElement{elem.index, mappingFunc(elem.value)}
  }
}

func Map(mappable Mappable, mappingFunc mappingFuncType) []interface{} {
  cpus := runtime.NumCPU()

  inChannel := make(chan sliceElement, mappable.Len())
  go func() {
    for i := 0; i < mappable.Len(); i++ {
      inChannel <- sliceElement{i, mappable.At(i)}
    }

    close(inChannel)
  }()

  outChannel := make(chan sliceElement, mappable.Len())
  for i:= 0; i < cpus; i++ {
    go worker(mappingFunc, inChannel, outChannel)
  }

  result := make([]interface{}, mappable.Len())

  for i := 0; i < mappable.Len(); i++{
    v := <- outChannel
    result[v.index] = v.value
  }

  return result
}

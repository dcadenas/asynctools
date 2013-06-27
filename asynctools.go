package asynctools

import "runtime"

type sliceElement struct {
  index int
  value interface{}
}

func worker(mappingFunc func(interface{})interface{}, inChannel, outChannel chan sliceElement) {
  for elem := range inChannel {
    outChannel <- sliceElement{elem.index, mappingFunc(elem.value)}
  }
}

func Map(mappable []interface{}, mappingFunc func(interface{})interface{}) []interface{} {
  cpus := runtime.NumCPU()

  inChannel := make(chan sliceElement, len(mappable))
  go func() {
    for i, v := range mappable {
      inChannel <- sliceElement{i, v}
    }

    close(inChannel)
  }()

  outChannel := make(chan sliceElement, len(mappable))
  for i:= 0; i < cpus; i++ {
    go worker(mappingFunc, inChannel, outChannel)
  }

  result := make([]interface{}, len(mappable))

  for i := 0; i < len(mappable); i++{
    v := <- outChannel
    result[v.index] = v.value
  }

  return result
}

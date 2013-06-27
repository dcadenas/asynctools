package asynctools

import (
  "testing"
  "runtime"
  "time"
)

func init() {
  runtime.GOMAXPROCS(runtime.NumCPU())
}

func MapOneCpu(mappable Mappable, mappingFunc mappingFuncType) []interface{} {
  result := make([]interface{}, mappable.Len())
  for i := 0; i < mappable.Len(); i++ {
    result[i] = mappingFunc(mappable.At(i))
  }

  return result
}

func expensiveCPUBoundFunction(val interface{}) interface{} {
  value := val.(int)

  for i := 1; i < 10000; i++ {
    value += value
  }

  return value
}

func expensiveIOBoundFunction(val interface{}) interface{} {
  time.Sleep(1 * time.Millisecond)
  return val
}

func benchmark(b *testing.B, mapFunction func(Mappable, mappingFuncType) []interface{} ,expensiveFunction func(interface{}) interface{}) {
  sliceSize := 100
  mappable := make(intMappable, sliceSize)

  for i := 0; i < sliceSize; i++ {
    mappable[i] = i
  }

  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    mapFunction(mappable, expensiveFunction)
  }
}

func BenchmarkCPUBoundSingle(b *testing.B) {
  benchmark(b, MapOneCpu, expensiveCPUBoundFunction)
}

func BenchmarkCPUBoundMulti(b *testing.B) {
  benchmark(b, Map, expensiveCPUBoundFunction)
}

func BenchmarkIOBoundSingle(b *testing.B) {
  benchmark(b, MapOneCpu, expensiveIOBoundFunction)
}

func BenchmarkIOBoundMulti(b *testing.B) {
  benchmark(b, Map, expensiveIOBoundFunction)
}

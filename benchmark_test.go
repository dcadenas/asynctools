package asynctools

import (
  "testing"
  "runtime"
  "time"
)

func init() {
  runtime.GOMAXPROCS(runtime.NumCPU())
}

func MapOneCpu(mappable []interface{}, mappingFunc func(interface{})interface{}) []interface{} {
  result := make([]interface{}, len(mappable))
  for i, v := range mappable {
    result[i] = mappingFunc(v)
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

func benchmark(b *testing.B, mapFunction func([]interface{}, func(interface{})interface{}) []interface{} ,expensiveFunction func(interface{}) interface{}) {
  sliceSize := 100
  mappable := make([]interface{}, sliceSize)

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

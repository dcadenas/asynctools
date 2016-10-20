asynctools
==========

Playing with some ideas to abstract parallelization.

Map
---

Implements a functional map that splits the traversed array in chunks, one per available CPU.

Usage:

```
	mappable := intMappable{1, 2, 3, 4, 5}

	result := Map(mappable, func(val interface{}) interface{} {
		return val.(int) * 2
	})

	assertEqual(t, result[0], 2) 
	assertEqual(t, result[1], 4)
	assertEqual(t, result[2], 6)
	assertEqual(t, result[3], 8)
	assertEqual(t, result[4], 10)
```

Unscientific benchmarks
-----------------------

Why can I only see a clear benefit in IO bound chunks?

```
$ go test -bench=. -benchtime=20s
BenchmarkCPUBoundSingle-4        1000000             41609 ns/op
BenchmarkCPUBoundMulti-4         1000000             36404 ns/op
BenchmarkIOBoundSingle-4             200         121943497 ns/op
BenchmarkIOBoundMulti-4             1000          30969379 ns/op
PASS
ok      github.com/dcadenas/asynctools  150.100s
```

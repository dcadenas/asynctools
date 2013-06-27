package asynctools

import (
  "testing"
)

func assertEqual(t *testing.T, actual, expected interface{}) {
  if actual != expected {
    t.Error("Should be", expected, "but was", actual)
  }
}

func assert(t *testing.T, actual bool) {
  if !actual {
    t.Error("Should be true")
  }
}

type intMappable []int
func (im intMappable) At(i int) interface{} {
  return im[i]
}

func (im intMappable) Len() int {
  return len(im)
}

func TestEmptySliceDoesNothing(t *testing.T) {
  mappable := intMappable{}
  Map(mappable, func(val interface{}) interface{} {
    t.Error("Should not raise this")
    return nil
  })
}

func TestIdentityMapping(t *testing.T) {
  mappable := intMappable{1, 2, 3, 4, 5}
  result := Map(mappable, func(val interface{}) interface{} {
    return val
  })

  assertEqual(t, result[0], 1)
  assertEqual(t, result[1], 2)
  assertEqual(t, result[2], 3)
  assertEqual(t, result[3], 4)
  assertEqual(t, result[4], 5)
}

func TestDoubleMap(t *testing.T) {
  mappable := intMappable{1, 2, 3, 4, 5}
  result := Map(mappable, func(val interface{}) interface{} {
    return val.(int) * 2
  })

  assertEqual(t, result[0], 2)
  assertEqual(t, result[1], 4)
  assertEqual(t, result[2], 6)
  assertEqual(t, result[3], 8)
  assertEqual(t, result[4], 10)
}


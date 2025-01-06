package main

import (
	"fmt"
)

func main() {
	// Integer types
	var a int = 32
	var b int8 = 8
	var c int16 = 16
	var d int32 = 32
	var e int64 = 64

	fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n", a, b, c, d, e)

	// Unsigned integer types
	var f uint = 32
	var g uint8 = 8
	var h uint16 = 16
	var i uint32 = 32
	var j uint64 = 64
	var k uintptr = 0x123456

	fmt.Printf("uint: %d, uint8: %d, uint16: %d, uint32: %d, uint64: %d, uintptr: %#x\n", f, g, h, i, j, k)

	// Float types
	var m float32 = 3.14
	var n float64 = 2.71828

	fmt.Printf("float32: %f, float64: %f\n", m, n)

	// Complex types
	var o complex64 = complex(1, 2)
	var p complex128 = complex(1, 2)

	fmt.Printf("complex64: %v, complex128: %v\n", o, p)

	// Demonstrate type conversion
	var temperatureFloat float64 = 88.26
	var temperatureInt int64 = int64(temperatureFloat)

	fmt.Printf("Original float: %f, Converted to int64 (truncated): %d\n", temperatureFloat, temperatureInt)

	// Converting between integer types
	var originalInt int = 1024
	var smallerInt int8 = int8(originalInt)

	fmt.Printf("Original int: %d, Converted to int8 (potential overflow): %d\n", originalInt, smallerInt)
}

package main

import "fmt"

// Example 1: Enum-like constants with iota
const (
	ReadPermission  = iota // 0
	WritePermission        // 1
	ExecPermission         // 2
)

// Example 2: Bitmasking with iota (permissions as flags)
const (
	ReadFlag  = 1 << iota // 1 << 0 = 1
	WriteFlag             // 1 << 1 = 2
	ExecFlag              // 1 << 2 = 4
)

// Example 3: File size units using iota
const (
	_  = iota             // Ignore the first value (0)
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1048576
	GB                    // 1 << 30 = 1073741824
	TB                    // 1 << 40 = 1099511627776
)

// Example 4: Custom values using iota
const (
	ErrorLevelNone   = iota * 10 // 0 * 10 = 0
	ErrorLevelLow                // 1 * 10 = 10
	ErrorLevelMedium             // 2 * 10 = 20
	ErrorLevelHigh               // 3 * 10 = 30
)

// Example 5: Skipping values with iota
const (
	FirstValue  = iota // 0
	_                  // Skip 1
	SecondValue        // 2
	ThirdValue         // 3
)

func main() {
	// Example 1: Enum-like constants
	fmt.Println("Permissions:")
	fmt.Println("ReadPermission:", ReadPermission)
	fmt.Println("WritePermission:", WritePermission)
	fmt.Println("ExecPermission:", ExecPermission)

	// Example 2: Bitmasking
	fmt.Println("\nBitmasking:")
	myPermissions := ReadFlag | ExecFlag
	fmt.Printf("MyPermissions: %b (Readable: %t, Executable: %t, Writable: %t)\n",
		myPermissions,
		myPermissions&ReadFlag != 0,
		myPermissions&ExecFlag != 0,
		myPermissions&WriteFlag != 0,
	)

	// Example 3: File size units
	fmt.Println("\nFile Sizes:")
	fmt.Println("1 KB =", KB, "bytes")
	fmt.Println("1 MB =", MB, "bytes")
	fmt.Println("1 GB =", GB, "bytes")
	fmt.Println("1 TB =", TB, "bytes")

	// Example 4: Custom values
	fmt.Println("\nError Levels:")
	fmt.Println("None:", ErrorLevelNone)
	fmt.Println("Low:", ErrorLevelLow)
	fmt.Println("Medium:", ErrorLevelMedium)
	fmt.Println("High:", ErrorLevelHigh)

	// Example 5: Skipping values
	fmt.Println("\nSkipping Values:")
	fmt.Println("FirstValue:", FirstValue)
	fmt.Println("SecondValue:", SecondValue)
	fmt.Println("ThirdValue:", ThirdValue)
}

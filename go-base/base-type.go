package go_base

import "fmt"

//bool
//string
//int、int8、int16、int32、int64
//uint、uint8、uint16、uint32、uint64、uintptr
//byte // uint8 的别名
//rune // int32 的别名 代表一个 Unicode 码
//float32、float64
//complex64、complex128

// go 的基本类型
func BaseType() {
	// 布尔类型
	var b bool
	b = true
	fmt.Println("布尔类型", b)

	// string 类型
	var str string
	str = "字符串"
	fmt.Println("字符串类型", str)

	// int 类型
	var i int
	i = 1
	fmt.Println("int", i)

	// uint8 类型
	var ui8 uint8
	ui8 = 8
	fmt.Println("uint8", ui8)

	// uint16 类型
	var ui16 uint8
	ui16 = 16
	fmt.Println("uint16", ui16)

	// uint32类型
	var ui32 uint8
	ui32 = 32
	fmt.Println("ui32", ui32)

	// uint64 类型
	var ui64 uint8
	ui64 = 64
	fmt.Println("uint64", ui64)

	// byte uint8 的别名
	var by byte
	by = 'a'
	fmt.Println("byte类型", by)

	// rune int32 的别名 代表一个 Unicode 码
	var ru rune
	ru = 11
	fmt.Println("rune类型", ru)

	// float32 类型
	var f32 float32
	f32 = 1.2
	fmt.Println("float32类型", f32)

	var f64 float64
	// float32 类型
	f64 = 1.4
	fmt.Println("float64类型", f64)

	var c64 complex64
	c64 = 122
	fmt.Println(c64)

	var c128 complex128
	c128 = 1222
	fmt.Println(c128)
}

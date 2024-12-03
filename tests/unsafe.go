package tests

import (
	"fmt"
	"reflect"
	"unsafe"
)

func SizeWithoutPadding(s interface{}) uintptr {
	v := reflect.ValueOf(s)
	t := v.Type()
	size := uintptr(0)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i).Type
		fieldSize := unsafe.Sizeof(reflect.New(fieldType).Elem().Interface())
		size += fieldSize
	}

	return size
}

func Convert() {
	var x int = 42
	var p *int = &x

	// Convert *int to *float64
	var pf *float64 = (*float64)(unsafe.Pointer(p))

	// Dereference the pointer (this is unsafe and may lead to undefined behavior)
	fmt.Println(*pf) // This will print a garbage value
}

func Accessing() {
	var x int = 42
	var p *int = &x

	// Convert *int to unsafe.Pointer
	up := unsafe.Pointer(p)

	// Convert unsafe.Pointer back to *int
	p2 := (*int)(up)

	// Dereference the pointer
	fmt.Println(*p2) // This will print 42
}

func PointerArithmetic() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var p *int = &arr[0]

	// Convert *int to uintptr
	up := uintptr(unsafe.Pointer(p))

	// Perform pointer arithmetic
	up += unsafe.Sizeof(arr[0]) * 2 // Move to the third element
	up++                            // Another move
	up <<= 10

	// Convert uintptr back to *int
	p2 := (*int)(unsafe.Pointer(up))

	// Dereference the pointer
	fmt.Println(*p2) // This will print 3
}

func ConvertingPointerToSlice() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var p *int = &arr[0]

	// Convert *int to unsafe.Pointer
	up := unsafe.Pointer(p)

	// Convert unsafe.Pointer to a slice
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	sliceHeader.Data = uintptr(up)
	sliceHeader.Len = 5
	sliceHeader.Cap = 5

	slice := *(*[]int)(unsafe.Pointer(sliceHeader))

	// Print the slice
	fmt.Println(slice) // This will print [1 2 3 4 5]
}

func ConvertingSliceToPointer() {
	var slice []int = []int{1, 2, 3, 4, 5}

	// Convert slice to unsafe.Pointer
	up := unsafe.Pointer(&slice[0])

	// Convert unsafe.Pointer back to *int
	p := (*int)(up)

	// Dereference the pointer
	fmt.Println(*p) // This will print 1
}

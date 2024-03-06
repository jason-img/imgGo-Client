package main

import (
	"fmt"
	"testing"
	"time"
)

func TestUserLogin(t *testing.T) {
	fmt.Println(time.Now().Unix())
}

func TestUserReg(t *testing.T) {

}

// 泛型
type Numeric interface {
	float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | complex64 | complex128
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}

func TestTest233(t *testing.T) {
	fmt.Println(AddInt(1, 2))
	fmt.Println(AddFloat(1.1, 2.2))
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1.1, 2.2))
}

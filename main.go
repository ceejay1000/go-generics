package main

import "fmt"

type CustomeInteger int

// type Subtractable interface {
// 	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | CustomeInteger
// }

type Subtractable interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64 | uint
}

type Movable[S Subtractable] interface {
	Move(S)
}

func Move[V Movable[S], S Subtractable](v V, distance, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}

type Person[S Subtractable] struct {
	Name string
}

func (p *Person[S]) Move(meter S) {
	fmt.Printf("%s moved %d meters\n", p.Name, meter)
}

type Car[S Subtractable] struct {
	Name string
}

func (c *Car[S]) Move(meter S) {
	fmt.Printf("%s moved %d meters\n", c.Name, meter)
}

// Slice of results that are subtractable0
// type Results[T Subtractable] []T
type Results[T any] []T
type ResultsComparable[T comparable] []T

func main() {
	res1 := Subtract(1, 2)
	fmt.Println(res1)

	res2 := Subtract(2.5, 2.2)
	fmt.Println(res2)

	res3 := Subtract[int](20, 2)
	fmt.Println(res3)

	res4 := Subtract[CustomeInteger](10, 5)
	fmt.Println(res4)

	var resultStorage Results[int]
	resultStorage = append(resultStorage, res1)

	// ---------------------------------------------
	p := Person[int]{Name: "John"}
	c := Car[int]{Name: "Subaru"}

	val := Move[*Person[int], int](&p, 20, 29)
	Move(&c, 20, 30)
	fmt.Println(val)
}

// func Subtract(a, b int) int {
// 	return a - b
// }

// func Subtract[V int | float64](a, b V) V {
// 	return a - b
// }

func Subtract[V Subtractable](a, b V) V {
	return a - b
}

func SubtractFloat(a, b float32) float32 {
	return a - b
}

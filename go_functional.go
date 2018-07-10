package main

import "fmt"

type any interface{}

func main() {
	filterNumGreaterThan10 := func(v any) bool { return (v).(int) < 10 }
	squareNum := func(v any) any { return (v).(int) * (v).(int) }
	addNum := func(v1, v2 any) any { return (v1).(int) + (v2).(int) }

	vs := []any{1, 2, 3, 12, 20}

	/*	nvs1 := Filter(filterNumGreaterThan10, vs...)

		nvs2 := Map(squareNum, nvs1...)

		nvs3 := Reduce(addNum, nvs2...)

		fmt.Println(nvs3) */

	fmt.Println(Reduce(addNum, Map(squareNum, Filter(filterNumGreaterThan10, vs...)...)...))
}

func Filter(f func(v any) bool, vs ...any) []any {
	nvs := make([]any, 0)
	for _, value := range vs {
		if f(value) {
			nvs = append(nvs, value)
		}
	}
	return nvs
}

func Map(f func(v any) any, vs ...any) []any {
	nvs := make([]any, len(vs))
	for i, value := range vs {
		nvs[i] = f(value)
	}
	return nvs
}

func Reduce(f func(v1, v2 any) any, vs ...any) any {
	var ret any
	for i, value := range vs {
		if i == 0 {
			ret = value
		} else {
			ret = f(ret, value)
		}
	}
	return ret
}

package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		n      int32
		A      [][]int32
		result int
	}{
		{
			nombre: "Case 1:",
			n:      10,
			A:      [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}},
			result: 10,
		},

		{
			nombre: "Case 2:",
			n:      5,
			A:      [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}},
			result: 200,
		},

		{
			nombre: "Case 3:",
			n:      10,
			A:      [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}},
			result: 10,
		},

		{
			nombre: "Case 4:",
			n:      10,
			A:      [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}},
			result: 31,
		},

		{
			nombre: "Case 5:",
			n:      11,
			A:      [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}, {10, 10, 12}, {10, 11, 1}},
			result: 13,
		},

		{
			nombre: "Case 6:",
			n:      14,
			A:      [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}, {10, 10, 12}, {10, 11, 1}, {13, 13, 14}},
			result: 14,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := arrayManipulation(test.n, test.A)

			if int(result) != test.result {
				t.Errorf("Se esperaba un retorno de %d, y se retornó %d", test.result, result)
			}
		})
	}

}

func BenchmarkArr1(b *testing.B) {
	fmt.Println("Benchmark algoritmo A")
	data := make([][]int32, 0)
	n := int32(math.Pow(10, 7))
	m := int(2 * math.Pow(10, 5))
	for i := 1; i <= m; i++ {
		a := int32(1)
		b := int32(n)
		k := int32(math.Pow(10, 9))
		data = append(data, []int32{a, b, k})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyData := make([][]int32, len(data))
		copy(copyData, data)
		arrayManipulation(n, copyData)
	}
}

func BenchmarkArr2(b *testing.B) {
	fmt.Println("Benchmark algoritmo B")
	data := make([][]int32, 0)
	n := int32(math.Pow(10, 8))
	m := int(2 * math.Pow(10, 6))
	for i := 1; i <= m; i++ {
		a := int32(1)
		b := int32(n)
		k := int32(math.Pow(10, 9))
		data = append(data, []int32{a, b, k})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copyData := make([][]int32, len(data))
		copy(copyData, data)
		arrayManipulation2(n, copyData)
	}
}

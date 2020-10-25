package main

import (
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		s      string
		result int32
	}{
		// {
		// 	nombre: "Case 1:",
		// 	s:      "abba",
		// 	result: 4,
		// },

		{
			nombre: "Case 2:",
			s:      "abcd",
			result: 0,
		},

		// {
		// 	nombre: "Case 3:",
		// 	s:      "ifailuhkqq",
		// 	result: 3,
		// },

		// {
		// 	nombre: "Case 4:",
		// 	s:      "kkkk",
		// 	result: 10,
		// },

		// {
		// 	nombre: "Case 5:",
		// 	s:      "cdcd",
		// 	result: 5,
		// },
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := sherlockAndAnagrams(test.s)

			if result != test.result {
				t.Errorf("Se esperaba un retorno de %d, y se retornó %d", test.result, result)
			} else {
				t.Logf("Pasa test %s", test.nombre)
			}
		})
	}

}

/*
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
*/

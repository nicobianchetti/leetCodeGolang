package main

/*
https://www.hackerrank.com/challenges/new-year-chaos/problem

It's New Year's Day and everyone's in line for the Wonderland rollercoaster ride! There are a number of people queued up, and each person wears a sticker indicating their initial position in the queue. Initial positions increment by  from  at the front of the line to  at the back.

Any person in the queue can bribe the person directly in front of them to swap positions. If two people swap positions, they still wear the same sticker denoting their original places in line. One person can bribe at most two others. For example, if  and  bribes , the queue will look like this: .

Fascinated by this chaotic queue, you decide you must know the minimum number of bribes that took place to get the queue into its current state!

*/

func minimumBribes(q []int32) (int, bool) {
	// func minimumBribes(q []int32) {
	length := len(q)
	isError := false
	var isOrder bool
	moves := 0
	helper := make([]int, length+1)

	for i, v := range q {
		if (int(v) - (i + 1)) > 2 {
			isError = true
			moves = 0
			return moves, isError
		}
		helper[v] = i
	}

	for i := 1; i < length; i++ {
		isOrder = false
		for isOrder == false {
			if i < (helper[i] + 1) {
				moves++
				qOrigen := q[helper[i]]
				qDestino := q[helper[i]-1]
				helperOrigen := helper[qOrigen]
				helperDestino := helper[qDestino]
				q[helper[i]], q[helper[i]-1] = qDestino, qOrigen
				helper[qOrigen], helper[qDestino] = helperDestino, helperOrigen
			} else {
				isOrder = true
			}
		}
	}

	// if isError {
	//     fmt.Println("Too chaotic")
	// } else {
	//     fmt.Println(moves)
	// }

	return moves, isError
}

/*
Solution 1 , Error Timeout
func minimumBribes(q []int32) (int, bool) {
	// func minimumBribes(q []int32) {
	qSort := make([]int32, len(q))
	copy(qSort, q)
	length := len(q)
	moves := 0
	isError := false
	ArrayIndex := make([]int, 10)

	//https://forum.golangbridge.org/t/how-to-sort-int32/10529/2
	sort.Slice(qSort, func(i, j int) bool { return qSort[i] < qSort[j] })

	// fmt.Println(qSort)
	isOrder := false

	for isOrder == false {
		if reflect.DeepEqual(q, qSort) {
			isOrder = true
			break
		}
		for i := 0; i < length; i++ {
			if q[i] > q[i+1] {
				if ArrayIndex[q[i]] >= 2 || ArrayIndex[q[i+1]] >= 2 {
					isOrder = true
					moves = 0
					isError = true
				} else {
					ArrayIndex[q[i]]++
					// ArrayIndex[q[i+1]]++
					//Permuto
					origen := q[i]
					destino := q[i+1]
					q[i] = destino
					q[i+1] = origen
					moves++
				}
				break
			}
		}
	}

	return moves, isError
}

*/

/*
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	qSort := make([]int32, len(q))
	copy(qSort, q)
	length := len(q)
	moves := 0
	isError := false
	ArrayIndex := make([]int, 10)

	//https://forum.golangbridge.org/t/how-to-sort-int32/10529/2
	sort.Slice(qSort, func(i, j int) bool { return qSort[i] < qSort[j] })

	// fmt.Println(qSort)
	isOrder := false

	for isOrder == false {
		if reflect.DeepEqual(q, qSort) {
			isOrder = true
			break
		}
		for i := 0; i < length; i++ {
			if q[i] > q[i+1] {
				if ArrayIndex[q[i]] >= 2 || ArrayIndex[q[i+1]] >= 2 {
					isOrder = true
					moves = 0
					isError = true
				} else {
					ArrayIndex[q[i+1]]++
					//Permuto
					origen := q[i]
					destino := q[i+1]
					q[i] = destino
					q[i+1] = origen
					moves++
				}
				break
			}
		}
	}

	if isError {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(moves)
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

*/

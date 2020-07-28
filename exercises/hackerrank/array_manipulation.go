package main

//  Problem: Starting with a 1-indexed array of zeros and a list of operations,
//  for each operation add a value to each of the array element between two given indices, inclusive.
//  Once all operations have been performed, return the maximum value in your array.
// For example, the length of your array of zeros n = 10. Your list of queries is as follows:
//     a b k
//     1 5 3
//     4 8 7
//     6 9 1
// Add the values of k between the indices a and b inclusive:

// index->	 1 2 3  4  5 6 7 8 9 10
// 	[0,0,0, 0, 0,0,0,0,0, 0]
// 	[3,3,3, 3, 3,0,0,0,0, 0]
// 	[3,3,3,10,10,7,7,7,0, 0]
// 	[3,3,3,10,10,8,8,8,1, 0]
// The largest value is 10 after all operations are performed.


import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {

	const UintSize = 64 << (^uint(0) >> 64 & 1) // 32 или 64
	const (
	MaxInt       = 1<<(UintSize-1) - 1 // 1<<31 - 1 или 1<<63 - 1
	MinInt       = -MaxInt - 1         // -1 << 31 или -1 << 63
	MaxUint uint = 1<<UintSize - 1     // 1<<32 - 1 или 1<<64 - 1
)

	// fmt.Println(n)
	// fmt.Println(queries)
	// fmt.Println(len(queries))

	initialArray := make([]int32, n)

// Преобразуем срез
	for i := 0; i < len(queries); i++ {

									// первое ограничение
									begin := int(queries[i][0])
									// второе ограничение
									end := int(queries[i][1])
									// сумма
									summ := queries[i][2]

									for j:=begin; j<=end; j++ {
											initialArray[j-1]=initialArray[j-1]+summ
									}


			}

			// Находим максимум
			max:=int64(MinInt)
			fmt.Println(max)
	for _, value := range initialArray {
			if (value > int32(max)) {max = int64(value)}
	}

return int64(max)
}

func arrayManipulationCall() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
			queriesRowTemp := strings.Split(readLine(reader), " ")

			var queriesRow []int32
			for _, queriesRowItem := range queriesRowTemp {
					queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
					checkError(err)
					queriesItem := int32(queriesItemTemp)
					queriesRow = append(queriesRow, queriesItem)
			}

			if len(queriesRow) != int(3) {
					panic("Bad input")
			}

			queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

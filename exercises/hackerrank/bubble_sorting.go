package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(arr []int32, comp func(int32, int32) bool) ([]int32, int32) {
	size := len(arr)
	var count int32
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
				count++
			}
		}
	}
	return arr, count
}

// resulting output will be in descending order.
func less(value1 int32, value2 int32) bool {
	return value1 < value2
}

// sorted output will be increasing in order
func more(value1 int32, value2 int32) bool {
	return value1 > value2
}

func BubbleSortCall() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	fmt.Println(s)

	n, _ := strconv.Atoi(s)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < n; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	fmt.Println(arr)

	result, count := BubbleSort(arr, more)

	// result := someFunction(s, n)

	fmt.Fprintf(writer, "Array is sorted in %v swaps.\n", count)
	fmt.Fprintf(writer, "First Element: %v\n", result[0])
	fmt.Fprintf(writer, "Last Element: %v\n", result[len(arr)-1])

	writer.Flush()
}

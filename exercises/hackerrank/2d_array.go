// Problem https://www.hackerrank.com/challenges/2d-array/problem
// Given a 6x6 2D Array
// We define an hourglass in A to be a subset of values with indices falling in this pattern in arr's graphical representation:
// a b c
//   d
// e f g
// Calculate the hourglass sum for every hourglass in arr, then print the maximum hourglass sum.
package main

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {

fmt.Println(arr)

// var tmparr [][]int32
var summ,result int32

var init bool

for i:=0; i<4; i=i+1 {

    for j:=0; j<4; j=j+1 {

        // create empty collection.
    values := [][]int32{}

        row1 := []int32{arr[i][j],arr[i][j+1],arr[i][j+2]}
        row2 := []int32{arr[i+1][j+1]}
        row3 := []int32{arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2]}
        values = append(values, row1)
        values = append(values, row2)
        values = append(values, row3)
        fmt.Println(values)
        summ = 0

        for i := range values {
            for _,v := range values[i] {
                summ = summ + v
            }
        }

// result инициализируется нулём, но может быть и отрицательным, поэтому нужно его инициализировать первым значением суммы.
        if !init {
            result = summ
            init = true
        }

        fmt.Println(summ, result)

        if summ>result { result = summ }

    }

}

return result
}

func twoDArray() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

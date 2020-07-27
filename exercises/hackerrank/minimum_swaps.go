package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates. You are allowed to swap any two elements. You need to find the minimum number of swaps required to sort the array in ascending order.

// Sample Input 1
// 5
// 2 3 4 1 5
// Sample Output 1
// 3
// swapping (2;3)->(0;1)->(0;2)

// Sample Input 2
// 7
// 1 3 5 2 4 6 7
// Sample Output 2
// 3
// swapping (1;3)->(2;3)->(3;4)


// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {

        var t, i, swap, temp int32

        for i=0;int(i)<len(arr);i++{
            if((i+1)!=arr[i]){
                t=i
                for ; arr[t]!=(i+1); {
                    t++
                }
                temp=arr[t]
                arr[t]=arr[i]
                arr[i]=temp
                swap++
            }
        }
        return swap

}


func minimumSwapsCall() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

    writer.Flush()
}

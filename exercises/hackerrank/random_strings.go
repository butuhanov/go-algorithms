package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func main() {

	// Debug
	input := generateInput(20000)
	//fmt.Println(input)

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

// randStringBytesMaskImprSrcUnsafe generates random strings from a specified range of letters of a specified length
func randStringBytesMaskImprSrcUnsafe(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// generateInput generates debug input
func generateInput(n int) string {
	t0 := time.Now() // замеряем производительность
	var result string
	result = strconv.Itoa(n) + "\n" // первая строка с количеством строк
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(n); i++ {
		result = result + randStringBytesMaskImprSrcUnsafe(4) + " " + strconv.Itoa(rand.Intn(100)) + "\n"
	}
	// Строки для проверки
	for i := 0; i < int(n); i++ {
		result = result + randStringBytesMaskImprSrcUnsafe(4) + "\n"
	}
	t1 := time.Now()
	fmt.Printf("Elapsed time generate input: %v \n", t1.Sub(t0))
	return result
}

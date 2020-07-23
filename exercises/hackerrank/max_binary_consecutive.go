package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func maximumBynaryConsecutive() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	// n = 439 // want 3
	// Конвертируем число в двоичный формат
	bin, err := convertInt(strconv.Itoa(int(n)), 10, 2)
	checkError(err)
	fmt.Println(bin)

	fmt.Fprintf(writer, "%d\n", maximumNumberOfConsecutive(bin, "1"))
	writer.Flush()

}

// ConvertInt конвертирует значение из одной системы счисления
// в другую, которая указана в toBase
func convertInt(val string, base, toBase int) (string, error) {
	// const bin = "10111"
	// const hex = "1A"
	// const oct = "12"
	// const dec = "10"
	// const floatNum = 16.123557
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func maximumNumberOfConsecutive(val string, symbol string) int {
	var count, maxcount int

	// Для учета символов юникода работаем с рунами
	for i, w := 0, 0; i < len(val); i += w {
		runeValue, width := utf8.DecodeRuneInString(val[i:])
		s := fmt.Sprintf("%c", runeValue)

		// Поиск максимальной последовательности
		// Увеличиваем счетчик при каждом попадании искомого числа и сбрасываем при непопадании (начинаем отчет заново)
		if s == symbol {
			count++
			if maxcount < count {
				maxcount = count
			}
		} else {
			count = 0
		}
		// fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	return maxcount
}

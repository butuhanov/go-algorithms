// Problem: Given  names and phone numbers, assemble a phone book that maps friends' names to their respective phone numbers. You will then be given an unknown number of names to query your phone book for. For each name queried, print the associated entry from your phone book on a new line in the form name=phoneNumber; if an entry for name is not found, print Not found instead.

package main
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"

)

func main() {

    //Enter your code here. Read input from STDIN. Print output to STDOUT


    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)


    n, err := strconv.ParseInt(readLine(reader), 10, 32) // считываем первую строку, получаем количество элементов
    checkError(err)

    //m := make(map[string]int)

    m := make(map[string]string)

    t0 := time.Now();   // замеряем производительность
    // читаем входные строки
    for i:=0; i<int(n); i++ {

        // разделяем строку по пробелам
        arrRowTemp := strings.Split(readLine(reader), " ")
        // переводим в int, хотя это в этой задаче можно было не делать
        // n, err := strconv.ParseInt(arrRowTemp[1], 10, 64)
        // checkError(err)
        // m[arrRowTemp[0]] = int(n)

        m[arrRowTemp[0]] = arrRowTemp[1]

    }

    t1 := time.Now();
    fmt.Printf("Elapsed time loop 1: %v \n", t1.Sub(t0));

    t0 = time.Now();   // замеряем производительность

var result strings.Builder // оптимизация конкатенации

    for { // продолжаем читать строки пока они не закончатся
        s := readLine(reader)

        // если достигли конец ввода
        if s =="" {
            break
        }

        // Делаем перенос строки, но он не нужен в начале и в конце
        if len(result.String())!=0 {
            result.WriteString("\n")
        }

        // находим значение для ключа
        _, ok := m[s] // ok равен true, если ключ найден

        if ok {
            // result = result  + s + "=" + strconv.Itoa(m[s])
            result.WriteString(s)
            result.WriteString("=")
            result.WriteString(m[s])
        } else {
           result.WriteString("Not found")
        }


    }

    t1 = time.Now();
    fmt.Printf("Elapsed time loop 2: %v", t1.Sub(t0));



   // fmt.Fprintf(writer, "%s", result)
   fmt.Fprintf(writer, "%s", result.String())


    writer.Flush()

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

package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    l := len(arr)
    var zero, pos, neg float64
    for i := 0; i < l; i++{
        if arr[i] == 0{
            zero = zero + float64(1/float64(l))
        }
        if arr[i] > 0 {
            pos = pos + float64(1/float64(l))
        }

        if arr[i] < 0 {
            neg = neg + float64(1/float64(l))
        }
    }

    fmt.Printf("%.6f\n",pos)
    fmt.Printf("%.6f\n",neg)
    fmt.Printf("%.6f\n",zero)
}

func main() {
   arr := []int32{5,-4,3,-9,0,4,1}

    plusMinus(arr)
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

package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    var min, max,sum uint64
    for j := 0; j < len(arr); j++{
        sum = 0
        for i := 0; i < len(arr); i++{
            if i != j{
                sum = sum + uint64(arr[i])
            }
        }

         if max < sum{
            max = sum
         }
         if(min == 0 || min > sum){
            min = sum
         }

    }
    fmt.Println(min,max)
}

func main() {
   arr := []int32{769082435,210437958,673982045,375809214,380564127}
    miniMaxSum(arr)
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

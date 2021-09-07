package main

import (
        "fmt"
        "time"
        "sort"
)

// arr = [8,7,2,5,3,1]
// target = 10
// o/p = (8,2)

// arr = [5,2,6,8,1,9]
// target = 12
// o/p = pair not found
// 27.481µs
func findPair0(arr []int, target int){
    check := 0
    for i :=0; i < len(arr); i++{
        for j := 0; j < len(arr); j++{
            if j != i{
                t := arr[i] + arr[j]
                if target == t {
                    fmt.Printf("(%d,%d)\n", arr[i], arr[j])
                    check = 1
                    break
                }
            }
        }
    }
    if check == 0{
        fmt.Println("pair not found")
    }
}

//33.391µs
func findPair1(arr []int, target int){
    check := 0
    for i :=0; i < len(arr); i++{
        for j := i + 1; j < len(arr); j++{
           t := arr[i] + arr[j]
           if target == t {
               fmt.Printf("(%d,%d)\n", arr[i], arr[j])
               check = 1
               break
           }
        }
    }
    if check == 0{
        fmt.Println("pair not found")
    }
}
//d/q 33.764µs
func findPair2(arr []int, target int){
    sort.Ints(arr)
    mid := (len(arr)-1)/2
    for i := 0; i < mid; i++{
        if target == (arr[i] + arr[mid]){
            fmt.Printf("(%d,%d)\n", arr[i], arr[mid])
            return
        }

        if target > (arr[i] + arr[mid]){
            mid = mid + 1
        }else{
            mid = mid - 1
        }
    }
    fmt.Println("pair not found")
}

func main(){
    arr1 := []int{8,7,2,5,3,1}
    arr2 := []int{5,2,6,8,1,9}

    now1 := time.Now()
    findPair1(arr1, 10)
    duration1 := time.Since(now1)
    fmt.Println("Duration 1")
    fmt.Println(duration1)


    now2 := time.Now()
    findPair1(arr2, 12)
    duration2 := time.Since(now2)
    fmt.Println("Duration 2")
    fmt.Println(duration2)
}
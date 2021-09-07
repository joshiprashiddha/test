

package main

import (
        "fmt"
        "time"
        )

func duplicate(items map[int]int, element int) bool{
    _, ok := items[element]
    return ok
}

func appender(collection map[int]int, n []int, i int) int{
    if duplicate(collection, n[i]){
        return len(collection)
    }
    collection[n[i]] = n[i]
    return appender(collection, n, n[i])
}

func arrayNesting(nums []int) int {
    longest := 0
    for i := 0; i < len(nums); i++{
        l := appender(map[int]int{},nums, i)
        if longest < l {
            longest = l
        }
    }
    return longest
}
//5,4,0,3,1,6,2
//len = 7
func arrayNestingg(a []int) int {
        maxlen := 0
        for i := 0; i < len(a); i++{
            fmt.Println(a)
            size := 0;
            for j := i ; a[j] >= 0; size++{
                aj := a[j]
                a[j] = -1
                j = aj
            }
            fmt.Println(size)
            if maxlen < size {
                maxlen = size
            }
        }
        return maxlen
    }

func main(){
    nums := []int{5,4,0,3,1,6,2}
    start  := time.Now()
    f := arrayNestingg(nums)
    duration := time.Since(start)
    fmt.Println(f)
    fmt.Println(duration)
}
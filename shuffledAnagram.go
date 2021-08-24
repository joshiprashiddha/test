package main

import "fmt"

func checkAnagram(input string, possibleAnagram string) bool{
    if len(input) != len(possibleAnagram){
        return false
    }

    for j:= 0; j<len(input);j++{
        if(input[j] == possibleAnagram[j]){
            return false
        }
    }

    return true
}

func anagram(input string) string{
    for i := range(input){
        possibleAnagram := input[i+1:]+input[:i+1]
        if(checkAnagram(input, possibleAnagram)){
            return possibleAnagram
        }
    }
    return "IMPOSSIBLE"
}

func main(){
    var numOfCases int
    fmt.Scanln(&numOfCases)
    for i := 0; i < numOfCases; i++{
        var testcase string
        fmt.Scanln(&testcase)
        fmt.Printf("Case #%d: %s\n", i+1 , anagram(testcase))
    }
}
package main
/*
Problem
Let S be a string containing only letters of the English alphabet. An anagram of S is any string that contains exactly the same letters as S (with the same number of occurrences for each letter), but in a different order. For example, the word kick has anagrams such as kcik and ckki.

Now, let S[i] be the i-th letter in S. We say that an anagram of S, A, is shuffled if and only if for all i, S[i]≠A[i]. So, for instance, kcik is not a shuffled anagram of kick as the first and fourth letters of both of them are the same. However, ckki would be considered a shuffled anagram of kick, as would ikkc.

Given an arbitrary string S, your task is to output any one shuffled anagram of S, or else print IMPOSSIBLE if this cannot be done.

Input
The first line of the input gives the number of test cases, T. T test cases follow. Each test case consists of one line, a string of English letters.

Output
For each test case, output one line containing Case #x: y, where x is the test case number (starting from 1) and y is a shuffled anagram of the string for that test case, or IMPOSSIBLE if no shuffled anagram exists for that string.

Limits
Memory limit: 1 GB.
1≤T≤100.
All input letters are lowercase English letters.

Test Set 1
Time limit: 20 seconds.
1≤ the length of S ≤8.

Test Set 2
Time limit: 40 seconds.
1≤ the length of S ≤104.

Sample Input
2
start
jjj
cccdcada

Output
Case #1: tarts
Case #2: IMPOSSIBLE
Case #3: dadcaccc

*/

import (
        "fmt"
        "sort"
       )

type MyRune struct{
    pos int
    value rune
}

type ByPos []MyRune
func (p ByPos) Len() int           { return len(p) }
func (p ByPos) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPos) Less(i, j int) bool { return p[i].pos < p[j].pos }

type ByValue []MyRune
func (v ByValue) Len() int           { return len(v) }
func (v ByValue) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByValue) Less(i, j int) bool { return v[i].value < v[j].value }

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
    myRune := []MyRune{}
    for i, c := range(input) {
        r := MyRune{pos: i, value: c}
        myRune = append(myRune,r)
    }
    sort.Sort(ByValue(myRune))
    mid := len(myRune)/2

    if len(myRune)%2 != 0{
        temp := myRune[len(myRune) - 1].value
        myRune[len(myRune) - 1].value = myRune[mid].value
        myRune[mid].value = temp
    }

    for j := 0; j < mid; j++{
        temp := myRune[j].value
        myRune[j].value = myRune[mid+j].value
        myRune[mid+j].value = temp
    }
    finalString := ""
    sort.Sort(ByPos(myRune))

    for j, _ := range(myRune){
        finalString = finalString + string(myRune[j].value)
    }
    if( !checkAnagram(input, finalString) ){
       return "IMPOSSIBLE"
    }
    return finalString
}

func main(){
    var numOfCases int
    fmt.Scanln(&numOfCases)
    var input []string
    for i := 0; i < numOfCases; i++{
        var testcase string
        fmt.Scanln(&testcase)
        input = append(input, fmt.Sprintf("Case #%d: %s", i+1 , anagram(testcase)))
    }
    for j := 0; j < numOfCases; j++{
        fmt.Printf("%s\n",input[j])
    }
}
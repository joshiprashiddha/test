package main

import (
            "fmt"
            "errors"
            )
/* test 2
Rune open/close formatting check, the possible runes are {[()]}, each opening rune should have a closing one.
After all runes are closed there should be no open runes left.

 {}{} == good
{{}} == good
{[]} == good
{[()]} == good
{{[()()]}} == good
{}[] == good
{[}] == bad, not in sequence
{{[()]]} == bad, opening { is not closed when closing ] was given
{{[] == bad, no closing rune for {
{}] == bad, no opening rune for ]
*/

// rune items will be stacked into rune stack
// stack is a LIFO data structure
type RuneStack struct{
    items []string
}

// push adds items to end of stack
func (rs *RuneStack) push(r string){
    rs.items = append(rs.items, r)
}

// pop removes items from end of stack
func (rs *RuneStack) pop() (string, error){
    l := len(rs.items) - 1
    if l < 0{
        return "", errors.New("pop empty stack nothing to remove")
    }
    item := rs.items[l]
    rs.items = rs.items[:l]
    return item, nil
}

type runecase struct{
    input string
    output string
}

func RuneChecker(i string) string{
    runeStack := RuneStack{}
    result := ""
    for _,r := range(i){
        result = "good"
        if string(r) == "[" || string(r) == "{" || string(r) == "(" {
            runeStack.push(string(r))
            result = "bad"
        }
        if string(r) == "]" {
            rc, _ := runeStack.pop()
            if(rc != "["){
                result = "bad"
                return result
            }
        }
        if string(r) == "}" {
            rc, _ := runeStack.pop()
            if(rc != "{"){
                result = "bad"
                return result
            }
        }
        if string(r) == ")" {
            rc, _ := runeStack.pop()
            if(rc != "("){
                result = "bad"
                return result
            }
        }
    }
    return result
}

func TestRuneChecker(ces []runecase){
    for _,r:= range(ces){
        output := RuneChecker(r.input)
        if output != r.output{
            fmt.Println(fmt.Sprintf("FAIL::Expected output: %s for input %s but got %s", r.output, r.input, output))
        }else{
            fmt.Printf("==PASS::Expected output: %q got %q\n", r.output, output)
        }
    }
}

func main(){
    rc := []runecase {
            {
                input: "{}{}",
                output: "good",
            },
            {
                input: "{{}}",
                output: "good",
            },
            {
                input: "{[]}",
                output: "good",
            },
            {
                input: "{[()]}",
                output: "good",
            },
            {
                input: "{{[()()]}}",
                output: "good",
            },
            {
                input: "{}[",
                output: "bad",
            },
            {
                input: "{{[()]]}",
                output: "bad",
            },
       }
    TestRuneChecker(rc)
}
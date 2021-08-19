package main

import (
        "fmt"
        "errors"
        )


// Stack
type Stack struct{
    items []int
}
//push to the top of stack
func (s *Stack) push(i int){
    s.items = append(s.items,i)
}

//pop removes items from top of stack
func (s *Stack) pop() (int, error){
    l := len(s.items)-1
    if l < 0 {
        return 0, errors.New("pop empty stack nothing to remove")
    }
    iToRemove := s.items[l]
    s.items = s.items[:l]
    return iToRemove, nil
}

func main(){
    myStack := Stack{}
    removedItem, err := myStack.pop()
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(myStack)
    myStack.push(10)
    fmt.Println(myStack)
    removedItem, err = myStack.pop()
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(myStack)
    fmt.Println("Removed Item", removedItem)
}
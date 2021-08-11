package main

import "fmt"

// Node is basic unit of singly linked list
type Node struct{
    data string
    next *Node
}

// insertion operation
func (l *Node) append(newData string) *Node{
    last := l

    newNode := Node{
        data: newData,
        next: nil,
    }

    // linkedlist is empty
    if (Node{} == *l) {
        l = &newNode
    }else{
        for last.next != nil{
            last = last.next
        }
        last.next = &newNode
    }

    return l
}

// display prints node data present in singly linked list
func (l *Node) display(){
    last := l

    if (Node{} == *last) {
        fmt.Println("Linked List is Empty")
        return
    }

    for ok := true; ok; ok = (last != nil) {
        fmt.Println(last.data)
        last = last.next
    }

}

func main(){
    list := &Node{}
    list = list.append("6")
    list = list.append("7")
    list = list.append("8")
    list = list.append("9")
    list = list.append("10")
    list = list.append("11")
    list.display()
    list = list.append("100")
    list = list.append("110")
    list.display()
}
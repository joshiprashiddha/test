package main

import "fmt"

// Node is basic unit of singly linked list
type Node struct{
    data string
    next *Node
}

// insertion operation with pointer receiver
func (l *Node) append(newData string){
    last := l

    newNode := Node{
        data: newData,
        next: nil,
    }

    // linkedlist is empty
    if (Node{} == *l) {
        l.data = newData
        l.next = nil
    }else{
        for last.next != nil{
            last = last.next
        }
        last.next = &newNode
    }
}

func (l *Node) appendBeginning(newData string) *Node{
    newNode := Node{
        data: newData,
        next: nil,
    }

    if (Node{} == *l){
        l.data = newData
        l.next = nil
    }else{
        newNode.next = l
    }

    return &newNode
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
    list.append("6")
    list.append("7")
    list.append("8")
    list.append("9")
    list.append("10")
    list.append("11")
    list = list.appendBeginning("100")
    list.append("110")
    list.display()
}
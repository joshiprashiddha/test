package main

import (
        "fmt"
        "time"
        )

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// Node, Left, Right
func (tn *TreeNode) preOrderTravesal(preOrder *[]int){
    //Empty Tree
    if (tn == nil || TreeNode{} == *tn){
        *preOrder = append(*preOrder, -1)
        return
    }
    //Print Node Val
    *preOrder = append(*preOrder, tn.Val)
    //check if leaf node reached.
    if tn.Left != nil || tn.Right != nil{
        tn.Left.preOrderTravesal(preOrder)
        tn.Right.preOrderTravesal(preOrder)
    }
}

func (tn *TreeNode) Add(i int){
    //empty tree condition
    if (TreeNode{} == *tn){
        tn.Val = i
        return
    }

    if (i < tn.Val){
        if (tn.Left == nil){
            newTn := TreeNode{
                Val: i,
            }
            tn.Left = &newTn
        }else{
            tn.Left.Add(i)
        }
    } else{
        if (tn.Right == nil){
            newTn := TreeNode{
                Val: i,
            }
            tn.Right = &newTn
        }else{
            tn.Right.Add(i)
        }
    }
}

func main(){
    now := time.Now()
    bst := TreeNode{}

    bst.Add(1)
    bst.Add(2)
    bst.Add(3)

    var preOrder []int
    bst.preOrderTravesal(&preOrder)
    fmt.Println(preOrder)
    duration := time.Since(now)
    fmt.Println(duration)
}
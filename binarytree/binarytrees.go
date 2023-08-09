package main

import (
    
    "fmt"
    "os"
)

type Node struct{
    data int 
    left *Node
    right *Node
}
var N *Node
var Flag int=0
func (tre *Node)append(ele int){
    createnode:=&Node{data:ele,left:nil,right:nil}
    if N==nil{
      N=createnode
      return
      //current:=n
    }
    current:=N
    for current != nil{
        if ele < current.data{
            if current.left==nil{
                current.left=createnode
                return
            }
            current=current.left
            }else if ele>current.data{
                if current.right==nil{
                current.right = createnode
                return
                }
                current =current.right
            }else{
                return
            }

        
        }
    }
    func preordertraversal(temp *Node){
        if temp==nil{
            return
        }
        fmt.Print(temp.data," ")
        preordertraversal(temp.left)
        preordertraversal(temp.right)
    }
    func inordertraversal(temp *Node){
        if temp==nil{
            return
        }
        
        inordertraversal(temp.left)
        fmt.Print(temp.data," ")
        inordertraversal(temp.right)
    }
    func postordertraversal(temp *Node){
        if temp==nil{
            return
        }
        
        postordertraversal(temp.left)
        postordertraversal(temp.right)
        fmt.Print(temp.data," ")
    }
    func search(temp *Node,sear int){
        
        if temp==nil{
           
            return
        }
        if temp.data ==sear{
            fmt.Print("Element found")
            Flag=1
            return
        }
        search(temp.left,sear)
        search(temp.right,sear)
        
        //return Flag
    }



func main(){
    tr:=Node{}
    for true{
    fmt.Println("Enter your choice")
    fmt.Println("1. add element to tree ")
    fmt.Println("2. inorder taversal ")
    fmt.Println("3. Preorder traversal")
    fmt.Println("4. postorder traversal ")
    fmt.Println("5. search element")
    fmt.Println("6. Exit")
    var choice int
    fmt.Scanln(&choice)
    switch choice{
    case 1: fmt.Println("enter the element")
            var ele int
            fmt.Scanln(&ele)
            tr.append(ele)
    case 2: inordertraversal(N)
    case 3: preordertraversal(N)
    case 4: postordertraversal(N)
    case 5: fmt.Println("enter element to search")
            var sea int
            fmt.Scanln(&sea)
            Flag =0
             search(N,sea)
            if Flag ==0{
                fmt.Println("Element not found")
             }
    case 6: os.Exit(0)



    }
}
}
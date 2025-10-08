package goFor

import (
    "fmt"
)


func RunANormalForLoops(){
    for i :=0; i<7; i++ {
        fmt.Println(i)
    }
}

func ArrayIterativeForLoops(arr []string){
    for index, ele := range arr{
        fmt.Printf("%v : %v\n", index, ele)
    }
}
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
func ArrayIndexIterativeForLoops(arr []string){
    for index, _ := range arr{
        fmt.Printf("Index : %v\n", index)
    }
}
func ArrayValuesIterativeForLoops(arr []string){
    for _, ele := range arr{
        fmt.Printf("Values : %v\n", ele)
    }
}

func WhileLoopUsingFor(arr []string){
    i := len(arr)-1
    for i >= 0 {
        fmt.Println("The ele :", arr[i])
        i--;
    }
}

func TrueLoop(arr []string){
    i := 0
    for {
        fmt.Println(arr[i])
        i++;
        if i==len(arr) {
            break;
        }
    }
}
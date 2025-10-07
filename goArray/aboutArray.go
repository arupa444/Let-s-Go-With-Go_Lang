package goArray

import (
    "fmt"
)

func AcceptArray(){
    var storeArray = [5]int{1,2,4,5,7}
    fmt.Printf("This is a %T and it stores %v", storeArray, storeArray)
}
package goFunc

import (
    "fmt"
)

func NormalOneWithNoReturn(){
    fmt.Println("Let's start with functions...")
}

func NormalOneWithNoReturnButWithParameter(ele int){
    fmt.Printf("Starting no: %v\n",ele)
}

func WithReturn1ButWithParameter(ele int) int{
    return ele+1
}

func WithReturn2ButWithParameter(ele int, ele1 int) int{
    return ele+ele1
}

func WithBarReturn(ele int, ele1 int) (result int){
    result = ele+ele1
    return
}
func MultiValReturn(ele int, ele1 int) (result int, result1 int){
    result = ele + ele1 -2
    result1 = ele - ele1 +2
    return
}

func FactorialRecur(count int) int{
    if count == 1{
        return 1
    }
    return count + FactorialRecur(count - 1)
}
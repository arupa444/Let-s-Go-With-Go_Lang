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
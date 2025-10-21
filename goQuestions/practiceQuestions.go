package goQuestions

import (
    "fmt"
    "math"
)

// Write a program that prints numbers from 1 to 100, but for multiples of 3 print "Fizz",
// for multiples of 5 print "Buzz", and for multiples of both print "FizzBuzz".


func GoFirstQue(){
    for i := 1; i<=100; i++{
        if i%3 == 0 && i%5 == 0{
            fmt.Println("FizzBuzz")
        }else if i%3 == 0{
            fmt.Println("Fizz")
        }else if i%4 == 0{
            fmt.Println("Buzz")
        }else{
            fmt.Println(i)
        }
    }
}

// Write a function that takes a string and returns it reversed.

func GoSecondQue(str string)string{
    storeStr := []rune(str)
//     storeStr1 := []byte(str)
//     fmt.Println(storeStr1)
    for i, j := 0, len(storeStr)-1; i<j; i, j = i+1, j-1{
        storeStr[i], storeStr[j] = storeStr[j], storeStr[i]
    }
    return string(storeStr)
}


// Write a program that checks if a given number is prime or not.

func GoThirdQue(ele int)string{
    ele1 := math.Sqrt(float64(ele))
    for i:=2; i<int(ele1); i++{
        if ele%i == 0{
            return "Not a Prime No."
        }
    }
    return "Prime No."
}
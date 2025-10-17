package goQuestions

// import (
//     "fmt"
// )

// Write a program that prints numbers from 1 to 100, but for multiples of 3 print "Fizz",
// for multiples of 5 print "Buzz", and for multiples of both print "FizzBuzz".


func GoFirstQue() string{
    for i := 1; i<=100; i++{
        if i%3 == 0 && i%5 == 0{
            return "FizzBuzz"
        }else if i%3 == 0{
            return "Fizz"
        }else if i%4 == 0{
            return "Buzz"
        }else{
            return string(i)
        }
    }
}
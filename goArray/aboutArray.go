package goArray

import (
    "fmt"
)

func AcceptArray(){
// defined length
    var storeArray = [5]int{1,2,4,5,7}
// inferred length
    storeInferredArray := [...]int{2,32,3423,5,34,645,7,657,567}
    storeStringInferredArray := [...]string{"arupa", "deni", "ashutosh", "subham", "i don't know"}
    fmt.Printf("This is a %T and it stores %v\n", storeArray, storeArray)
    fmt.Printf("This is a %T and it stores %v\n", storeInferredArray, storeInferredArray)
    fmt.Printf("This is a %T and it stores %v\n", storeStringInferredArray, storeStringInferredArray)
    fmt.Printf("the length is %v\n", len(storeStringInferredArray))
    newVar := storeStringInferredArray[2:4]
    fmt.Printf("the Capacity is %v\n", newVar)
    fmt.Printf("the Capacity is %v\n", cap(newVar))
    fmt.Printf("the length is %v\n", len(newVar))
    fmt.Printf("%v\n", storeStringInferredArray[0:3])
    storeStringInferredArray[0] = "Arupa"
    fmt.Printf("updated list : %v\n", storeStringInferredArray)
}
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
    newVar := storeStringInferredArray[1:2]
    fmt.Printf("the array is %v\n", newVar)
    fmt.Printf("the Capacity is %v\n", cap(newVar))
    fmt.Printf("the length is %v\n", len(newVar))
    fmt.Printf("%v\n", storeStringInferredArray[0:3])
    storeStringInferredArray[0] = "Arupa"
    fmt.Printf("updated list : %v\n\n\n\n\n\n\n", storeStringInferredArray)

//     create a null array
    fmt.Println("Check how to make null array... ")
    buildAnNUllArray := make([]int,5,8)
    fmt.Println(buildAnNUllArray)
    fmt.Println("the length of the null array : ",len(buildAnNUllArray))
    fmt.Println("The capacity of the null array : ",cap(buildAnNUllArray))

    buildAnNUllArray = make([]int,3)
    fmt.Println(buildAnNUllArray)
    fmt.Println("the length of the null array : ",len(buildAnNUllArray))
    fmt.Println("The capacity of the null array : ",cap(buildAnNUllArray))
    buildAnNUllArray = append(buildAnNUllArray, 23)
    fmt.Println("the array after append values in array : ",buildAnNUllArray)
    fmt.Println("the length of the null array : ",len(buildAnNUllArray))
    fmt.Println("The capacity of the null array : ",cap(buildAnNUllArray))


//     memory efficient
    effVar := make([]int,6, 10)
    fmt.Println("The array : ",effVar)
    fmt.Println("The Capacity : ",cap(effVar))
    fmt.Println("The Length ; ",len(effVar))
    newEffVar := effVar[0:4]
    fmt.Println("The array : ",newEffVar)
    fmt.Println("The Capacity : ",cap(newEffVar))
    fmt.Println("The Length ; ",len(newEffVar))

    aBigArray := make([]int, 15)
    fmt.Println("The array : ",aBigArray)
    fmt.Println("The Capacity : ",cap(aBigArray))
    fmt.Println("The Length ; ",len(aBigArray))

    slicedBigArray := aBigArray[:len(aBigArray)-5]

    fmt.Println("The array : ",slicedBigArray)
    fmt.Println("The Capacity : ",cap(slicedBigArray))
    fmt.Println("The Length ; ",len(slicedBigArray))

    copyArrayForMemoryManagement := make([]int,len(slicedBigArray))

    copy(copyArrayForMemoryManagement, slicedBigArray)

    fmt.Println("The array : ",copyArrayForMemoryManagement)
    fmt.Println("The Capacity : ",cap(copyArrayForMemoryManagement))
    fmt.Println("The Length ; ",len(copyArrayForMemoryManagement))

    slicedBigArray = nil

    fmt.Println("The array : ",slicedBigArray)
    fmt.Println("The Capacity : ",cap(slicedBigArray))
    fmt.Println("The Length ; ",len(slicedBigArray))
}
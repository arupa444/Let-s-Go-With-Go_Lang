package goString

import(
    "fmt"
)
// outsideTry := "ouuuutttt" // can only be assigned inside the functions....
var globeVar string = "Bahar hu"

const (
    PI = "3.14"
)

func ChString(){
    var storeString = "Arupa Nanda Swain"
    fmt.Println(storeString)
    fmt.Println("the value of pi is "+PI)
    fmt.Println(globeVar+"he he he \n")
}

package main
import (
    "fmt"
    "startGoLLM/inventoryDetails"
    "startGoLLM/goIfElse"
    "startGoLLM/goString"
//     "bufio"
//     "os"
)


func main(){
    fmt.Println("This is the inventory......")

//     var age int
//     var age1 int;
//
//     fmt.Scan(&age) // simplest way to get input
//
//     fmt.Printf("hey, %d\n",age)

    inventoryDetails.PrintDetails() // package calling
//
//     reader := bufio.NewReader(os.Stdin) // another way to get input
//
//     name, _ := reader.ReadString('\t')
//
//     fmt.Println("%s", name)
    var storeInt int = 23
    var storeStr string = "Arupa Nanda Swain"
    var storeBool bool = true
    var storeFloat float64 = 69.6969696969

    fmt.Printf("%q %v %.2v %v\n", storeStr, storeBool, storeFloat, storeInt)

    i, j := 0.21, "Arupa"

    fmt.Printf("%T: %v, %T: %v%v", i, i, j, j, "\n")
    storeMessage := ""

    storeMessage = fmt.Sprintf("The name is %v and the age is %v",storeStr, storeInt)


    fmt.Println(storeMessage)

    goIfElse.IfFunction()

}
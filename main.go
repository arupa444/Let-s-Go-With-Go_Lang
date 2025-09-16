package main
import (
    "fmt"
    "startGoLLM/inventoryDetails"
    "bufio"
    "os"
)


func main(){
    fmt.Println("This is the inventory......")

//     var age int
//
//     fmt.Scan(&age) // simplest way to get input
//
//     fmt.Printf("hey, %d\n",age)

    inventoryDetails.PrintDetails() // package calling

    reader := bufio.NewReader(os.Stdin) // another way to get input

    name, _ := reader.ReadString('\t')

    fmt.Println("%s", name)

}
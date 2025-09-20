package main
import (
    "fmt"
    "startGoLLM/inventoryDetails"
//     "bufio"
//     "os"
)
func mod(x int, y int) int {
    return x%y
}

func main(){
    fmt.Println("This is the inventory......")

//     var age int
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


    allowGuy := "Arupa"

    if allowGuy == "Arupa"{
        fmt.Printf("Hey, %v\n", allowGuy)
    }

    if allowGuy := "Ashutosh"; allowGuy == "Ashutosh"{
        fmt.Println("Hi, i am",allowGuy,"\nhow was your day?",storeMessage)
    }
    fmt.Printf("%v",allowGuy)


    x, y := 20,6
    fmt.Printf("The mod of x(%v) -> y(%v) is %v",x, y, mod(x,y))



    if check := 43; check != 20{
        fmt.Printf("\n%v is a right age...\n",check)
    }

}
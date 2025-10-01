package goIfElse
import(
    "fmt"
)

func IfFunction(){
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

    if check := 1; check > 0{
        fmt.Printf("\n%v is a right age...\n",check)
    }

    var ifString string = "Arupa"
    if ifString == "Arupa"{
        fmt.Printf("Hey, this is %v",ifString);
    }else if ifString == "Ashutosh"{
        fmt.Printf("Hey, this isn't Arupa!! this is me(%v)",ifString)
    }else{
        fmt.Printf("Danish bhaiya ko bhul gaye kya....")
    }

    // getLength variable bro
}
package main
import (
    "fmt"
    "startGoLLM/inventoryDetails"
    "startGoLLM/goIfElse"
    "startGoLLM/goString"
    "startGoLLM/goArray"
    "startGoLLM/goSwitch"
    "startGoLLM/goFor"
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

//     Multi variable declaration
    var a, b, c, d string = "a","b","c","d"
    var aa, bb = 3, "aabb"
    cc, dd := "Arupa", false

    var(
        aaa int = 46
        bbb string = "fortnight"
    )


    fmt.Printf("\n\n\n\n%v %v %v %v\n\n\n\n",a,b,c,d)
    fmt.Printf("\n\n\n\n%v %v %v %v\n\n\n\n",aa,bb,cc,dd)
    fmt.Printf("\n\n\n\n%v %v\n\n\n\n",aaa,bbb)

    var bitManipulation int = 15

    fmt.Printf("%b\n", bitManipulation)
    fmt.Printf("%d\n", bitManipulation)
    fmt.Printf("%+d\n", bitManipulation)
    fmt.Printf("%o\n", bitManipulation)
    fmt.Printf("%O\n", bitManipulation)
    fmt.Printf("%x\n", bitManipulation)
    fmt.Printf("%X\n", bitManipulation)
    fmt.Printf("%#x\n", bitManipulation)
    fmt.Printf("%4d\n", bitManipulation)
    fmt.Printf("%-4d\n", bitManipulation)
    fmt.Printf("%04d\n", bitManipulation)
    
    var floatVar = 3.141
    
    fmt.Printf("%e\n", floatVar)
    fmt.Printf("%f\n", floatVar)
    fmt.Printf("%.2f\n", floatVar)
    fmt.Printf("%50.2f\n", floatVar)
    fmt.Printf("%g\n", floatVar)

//     print using all type of print options
    const PrintExample string = "hui hui\nguys hui hui"

    fmt.Println(PrintExample)
    fmt.Printf("%v\n",PrintExample)
    fmt.Print(PrintExample,"\n")

    fmt.Printf("%v%%\n", PrintExample)
    fmt.Printf("%#v\n", PrintExample)

    fmt.Println("Boolean:\t", true)
    fmt.Println("Integer:\t", 23)
    fmt.Println("Float:\t", 3.67)
    fmt.Println("String:\t", "yyuuoo")


    goIfElse.IfFunction()
    goString.ChString()
    goArray.AcceptArray()
    goSwitch.Tier1Switch(2)
    goSwitch.MultiTierSwitch(51)
    goFor.RunANormalForLoops()
    storeStringInferredArray := []string{"arupa", "deni", "ashutosh", "Subham", "i don't know"}
    goFor.ArrayIterativeForLoops(storeStringInferredArray)
    goFor.ArrayIndexIterativeForLoops(storeStringInferredArray)
    goFor.ArrayValuesIterativeForLoops(storeStringInferredArray)
    goFor.WhileLoopUsingFor(storeStringInferredArray)
    goFor.TrueLoop(storeStringInferredArray)
    goFor.WhileLoopUsingFor(storeStringInferredArray)
}
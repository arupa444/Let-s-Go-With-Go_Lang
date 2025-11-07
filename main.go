package main
import (
    "fmt"
    "startGoLLM/inventoryDetails"
    "startGoLLM/goIfElse"
    "startGoLLM/goString"
    "startGoLLM/goArray"
    "startGoLLM/goSwitch"
    "startGoLLM/goFor"
    "startGoLLM/goFunc"
    "startGoLLM/goStruct"
    "startGoLLM/goMap"
    "startGoLLM/goQuestions"
    "startGoLLM/goDirectory"
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
    storeStringInferredArray := []string{"arupa", "pinki", "jersey", "megha", "deni", "ashutosh", "Subham", "i don't know"}
    goFor.ArrayIterativeForLoops(storeStringInferredArray)
    goFor.ArrayIndexIterativeForLoops(storeStringInferredArray)
    goFor.ArrayValuesIterativeForLoops(storeStringInferredArray)
    goFor.WhileLoopUsingFor(storeStringInferredArray)
    goFor.TrueLoop(storeStringInferredArray)
    goFor.WhileLoopUsingFor(storeStringInferredArray)

    funcVar := 29
    goFunc.NormalOneWithNoReturn()
    goFunc.NormalOneWithNoReturnButWithParameter(funcVar)
    fmt.Print(goFunc.WithReturn1ButWithParameter(funcVar),"\n")
    fmt.Print(goFunc.WithReturn2ButWithParameter(funcVar, funcVar),"\n")
    fmt.Print(goFunc.WithBarReturn(funcVar, funcVar),"\n")
    ele, ele1 := goFunc.MultiValReturn(funcVar, funcVar)
    fmt.Println("Multi return values :",ele,ele1)
    fmt.Println(goFunc.MultiValReturn(funcVar, funcVar))
    fmt.Println(goFunc.FactorialRecur(5))
    fmt.Println(goStruct.StructIt("Arupa", "Ai Developer", 21, 417000))
    storeStruct := goStruct.StructItTwo("Ashutosh", "Ai Developer", 22, 417000)
    fmt.Println(storeStruct)
    goStruct.PrintStruct(storeStruct)
    goMap.CreateStaticMap()
    fmt.Println(goMap.CreateDynamicBMWMap())
    fmt.Println(goMap.CreateDynamicRRMap())
    fmt.Println(goMap.CreateDynamicTATAMap())

//     Go Practice Question calling/invoking
    goQuestions.GoFirstQue()
    fmt.Println(goQuestions.GoSecondQue("his❤️"))
    fmt.Println(goQuestions.GoThirdQue(101))
//     goQuestions.GoThirdQue(21)
    tempArray := []int{1,2,34,32,973596,52,35,235,235}
    fmt.Println(goQuestions.GoForthQue(tempArray))
    tempArray1 := []float64{13.34,4353.564,4564.234,3464.5645,33.43,346543.34}
    fmt.Println(goQuestions.GoForthQue(tempArray1))
    tempStr := "Arupa nanda swain"
    fmt.Println(goQuestions.GoFifthQue1(tempStr))
    fmt.Println(goQuestions.GoFifthQue2(tempStr))
    tempArray1 = []float64{112,323,423,4,3,43,43,43,23,0,321,2,1,21,1,21,12}
    fmt.Println(goQuestions.GoSixthQue(tempArray1))

    tempSortedArray1 := []int{1,23,54,345,67567,5867868}
    tempSortedArray2 := []int{1,2,3,3,3,4,23,44,55,643,54656,3453336}

    fmt.Println(goQuestions.GoSevenQue(tempSortedArray1, tempSortedArray2))

    tempArray71 := []int{1,34,53,45,46,456,75,67,568765,823,54,345,67567,5867868}
    tempArray72 := []int{243534,645,7658,61,24,456,56,3,3,3,4,23,44,55,643,54656,45645645}
    fmt.Println(goQuestions.GoSevenQue1(tempArray71, tempArray72))
    fmt.Println(tempArray71)
    fmt.Println(goQuestions.GoEightQue(tempArray71, 0))
    fmt.Println(goQuestions.GoNinthQue(tempArray72))
    fmt.Println(goQuestions.GoNinthQue1(tempArray72))
    fmt.Println(goQuestions.GoTenthQue(tempArray72, tempArray72))
    fmt.Println(goQuestions.GoElevenQue("HI arupa how are you you good??"))
    fmt.Println(goQuestions.GoTwelveQue("Bad Touch and Good Touch", []string{"Arupa", "Ashutosh", "Daniel"}, 6969))
    storeForFunc12 := goQuestions.GoTwelveQue("Bad Touch and Good Touch", []string{"Arupa", "Ashutosh", "Daniel"}, 6969)
    goQuestions.CallBook(storeForFunc12)
    fmt.Println(goQuestions.GoThirteenQue("Arupa Nanda Swain"))
//     fmt.Println(goQuestions.GoFourteenQue("Arupa"))
    fmt.Println(goQuestions.GoFifteenQue(storeStringInferredArray))
    fmt.Println(goQuestions.GoSixteenQue(7))
    fmt.Println(goQuestions.GoLinearSearch(tempArray71, 99))
    fmt.Println(goQuestions.GoBinarySearch(tempArray71, 67567))
    fmt.Println(goQuestions.GoPointers(75))
    var daniel int = 75
    var arupa int = 69
    goQuestions.GoPointers1(&daniel, &arupa)
    fmt.Println(daniel, arupa)
    fmt.Println(goQuestions.GoVariadicNo(12,34,23,43,54,6457))



//     About Directory manipulation
    goDirectory.CreateDir("Hella Deni")
    goDirectory.CreateNestedDir("app","model","images")
    goDirectory.RemoveDir("Hella Deni")
}
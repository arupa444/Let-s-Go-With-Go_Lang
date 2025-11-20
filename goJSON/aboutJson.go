package goJSON

import(
    "fmt"
    "os"
)


type SavingsAccount struct{
    Name :
}

func FetchJsonData(){
    jsonFile, err := os.Open("try.json")
    if err != nil{
        fmt.Errorf("The error : ", err)
    }
    fmt.Println("Successfully opened the .json file")
    defer jsonFile.Close()


}
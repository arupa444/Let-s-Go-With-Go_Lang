package goJSON

import(
    "fmt"
    "os"
)

type AccountDetailsStruct struct{
    Balance float64
    Currency string
}

type SavingsAccount struct{
    Name string
    AccountDetails AccountDetailsStruct
    Age int
}

func FetchJsonData(){
//     var saveAccount SavingsAccount
    jsonFile, err := os.Open("try.json")
    if err != nil{
        fmt.Errorf("The error : ", err)
    }
    fmt.Println("Successfully opened the .json file")
    defer jsonFile.Close()


}
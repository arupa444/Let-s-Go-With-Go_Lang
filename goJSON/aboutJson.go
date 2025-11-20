package goJSON

import(
    "fmt"
    "os"
)

type AccountDetails struct{
    Balance             float64
    Currency            string
}

type SavingsAccount struct{
    Name                string
    AccountDetails      AccountDetails
    Age                 int
}

type Accounts struct{
    ID                  string
    SavingsAccount      SavingsAccount
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
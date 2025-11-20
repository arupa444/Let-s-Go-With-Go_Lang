package goJSON

import(
    "fmt"
    "os"
)

type AccountDetails struct{
    Balance             float64             `json:"balance"`
    Currency            string              `json:"currency"`
}

type SavingsAccount struct{
    Name                string              `json:"name"`
    AccountDetails      AccountDetails      `json:"accountDetails"`
    Age                 int                 `json:"age"`
}

type Accounts struct{
    ID                  string              `json:"id"`
    SavingsAccount      SavingsAccount      `json:"savingsAccount"`
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
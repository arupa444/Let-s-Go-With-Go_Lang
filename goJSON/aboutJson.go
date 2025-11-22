package goJSON

import(
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
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
    var accounts Accounts
    jsonFile, err := os.Open("goJSON/try.json")
    if err != nil{
        fmt.Errorf("The error : ", err)
    }
    fmt.Println("Successfully opened the .json file")
    defer jsonFile.Close()

    byteValues, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal(byteValues, &accounts)
    fmt.Println(accounts)
}



package goJSON

import(
    "fmt"
    "os"
)

func FetchJsonData(){
    jsonFile, err := os.Open("aboutJson.go")
    fmt.Println(jsonFile)
}
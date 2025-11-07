package goDirectory

import (
    "fmt"
    "os"
)

func CreateDir(dirName string){
    err := os.Mkdir(dirName, 0755)
    if err != nil{
        fmt.Println("Error while directory creation: ", err)
        return
    }
    fmt.Println("Successfully created a Directory....")
}
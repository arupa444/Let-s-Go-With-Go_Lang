package goDirectory

import (
    "fmt"
    "os"
)

func createDir(name str){
    err := os.Mkdir(name, 0755)
    if err != nil{
        fmt.Println("Error while directory creation: ", err)
        return
    }
    fmt.Println("Successfully created a Directory....")
}
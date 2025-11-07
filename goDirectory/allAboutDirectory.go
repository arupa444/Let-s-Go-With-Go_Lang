package goDirectory

import (
    "fmt"
    "os"
    "strings"
)

func CreateDir(dirName string){
    err := os.Mkdir(dirName, 0755)
    if err != nil{
        fmt.Println("Error while directory creation: ", err)
        return
    }
    fmt.Println("Successfully created a Directory....")
}

func CreateNestedDir(nums ...string){
    err := os.MkdirAll(strings.Join(nums, "/"), 0755)
    if err != nil{
        fmt.Println("Error while creating a nested dir: ", err)
        return
    }
    fmt.Println("successfully created a directory")
}

func RemoveDir(dirName string){
    err := os.Remove(dirName)
    if err != nil{
        fmt.Println("Error while directory deletion: ", err)
        return
    }
    fmt.Println("successfully deleted a directory")
}
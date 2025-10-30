package goSwitch

import (
    "fmt"
)
// its about witch case in GoLang
func Tier1Switch(ele int){
    switch ele {
        case '+':
            fmt.Println(ele,"reeeeeeeeee")
        case 2:
            fmt.Println(ele,"reeeeeeeeee")
        case 3:
            fmt.Println(ele,"reeeeeeeeee")
        case 4:
            fmt.Println(ele,"reeeeeeeeee")
        default:
            fmt.Println("there is no element name :",ele)
    }
}
func MultiTierSwitch(ele int){
    switch ele {
        case 1, 11:
            fmt.Println(ele,"reeeeeeeeee")
        case 2, 21:
            fmt.Println(ele,"reeeeeeeeee")
        case 3, 31:
            fmt.Println(ele,"reeeeeeeeee")
        case 4, 41:
            fmt.Println(ele,"reeeeeeeeee")
        default:
            fmt.Println("there is no element name :",ele)
    }
}
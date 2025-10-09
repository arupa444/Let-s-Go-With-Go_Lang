package goStruct
import (
    "fmt"
)

type employee struct{
    name string
    domain string
    age int
    salary int
}

func StructIt(name string, domain string, age string, salary string)(emp1 struct){
    var emp1 employee
    emp1.name = name
    emp1.domain = domain
    emp1.salary = salary
    emp1.age = age

    return
}
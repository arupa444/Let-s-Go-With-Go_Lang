package goStruct

import (
    "fmt"
)

type Employee struct{
    name string
    domain string
    age int
    salary int
}

func StructIt(name string, domain string, age int, salary int) Employee{
    var emp1 Employee
    emp1.name = name
    emp1.domain = domain
    emp1.salary = salary
    emp1.age = age

    return emp1
}


func StructItTwo(name string, domain string, age int, salary int) (emp1 Employee){
    emp1.name = name
    emp1.domain = domain
    emp1.salary = salary
    emp1.age = age

    return
}

func PrintStruct(emp1 Employee){
    var storeKeys = []string{"name", "domain", "age", "salary"}
    fmt.Println("The employee documents that we need : ", storeKeys, emp1)
    fmt.Println(storeKeys[0],":",emp1.name)
    fmt.Println(storeKeys[1],":",emp1.domain)
    fmt.Println(storeKeys[2],":",emp1.age)
    fmt.Println(storeKeys[3],":",emp1.salary)
}
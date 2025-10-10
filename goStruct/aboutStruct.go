package goStruct


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
    var storeKeys := make([]string, "name", "domain", "age", "salary")
}
package goStruct


type employee struct{
    name string
    domain string
    age int
    salary int
}

func StructIt(name string, domain string, age int, salary int) employee{
    var emp1 employee
    emp1.name = name
    emp1.domain = domain
    emp1.salary = salary
    emp1.age = age

    return emp1
}


func StructItTwo(name string, domain string, age int, salary int) (emp1 employee){
    emp1.name = name
    emp1.domain = domain
    emp1.salary = salary
    emp1.age = age

    return
}
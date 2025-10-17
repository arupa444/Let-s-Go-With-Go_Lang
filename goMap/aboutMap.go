package goMap

import (
    "fmt"
)

func CreateStaticMap(){
    var storeCarPrices = map[string]int{"BMW":50000000, "Lambo": 80000000, "Suzuki": 3000000, "Toyota": 20000000, "RR": 2500000000}
    storeTeam := map[string]string{"Arupa": "AI", "Daniel": "AI"}
    fmt.Printf("%#v\n",storeCarPrices)
    fmt.Printf("%v\n",storeCarPrices)
    fmt.Printf("%#v\n",storeTeam)
    fmt.Printf("%v\n",storeTeam)
}

func CreateDynamicMap()map[string]string{
    var storeCarPrices = make(map[string]string)
    storeCarPrices["name"] = "BMW"
    storeCarPrices["price"] = "90000000"
    storeCarPrices["location"] = "Berhampur"
    return storeCarPrices
}
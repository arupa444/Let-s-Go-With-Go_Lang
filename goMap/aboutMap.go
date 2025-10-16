package goMap

import (
    "fmt"
)

func CreateMap(){
    var storeCarPrices = map[string]int{"BMW":50000000, "Lambo": 80000000, "Suzuki": 3000000, "Toyota": 20000000, "RR": 2500000000}
    storeTeam := map[string]string{"Arupa": "AI", "Daniel": "AI"}
    fmt.Printf("%#v\n",storeCarPrices)
    fmt.Printf("%v\n",storeCarPrices)
    fmt.Printf("%#v\n",storeTeam)
    fmt.Printf("%v\n",storeTeam)
}
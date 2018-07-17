package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}


func main() {
    //var myMap map[string]Vertex
    myMap := make(map[string]Vertex)
    myMap["Umass Amherst"] = Vertex{1.23, 2.34}
    fmt.Println(myMap["Umass Amherst"])
    var m = map[string]int {
        "Amherst College" : 1,
        "Umass Amherst" : 2,
    }
    fmt.Println(m)
}

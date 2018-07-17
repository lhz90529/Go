package main

import (
        "fmt"
       )

type Vertex struct {
    X, Y float64
}

func myFunc(v Vertex, f float64){
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v Vertex) m1(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v* Vertex) m2(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    ptr := &v
    myFunc(v, 2)
    fmt.Println(v)

    //m1 is a method with normal receiver argument 
    v.m1(2)
    fmt.Println(v)
    ptr.m1(2)
    fmt.Println(v)
    //m2 is a method with pointer receiver argument
    v.m2(2)
    fmt.Println(v)
    ptr.m2(2)
    fmt.Println(v)

}

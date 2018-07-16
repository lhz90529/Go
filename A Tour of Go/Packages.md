#### Packages
- Every Go program is made up of `packages`
- Program entry: `package main`
    ```go
    package main//example program

    import (//import with the open-closing parenthesis
        "fmt"
        "math/rand"
    )

    func main() {
        fmt.Println("My favorate number is ", rand.Intn(10))//note: no terminate symbol ";"
    }
    ```
#### Imports
- This code groups the imports into a parenthesized, "factored" import statement.
    ```go
    import (//factored import statement
        "a"//import a
        "b"//import b
        .
        .
        .
        "and so on"//...
    )
    ```
    ```go
    //multiple import statements
    import "a"
    import "b"
    ```
- ***It is good style to use the factored import statement.***

#### Exported names
- In Go, a name is exported **if it begins with a capital letter**
    - e.g. `Pi` from `math`
- You have **access only** to `exported names`
    - which means any `unexported names` **are not accessible** from outside the package

    ```go 
    package main

    import (
        "fmt"
        "math"
    )

    func main() {
        fmt.Println(math.pi)
        //If you compile this code, you will get the error message: "cannot refer to unexported name math.pi"
        //since pi is not an exported name from the package math
    }
    ```
#### Functions
- A function can take `zero` or `more arguments`
- `data type` is declared **after** `variable name`
    - ***very different from the programming language I have learned sofar***

    ```go
    package main

    import "fmt"

    func add(x int, y int) int{//Note where do you put the data type of return value
        return x + y
    }

    func main() {
        fmt.Println(add(2, 4))
    }
    ```

- When 2 or more arguments **share a same type**, you can **omit** the type from all but the last
    ```go
    func add(x int, y int) int {
        return x + y
    }
    //Can be written as 
    func add (x, y int) int {//omit the type from all but the last
        return x + y
    }
    ```

#### Multiple results
> Surprising feature!!!
- A `function` can return any number of results
    ```go 
    package main

    import "fmt"

    func swap (a, b string) (string, string) {//note how should we write the data types of multiple return values
        return b, a;//return multiple result
    }

    func main() {
        fmt.Println(swap("hello", "world"))
    }
    ```

#### Named return value
- `return value` can be **named**
- A `return` statement without arguments will return all the `named variables`
    - This is known as **`naked return`**
    ```go
    func split(c int) (a, b int) {//named return value
        a = c * 4 / 9
        b = c - a
        return//naked return
    }
    ```
#### Variables
The `var` statement can be used to declare one or more varaibles.
- **Basic Syntax**
    ```go
    var v T//declare a variable called v with a data type of T
    var v1, v2, ... vn T//declare a list of variables v1, v2, ... vn with data type of T
    ```
- **Variables with initializers**
    - A `var` declaration can include `initializers`, one per variable
    ```go
    package main

    import "fmt"

    func main() {
        var a, b, c = true, 4, "no"
        fmt.Println(a, b, c)
    }
    ```
- **Short variable declarations**
    - the `:=` short assignment can be used in place of a `var` declaration with ***implicit type***
    - the `:=` short assignment **cannot be used outside of function**
    ```go
    package main

    import "fmt"

    func main() {
        var i, j int = 1, 5
        k := 3//short assignment for one variable
        x, y, z := true, false, "no"//short assignment for multiple varaibles
        //Note: the data type can be imferred
        fmt.Println(i, j, k, x, y, z)
    }
    ```
- **factored**
    - `variable declarations` can be **factored** into blocks, as same as `import statements`
        ```go
        var (//factored variable declarations using open-closing parenthesis
            a bool = true
            b int = 3
            c string = "No"
        )
        ```

#### Basic types
|    Type    |     Size (bits)    |           Range          | Alias |                                      Demonstration                                      |
|:----------:|:------------------:|:------------------------:|:-----:|:---------------------------------------------------------------------------------------:|
|    bool    |   8, i.e. 1 byte   |        true, false       |       |                                                                                         |
|   string   |                    |                          |       |                                   a sequence of bytes                                   |
|     int    | Platform dependent |    Platform dependent    |       |                     32-bits system: 32 bits 64-bits system: 64 bits                     |
|    int8    |          8         |        -128 - 127        |       |                                                                                         |
|    int16   |         16         |      -32768 - 32767      |       |                                                                                         |
|    int32   |         32         | -2147483648 - 2147483647 |       |                                                                                         |
|    int64   |         64         |                          |       |                                                                                         |
|    uint    | Platfrom dependent |    Platfrom dependent    |       |                     32-bits system: 32 bits 64-bits system: 64 bits                     |
|    uint8   |          8         |          0 - 255         |       |                                                                                         |
|   uint16   |         16         |         0 - 65535        |       |                                                                                         |
|   uint32   |         32         |                          |       |                                                                                         |
|   uint64   |         64         |                          |       |                                                                                         |
|    byte    |                    |                          | uint8 |                                represent ASCII characters                               |
|    rune    |                    |                          | int32 |                  represent UTF-8 characters default type for characters                 |
|   float32  |         32         |                          |       |                                     single-precision                                    |
|   float64  |         64         |                          |       |                double-precision  default type for floating number in Go                 |
|  complex64 |         64         |                          |       |                    both real and imaginary parts are of float32 type                    |
| complex128 |         128        |                          |       | both real and imaginary parts are of float64 type default type for complex number in Go |

#### Zero values
variable declaration without initialization will be assigned with `zero value`
- `0` for numeric value
- `false` for boolean value
- `"" (empty string)` for string

#### Type conversions
`T(v)` will convert the varabile `v` to the type `T`
```go
var x int = 13//declare a interger called x with value of 13
var y float64 = float64(x)//convert a integer to the data type of float
var z uint = uint(y)//convert a double precision floating number to the data type of unsigned integer
```
- Assignment between items of different type **requires a explicit conversion**

#### Type inference
When declaring a variable without specifying an explicit type, the variable's type will be **inferred from the value on the right hand side**


#### Constants
- declared with keyword `const`
- cannot be declared using Short variable declarations `const`
    ```go
    const Pi = 3.14
    ```
#### Numeric Constants
- high-precision 


#### for
- Go has **only one loop construct:** `for`
    - **inital statement:** executed before the start of the loop
    - **condition expression:** evaluated before every iteration
    - **post statement:** executed at the end of each iteration
- Syntax
    ```go
    for initial statement; condition; post statement {//no parenthesis "()"is need to enclose the 3 componenets 
        //do something
    }
    ```
- `initial and post statement` are **optional**
    ```go
    for ; condition; {//note: the omitted initial and post statement
        //do something
    }
    ```
- You can even drop the semicolons `;`
    ```go
    for condition {//This is what exactyly "while loop" in C
        //do something
    }
- Forever
    ```go
    for {//loops forever
        //do something
    }
    ```
#### if
- no parenthesis `()` is needed
    ```go
    if condition {

    }
    ```
- `if` with a short statement
    - Like `for`, the `if` statement can start with a short statement to execute short statement before the condition
    - `variables` declared in the statement are only visible to the scope of `if`
        ```go
        if inital statement; condition {
            //do something    
        }
        ```

#### if and else
- `variables` declared inside the short statement are also visible to the section `else`
    ```go
    if initial statement; condition {
        //do somehitng
    } else {
        //do something else
    }
    ```
#### siwtch
- **Basic syntax** 
```go
    switch intial statement; v {
        case a:
            //do something
        case b:
            //do something
        .
        .
        .
        default:
            //do something
    }
```

- **Evaluation order**
    - `switch` evaluates `cases` from top to bottom 
    - stop when a `case` succeeds

- **Switch with no condition**
    - a clear way to write long `if-else` statement

#### Defer
A defer statement defers the execution of a function until the surrounding function returns.
```go
package main 

import "fmt"

func main() {
    defer fmt.Println("world")//This line will be executed until the function main() return
    fmt.Println("hello")
}

/*  
Output:
    hello
    world
*/
```
- **Stacking defer**
    - `deferred function calls` are pushed onto stack
    - When the `surrouding function call` returns, the `deferred function calls` will be poped out of stack and executed

#### Pointers
- Hold the `memory address` of a value
- `*T` is a pointer that point to a value of type `T`
    - zero value of pointer: `nil`
    ```go
    var ptr *T//a pointer named "ptr" that points to a value of type T
    ```
- `&`: get the address of a variable
- `*`: dereferencing a value from a pointer
- **Go has no pointer arithmetic**


#### Struct
A `struct` is a collection of `fields`
- **Basic syntax**
    ```go
    type name struct {
        //field1
        //field2
        .
        .
        .
        //an so on
    }
    ```
- `field` is accessed using dot `.`

- **Struct Pointer**
    - `field` can be access through a struct pointer
    ```go
    package main

    import "fmt"

    type Vertex struct {
        X int
        Y int
    }

    func main() {
        v := Vertex{1, 2}//construct a Vertex
        p := &v//a pointer point to the Vertex
        
        /*
        Note here: we can access field through struct pointer without dereferencing first
        In C++ we achieve this by doing: 
        p->X = 1e9
        */
        p.X = 1e9

        fmt.Println(v)
    }

#### Struct Literals
- TODO 

#### Arrays
- `[n]T` is an array of `n` values of type `T`
- **Array declaration:**
    ```go
    var name[n]T//create an array of n elements of type T
    ```
- **Array declaration and initialization**
    ```go
    array := [size]{ele1, ele2, ... elen}//declaration and initialization
- **Example**
    ```go
    package main

    import "fmt"

    func main() {
        var a[2]string
        b := [4]int {0,1,2,3}
        a[0] = "hello"
        a[1] = "world"
        fmt.Println(a)
        fmt.Println(b)
    }
    ```

#### Slices
- dynamically-sized
- a[low : high] is a range of `[low, high)`
    ```go
    package main

    import "fmt"

    func main() {
        b := [4]int {0,1,2,3}
        c := b[0:3]
    }
    ```

- **Slices are references to arrays**
    - Change of elements of a slice modifies the corresponding elements of its underlying array
        ```go
        import "fmt"

        func main() {
            names := [4]string{
                "John",
                "Paul",
                "George",
                "Ringo",
            }
            fmt.Println(names)

            a := names[0:2]//a is [John, Paul]
            b := names[1:3]//b is [Paul, George]
            fmt.Println(a, b)

            //Note the reference property of slice
            b[0] = "XXX"//Paul ---> XXX
            fmt.Println(a, b)//[John, XXX, XXX, George]
            fmt.Println(names)//[John, Paul, George, Ringo]
        }
        ```
- **Slice default**
    ```go 
    a[:x]//a slice of array "a" with the range [0 : x)
    a[x:]//a slice of array "a" with the range [x : n] where n is the length of array "a"
    a[:]//It is "a" itself
    ```

- **Slice has both `lenght` and `capacity`**
- **`length`**:The number of elements that the slice contains
    - `length` can be obtained by `len(slice)`
- **`capacity`**:The number of elements that the underlaying array has
    - `capacity` can be obtained by `cap(slice)`
- `slice` can be re-slicing, 

- a `nil slice` has
    - `lenght = 0`
    - `capacity = 0`
    - `no underlying array`
- **Creating a slice with `make`**
    - where the thrid arguments `capacity` is **optional**
    ```go
    a := make([]T, len, cap)//creat a zeroed array and return slice that refers to that array
    //This is how we create dynamically allocated array
    ```
- **slice of slice**
    - very much like a multi-dimensional array
- **append new element to slice**
    - `func(s []T, ele...T)`
        - append a list of elements of type `T` onto slice `s`
        - the return is `s + appending elements`

#### Range
- `range` is a way to iterate over a slice or map
- `range` returns two value
    - the index `i`
    - a copy of element at index `i`
- You can skip the index of value by assigning `_`
- If you want the index only, drop the `element` part
    ```go
    //example
    package main

    import "fmt"

    func main() {
        array := make([]int, 10)
        for i := range array {//simply drop the `element` part
            pow[i] = 1 << uint(i)
            /* 
            pow[0] = 1 << 0
            pow[1] = 1 << 1
            .
            .
            .
            and so on
            */
        }

        for _, value := range pow {//simply put `_` to drop the index
            fmt.Println(value)//?
        }
    }
    ```





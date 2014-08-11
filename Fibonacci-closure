package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    p, p2 := 0, 0
    
    return func () int {
        
        p, p2 = p2, p + p2
        
        if p + p2 < 2 {
            p = 1
            return p
        } else {
            return p2
        }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

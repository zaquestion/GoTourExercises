package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    
    sa := strings.Split(s, " ")
    
    for _, v := range sa {
       m[v]++;   
    }
    
    return m
}

func main() {
    wc.Test(WordCount)
}

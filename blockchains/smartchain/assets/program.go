package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)
type Symbol struct {
    Symbol   string
    
}
func main() {
    symbolFile, err := os.Open("info.json")
    if err != nil {
        log.Fatal(err)
    }
    defer symbolFile.Close()
    var studentDecoder *json.Decoder = json.NewDecoder(symbolFile)
    if err != nil {
        log.Fatal(err)
    }
    var studentList []Symbol
    err = studentDecoder.Decode(&studentList)
    if err != nil {
        log.Fatal(err)
    }
    for i, student := range studentList {
        fmt.Println("Symbol", i+1)
        fmt.Println("Symbol:", student.Symbol)
    }
}
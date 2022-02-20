package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    files, err := ioutil.ReadDir(".")
    if err != nil {
        panic(err)
    }

    fmt.Println("Файлы в репозитории")

    for _, fileInfo := range files {
        fmt.Println(fileInfo.Name())
    }
}
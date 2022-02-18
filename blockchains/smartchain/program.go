package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Website struct {
      Website   string
    
}

func main() {
    // open the file pointer
    websiteFile, err := os.Open("info.json")
    if err != nil {
        log.Fatal(err)
    }
    defer websiteFile.Close()

    // create a new decoder
    var studentDecoder *json.Decoder = json.NewDecoder(websiteFile)
    if err != nil {
        log.Fatal(err)
    }

    // initialize the storage for the decoded data
    var studentList []Website
    
    // decode the data
    err = studentDecoder.Decode(&studentList)
    if err != nil {
        log.Fatal(err)
    }

    for i, student := range studentList {
        fmt.Println("Student", i+1)
        fmt.Println("Student name:", student.Website)
    
    }
}
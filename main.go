//
// Refactor Pycut tool using Golang
//
//      TODO:
//       - Maybe use channels for printing to output?
//          According to the following GH page
//          https://github.com/spkg/bom/blob/master/bom.go
//          The compiler can work with BOMs. Perhaps marshalling the data into
//          a JSON channel is the best approach?
//       - Add argument parsing
//       - Use Stream instead of Reader
//

package main

import (
//    "bufio"
    "encoding/csv"
//    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
    "github.com/bom"
)

func main() {
    csvFile, _ := os.Open("test.csv")
    reader := csv.NewReader(bom.NewReader(csvFile))

    // I SMELL A GOROUTINE?
    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        fmt.Println(line)
    }
}

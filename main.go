//
// Refactor Pycut tool using Golang
//
//      TODO:
//       - Maybe use channels for printing to output?
//       - Add argument parsing
//       - Use Stream instead of Reader
//

package main

import (
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
)

func doCheckLine(reader bufio.Reader, field int, delimeter string) {

}

func printLines(reader bufio.Reader, field int, delimeter string) {

}

func main() {
    csvFile, _ := os.Open("people.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))

    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }

        //people = append(people, Person{
        //    Firstname: line[0],
        //    Lastname:  line[1],
        //    Address: &Address{
        //        City:  line[2],
        //        State: line[3],
        //    },
        //})
    }
    peopleJson, _ := json.Marshal(people)
    fmt.Println(string(peopleJson))
}

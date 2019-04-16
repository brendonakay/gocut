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
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
)

const (
	bom0 = 0xef
	bom1 = 0xbb
	bom2 = 0xbf
)

func doCheckLine(line bufio.Reader, field int, delimeter string) {
    b, error := line.Peak(3)
	if len(b) >= 3 &&
		b[0] == bom0 &&
		b[1] == bom1 &&
		b[2] == bom2 {
		return b[3:]
	}
	return b

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

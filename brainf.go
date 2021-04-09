package main

import (
    "flag"
    "fmt"
    "os"
    "time"
    "io/ioutil"
)

func render_memory(cells [8]byte, pntr int) {
    for k := 0; k < len(cells); k++ {
        fmt.Printf("|-%03d-", cells[k])
    }

    fmt.Println("|")
    fmt.Printf(" ")
    for i := 0; i < pntr; i++ {
        fmt.Printf("      ")
    }
    fmt.Println(" (^)                                             ")
}

func main() {
    fmt.Printf("\x1b[2J")
    output := []string{}
    var programFile string
    var input string
 //   var cellNumber int
    const cellNumber = 8
    flag.StringVar(&programFile, "file", "", "path to program file")
    flag.StringVar(&input, "input", "", "(optional) input for program")
//    flag.IntVar(&cellNumber, "cells", 8, "max amount of cells")
    flag.Parse()
    inputArray = []byte(input)
    if len(programFile) == 0 {
        os.Exit(1)
    }
    buf, _ := ioutil.ReadFile(programFile)
    program := string(buf)
    cells := [cellNumber]byte{}
    pntr := 0

    for i := 0; i< len(program); i++ {
        switch program[i] {
        case '+':
            cells[pntr] = cells[pntr] + 1

        case '-':
            cells[pntr] = cells[pntr] - 1

        case '>':
            pntr = (pntr + 1) % 8

        case '<':
            pntr = (pntr - 1) % 8

        case '.':
            output = append(output, string(cells[pntr]))

        case ',':
            cells[pntr] = inputArray[0]
            inputArray = inputArray[:(len(stack) - 1)]

        case '[':
            if (cells[pntr] == 0) {
                openBraces := 1
                for (openBraces > 0) {
                    i = i + 1
                    if (program[i] == '[') {
                        openBraces = openBraces + 1
                    }

                    if (program[i] == ']') {
                        openBraces = openBraces - 1
                    }
                }
            }

        case ']':
                openBraces := 1
                for (openBraces > 0) {
                    i = i - 1
                    if (program[i] == '[') {
                        openBraces = openBraces - 1
                    }

                    if (program[i] == ']') {
                        openBraces = openBraces + 1
                    }
                }
                i = i - 1
        }
        if (pntr < 0) {pntr = pntr + 8}
        if (pntr > 7) {pntr = pntr - 8}
        fmt.Print("\x1b[H")
        render_memory(cells, pntr)
        fmt.Printf("Output: %s", output)
        time.Sleep((1 * time.Second)/20)
        fmt.Print("\x1b[H")
    }
//    render_memory(cells)
    fmt.Print("\x1b[4B")
    fmt.Scanln()
}

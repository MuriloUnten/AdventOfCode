package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
)


func openFile(path string) *os.File {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    return file
}


func part1(path string) int {
    file := openFile(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var score int
    for scanner.Scan() {
        firstChar := scanner.Text()[0]
        secondChar := scanner.Text()[2]

        switch firstChar {
            case 'A':
                switch secondChar {
                    case 'X':
                        score += 4
                    case 'Y':
                        score += 8
                    case 'Z':
                        score += 3
                }
            case 'B':
                switch secondChar {
                    case 'X':
                        score += 1
                    case 'Y':
                        score += 5
                    case 'Z':
                        score += 9
                }
            case 'C':
                switch secondChar {
                    case 'X':
                        score += 7
                    case 'Y':
                        score += 2
                    case 'Z':
                        score += 6
                }
        }
    }
    return score
}


func part2(path string) int {
    file := openFile(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var score int
    for scanner.Scan() {
        firstChar := scanner.Text()[0]
        secondChar := scanner.Text()[2]

        switch firstChar {
            case 'A':
                switch secondChar {
                    case 'X':
                        score += 3
                    case 'Y':
                        score += 4
                    case 'Z':
                        score += 8
                }
            case 'B':
                switch secondChar {
                    case 'X':
                        score += 1
                    case 'Y':
                        score += 5
                    case 'Z':
                        score += 9
                }
            case 'C':
                switch secondChar {
                    case 'X':
                        score += 2
                    case 'Y':
                        score += 6
                    case 'Z':
                        score += 7
                }
        }
    }
    return score
}


func main() {
    fmt.Println(part1("input.txt"))
    fmt.Println(part2("input.txt"))
}

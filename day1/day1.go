package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
)


func openFile(path string) *os.File {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return file
}


func readCurrentLine(scanner *(bufio.Scanner)) (int, error) {
    value, err := strconv.Atoi(scanner.Text())
    return value, err
}


func part1(filePath string) int {
    file := openFile(filePath)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    biggest, current := 0, 0

    for scanner.Scan() {
        value, err := readCurrentLine(scanner)
        if err != nil {
            if current > biggest {
                biggest = current
            }
            current = 0
            continue
        }
        current += value
    }
    if current > biggest {
        biggest = current
    }

    return biggest
}


func part2(filePath string) int {
    file := openFile(filePath)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    first, second, third, current := 0, 0, 0, 0

    for scanner.Scan() {
        value, err := readCurrentLine(scanner)
        if err != nil {
            if current > third {
                third = current
                if current > second {
                    second, third = third, second
                    if current > first {
                        first, second = second, first
                    }
                }
            }
            current = 0
            continue
        }
        current += value
    }

    if current > third {
        third = current
        if current > second {
            second, third = third, second
            if current > first {
                first, second = second, first
            }
        }
    }
    
    return (first + second + third)
}


func main() {
    fmt.Println(part1("input.txt"))
    fmt.Println(part2("input.txt"))
}

package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)


func findMisplacedItem(scanner *bufio.Scanner) byte {
    text := scanner.Text()
    half := len(text) / 2

    for i := 0; i < half; i++ {
        for j := half; j < 2 * half; j++ {
            if text[i] == text[j] {
                return text[i]
            }
        }
    }

    return 0
}


func calculatePriority(char byte) int {
    if char <= 'Z' {
        return (int(char) - 38)
    }
    return (int(char) - 96)
}


func  findCommonItemPriority(first string, second string, third string) int {
    var count [52]int
    for i := 0; i < len(first); i++ {
        charPosition := calculatePriority(first[i]) - 1
        if count[charPosition] == 0 {
            count[charPosition] = 1
        }
    }
    for i := 0; i < len(second); i++ {
        charPosition := calculatePriority(second[i]) - 1
        if count[charPosition] == 1 {
            count[charPosition] = 2
        }
    }
    for i := 0; i < len(third); i++ {
        charPosition := calculatePriority(third[i]) - 1
        if count[charPosition] == 2 {
            return charPosition + 1
        }
    }
    return 0
}


func part1(path string) int {
    file, err := os.Open(path)
    if err != nil { log.Fatal(err) }
    defer file.Close()

    priority := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        char := findMisplacedItem(scanner)
        priority += calculatePriority(char)
    }
    return priority
}


func part2(path string) int {
    file, err := os.Open(path)
    if err != nil { log.Fatal(err) }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    sumOfPriorities := 0

    for scanner.Scan() {
        first := scanner.Text()
        scanner.Scan()
        second := scanner.Text()
        scanner.Scan()
        third := scanner.Text()

        sumOfPriorities += findCommonItemPriority(first, second, third)
    }
    return sumOfPriorities
}


func main() {
    fmt.Println(part1("input.txt"))
    fmt.Println(part2("input.txt"))
}

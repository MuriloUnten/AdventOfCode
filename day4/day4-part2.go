package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
    "strconv"
)

func main() {
    begin := time.Now()
    answer := part2("input.txt")
    end := time.Now()
    fmt.Println(answer, end.Sub(begin))
}


func part2(path string) int {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    overlaps := 0
    var pairAssignments []string
    var firstElfAssignment []string
    var secondElfAssignment []string

    for scanner.Scan() {
        pairAssignments = strings.Split(scanner.Text(), ",")

        firstElfAssignment = strings.Split(pairAssignments[0], "-")
        first, _ := strconv.Atoi(firstElfAssignment[0])
        second, _ := strconv.Atoi(firstElfAssignment[1])

        secondElfAssignment = strings.Split(pairAssignments[1], "-")
        third, _ := strconv.Atoi(secondElfAssignment[0])
        fourth, _ := strconv.Atoi(secondElfAssignment[1])

        if (third >= first && third <= second) ||
        (second >= third && second <= fourth) ||
        (first <= fourth && first >= third) {
            overlaps++
        }
    }

    return overlaps
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "time"
    "strconv"
)

type stack []byte

func (stack *stack) Push(char byte) {
    *stack = append(*stack, char)
}

func (stack *stack) Pop() byte {
    index := len(*stack) -1
    element := (*stack)[index]
    *stack = (*stack)[:index]
    return element
}

func (stack * stack) IsEmpty() bool {
    return len(*stack) == 0
}

func main() {
    fmt.Println(part1("input.txt"))
}


func part1(path string) string {
    file, _ := os.Open(path)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var invertedStacks []stack
    var correctedStacks []stack

    for scanner.Scan() && scanner.Text() != "" {
        for i, char := range scanner.Text() {
            if char >= 'A' && char <= 'Z' {
                var index int = i / 4
                if index > len(invertedStacks) - 1 {
                    difference := index - len(invertedStacks) + 1
                    for j := 0; j < difference; j++ {
                        invertedStacks = append(invertedStacks, make(stack, 0))
                    }
                }
                invertedStacks[index].Push(byte(char))
            }
        }
    }
    for i, currentStack := range invertedStacks {
        correctedStacks = append(correctedStacks, make(stack, 0))
        for !currentStack.IsEmpty() {
            correctedStacks[i].Push(currentStack.Pop())
        }
    }

    for scanner.Scan() {
        lineTokens := strings.Split(scanner.Text(), " ")
        numMovements, _ := strconv.Atoi(lineTokens[1])
        origin, _ := strconv.Atoi(lineTokens[3])
        destination, _ := strconv.Atoi(lineTokens[5])
        for i := 0; i < numMovements; i++ {
            correctedStacks[destination - 1].Push(correctedStacks[origin - 1].Pop())
        }
    }

    var answer string
    for _, currentStack := range correctedStacks {
        answer += string(currentStack.Pop())
    }
    return answer
}

package main

import (
    "fmt"
    "bufio"
    "os"
    "time"
)

func main() {
    begin := time.Now()
    answer := part1("input.txt")
    end := time.Now()
    fmt.Println(answer, end.Sub(begin))
}


func part1(path string) int {
    file, _ := os.Open(path)
    defer file.Close()
    reader := bufio.NewReader(file)
    buffer := make([]byte, 4096)
    n, _ := reader.Read(buffer)
    var helperBuffer [14]byte

    for i := 0; i < 14; i++ {
        helperBuffer[i] = buffer[i]
    }

    for i := 4; i < n; i++ {
        if !checkForRepeatedLetters(helperBuffer) {
            return i
        }
        helperBuffer = cicleLetters(helperBuffer, buffer, i)
    }
    return 0
}


func checkForRepeatedLetters(buffer [14]byte) bool {
    for i := 0; i < 14; i++ {
        for j := i + 1; j < 14; j++ {
            if buffer[i] == buffer[j] {
                return true
            }
        }
    }
    return false
}


func cicleLetters(helper [14]byte, mainBuffer []byte, position int) [14]byte {
    for i := 0; i < 13; i++ {
        helper[i] = helper[i + 1]
    }
    helper[13] = mainBuffer[position]
    return helper
}

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
    var helperBuffer [4]byte

    helperBuffer[0] = buffer[0]
    helperBuffer[1] = buffer[1]
    helperBuffer[2] = buffer[2]
    helperBuffer[3] = buffer[3]
    for i := 4; i < n; i++ {
        if !checkForRepeatedLetters(helperBuffer) {
            return i
        }
        helperBuffer = cicleLetters(helperBuffer, buffer, i)
    }
    return 0
}


func checkForRepeatedLetters(buffer [4]byte) bool {
    for i := 0; i < 4; i++ {
        for j := i + 1; j < 4; j++ {
            if buffer[i] == buffer[j] {
                return true
            }
        }
    }
    return false
}


func cicleLetters(helper [4]byte, mainBuffer []byte, position int) [4]byte {
    for i := 0; i < 3; i++ {
        helper[i] = helper[i + 1]
    }
    helper[3] = mainBuffer[position]
    return helper
}

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// Borrowed from Leetcode Solution here:
//   https://leetcode.com/problems/sudoku-solver/discuss/1586117/Go-Solution%3A-4-ms-(86.32)-2.6-MB-(11.58)
func nextUnknownSquare(board [][]byte) (int, int) {
    for y := 0; y < len(board); y++ {
        for x := 0; x < len(board[y]); x++ {
            if board[y][x] == '.' {
                return y, x
            }
        }
    }

    return -1, -1
}

func possibleDigits(board [][]byte, y int, x int) []int {
    result := make([]int, 0)

    digits := make([]bool, 9)

    // Horizontal
    for xPos := 0; xPos < len(board[y]); xPos++ {
        if board[y][xPos] != '.' {
            digits[board[y][xPos]-'1'] = true
        }
    }

    // Vertical
    for yPos := 0; yPos < len(board); yPos++ {
        if board[yPos][x] != '.' {
            digits[board[yPos][x]-'1'] = true
        }
    }

    // Square
    for yPos := y / 3 * 3; yPos < y/3*3+3; yPos++ {
        for xPos := x / 3 * 3; xPos < x/3*3+3; xPos++ {
            if board[yPos][xPos] != '.' {
                digits[board[yPos][xPos]-'1'] = true
            }
        }
    }

    for index, value := range digits {
        if !value {
            result = append(result, index+1)
        }
    }

    return result
}

func solveSudokuHelper(board [][]byte) bool {
    y, x := nextUnknownSquare(board)

    if (y == -1) || (x == -1) {
        return true
    }

    for _, digit := range possibleDigits(board, y, x) {
        board[y][x] = byte(digit) + '0'

        if solveSudokuHelper(board) {
            return true
        }

        board[y][x] = '.'
    }

    return false
}

func solveSudoku(board [][]byte) {
    solveSudokuHelper(board)
}

func printBoard(board [][]byte) {
    for y := 0; y < len(board); y++ {
        fmt.Println(string(board[y]))
    }
}
func readFile(filename string) [][]byte {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var contents [][]byte
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        line := []byte(scanner.Text())
        contents = append(contents, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return contents
}

func main() {
    board := readFile("board.txt")

    solveSudoku(board)

    printBoard(board)
}

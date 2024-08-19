package main

import (
    "fmt"
    "time"
    "math/rand"
    "strconv"
)

func display (board [][]string) {
    for i :=0; i < len(board); i++ {
        fmt.Print("\033[1A\033[K")
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            fmt.Printf("%s ", board[i][j])
        }
        fmt.Println()
    }
}

func first_board_display () {
    board := [][]string{
        {"二桁の番号を入力するのじゃ"},
        {"　","１","２","３","４","５","６","７","８","９"},
        {"１","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"２","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"３","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"４","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"５","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"６","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"７","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"８","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"９","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            fmt.Printf("%s ", board[i][j])
        }
        fmt.Println()
    }
}

func board_display (stones [][]int) {

    board := [][]string{
        {"１","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"２","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"３","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"４","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"５","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"６","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"７","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"８","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
        {"９","＋","＋","＋","＋","＋","＋","＋","＋","＋",},
    }

    stone := []string {"＋", "⚪︎", "⚫︎"}
    for i := 0; i < len(stones); i++ {
        for j:= 0; j < len(stones[i]); j++{
            board[i][j+1] = stone[stones[i][j]]
        }
    }

    display(board)
}

func com_set_stones (stones [][]int) [][]int {
    var positions []struct{ i, j int }
    for i, row := range stones {
        for j, value := range row {
            if value == 0 {
                positions = append(positions, struct{ i, j int }{i, j})
            }
        }
    }
    rand.Seed(time.Now().UnixNano())
    pos := positions[rand.Intn(len(positions))]
    stones[pos.i][pos.j] = 2
    return stones
}

func user_set_stones (stones [][]int) [][]int {
    var input string
    for {
        // input
        fmt.Scanln(&input)
        fmt.Print("\033[1A\033[K")

        if value, err := strconv.Atoi(input); err != nil {
            // check type
            continue
        } else if value < 11 {
            // check length
            continue
        } else if value > 99 {
            // check length
            continue
        } else if i := value % 10; i == 0 {
            // check i not 0
            continue
        } else if j := int(value/10); j == 0 {
            // check j not 0
            continue
        } else if stones[i-1][j-1] != 0 {
            // check already stone
            continue
        } else {
            // set
            stones[i-1][j-1] = 1
            return stones
        }
    }
}

func main() {
    stones := [][]int{
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0},
    }

    first_board_display()

    for i := 0; i < 40; i++{
        stones = user_set_stones(stones)
        board_display(stones)

        time.Sleep(2 * time.Second)
        stones = com_set_stones(stones)
        board_display(stones)
    }
}


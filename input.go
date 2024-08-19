package main

import (
    "fmt"
    "time"
    "math/rand"
    "strconv"
)

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
        clear_row()

        // validation
        if value, err := strconv.Atoi(input); err != nil {
            // check type
            continue
        } else if !(10 < value && value < 100) {
            // check length
            continue
        } else if i, j := value%10, int(value/10); i*j == 0 {
            // check both i and j are not 0
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

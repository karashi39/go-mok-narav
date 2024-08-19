package main

import (
    "fmt"
    "time"
    "math/rand"
    "strconv"
)

func com_set_stones (s *stones) (int,int) {
    // list up empty place
    var positions []struct{ i, j int }
    for i, row := range s {
        for j, value := range row {
            if value == 0 {
                positions = append(positions, struct{ i, j int }{i, j})
            }
        }
    }

    // set com stone random
    rand.Seed(time.Now().UnixNano())
    pos := positions[rand.Intn(len(positions))]
    return pos.i, pos.j
}

func user_set_stones (s *stones) (int,int) {
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
        } else if s[i-1][j-1] != 0 {
            // check already stone
            continue
        } else {
            // set
            return i-1, j-1
        }
    }
}

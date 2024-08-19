package main

import (
    "time"
)

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


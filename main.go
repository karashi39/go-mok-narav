package main

import (
    "time"
)

func main() {
    s := new(stones)
    first_board_display(s)

    for i := 0; i < 40; i++{
        s = user_set_stones(s)
        display_board(s)

        time.Sleep(1 * time.Second)
        s = com_set_stones(s)
        display_board(s)
    }
}


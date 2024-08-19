package main

import (
    "time"
)

func main() {
    s := new(stones)
    first_board_display(s)

    var stone, i, j int
    for t := 0; t < 40; t++{
        //先攻
        stone = 1
        i, j = user_set_stones(s)
        update_stones(s,i,j,stone)
        display_board(s)

        //後攻
        stone = 2 //後攻
        time.Sleep(1 * time.Second)
        i, j = com_set_stones(s)
        update_stones(s,i,j,stone)
        display_board(s)
    }
}


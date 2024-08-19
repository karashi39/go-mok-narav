package main

import (
    "time"
)

func main() {
    s := new(stones)
    d := new(display)
    d.init(s)

    var stone, i, j int
    for t := 0; t < 5; t++{
        //先攻
        stone = 1
        i, j = user_set_stones(s)
        update_stones(s,i,j,stone)
        d.update_board(s)

        time.Sleep(1 * time.Second)
        if t == 4 {
            break
        }

        //後攻
        stone = 2 //後攻
        i, j = com_set_stones(s)
        update_stones(s,i,j,stone)
        d.update_board(s)
    }
    d.winner()
}


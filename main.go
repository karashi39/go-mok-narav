package main

import (
    "time"
)

func main() {
    s := new(stones)
    d := new(display)
    d.init(s)

    var i, j, p int
    for t := 0; t < 5; t++{
        //先攻
        p = 1
        i, j = user_set_stones(s)
        s.update(i,j,p)
        d.update_board(s)

        time.Sleep(1 * time.Second)

        if s.judge(p) {
            d.winner()
            break
        }

        //後攻
        p = 2 //後攻
        i, j = com_set_stones(s)
        s.update(i,j,p)
        d.update_board(s)
        if s.judge(p) {
            d.loser()
            break
        }
    }
}


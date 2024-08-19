package main

import (
    "fmt"
    "strings"
)

func clear_row () {
    // stdoutのカーソルを1行上に戻す
    fmt.Print("\033[1A\033[K")
}

type display struct {
    lines [12]string
    stone []string
    index []string
}

func (d *display) init(s *stones) {
    stone := []string {"＋", "⚪︎", "⚫︎"}
    index := []string{"１","２","３","４","５","６","７","８","９"}
    d.index, d.stone = index, stone

    for i:=0; i<len(d.lines); i++{
        fmt.Println()
    }

    d.lines[0] = "二桁の番号を入力するのじゃ"
    d.lines[1] = fmt.Sprintf("　 %s", strings.Join(index, " "))
    for i := 2; i < 11; i++ {
        d.lines[i] = fmt.Sprintf("%s ＋ ＋ ＋ ＋ ＋ ＋ ＋ ＋ ＋ ", d.index[i-2])
    }
    d._update()
}

func (d *display) _update() {
    for i := 0; i < len(d.lines); i++{
        clear_row()
    }
    for i := 0; i < len(d.lines); i++{
        fmt.Println(d.lines[i])
    }
}

func (d *display) update_board (s *stones) {
    for i := 0; i < len(d.index); i++ {
        line := fmt.Sprintf("%s", d.index[i])
        for j:= 1; j < 10; j++{
            line = fmt.Sprintf("%s %s", line, d.stone[s[i][j-1]])
        }
        d.lines[i+2] = line
    }
    d._update()
}

func (d *display) winner () {
    d.lines[11] = "           YOU WIN"
    d._update()
}

func (d *display) loser () {
    d.lines[11] = "           YOU LOSE"
    d._update()
}

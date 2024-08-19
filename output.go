package main

import (
    "fmt"
)

func clear_row () {
    // stdoutのカーソルを1行上に戻す
    fmt.Print("\033[1A\033[K")
}

func display_board (s *stones) {
    y_index := []string{"１","２","３","４","５","６","７","８","９"}
    stone := []string {"＋", "⚪︎", "⚫︎"}

    for i := 0; i < 9; i++ {
        clear_row()
    }
    for i := 0; i < 9; i++ {
        fmt.Printf("%s ", y_index[i])
        for j:= 1; j < 10; j++{
            fmt.Printf("%s ", stone[s[i][j-1]])
        }
        fmt.Println()
    }
}

func first_board_display (s *stones) {
    fmt.Println("二桁の番号を入力するのじゃ")
    fmt.Println("　 １ ２ ３ ４ ５ ６ ７ ８ ９")
    for i := 0; i < 9; i++ {
        fmt.Println()
    }
    display_board(s)
}

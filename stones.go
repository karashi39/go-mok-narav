package main

import (
    "strconv"
    "strings"
    "regexp"
    "fmt"
)

type stones [9][9]int

func (s *stones) update (i, j, p int) {
    s[i][j] = p
}

func (s *stones) judge (p int) bool{
    win_pattern := strings.Repeat(strconv.Itoa(p), 5)

    // 横方向の判定
    for _, line := range s{
        var line_str string
        for _, num := range line {
            line_str = fmt.Sprintf("%s%d", line_str, num)
        }
        if match, _ := regexp.MatchString(win_pattern,line_str); match {
            return true
        }
    }

    // 縦方向の判定
    for j:=0; j < len(s[0]); j++{
        var line_str string
        for i:=0; i < len(s); i++{
            line_str = fmt.Sprintf("%s%d", line_str, s[i][j])
        }
        if match, _ := regexp.MatchString(win_pattern,line_str); match {
            return true
        }
    }
    // 斜め方向の判定
    return false
}

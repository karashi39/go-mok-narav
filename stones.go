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
        match, _ := regexp.MatchString(win_pattern,line_str)
        if match {
            return true
        }
    }
    // 縦方向の判定
    // 斜め方向の判定
    return false
}

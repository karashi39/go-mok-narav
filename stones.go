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
    var horizontal string
    for _, line := range s{
        for _, num := range line {
             horizontal= fmt.Sprintf("%s%d", horizontal, num)
        }
        horizontal= fmt.Sprintf("%s%s", horizontal, "-")
    }
    if match, _ := regexp.MatchString(win_pattern,horizontal); match {
        return true
    }

    // 縦方向の判定
    var vertical string
    for j:=0; j < len(s[0]); j++{
        for i:=0; i < len(s); i++{
            vertical = fmt.Sprintf("%s%d", vertical, s[i][j])
        }
        vertical = fmt.Sprintf("%s%s", vertical, "-")
    }
    if match, _ := regexp.MatchString(win_pattern,vertical); match {
        return true
    }

    // 斜め方向の判定
    var oblique string
    for i:= -4; i<5; i++{
        for j:=0; j<len(s[0]); j++{
            if i + j < 0 || 8 < i + j  {
                continue
            } else {
                oblique = fmt.Sprintf("%s%d", oblique, s[i+j][j])
            }
        }
        oblique = fmt.Sprintf("%s%s", oblique, "-")
    }

    if match, _ := regexp.MatchString(win_pattern,oblique); match {
        return true
    }

    // 逆斜め方向の判定
    var oblique_r string
    for i:= 4; i<13; i++{
        for j:=0; j<len(s[0]); j++{
            if i - j < 0 || 8 < i - j  {
                continue
            } else {
                oblique_r = fmt.Sprintf("%s%d", oblique_r, s[i-j][j])
            }
        }
        oblique_r = fmt.Sprintf("%s%s", oblique_r, "-")
    }
    if match, _ := regexp.MatchString(win_pattern,oblique_r); match {
        return true
    }

    return false
}

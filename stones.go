package main

type stones [9][9]int

func update_stones(s *stones, i, j, p int) {
    s[i][j] = p
}

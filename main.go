package main

import (
	"fmt"
	"time"
        "math/rand"
)

var social [20][20]int
var length = 20

func display() {
	for _, k := range social {
		fmt.Println(k)
	}
}

func generate() {
	newSoc := social
	for i:=0;i<length;i++ {
		for j:=0;j<length;j++ {
			num := count(i, j)
			if newSoc[i][j] == 0 && num == 3 {
				newSoc[i][j] = 0
			} else if newSoc[i][j] == 1 {
				if num != 2 && num != 3 {
					newSoc[i][j] = 0
				}
			}
		}
	}
	social = newSoc
}

func genRand() {
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			if rand.Intn(3) < 2 {
				social[i][j] = 1
			}
		}
	}
}

func count(row, col int) int {
	cot := 0
	for i:=-1;i<2;i++{
		for j:=-1;j<2;j++{
			tRow := row+i
			tCol := col+j
			tRow = checkRange(tRow)
			tCol = checkRange(tCol)
			if (tRow == row && tCol == col) || tRow == -1 || tCol == -1 {
				continue
			}
			if social[tRow][tCol] == 1 {
				cot += 1
			}
		}
	}
	return cot
}

func checkRange(num int) int {
	if num < 0 || num > length-1 {
		return -1
	}
	return num
}

func showTimely() {
	for {
		time.Sleep(time.Second)
		generate()
		display()
                fmt.Println()
	}
}

func main() {
	fmt.Println(1)
        genRand()
        display()
        fmt.Println()
	showTimely()
}




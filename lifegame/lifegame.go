package lifegame

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"time"
)

var social [10][10]int
var length = 10

type LifegameController struct {
	beego.Controller
}


func (c *LifegameController) GetIndex() {
	c.TplName="index.html"
}


func (c *LifegameController) Get() {
	c.Data["json"] = updateSocial()
	c.TplName="draw.html"
}


func init() {
	genRand()
	//test2()
}

func display() {
	for _, k := range social {
		fmt.Println(k)
	}
	fmt.Println()
}

func genRand() {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			if rand.Intn(2) == 0 {
				social[i][j] = 1
			}
		}
	}
}

func generate() {
	newSoc := social
	for i:=0;i<length;i++ {
		for j:=0;j<length;j++ {
			num := count(i, j)
			if newSoc[i][j] == 0 && num == 3 {
				newSoc[i][j] = 1
			} else if newSoc[i][j] == 1 {
				if num != 2 && num != 3 {
					newSoc[i][j] = 0
				}
			}
		}
	}
	social = newSoc
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
	}
}

func test0() {
	social[0][0] = 1
	social[1][0] = 1
	social[0][1] = 1
	social[1][1] = 1
}

func test1() {
	social[2][2] = 1
	social[2][3] = 1
	social[2][4] = 1
}

func test2() {
	social[2][3] = 1
	social[2][4] = 1
	social[2][5] = 1
	social[3][2] = 1
	social[3][3] = 1
	social[3][4] = 1
}

func updateSocial() [10][10]int {
	generate()
	display()
	fmt.Println()
	return social
}



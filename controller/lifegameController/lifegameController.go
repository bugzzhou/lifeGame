package lifegameController

import (
	"beego/controller/baseController"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var social [10][10]int
var length = 10

type LifegameController struct {
	baseController.BaseController
}

func (c *LifegameController) FetchAutofresh() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Data["json"] = updateSocial()
	c.ServeJSON()
//	c.TplName = "empty.html"
}

func (c *LifegameController) FetchReset() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	genRand()
	c.Data["json"] = social
	c.ServeJSON()
//	c.TplName = "empty.html"
}

func (c *LifegameController) TestPost() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	input := c.Ctx.Input.RequestBody

	var inputMap map[string]int
	err := json.Unmarshal(input, &inputMap)
	if err != nil {
		log.Printf("failed to unmarshal, and err is: %s\n", err.Error())
		c.Data["json"] = "error"
		c.ServeJSON()
		return
	}

	count := inputMap["number"]
	res := make([][10][10]int, 0)
	for i := 1; i <= count; i++ {
		tmp := updateSocial()
		res = append(res, tmp)
	}

	c.Data["json"] = res
	c.ServeJSON()
//	c.TplName = "test.html"
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
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			t := rand.Intn(2)
			if t != 0 {
				social[i][j] = 1
			} else {
				social[i][j] = 0
			}
		}
	}
}

func generate() {
	newSoc := social
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
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
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			tRow := row + i
			tCol := col + j
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

//view 自用，暂不使用，保留
//func (c *LifegameController) Autofresh() {
//	c.Data["json"] = updateSocial()
//	c.TplName = "drawbak.html"
//}
//
//func (c *LifegameController) Reset() {
//	genRand()
//	c.Data["json"] = updateSocial()
//	c.TplName = "drawbak.html"
//}
//
//func (c *LifegameController) Test() {
//	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
//	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
//	//res := map[string]interface{}{
//	//	"result": []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
//	//}
//	c.Data["json"] = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
//	c.ServeJSON()
//	c.TplName = "test.html"
//}
//
//func (c *LifegameController) GetMap() {
//	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
//	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
//	res := map[string]interface{}{
//		"result": []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
//	}
//	c.Data["json"] = res
//
//	c.ServeJSON()
//	c.TplName = "test.html"
//}
//
//func (c *LifegameController) GetSlice() {
//	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
//	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
//	//res := map[string]interface{}{
//	//	"result": []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
//	//}
//	c.Data["json"] = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
//	c.ServeJSON()
//	c.TplName = "test.html"
//}

package service

import (
	"goLog/logs"
	"net/http"

	"github.com/labstack/echo"
)

//not recommend
func funcOne(a int, b int) (int, int) {
	return a + b, a * b
}

//recommend
func funcTwo(a, b int) (add int, multiplied int) {
	add = a + b
	multiplied = a * b
	return
}

func funcThree(arg ...int) {
	for i, n := range arg {
		logs.Debug.Printf("\nindex:%d", i)
		logs.Debug.Printf("\nnum:%d", n)
	}
}

func Getfunc(c echo.Context) error {
	a, b := funcOne(1, 2)
	a2, b2 := funcOne(3, 4)
	logs.Debug.Printf("\nadd:%v,multi:%v", a, b)
	logs.Debug.Printf("\nadd:%v,multi:%v", a2, b2)
	funcThree(1, 2, 5, 6)
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: "Hello func"})
}

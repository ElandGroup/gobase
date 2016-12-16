package service

import (
	"goLog/logs"
	"net/http"

	"github.com/labstack/echo"
)

func ifOne(x int) string {
	r := "greater than 10"
	if x > 10 {

	} else if x > 5 {
		r = "greater than 5"
	} else {
		r = "less than 5"
	}
	return r
}

func computerValue(x int, y int) int {
	return x - y
}

func ifTwo(x int, y int) bool {
	r := false
	if z := computerValue(x, y); z > 10 {
		r = true
	}
	return r
}

func Getif(c echo.Context) error {
	logs.Debug.Printf("\n8:%v", ifOne(8))
	logs.Debug.Printf("\nif2:%v", ifTwo(18, 3))
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: "Hello if"})
}

func forOne() {

	for index := 10; index > 0; index-- {
		logs.Debug.Printf("\ni:%d", index)
	}
}

func forTwo() {
	m := map[string]int{"X": 5, "C": 6, "V": 7}
	for k, v := range m {
		logs.Debug.Printf("\nkey:%v", k)
		logs.Debug.Printf("\nvalue:%v", v)
	}
}

func Getfor(c echo.Context) error {
	forOne()
	forTwo()
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: "Hello for"})
}

func switchOne() {
	i := 10
	switch i {
	case 1:
		logs.Debug.Printf("\nnumber:%d", i)
	case 2, 3, 10:
		logs.Debug.Printf("\nnumber:%d", i)
	case 8:
		logs.Debug.Printf("\nnumber:%d", i)
	default:
		logs.Debug.Printf("\nnumber:%d", i)
	}
}

func Getswitch(c echo.Context) error {
	switchOne()
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: "Hello switch"})
}

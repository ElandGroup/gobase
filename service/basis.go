package service

import (
	"net/http"

	"gorouter/logs"

	"github.com/labstack/echo"
)

var (
	gValue = "I am a global variable"
)

const (
	cValue = "I am a const"
)

func GetBasis(c echo.Context) error {
	//declare variable
	result := "basis practice"
	logs.Debug.Println("gValue:", gValue)
	lValue := "I am a local variable"
	logs.Debug.Println("\nlValue:", lValue)
	logs.Debug.Println("\ncValue:", cValue)
	//int , bool
	var i = 10
	logs.Debug.Println("\nint :", i)
	b := true
	logs.Debug.Println("\nbool:", b)

	//string
	s := "hello"
	m := " world"
	a := s + m
	logs.Debug.Printf("\n%s", a)

	ml := `Hello 
            world multi`
	logs.Debug.Printf("\nmulti:%s", ml)
	//slice
	sl := "hello"
	sl = "c" + sl[2:3]
	logs.Debug.Println("\nslice new word:", sl)

	//array
	ar := [3]int{1, 2, 3}
	br := [...]int{4, 5, 6}
	cr := []int{7, 8, 9}
	logs.Debug.Printf("\narray1:%v\n,array2:%v\n,slice:%v\n ", ar, br, cr)

	ar2 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	logs.Debug.Printf("\narray2:%v", ar2)

	//slice 2
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := array[2:5]
	logs.Debug.Println("\nslice:", slice)

	//map
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 10
	numbers["thress"] = 3

	nm2 := map[string]int{"C": 5, "go": 6, "C++": 7}
	logs.Debug.Printf("\nmap1:%v,map2:%v", numbers, nm2)

	return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})

}

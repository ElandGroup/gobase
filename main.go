package main

import "github.com/labstack/echo"

var (
	e = echo.New()
)

func main() {

	RegApi()

	e.Start(":1115")
}

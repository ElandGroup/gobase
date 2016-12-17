package service

import (
	"fmt"
	"gorouter/logs"
	"net/http"

	"github.com/labstack/echo"
)

type dowork func(name string) string

func teacherWork(name string) string {
	return fmt.Sprintf("\nTeacher:%v's work is teaching student", name)
}

func studentWork(name string) string {
	return fmt.Sprintf("\nStudent:%v's work is listening teacher's word", name)
}

func workDetail(name string, work dowork) string {
	return work(name)
}

func GetFuncPass(c echo.Context) error {
	logs.Debug.Printf("\n%v", workDetail("zhang", teacherWork))
	logs.Debug.Printf("\n%v", workDetail("xiao", studentWork))
	return c.JSON(http.StatusOK, APIResult{Success: true, Result: "Hello func param pass"})
}

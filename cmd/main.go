package main

import (
	"github.com/phipsp/todo"
	"github.com/sirupsen/logrus"
)

func main() {
	todo.SetLogger(logrus.New())
	err := todo.UntilDate("2022/06/07", "Expand todo API")
	if err != nil {
		return
	}
}

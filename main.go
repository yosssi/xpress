package main

import (
	"fmt"
	"github.com/yosssi/xpress/app/models"
)

var (
	application models.Application
)

// init executes initial processes.
func init() {
	application, err := models.NewApplication()
	if err != nil {
		panic(err)
	}
	fmt.Println(application)
}

// main executes main processes.
func main() {

}

package main

import "github.com/configure/config"

func main() {
	// calling micro services need to pass env + vendor
	c,err := config.GetConfig("dev","p52")
	if err != nil {
		println("dfgdf")
	}
	println(c)
	println(c.GetString("mongo"))
}
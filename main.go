package main

import (
	"fmt"
	"time"
)

type hello struct {
	Name string
	Time time.Time
}

func main() {
	val := map[string]hello{
		"a": hello{Name: "ssas"},
	}
	fmt.Printf("%s", val["a"])
}

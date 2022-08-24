package main

import (
	"fmt"
	"m8"
	"time"
)

func main() {
	fmt.Println(m8.Response(time.Now().Unix()))
}

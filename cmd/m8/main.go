package main

import (
	"fmt"
	"time"

	"github.com/reinaldoleon/m8"
)

func main() {
	fmt.Println(m8.Response(time.Now().Unix()))
}

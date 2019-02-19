package main

import (
	"fmt"

	"github.com/travelgateX/go-jwt-tools/jwt"
)

func main() {
	c := jwt.ParserConfig{}
	fmt.Println("hello dep world!", c.AdminGroup)
}

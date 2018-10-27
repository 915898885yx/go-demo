package main

import (
	"fmt"
)

type SuperHreo struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Num    int
	Street string
	City   string
}

func main() {
	e := SuperHreo{}
}

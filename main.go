package main

import (
	"fmt"
	"github.com/timojarv/orjabot/api"
)
func main() {
	fmt.Println("== ORJA ==")
	fmt.Println("Koska tietojohtaminen on helppoa ja kivaa!")
	go api.RunRouter()
	RunBot()
}

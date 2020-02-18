package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/timojarv/orjabot/api"
)

func main() {
	listFood := flag.Bool("food", false, "")
	flag.Parse()

	fmt.Println("== ORJA ==")
	fmt.Println("Koska tietojohtaminen on helppoa ja kivaa!")

	if *listFood {
		rs, err := FetchRestaurants()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rs.String())

		return
	}

	go api.RunRouter()
	RunBot()
}

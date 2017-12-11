package main

import ("fmt"
   "github.com/davecgh/go-spew/spew"

)

type TravelAgency struct {
	id int
	name string
}

func main() {

	
	ta := TravelAgency{id: 1, name: "Steves Travel"}

	spew.Dump(ta)

}
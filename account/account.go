package main

import ("fmt"
   "github.com/davecgh/go-spew/spew"

)

type Account struct {
	id int
	name string
	travelAgencyId int
	accountNbr string
}

func main() {
    fmt.Println("Starts")
	
	entity := Account{id: 1000000, name: "Steves Travel- CRM", travelAgencyId: 1, 
	                  accountNbr: "150568"}

	spew.Dump(entity)

}
package main

import (
	"fmt"
	"os"
)

type Brother struct {
	id         uint
	name       string
	address    string
	occupation string
	reason     string
}

var brothers []Brother = []Brother{
	{
		id:         1,
		name:       "Eman",
		address:    "Ambon",
		occupation: "Entrepreneur",
		reason:     "Karena Golang ez",
	},
	{
		id:         2,
		name:       "Rio",
		address:    "Rawa Lumbu",
		occupation: "Mobile Developer",
		reason:     "Karena Golang mudah untuk dipahami",
	},
	{
		id:         3,
		name:       "Bayu",
		address:    "Pamulang",
		occupation: "Software Engineer",
		reason:     "Karena Golang sangat terkenal",
	},
	{
		id:         4,
		name:       "Fajri",
		address:    "Pemalang",
		occupation: "Web Developer",
		reason:     "Karena dipaksa nyobain Golang",
	},
	{
		id:         5,
		name:       "Alfan",
		address:    "Warung Bambu",
		occupation: "Security Engineer",
		reason:     "Karena ngeliat dari temen",
	},
}

func main() {
	var arg = os.Args

	fmt.Printf("%#v\n", arg)
	
	if len(arg) < 1 {
		fmt.Println("Usage: program_name name")
		return
	}
	
	nameToFind := arg[1]
	
	found := findMyBrother(nameToFind)
	
	if !found {
		fmt.Println("Not Found Your Brother!")
	}  
}

func findMyBrother(name string) bool {
	found := false
	for _, eachBrother := range brothers {
		if name == eachBrother.name {
			fmt.Println("ID:", eachBrother.id)
			fmt.Println("Nama:", eachBrother.name)
			fmt.Println("Address:", eachBrother.address)
			fmt.Println("Occupation:", eachBrother.occupation)
			fmt.Println("Reason:", eachBrother.reason)
			found = true
			break
		}
	}
	return found
}
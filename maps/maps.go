package main

import "fmt"

func main() {

	StudentAge := make(map[string]int)

	StudentAge["teterkin"] = 42
	fmt.Println(StudentAge["teterkin"])
	fmt.Println(len(StudentAge))

	StudentAge["vasya"] = 43
	fmt.Println(StudentAge["vasya"])
	fmt.Println(len(StudentAge))

	fmt.Println(StudentAge)

	superHero := map[string]map[string]string{
		"Superman": {
			"realName": "Clark Kent",
			"city":     "Metropolis",
		},
		"Batman": {
			"realName": "Bruce Wayne",
			"city":     "Gotham City",
		},
	}
	fmt.Println(superHero)
	if temp, hero := superHero["Superman"]; hero {
		fmt.Println("Superman's real name is", temp["realName"], "and he lives in", temp["city"])
	}

}

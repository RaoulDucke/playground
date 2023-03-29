package main

import "fmt"

func main() {
	// 1 способ инициализации Мап
	m := make(map[string]string)
	m["RU"] = "+7"
	m["US"] = "+1"
	m["UZ"] = "+93"
	m["CN"] = "+86"

	for k, v := range m {
		fmt.Printf("%s = %s \n", k, v)
	}
	print("\n")
	delete(m, "RU")

	for k, v := range m {
		fmt.Printf("%s = %s \n", k, v)
	}
	// 2 cпособ инициализации Мап
	m1 := map[string]string{
		"RU": "+7",
		"US": "+1",
		"UA": "+338",
	}
	print("\n")
	for k, v := range m1 {
		fmt.Printf("%s = %s \n", k, v)
	}

}

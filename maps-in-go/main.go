package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Twitter":  "https://www.twitter.com",
	}

	fmt.Println("Google:", websites["Google"])
	fmt.Println("Facebook:", websites["Facebook"])
	fmt.Println("Twitter:", websites["Twitter"])

	websites["Google"] = "https://www.google.com/new"
	fmt.Println("Google:", websites["Google"])

	delete(websites, "Twitter")
	fmt.Println("Twitter:", websites["Twitter"])

	fmt.Println("Websites:", websites)

	_, isFound := websites["Twitter"]
	fmt.Println("Twitter found:", isFound)

	for key, value := range websites {
		fmt.Println(key, value)
	}
}

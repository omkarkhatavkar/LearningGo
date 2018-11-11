package main

import "fmt"

func main() {
	emails := make(map[int]string)

	emails[1] = "okhatavkar@gmail.com"
	emails[1] = "ramesh@gmail.com"
	emails[2] = "okhatavkar@gmail.com"

	fmt.Println(emails)
	delete(emails, 3)
	fmt.Println(emails)
	fmt.Println(getCount("okhatavkar@gmail.com", emails))
	for k := range emails {
		fmt.Printf("ID = %d \n", k)
	}

}

func getCount(email string, emails map[int]string) int {
	count := 0
	for _, emailID := range emails {
		if emailID == email {
			count = count + 1
		}
	}
	return count
}

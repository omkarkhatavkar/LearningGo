package main

import "fmt"

func main() {
	canPlayGames(32, "Ramesh")
	canPlayGames(1, "Rakesh")
	fmt.Println(tellMeTheTeam("Yellow"))
	fmt.Println(tellMeTheTeam("Green"))
	var color string
	fmt.Println("Give me the new color ? ")
	fmt.Scanln(&color)
	fmt.Println(tellMeTheTeam(color))
}

func canPlayGames(age int16, name string) bool {
	if age >= 10 {
		fmt.Println(name + " can play the video game :)")
		return true
	} else {
		fmt.Println(name + " can not play the game :(")
		return false
	}
}

func tellMeTheTeam(color string) string {
	switch color {
	case "Yellow":
		return "This is an Austrelia Team !"
	case "Blue":
		return "This is an Indian Team !"
	default:
		return "There is no team with that color !"

	}
}

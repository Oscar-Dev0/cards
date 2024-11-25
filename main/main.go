package main

import "github.com/Oscar-Dev0/cards"

func main() {
	car := cards.NewMemberCard()

	car.SetBackground("https://i.imgur.com/POQQ48c.png")
	car.SetUser("世界🎉", nil)


	bu, err := car.Buffer()
	if err != nil {
		return
	}

	cards.SavePNG(*bu, "./images/Membercard.png")
}
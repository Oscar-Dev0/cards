package examples

import "github.com/Oscar-Dev0/cards"

func init() {
	car := cards.NewMemberCard();

	car.SetUser("世界💀", nil);

	bu, err := car.Buffer();
	if err != nil {
		return;
	}

	cards.SavePNG(*bu, "../images/Membercard.png")

}
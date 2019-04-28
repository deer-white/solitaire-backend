package solitaire

type Cards struct {
	//     z y x
	Coords [][][]int
}

func NewGame(boardTemplate string) Cards {

	var cards Cards

	if boardTemplate == "default" {
		//cards.Coords = [][][]int{ { {0} } }

		cards.Coords = make([][][]int, 5)

		for z := range cards.Coords {
			cards.Coords[z] = make([][]int, 16)
			for y := range cards.Coords[z] {
				cards.Coords[z][y] = make([]int, 30)
			}
		}




	}

	return cards
}
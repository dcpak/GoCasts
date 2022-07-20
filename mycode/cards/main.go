package main

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five od Diamonds"

	//card := newCard()
	//fmt.Println(card)

	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	//cards := newDeck()

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	//cards[:2].print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
	//greeting := "Hi there"
	//fmt.Print([]byte(greeting))

	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards.txt")

	//cards := newDeckFromFile("my_acards.txt")
	//cards := newDeckFromFile("my_cards.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

//func newCard() string {
//	return "Five of Diamonds"
//}

package main

func main() {
	myCard := deckFromFile("Lidya.txt")
	myCard.shuffle()
	myCard.print()

}

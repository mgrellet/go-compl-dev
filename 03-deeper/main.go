package main


func main(){
	//cards := newDeck()
	//hand, remainingDeck := deal(cards,5)
	//cards.saveToFile("myFile.txt")
	//fmt.Println(cards.toString())
	//cards := newDeckFromFile("myFile.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print() 
}
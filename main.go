package main

func chooseRandomlyInitialTrigram(trigramMap TrigramMap) {

}

func main() {
	server := NewServer(NewTrigramStore())
	server.Run()
}

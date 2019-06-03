package main

func main() {
	server := NewServer(NewTrigramStore())
	server.Run()
}

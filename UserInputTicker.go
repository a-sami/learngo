package main

import "time"
import "fmt"

func main() {
	fmt.Print("Enter text: ")
	var input int
	fmt.Scanln(&input)

	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Second * 6)
	ticker.Stop()
	fmt.Println("Ticker stopped")
	fmt.Print(input)
}

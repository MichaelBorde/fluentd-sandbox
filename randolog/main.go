package main

import(
	"log"
	"math/rand"
	"time"
)

func main() {
	for {
		log.Printf("Some random log output %d\n", rand.Intn(1000))
		
		time.Sleep(1000 * time.Millisecond)
	}
}
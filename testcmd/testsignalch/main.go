package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	pokepoke "github.com/cheyinl/go-pokepoke"
)

func routine(waitGroup *sync.WaitGroup, sig *pokepoke.SignalCh) {
	defer waitGroup.Done()
	timeout := time.After(time.Second * 3)
	pokeCount := 0
	for {
		select {
		case <-timeout:
			log.Print("timeout, leave routine.")
			return
		case <-sig.Wait():
			pokeCount++
			log.Printf("got signal (%d).", pokeCount)
		}
	}
}

func main() {
	sig := pokepoke.NewSignalCh()
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go routine(&waitGroup, sig)
	for attempt := 0; attempt < 16; attempt++ {
		sig.Poke()
		log.Printf("poked (%d).", attempt)
		if rand.Intn(2) == 0 {
			log.Print("sleep 1 sec.")
			time.Sleep(time.Second)
		}
	}
	waitGroup.Wait()
	log.Print("stop.")
}

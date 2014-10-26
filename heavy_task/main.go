// main
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-sig:
			fmt.Println("Process was killed!")
			os.Exit(1)
		default:
		}
	}
	fmt.Println("Heavy task done! yay!!!!!")
}

package timeout

import (
	"fmt"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
)

// Exec executes timeout
func Exec(waitSecond int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	timeoutTime := time.Now().Add(time.Second * time.Duration(waitSecond))
	fmt.Printf("\rWaiting for %v seconds, press any key to quit...", waitSecond)

	var once sync.Once
	go func() {
		time.Sleep(time.Duration(waitSecond) * time.Second)
		once.Do(func() { done <- struct{}{} })
	}()

	go func() {
		if err := keyboard.Open(); err == nil {
			keyboard.GetKey()
		}
		once.Do(func() { done <- struct{}{} })
	}()
	defer keyboard.Close()

	for {
		select {
		case <-done:
			fmt.Println("")
			return
		case t := <-ticker.C:
			fmt.Printf("\rWaiting for %v seconds, press any key to quit...", timeoutTime.Sub(t).Round(time.Second).Seconds())
		}
	}
}

package timeout

import (
	"fmt"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
)

// Exec executes timeout
func Exec(waitSecond int) {
	fmt.Printf("\rWaiting for %v seconds, press any key to quit...", waitSecond)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	waitSecondDuration := time.Second * time.Duration(waitSecond)
	timeoutTime := time.Now().Add(waitSecondDuration)

	var once sync.Once
	go func() {
		time.Sleep(waitSecondDuration)
		once.Do(func() { close(done) })
	}()

	go func() {
		if err := keyboard.Open(); err == nil {
			keyboard.GetKey()
		}
		once.Do(func() { close(done) })
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

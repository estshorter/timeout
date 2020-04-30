package timeout

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

// Exec executes timeout
func Exec(waitSecond int) {
	fmt.Printf("\rWaiting for %v seconds, press any key to quit...", waitSecond)
	done := make(chan struct{})
	go func() {
		if err := keyboard.Open(); err == nil {
			keyboard.GetKey()
		}
		close(done)
	}()
	defer keyboard.Close()

	for i := waitSecond; i > 0; i-- {
		select {
		case <-done:
			fmt.Println("")
			return
		case <-time.After(time.Second):
			fmt.Printf("\rWaiting for %v seconds, press any key to quit...", i-1)
		}
	}
}

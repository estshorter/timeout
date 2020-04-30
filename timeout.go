package timeout

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func detectKeyInputAsync() <-chan error {
	done := make(chan error)
	go func() {
		err := keyboard.Open()
		if err == nil {
			keyboard.GetKey()
		}
		done <- err
		close(done)
	}()
	return done
}

// Exec executes timeout
func Exec(waitSecond int) error {
	fmt.Printf("\rWaiting for %v seconds, press any key to quit...", waitSecond)
	done := detectKeyInputAsync()

	defer fmt.Println("")
FOR:
	for i := waitSecond - 1; i >= 0; i-- {
		select {
		case err := <-done:
			if err != nil {
				return err
			}
			break FOR
		case <-time.After(time.Second):
			fmt.Printf("\rWaiting for %v seconds, press any key to quit...", i)
		}
	}
	return keyboard.Close()
}

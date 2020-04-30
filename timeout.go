package timeout

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

// Exec executes timeout
func Exec(waitSecond int) error {
	fmt.Printf("\rWaiting for %v seconds, press any key to quit...", waitSecond)
	inputComm, err := keyboard.GetKeys(0)
	if err != nil {
		return err
	}

	defer fmt.Println("")
FOR_LABEL:
	for i := waitSecond - 1; i >= 0; i-- {
		select {
		case <-inputComm:
			break FOR_LABEL
		case <-time.After(time.Second):
			fmt.Printf("\rWaiting for %v seconds, press any key to quit...", i)
		}
	}
	return keyboard.Close()
}

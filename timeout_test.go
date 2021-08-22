package timeout

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	// type args struct {
	// 	waitSecond int
	// }
	fmt.Println("start")
	Exec(3)
	fmt.Println("middle")
	Exec(3)
	fmt.Println("end")
}

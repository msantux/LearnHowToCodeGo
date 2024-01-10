package furtherexplored

import "fmt"

// this is alos "packge block" scope
// this is exported because the identifier "Fascinating" is captilized
func Fascinating() {
	fmt.Println("Planet speed:", planetSpeed)

	planetRotationSpeed := 1000
	fmt.Println("Planet spinning speed:", planetRotationSpeed)
}

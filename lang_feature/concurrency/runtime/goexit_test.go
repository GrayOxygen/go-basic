package runtime

import (
	"runtime"
	"testing"
)

func TestExit(t *testing.T) {
	exit := make(chan struct{}) //巧用于控制线程等待
	go func() {
		defer close(exit)
		defer println("a")
		func() {
			defer func() {
				println("b", recover() == nil) // recover
			}()

			func() { // Goexit
				println("c")
				runtime.Goexit() //
				println("c done.")
			}()

			println("b done.")
		}()

		println("a done.") //
	}()

	<-exit
	println("main exit.")
}

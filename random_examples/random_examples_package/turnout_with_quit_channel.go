package random_examples_package

func TurnoutWithQuitChannel(quit <-chan int, inA, inB, outA, outB chan int) {
	var data int
	for {
		select { //receive from first non blocking
		case data = <-inA:
		case data = <-inB:
		case <-quit: //Remember: closing always generates a message
			close(inA) //actually this is an anti-pattern...
			close(inB) //... but you can argue that quit acts as a delegate

			fanOut(inA, outA, outB) //flush the remaining data
			fanOut(inB, outA, outB)
			return
		}

		select { //send to first non blocking
		case outA <- data:
		case outB <- data:
		}
	}
}

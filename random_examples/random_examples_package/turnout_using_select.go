package random_examples_package

func Turnout(inA, inB <-chan int, outA, outB chan int) {
	var data int
	var more bool
	for {
		select { //receive from first non blocking
		case data, more = <-inA:
		case data, more = <-inB:
		}
		if !more {
			//...this can be better handled using a quit channel
			return
		}
		select { //send to first non blocking
		case outA <- data:
		case outB <- data:
		}
	}
}

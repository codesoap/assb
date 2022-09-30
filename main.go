package main

import (
	"fmt"
	"time"
)

func main() {
	returnPressed := make(chan bool)
	go reportReturnPresses(returnPressed)
	segment := 1
	isBreathing := false
	secondTicker := time.NewTicker(time.Second)
	secondsLeftInSegment := 4 * 60
	heldBreathFor := 0
	fmt.Printf("#1: Held breath for 0:00 min... press return to breathe.")
	for {
		select {
		case <-returnPressed:
			isBreathing = true
			if segment == 8 {
				fmt.Printf(
					"\r#%d: Held breath for %d:%02d min.                                             \n",
					segment, heldBreathFor/60, heldBreathFor%60)
				return
			}
		case <-secondTicker.C:
			secondsLeftInSegment--
			if !isBreathing {
				heldBreathFor++
				fmt.Printf(
					"\r#%d: Held breath for %d:%02d min... press return to breathe.",
					segment, heldBreathFor/60, heldBreathFor%60)
			} else {
				fmt.Printf(
					"\r#%d: Held breath for %d:%02d min. Wait %d:%02d min until the next breath hold.",
					segment,
					heldBreathFor/60, heldBreathFor%60,
					secondsLeftInSegment/60, secondsLeftInSegment%60)
			}
			if secondsLeftInSegment < 1 {
				if segment == 8 {
					fmt.Printf(
						"\r#%d: Held breath for %d:%02d min.                                             \n",
						segment, heldBreathFor/60, heldBreathFor%60)
					return
				}
				fmt.Printf(
					"\r#%d: Held breath for %d:%02d min, breathed for %d:%02d min.                   \n",
					segment,
					heldBreathFor/60, heldBreathFor%60,
					(4*60-heldBreathFor)/60, (4*60-heldBreathFor)%60)
				isBreathing = false
				secondsLeftInSegment = 4 * 60
				heldBreathFor = 0
				segment++
				fmt.Printf("#%d: Held breath for 0:00 min... press return to breathe.", segment)
			}
		}
	}
}

func reportReturnPresses(returnPressed chan bool) {
	for {
		fmt.Scanln()
		fmt.Print("\x1b[1A") // ANSI Escape code to move the cursor up one line.
		returnPressed <- true
	}
}

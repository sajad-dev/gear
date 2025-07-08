package developer

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func (a *AutoCompile) Loop() error {
	close := make(chan int)
	a.Run(close)
	if a.Test {
		a.RunTest()
	}

	color.Blue("Haha Babay run your project !!")
	for {
		chenge, err := a.Chenges()
		if err != nil {
			return err
		}

		if chenge {
			currentTime := time.Now()
			color.Blue(fmt.Sprintf("Last update at : %s", currentTime.Format("15:04:05")))

			close <- 1

			close = make(chan int)
			a.Run(close)
		}

		time.Sleep(time.Second )
	}
}

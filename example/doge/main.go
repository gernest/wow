package main

import (
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spinner"
)

func main() {
	w := wow.New(os.Stdout, spinner.Get(spinner.Dots), "Such Spins")
	w.Start()
	time.Sleep(2 * time.Second)
	w.Text("Very emojis").Spinner(spinner.Get(spinner.Hearts))
	time.Sleep(2 * time.Second)
	w.PersisWith(spinner.Spinner{
		Frames: []string{"👍"},
	}, " Wow!")

}

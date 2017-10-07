package main

import (
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spinner"
)

func main() {
	w := wow.New(os.Stdout, spinner.GetSpinner(spinner.Dots), "Such Spins")
	w.Start()
	time.Sleep(2 * time.Second)
	w.Text = "Very emojis"
	w.Spinner = spinner.GetSpinner(spinner.Hearts)
	time.Sleep(2 * time.Second)
	w.PersisWitht(spinner.Spinner{
		Frames: []string{"👍"},
	}, " Wow!")

}

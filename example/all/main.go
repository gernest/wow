package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spinner"
)

func main() {
	for _, v := range spinner.All {
		w := wow.New(os.Stdout, spinner.Get(v), " "+v.String())
		w.Start()
		time.Sleep(2)
		w.Persist()
		fmt.Print("\n")
	}
}

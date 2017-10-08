package main

import (
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

func main() {
	for _, v := range spin.All {
		w := wow.New(os.Stdout, spin.Get(v), " "+v.String())
		w.Start()
		time.Sleep(2)
		w.Persist()
	}
}

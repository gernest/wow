package main

import (
	"os"

	"github.com/gernest/wow"

	"github.com/gernest/wow/spinner"
)

func main() {
	w := &wow.Wow{}
	for _, v := range spinner.All {
		s := spinner.GetSpinner(v)
		w.AddLine(s, s.Name.String())
	}
	w.RenderTo(os.Stdout)
}

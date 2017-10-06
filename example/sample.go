package main

import (
	"os"

	"github.com/gernest/wow"

	"github.com/gernest/wow/spinner"
)

func main() {
	s := spinner.GetSpinner(spinner.Arrow)
	wow.Render(os.Stdout, s, " some fish")
}

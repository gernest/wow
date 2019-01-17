//DO NOT EDIT : this file was automatically generated.
package main

import (
	"os"
	"time"

	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

var all = []spin.Name{spin.Dots5, spin.Dots11, spin.Line2, spin.Arc, spin.Toggle2, spin.Star, spin.Circle, spin.Earth, spin.Toggle4, spin.Arrow2, spin.BouncingBar, spin.Monkey, spin.Dots8, spin.BoxBounce2, spin.SquareCorners, spin.Toggle11, spin.Moon, spin.Triangle, spin.Toggle10, spin.Grenade, spin.Dots2, spin.Balloon, spin.CircleQuarters, spin.BouncingBall, spin.Point, spin.Smiley, spin.Noise, spin.Arrow, spin.Arrow3, spin.Line, spin.Pipe, spin.SimpleDots, spin.SimpleDotsScrolling, spin.Star2, spin.Dqpb, spin.Layer, spin.Dots12, spin.Hamburger, spin.Bounce, spin.Weather, spin.Dots9, spin.GrowHorizontal, spin.Toggle3, spin.Hearts, spin.Shark, spin.Dots6, spin.BoxBounce, spin.CircleHalves, spin.Toggle, spin.Toggle5, spin.Dots7, spin.GrowVertical, spin.Toggle13, spin.Christmas, spin.Balloon2, spin.Toggle7, spin.Toggle9, spin.Clock, spin.Runner, spin.Dots10, spin.Flip, spin.Toggle6, spin.Dots, spin.Squish, spin.Toggle12, spin.Dots3, spin.Dots4, spin.Toggle8, spin.Pong}

func main() {
	for _, v := range all {
		w := wow.New(os.Stdout, spin.Get(v), " "+v.String())
		w.Start()
		time.Sleep(2)
		w.Persist()
	}
}

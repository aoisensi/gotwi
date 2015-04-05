package main

import (
	"io/ioutil"
	"log"

	"github.com/google/gxui"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := appInit(driver)
	window := theme.CreateWindow(800, 600, "gotwi")
	colomn := createColomn(theme)
	window.AddChild(colomn)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
	window.OnClose(driver.Terminate)
}

func appInit(driver gxui.Driver) gxui.Theme {
	theme := dark.CreateTheme(driver)
	if *font != "" {
		data, err := ioutil.ReadFile(*font)
		font, err := driver.CreateFont(data, 20)
		if err != nil {
			log.Print(err)
		} else {
			font.LoadGlyphs(32, 126)
			theme.SetDefaultFont(font)
		}
	}
	return theme
}

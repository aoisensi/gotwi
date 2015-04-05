package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/google/gxui"
)

func createTweetItem(theme gxui.Theme, tweet anaconda.Tweet) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.ExpandToContent)
	name := theme.CreateLabel()
	name.SetText(tweet.User.Name)
	layout.AddChild(name)
	text := theme.CreateLabel()
	text.SetText(tweet.Text)
	text.SetMultiline(true)
	layout.AddChild(text)
	return layout
}

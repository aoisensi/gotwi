package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type tweetAdapter struct {
	gxui.AdapterBase
	Tweets []anaconda.Tweet
}

func newTweetAdapter() *tweetAdapter {
	return &tweetAdapter{}
}

func (a *tweetAdapter) Count() int {
	return len(a.Tweets)
}

func (a *tweetAdapter) ItemAt(index int) gxui.AdapterItem {
	return index // This adapter uses integer indices as AdapterItems
}

func (a *tweetAdapter) ItemIndex(item gxui.AdapterItem) int {
	return item.(int) // Inverse of ItemAt()
}

func (a *tweetAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: 500, H: 100}
}

func (a *tweetAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	return createTweetItem(theme, a.Tweets[index])
}

package main

import (
	"log"

	"github.com/google/gxui"
)

func createColomn(theme gxui.Theme) gxui.Control {
	list := theme.CreateList()
	adapter := newTweetAdapter()
	tweets, err := apis[0].GetHomeTimeline(nil)
	if err != nil {
		log.Fatal(err)
	}

	adapter.Tweets = tweets
	list.SetOrientation(gxui.Vertical)
	list.SetAdapter(adapter)
	return list
}

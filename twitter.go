package main

import "github.com/ChimeraCoder/anaconda"

var apis []*anaconda.TwitterApi

func initTwitter() {
	anaconda.SetConsumerKey(config.Consumer.Key)
	anaconda.SetConsumerSecret(config.Consumer.Secret)
	apis = make([]*anaconda.TwitterApi, len(config.Access))
	for i, access := range config.Access {
		api := anaconda.NewTwitterApi(access.Token, access.Secret)
		apis[i] = api
	}
}

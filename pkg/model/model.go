package model

import "github.com/cvcio/twitter"

type TweetResponse struct {
	OK      bool            `json:"ok"`
	Message string          `json:"message"`
	Tweets  []twitter.Tweet `json:"tweets"`
}

package repository

import (
	"time"

	"github.com/cvcio/twitter"
)

type DatabaseRepo interface {
	InsertTweetsBySymbol(tweets *[]twitter.Tweet, Symbol string) error
	InsertTweetsWithEmotionsBySymbol(tweets *[]twitter.Tweet, Symbol string) error
	GetTweetsBySymbolByDate(Symbol string, start, end time.Time) ([]twitter.Tweet, error)
	GetAllTweetsWithEmotionsBySymbol(Symbol string) ([]twitter.Tweet, error)
}

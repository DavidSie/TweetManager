package dbrepo

import (
	"errors"
	"time"

	"github.com/cvcio/twitter"
)

// TriggerDBErrorSymbolOnTest causes errors in methods of testDBRepo type
const TriggerDBErrorSymbolOnTest = "TRIGGER-DB-ERROR"

func (t TestDBRepo) InsertTweetsBySymbol(tweets *[]twitter.Tweet, Symbol string) error {
	return nil
}
func (t TestDBRepo) InsertTweetsWithEmotionsBySymbol(tweets *[]twitter.Tweet, Symbol string) error {
	return nil
}
func (t TestDBRepo) GetTweetsBySymbolByDate(Symbol string, start, end time.Time) ([]twitter.Tweet, error) {
	if Symbol == TriggerDBErrorSymbolOnTest {
		return nil, errors.New("database Error")
	}
	return nil, nil
}
func (t TestDBRepo) GetAllTweetsWithEmotionsBySymbol(Symbol string) ([]twitter.Tweet, error) {
	if Symbol == TriggerDBErrorSymbolOnTest {
		return nil, errors.New("database Error")
	}
	return nil, nil
}

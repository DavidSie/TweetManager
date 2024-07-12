package dbrepo

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/cvcio/twitter"
)

// InsertTweetsBySymbol inserts tweets into database adding information about symbol, which can be company stock symbol or other important information
func (p postgresDBRepo) InsertTweetsBySymbol(tweets *[]twitter.Tweet, Symbol string) error {
	stmt := `insert into tweets (id, text, author_id, tweet_created_at, symbol, created_at, updated_at)
	values ( $1, $2, $3,$4, $5, $6, $7)`

	for _, tweet := range *tweets {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		tweetCreatedAt, err := tweet.CreatedAtTime()
		if err != nil {
			log.Printf("error while parsing tweet creation date %s: %v \n:", tweet.CreatedAt, err)
			continue
		}
		_, err = p.DB.ExecContext(ctx, stmt, tweet.ID, tweet.Text, tweet.AuthorID, tweetCreatedAt, Symbol, time.Now(), time.Now())
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			} else {
				log.Println("error for tweet", tweet, err)
				return err
			}
		}
	}
	return nil
}

// InsertTweetsWithEmotionsBySymbol insert tweets with emotions by a symbol, which can be company stock symbol or other important information
func (p postgresDBRepo) InsertTweetsWithEmotionsBySymbol(tweets *[]twitter.Tweet, Symbol string) error {
	stmt := `insert into tweets_with_emotions (id, text, author_id, tweet_created_at, created_at, updated_at)
	values ( $1, $2, $3,$4, $5, $6)`

	for _, tweet := range *tweets {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		tweetCreatedAt, err := tweet.CreatedAtTime()
		if err != nil {
			log.Printf("error while parsing tweet creation date %s: %v \n:", tweet.CreatedAt, err)
			continue
		}
		_, err = p.DB.ExecContext(ctx, stmt, tweet.ID, tweet.Text, tweet.AuthorID, tweetCreatedAt, time.Now(), time.Now())
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			} else {
				log.Println("error for tweet", tweet, err)
				return err
			}
		}
	}
	return nil
}

// GetAllTweetsBySymbol returns all tweets for a given symbol that have creation date within start and end date
func (p postgresDBRepo) GetTweetsBySymbolByDate(symbol string, start, end time.Time) ([]twitter.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tweets := []twitter.Tweet{}

	query := `
		select 
			id, text, author_id, tweet_created_at
		from 
			tweets
		where
			 symbol = $1
			 and $2 <= tweet_created_at 
			 and tweet_created_at <= $3
	`

	rows, err := p.DB.QueryContext(ctx, query, symbol, start, end)
	if err != nil {
		return tweets, err
	}
	for rows.Next() {
		var tweet twitter.Tweet
		err := rows.Scan(
			&tweet.ID,
			&tweet.Text,
			&tweet.AuthorID,
			&tweet.CreatedAt,
		)
		if err != nil {
			return tweets, err
		}
		tweets = append(tweets, tweet)
	}

	if err = rows.Err(); err != nil {
		return tweets, err
	}

	return tweets, nil
}

// GetAllTweetsWithEmotionsBySymbol returns all tweets with emotions for a given symbol
func (p postgresDBRepo) GetAllTweetsWithEmotionsBySymbol(symbol string) ([]twitter.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tweets := []twitter.Tweet{}

	query := `
		select 
			id, text, author_id, tweet_created_at
		from 
			tweets_with_emotions
		where
			symbol = $1
		
		
	`
	rows, err := p.DB.QueryContext(ctx, query, symbol)
	if err != nil {
		return tweets, err
	}
	for rows.Next() {
		var tweet twitter.Tweet
		err := rows.Scan(
			&tweet.ID,
			&tweet.Text,
			&tweet.AuthorID,
			&tweet.CreatedAt,
		)
		if err != nil {
			return tweets, err
		}
		tweets = append(tweets, tweet)
	}

	if err = rows.Err(); err != nil {
		return tweets, err
	}

	return tweets, nil
}

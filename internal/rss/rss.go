package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ahsanwtc/gator/internal/database"
	"github.com/google/uuid"
)

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error)  {
	requestWithContext, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		fmt.Println("error in forming the request")
		return nil, err
	}

	requestWithContext.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	response, err := client.Do(requestWithContext)
	if err != nil {
		fmt.Println("error in fetching the feeds")
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("error in reading the fetched data")
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		fmt.Println("error in unmarshaling the fetched data")
		return nil, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i:= 0; i < len(rssFeed.Channel.Item); i++ {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}

	return &rssFeed, nil
}

func ScrapeFeeds(ctx context.Context, db *database.Queries) error {
	nextFeed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Println("error in fetching feed data from the database")
		return err
	}

	err = db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		fmt.Println("error in updating the feed data: ", nextFeed.ID)
		return err
	}

	feeds, err := FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		fmt.Println("error in fetching feed for following url: ", nextFeed.Url)
		return err
	}

	for _, feed := range feeds.Channel.Item {
		_, err := db.CreatePost(ctx, database.CreatePostParams{
			ID: uuid.New(),
			Title: feed.Title,
			Description: feed.Description,
			Url: feed.Link,
			FeedID: nextFeed.ID,
			PublishedAt: parseDate(feed.PubDate),
		})

		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	fmt.Printf("Feed: `%s` added\n", feeds.Channel.Title)
	return nil
}

func parseDate(d string) sql.NullTime {
	timeLayouts := []string{
		time.RFC1123Z,           // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,            // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC3339,            // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,        // "2006-01-02T15:04:05.999999999Z07:00"
		time.RubyDate,           // "Mon Jan 02 15:04:05 -0700 2006"
		"Mon, 02 Jan 2006 15:04:05 MST", // custom variant
	}
	
	d = strings.TrimSpace(d)
	if d == "" {
		return sql.NullTime{}
	}

	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, d); err == nil {
			return sql.NullTime{Time: t.UTC(), Valid: true}
		}
	}
	return sql.NullTime{}
}
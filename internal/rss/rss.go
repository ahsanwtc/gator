package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
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
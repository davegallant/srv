package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/mmcdole/gofeed"
)

type RSS struct {
	Feeds []*gofeed.Feed
	c     *Controller
}

func (r *RSS) New(c *Controller) {
	r.c = c
}

// Update fetches all articles for all feeds
func (r *RSS) Update() {
	fp := gofeed.NewParser()
	var wg sync.WaitGroup
	var mux sync.Mutex
	r.Feeds = []*gofeed.Feed{}
	for _, f := range r.c.Config.Feeds {
		f := f
		wg.Add(1)
		go func() {
			defer wg.Done()
			feed, err := r.FetchURL(fp, f)
			if err != nil {
				log.Printf("error fetching url: %s, err: %v", f, err)
			}
			mux.Lock()
			if feed != nil {
				r.Feeds = append(r.Feeds, feed)
			}
			mux.Unlock()
		}()
	}
	wg.Wait()
}

// FetchURL fetches the feed URL and parses it
func (r *RSS) FetchURL(fp *gofeed.Parser, url string) (feed *gofeed.Feed, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	userAgent := browser.Firefox()

	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer func() {
			ce := resp.Body.Close()
			if ce != nil {
				err = ce
			}
		}()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Failed to get url %v, %v", resp.StatusCode, resp.Status)
	}

	return fp.Parse(resp.Body)
}

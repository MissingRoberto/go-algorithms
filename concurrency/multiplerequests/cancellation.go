package multiplerequests

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func makeRequestWithCancellationAndContext(ctx context.Context, url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(content))
}

// MaxInMultipleRequestsWithCancelation uses a separate channel for errors
func MaxInMultipleRequestsWithCancellation(urls []string) (int, error) {
	ch := make(chan int)
	errCh := make(chan error)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	for _, url := range urls {
		go func(url string) {
			code, err := makeRequestWithCancellationAndContext(ctx, url)
			if err != nil {
				errCh <- err
				return
			}
			ch <- code
		}(url)
	}

	var max int
	for i := 0; i < len(urls); i++ {
		select {
		case v := <-ch:
			if v > max {
				max = v
			}
		case <-errCh:
			return 0, errors.New("at least 1 failed")
		}
	}
	return max, nil
}

package multiplerequests

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func makeRequest(url string) (int, error) {
	resp, err := http.Get(url)
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

func makeRequestWithCancellation(url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
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

// MaxInMultipleRequestsSimple Fan-out algorithm
func MaxInMultipleRequestsSimple(urls []string) (int, error) {
	ch := make(chan int)

	for _, url := range urls {
		go func(url string) {
			code, _ := makeRequest(url)
			ch <- code
		}(url)
	}

	var max int
	for i := 0; i < len(urls); i++ {
		v := <-ch
		if v > max {
			max = v
		}
	}

	return max, nil
}

// MaxInMultipleRequestsWithErrors uses a separate channel for errors
func MaxInMultipleRequestsWithErrors(urls []string) (int, error) {
	ch := make(chan int)
	errCh := make(chan error)

	for _, url := range urls {
		go func(url string) {
			code, err := makeRequest(url)
			if err != nil {
				errCh <- err
				return
			}
			ch <- code
		}(url)
	}

	var max int
	var errors []error

	for i := 0; i < len(urls); i++ {
		select {
		case v := <-ch:
			if v > max {
				max = v
			}
		case err := <-errCh:
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return 0, fmt.Errorf("%d requests failed", len(errors))
	}

	return max, nil
}

// MaxInMultipleRequestsWithCancelation uses a separate channel for errors
func MaxInMultipleRequestsWithCancellation(urls []string) (int, error) {
	ch := make(chan int)
	errCh := make(chan error)

	for _, url := range urls {
		go func(url string) {
			code, err := makeRequestWithCancellation(url)
			if err != nil {
				errCh <- err
				return
			}
			ch <- code
		}(url)
	}

	var max int
	var errs []error

	for i := 0; i < len(urls); i++ {
		select {
		case v := <-ch:
			if v > max {
				max = v
			}
		case err := <-errCh:
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return 0, fmt.Errorf("%d requests failed", len(errs))
	}

	return max, nil
}

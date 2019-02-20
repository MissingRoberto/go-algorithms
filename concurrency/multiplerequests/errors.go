package multiplerequests

import (
	"fmt"
)

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

package multiplerequests

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

// MaxInMultipleRequestsWithErrorGroup uses errgroup to cancel all requests
func MaxInMultipleRequestsWithErrorGroup(urls []string) (int, error) {
	result := make([]int, len(urls))

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)
	for i, url := range urls {
		i, url := i, url // closure
		group.Go(func() error {
			code, err := makeRequestWithCancellationAndContext(ctx, url)
			if err != nil {
				return err
			}

			result[i] = code
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return 0, err
	}

	var max int
	for _, v := range result {
		if v > max {
			max = v
		}
	}

	return max, nil
}

package common

import (
	"math/rand"
	"time"
)

func Retry(attempts int, interval time.Duration, fun func() error) error {
	// avoid cold start and burst issues
	jitter := time.Duration(rand.Intn(200)) * time.Millisecond
	time.Sleep(jitter)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	var err error
	for i := attempts; i > 0; i-- {
		<-ticker.C
		attempts--
		err = fun()
		if err == nil {
			return nil
		}
	}
	return nil
}

func RetryGet[T any](attempts int, interval time.Duration, fun func() (*T, error)) (*T, error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	var err error
	var obj *T
	for i := attempts; i > 0; i-- {
		<-ticker.C
		attempts--
		obj, err = fun()
		if err == nil {
			return obj, nil
		}
	}
	// Return nil object if all attempts fail.
	return nil, err
}

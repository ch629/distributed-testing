package container

import (
	"fmt"
	"time"
)

type WaitForContainerOptions struct {
	HealthCheck func() bool
	Interval    time.Duration
	Retries     int
}

func WaitForContainer(options *WaitForContainerOptions) error {
	for i := 0; i < options.Retries; i++ {
		if options.HealthCheck() {
			return nil
		}
		time.Sleep(options.Interval)
	}
	return fmt.Errorf("failed to validate container running after %v tries", options.Retries)
}

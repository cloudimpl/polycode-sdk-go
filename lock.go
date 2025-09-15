package polycode

import "time"

type Lock interface {
	Acquire(expireIn time.Duration) error
	Release() error
}

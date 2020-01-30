package sync

import (
	"sync"
	"time"

	"github.com/sasha-s/go-deadlock"
)

func init() {
	deadlock.Opts.DeadlockTimeout = 10 * time.Second
	deadlock.Opts.OnPotentialDeadlock = func() {}
}

type Cond = sync.Cond
type Locker = sync.Locker
type Map = sync.Map
type Mutex = deadlock.Mutex
type Once = sync.Once
type Pool = sync.Pool
type RWMutex = deadlock.RWMutex
type WaitGroup = sync.WaitGroup

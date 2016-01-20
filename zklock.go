package main

import (
	"path"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/satori/go.uuid"
)

const (
	lockpath = "go-lock-"
)

type ZKMutex struct {
	c       *zk.Conn
	timeout time.Duration
	l       sync.Mutex
	ln      string // the locknode path
	guid    string
	hasLock bool
}

func NewZKMutex(zkservers []string, timeout time.Duration, lockpath string) (*ZKMutex, error) {
	c, _, err := zk.Connect(zkservers, time.Second)

	zkmutex := ZKMutex{}
	zkmutex.c = c
	zkmutex.timeout = timeout
	zkmutex.ln = lockpath

	if err != nil {
		return nil, err
	}

	return &zkmutex, nil
}

func (z *ZKMutex) Lock() error {
	z.l.Lock()
	defer z.l.Unlock()

	uuid := uuid.NewV4()
	fullpath := path.Join(z.ln, uuid.String())

	// create under lock node, ephemeral + sequqnce flags set, guid-lock-
	_, err := z.c.Create(fullpath, nil, zk.FlagEphemeral&zk.FlagSequence, nil)

	if err != nil {
		return err
	}

	// getChildren  without watch

	children, stat, err := z.c.Children(z.ln)

	if err != nil {
		return err
	}

	// if the pathname created in step 1 has the lowest sequence number suffix the client has the lock  and exits
	stat.

	// Client calls exiss with the watch flag set in the lock directory with the next lowest sequence number

	// If exists() returns false goto step 2. Otherweise wait for a notification for the pathname from the previous step before going t

	return nil
}

func (z *ZKMutex) Unlock() error {
	z.l.Lock()
	defer z.l.Unlock()

	return nil
}

func (z *ZKMutex) Close() error {
	return nil
}

func (z *ZKMutex) IsConnected() error {
	return nil
}

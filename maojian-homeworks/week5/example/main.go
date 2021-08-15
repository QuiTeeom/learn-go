package main

import (
	"errors"
	"fmt"
	slidingwindow "learn-go/maojian-homeworks/week5"
	"math/rand"
	"sync"
	"time"
)

func main() {
	l := NewLimiter(Rate{
		10, time.Second,
	})

	for i := 0; i < 11; i++ {
		go func() {
			for {
				err := l.Add()
				if err != nil {
					//	fmt.Printf("%s error\n", time.Now().Format("15:04:05.9999999"))
				} else {
					fmt.Printf("%s ok\n", time.Now().Format("15:04:05.9999999"))
				}
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}()
	}

	for {
		time.Sleep(time.Second)
	}
}

type Limiter struct {
	sw *slidingwindow.SlidingWindow
	m  sync.Mutex

	limit uint32
}

func (l *Limiter) Add() error {
	l.m.Lock()
	defer l.m.Unlock()
	c := l.sw.Get()
	if c >= l.limit {
		return errors.New("被限速啦～～")
	} else {
		l.sw.Add()
		return nil
	}
}

func NewLimiter(rate Rate) *Limiter {
	l := &Limiter{
		sw:    slidingwindow.NewWindow(rate.Unit),
		m:     sync.Mutex{},
		limit: rate.Count,
	}
	return l
}

type Rate struct {
	Count uint32
	Unit  time.Duration
}

var l = NewLimiter(Rate{
	10, time.Second,
})

func Add() error {
	err := l.Add()
	if err != nil {
		return err
	}
	return nil
}

package slidingwindow

import (
	"fmt"
	"sync"
	"time"
)

var (
	smoothingFactor = 10
)

type Bucket interface {
	Add()
	AddN(n uint32)
	Get() uint32
	Reset()
}

type bucket struct {
	count uint32
}

func (w *bucket) Add() {
	w.count++
}

func (w *bucket) Get() uint32 {
	return w.count
}

func (w *bucket) AddN(n uint32) {
	w.count += n
}

func (w *bucket) Reset() {
	w.count = 0
}

func NewBucket() Bucket {
	return &bucket{
		count: 0,
	}
}

type SlidingWindow struct {
	// smoothingFactor 平滑因子，即bucket的数量，越多越平滑
	factor int

	// start 开始时间
	start time.Time
	// size 窗口时间
	size time.Duration
	// 当前最大的时间
	max time.Time
	// sizePB 每个分片的时间
	sizePB int64

	// buckets 分片
	buckets []Bucket

	// current 当前分片index
	current int

	// preCount 之前分片的计数和
	preCount uint32

	// m 读写锁
	m sync.RWMutex
}

// NewWindow 创建滑动窗口
// size 窗口时间
func NewWindow(size time.Duration) *SlidingWindow {
	sw := &SlidingWindow{
		factor: smoothingFactor,
		size:   size,
	}
	sw.Init()
	return sw
}

func (w *SlidingWindow) Add() {
	w.m.Lock()
	defer w.m.Unlock()
	w.check()
	w.buckets[w.current].Add()
}

func (w *SlidingWindow) Get() uint32 {
	w.m.Lock()
	defer w.m.Unlock()
	w.check()
	return w.preCount + w.buckets[w.current].Get()
}

func (w *SlidingWindow) check() {
	if time.Now().After(w.max) {
		fmt.Printf("距离上次操作已经超过窗口时间 %d ms，进行 reset\n", w.size.Microseconds())
		w.reset()
		return
	}

	// 获取当前和时间差
	since := time.Since(w.start)
	if since < 0 {
		panic("")
	}
	// 计算move
	move := int(int64(since) / w.sizePB % int64(len(w.buckets)))

	if move > 0 {
		total := len(w.buckets)
		c := w.current
		for ; move > 0; move-- {
			w.preCount += w.buckets[c].Get()
			fi := c + 2
			if fi >= total {
				fi = fi - total
			}
			w.preCount = w.preCount - w.buckets[fi].Get()

			c++
			if c == total {
				c = c - total
			}
			w.buckets[c].Reset()

			w.max = w.max.Add(time.Duration(w.sizePB))
			w.start = w.start.Add(time.Duration(w.sizePB))
		}
		w.current = c
	}
}

func (w *SlidingWindow) Init() {
	buckets := make([]Bucket, smoothingFactor+1)
	for i := range buckets {
		buckets[i] = NewBucket()
	}
	w.buckets = buckets
	w.current = 0
	w.preCount = 0
	w.start = time.Now()
	w.max = time.Now().Add(w.size)
	w.sizePB = int64(w.size) / int64(w.factor)
}

func (w *SlidingWindow) reset() {
	for i := range w.buckets {
		w.buckets[i].Reset()
	}
	w.current = 0
	w.preCount = 0
	w.start = time.Now()
	w.max = time.Now().Add(w.size)
}

package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"strconv"
	"time"
)

func doSth(ctx *ReqCtx) {
	wg, c := errgroup.WithContext(ctx)

	max := rand.Intn(10)
	res := make([]string, max)
	for i := 0; i < max; i++ {
		wg.Go(func(slot int) func() error {
			return func() error {
				e, err := do(c, time.Duration(rand.Intn(8))*time.Second, strconv.Itoa(slot))
				if err != nil {
					return err
				}
				res[slot] = e
				logger.Info(e)
				return nil
			}
		}(i))
	}
	wg.Wait()
	b, _ := json.Marshal(res)
	ctx.writer.Write(b)
}

func do(ctx context.Context, d time.Duration, echo string) (string, error) {
	select {
	case <-time.After(d):
		if time.Now().Second()%10 == 4 {
			return "", errors.New("err lucky 4")
		}
		return fmt.Sprintf("after:%f -> echo %s", d.Seconds(), echo), nil
	case <-ctx.Done():
		return fmt.Sprintf("canceled:echo -> %s", echo), nil
	}
}

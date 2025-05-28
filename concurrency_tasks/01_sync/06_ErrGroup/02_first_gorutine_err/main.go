package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error { return task(ctx, "A", 4) })
	g.Go(func() error { return task(ctx, "B", 3) })
	g.Go(func() error { return task(ctx, "C", 1) })

	if err := g.Wait(); err != nil {
		fmt.Println("Завершено с ошибкой:", err)
	}
}

func task(ctx context.Context, name string, delay int) error {
	select {
	case <-ctx.Done():
		fmt.Printf("%s отменён\n", name)
		return ctx.Err()
	case <-time.After(time.Duration(delay) * time.Second):
		if delay == 1 { // Имитация ошибки
			return fmt.Errorf("%s: timeout", name)
		}
		fmt.Printf("%s успешно завершён\n", name)
		return nil
	}
}

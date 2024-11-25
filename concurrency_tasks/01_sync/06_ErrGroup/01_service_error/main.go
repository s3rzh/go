package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type Service struct {
	Name string
	Err  error
}

func main() {
	services := []Service{
		{
			Name: "service-1",
			Err:  nil,
		},
		{
			Name: "service-2",
			Err:  fmt.Errorf("service-2 error message"),
		},
    	{
			Name: "service-3",
			Err:  fmt.Errorf("service-3 error message also"),
		},
	}

	notify(services)
}

func notify(services []Service) {
	g := new(errgroup.Group)

	for _, service := range services {
		s := service // if go ver < 1.22
		g.Go(func() error {
			fmt.Printf("Starting to notifing %s\n", s.Name)
			if s.Err != nil {
				return s.Err
			}
			time.Sleep(5 * time.Second)
			fmt.Printf("Finished notifying %s\n", s.Name)

			return nil
		})
	}

	err := g.Wait() // обязательно дожидаемся выполнения всех горутин, даже если до этого уже есть ошибка.
	if err != nil { // те у нас есть возможность узнать о случившейся ошибке, причём только одной, первой вернувшейся, остальные недоступны (как с сервисом service-3)
		fmt.Printf("Error notifying services: %v\n", err)
		return
	}

	fmt.Println("All services notified!")
}

// Метод errorgroup.Group.Go() принимает функции, возвращающие ошибки, а метод errorgroup.Group.Wait() обрабатывает эти ошибки.
// Также стоит отметить, что использование пакета errgroup может привести к отличающимся результатам, поскольку возвращается первая обнаруженная ошибка. Каждый запуск горутин не гарантирует порядок выполнения, что одна будет выполняться раньше другой.

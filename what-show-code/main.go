package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// что показывает данный код
// что можно улучшить

func main() {
	fmt.Println(Do(context.Background(), []User{{Name: "aaa"}, {Name: "bbb"}, {Name: "ccc"}, {Name: "ddd"}, {Name: "eeee"}}))
}

type User struct {
	Name string
}

func getId(ctx context.Context) <-chan int {
	var ch chan int
	id := 1

	go func() {
		defer close(ch)
		time.Sleep(2000 * time.Millisecond)

		ch <- id
	}()

	return ch
}

func fetchByName(ctx context.Context, userName string) (int, error) {
	select {
	case <-ctx.Done():
		return 0, errors.New("timeout")
	case id := <-getId(ctx):
		return id, nil
	}
}

func Do(ctx context.Context, users []User) (map[string]int, error) {
	collected := make(map[string]int)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(len(users))

	for _, u := range users {

		go func(u User) {
			defer wg.Done()

			userID, err := fetchByName(ctx, u.Name)
			if err != nil {
				return
			}

			collected[u.Name] = userID
		}(u)
	}

	wg.Wait()
	for key, val := range collected {
		fmt.Printf("key: %s, value: %d ", key, val)
	}

	return collected, nil
}

// package whatshowcode

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// что показывает данный код
// что можно улучшить

// func main() {
// 	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eeee"}}))
// }

// type User struct {
// 	Name string
// }

// func fetchByName(ctx context.Context, userName string) (int, error) {
// 	time.Sleep(10 * time.Millisecond)

// 	return -1, nil
// }

// func Do(ctx context.Context, users []User) (map[string]int, error) {
// 	collected := make(map[string]int)
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(users))

// 	for _, u := range users {
// 		var collected map[string]int

// 		go func(u User) {
// 			defer wg.Done()
// 			userID, err := fetchByName(ctx, u.Name)
// 			if err != nil {
// 				return
// 			}

// 			collected[u.Name] = userID
// 		}(u)
// 	}

// 	wg.Wait()

// 	return collected, nil
// }

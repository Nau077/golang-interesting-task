package whatshowcode

import (
	"context"
	"fmt"
	"time"
)

// что показывает данный код
// что можно улучшить

func main() {
	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eeee"}}))
}

type User struct {
	Name string
}

func fetchByName(ctx context.Context, userName string) (int, error) {
	time.Sleep(10 * time.Millisecond)

	return -1, nil
}

func Do(ctx context.Context, users []User) (map[string]int, error) {
	collected := make(map[string]int)

	for _, u := range users {
		userID, err := fetchByName(ctx, u.Name)
		if err != nil {
			return collected, err
		}
		collected[u.Name] = userID
	}

	return collected, nil
}

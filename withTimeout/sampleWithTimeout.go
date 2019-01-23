package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	// タイムアウト付きのコンテキストを渡して、タイムアウトが経過した後にコンテキストをその機能を放棄することをブロッキング関数に伝えます。
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond) // 50msたったらタイムアウト
	defer cancel()

	select {
	case <-time.After(1 * time.Second): // 1s経ったら流れてくる。しかしctxは50msでキャンセルされるのでここは実行されない。
		fmt.Println("overslept")
	case <-ctx.Done(): //チャネルがキャンセルされたら実行される。
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}

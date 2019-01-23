package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // sync.WaitGroupによってゴルーチンの完了を検知できる
	done := make(chan interface{})
	defer close(done) //確実にチャンネルを閉じる

	wg.Add(1) //１つゴルーチンが起動したことをwgに登録

	go func() {
		defer wg.Done() //ゴルーチンの終了をwgに伝える
		if err := printGreeting(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Add(1) //もう一つゴルーチンが起動したことをwgに登録
	go func() {
		defer wg.Done() //ゴルーチンの終了をwgにつたあえる
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()
	wg.Wait() // すべてのゴルーチンが終了したと伝わるまでメインゴルーチンをブロック。

}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(10 * time.Second):
	}
	return "EN/US", nil
}

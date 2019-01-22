package main

import (
	"context"
	"fmt"
)

func main() {

	processRequest("jane", "abc123")

}

func processRequest(userID, authToken string) {
	// キーはGoでの比較可能性を満たさなければ行けない等値演算子==と!=が使える。
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	handleResponse(ctx)
}

func handleResponse(ctx context.Context) {
	// 返された値は複数のゴルーチンからアクセスされても安全でなければいけない
	fmt.Printf(
		"handling response for %v  (%v)",
		ctx.Value("userID"),
		ctx.Value("authToken"),
	)
}

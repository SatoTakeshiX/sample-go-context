package main

import (
	"context"
	"fmt"
)

func main() {

	processRequest("jane", "abc123")
}

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func userID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func authToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func processRequest(userID, authToken string) {
	// キーはGoでの比較可能性を満たさなければ行けない等値演算子==と!=が使える。
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	handleResponse(ctx)
}

func handleResponse(ctx context.Context) {
	// 返された値は複数のゴルーチンからアクセスされても安全でなければいけない
	fmt.Printf(
		"handling response for %v  (%v)",
		userID(ctx),
		authToken(ctx),
	)
}

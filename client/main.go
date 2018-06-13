package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"../pb"
)

func main() {
	// 連線到遠端 gRPC 伺服器。
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	// 建立新的 UserService 客戶端，所以等一下就能夠使用 UserService 的所有方法。
	c := pb.NewUserServiceClient(conn)

	// 傳送新請求到遠端 gRPC 伺服器 UserService 中，並呼叫 Get 函式，取得會員資料。
	r, err := c.Get(context.Background(), &pb.UserGetRequest{Id: 1})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s", r.Username)
}

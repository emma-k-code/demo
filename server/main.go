package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"../pb"
)

type user struct {
	Id       int
	Username string
}

var userData = [3]user{
	user{Id: 1, Username: "會員1"},
	user{Id: 2, Username: "會員2"},
	user{Id: 3, Username: "會員3"},
}

// server 實作 UserService 的 gRPC 伺服器。
type server struct{}

// Get 將傳入的會員ID資料回傳
func (s *server) Get(ctx context.Context, in *pb.UserGetRequest) (*pb.UserResponse, error) {
	// 會員ID
	id := in.Id
	name := "查無會員"
	users := userData

	if id <= 3 {
		name = users[id-1].Username
	}

	// 包裝成 Protobuf 建構體並回傳。
	return &pb.UserResponse{Id: id, Username: name}, nil
}

func main() {
	// 監聽指定埠口，這樣服務才能在該埠口執行。
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("無法監聽該埠口：%v", err)
	}

	// 建立新 gRPC 伺服器並註冊 UserService 服務。
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(s)

	// 開始在指定埠口中服務。
	if err := s.Serve(lis); err != nil {
		log.Fatalf("無法提供服務：%v", err)
	}
}

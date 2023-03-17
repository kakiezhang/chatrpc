package chat

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type ChatService struct {
	UnimplementedChatServiceServer
}

// grpcurl usage:
// grpcurl -plaintext -format json -d @ localhost:9000 chat.ChatService/SayHello
// grpcurl -plaintext -import-path . -proto chat.proto localhost:9000 list (could use without reflection.Register)
func (*ChatService) SayHello(ctx context.Context, req *SayRequest) (resp *SayResponse, err error) {
	log.Printf("req: %+v", req)
	resp = &SayResponse{
		Body: "Hakuku",
		Ts:   100001,
	}
	return
}

func SayHi(w http.ResponseWriter, r *http.Request) {
	log.Printf("req: %+v", r.Header)
	fmt.Fprintf(w, "%s", "helo")
}

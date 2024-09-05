package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {
	// Ollama 클라이언트 생성
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Ollama 클라이언트 생성 실패: %v", err)
	}
	// 컨텍스트 생성
	ctx := context.Background()

	// 서버 상태 확인
	if err := client.Heartbeat(ctx); err != nil {
		log.Fatalf("서버 연결 실패: %v", err)
	}
	fmt.Println("서버 연결 성공")

	// 버전 확인
	version, err := client.Version(ctx)
	if err != nil {
		log.Fatalf("버전 확인 실패: %v", err)
	}
	fmt.Printf("Ollama 버전: %s\n", version)

	// 모델 리스트 가져오기
	models, err := client.List(ctx)
	if err != nil {
		log.Fatalf("모델 리스트 가져오기 실패: %v", err)
	}
	fmt.Println("사용 가능한 모델:")
	for _, model := range models.Models {
		fmt.Printf("- %s\n", model.Name)
	}

	// 텍스트 생성
	generateReq := &api.GenerateRequest{
		Model:  "llama3",
		Prompt: "Go 언어의 장점을 한국어로 설명해주세요.",
	}
	err = client.Generate(ctx, generateReq, func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		return nil
	})
	if err != nil {
		log.Fatalf("텍스트 생성 실패: %v", err)
	}
	fmt.Println()

	// 채팅
	chatReq := &api.ChatRequest{
		Model: "llama3",
		Messages: []api.Message{
			{Role: "user", Content: "Go 언어로 'Hello, World!'를 출력하는 코드를 작성해주세요."},
		},
	}
	err = client.Chat(ctx, chatReq, func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	})
	if err != nil {
		log.Fatalf("채팅 실패: %v", err)
	}
	fmt.Println()

	// 임베딩 생성
	embedReq := &api.EmbeddingRequest{
		Model:  "llama3",
		Prompt: "Go 언어",
	}
	embedResp, err := client.Embeddings(ctx, embedReq)
	if err != nil {
		log.Fatalf("임베딩 생성 실패: %v", err)
	}
	fmt.Printf("임베딩 벡터 (처음 5개 요소): %v\n", embedResp.Embedding[:5])

	// 모델 정보 보기
	showReq := &api.ShowRequest{
		Name: "llama3",
	}
	showResp, err := client.Show(ctx, showReq)
	if err != nil {
		log.Fatalf("모델 정보 가져오기 실패: %v", err)
	}
	fmt.Printf("모델 정보:\n- 이름: %s\n", showResp.Modelfile)
}

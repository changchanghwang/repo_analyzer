package llm

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

type LLMClient struct {
	*api.Client
}

func NewLLMClient() *LLMClient {
	// Ollama 클라이언트 생성
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("Ollama 클라이언트 생성 실패: %v", err)
	}

	return &LLMClient{client}
}

func (c *LLMClient) Search(query string) (string, error) {
	// 컨텍스트 생성
	ctx := context.Background()

	// 서버 상태 확인
	if err := c.Heartbeat(ctx); err != nil {
		log.Fatalf("서버 연결 실패: %v", err)
	}
	fmt.Println("서버 연결 성공")

	// 버전 확인
	version, err := c.Version(ctx)
	if err != nil {
		log.Fatalf("버전 확인 실패: %v", err)
	}
	fmt.Printf("Ollama 버전: %s\n", version)

	// 텍스트 생성
	generateReq := &api.GenerateRequest{
		Model:  "llama3",
		Prompt: query,
	}

	r := ""
	c.Generate(ctx, generateReq, func(resp api.GenerateResponse) error {
		for _, c := range resp.Response {
			r += string(c)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("텍스트 생성 실패: %v", err)
	}

	return r, nil
}

func (c *LLMClient) ModelList(prompt string) error {
	ctx := context.Background()

	// 모델 리스트 가져오기
	models, err := c.List(ctx)
	if err != nil {
		log.Fatalf("모델 리스트 가져오기 실패: %v", err)
	}
	fmt.Println("사용 가능한 모델:")
	for _, model := range models.Models {
		fmt.Printf("- %s\n", model.Name)
	}

	return nil

}

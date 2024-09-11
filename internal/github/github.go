package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type SearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		FullName        string `json:"full_name"`
		Description     string `json:"description"`
		HtmlURL         string `json:"html_url"`
		ForksCount      int    `json:"forks_count"`
		OpenIssuesCount int    `json:"open_issues_count"`
		StargazersCount int    `json:"stargazers_count"`
	} `json:"items"`
}

type GithubClient struct {
	*http.Client
	apiURL string
	token  string
}

func NewGithubClient() *GithubClient {
	return &GithubClient{
		Client: &http.Client{},
		apiURL: "https://api.github.com/search/repositories",
		token:  os.Getenv("GITHUB_TOKEN"),
	}
}

func (client *GithubClient) SearchGitHub(query string) (*SearchResult, error) {
	// URL 인코딩
	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s?q=%s&per_page=5", client.apiURL, encodedQuery)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	// GitHub 개인 액세스 토큰 사용 (선택사항이지만 API 제한을 늘릴 수 있습니다)
	req.Header.Set("Authorization", "token "+client.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected error: %s", string(body))
	}

	var result SearchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

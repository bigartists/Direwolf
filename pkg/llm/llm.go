package llm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client interface {
	SendQuery(apiKey, baseURL, model, query string, params map[string]interface{}, responseChan chan<- string) error
	//SendQueryGetContent(apiKey, baseURL, maas, query string, responseChan chan<- string) error
	GetContentFromChunk(data string) (string, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) SendQuery(apiKey, baseURL, model, query string, params map[string]interface{}, responseChan chan<- string) error {

	url := fmt.Sprintf("%s/chat/completions", baseURL)

	fmt.Println("url=: ", url, "apiKey: ", apiKey, "maas: ", model, "query: ", query, "params: ", params)

	requestBody, err := json.Marshal(params)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Accept", "text/event-stream") // Set Except header for SSE

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Body)
		// 打印报错具体的信息
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return c.handleStreamResponse(resp.Body, responseChan)
}

func (c *client) handleStreamResponse(body io.ReadCloser, responseChan chan<- string) error {
	reader := bufio.NewReader(body)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "data: ") {
			if line == "data: [DONE]" {
				responseChan <- line // 发送 DONE 消息
				break
			}

			responseChan <- line
		}
	}

	close(responseChan)
	return nil
}

func (c *client) GetContentFromChunk(data string) (string, error) {
	// 首先去除 "data: " 前缀
	data = strings.TrimPrefix(strings.TrimSpace(data), "data: ")

	// 如果是 [DONE] 信号，直接返回空内容
	if data == "[DONE]" {
		return "", nil
	}

	// 解析 JSON
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", nil
	}

	delta, ok := choices[0].(map[string]interface{})["delta"].(map[string]interface{})
	if !ok {
		return "", nil
	}

	content, ok := delta["content"].(string)
	if !ok {
		return "", nil
	}

	return content, nil
}

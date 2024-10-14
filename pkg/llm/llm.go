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
	SendQueryGetContent(apiKey, baseURL, model, query string, responseChan chan<- string) error
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) SendQuery(apiKey, baseURL, model, query string, params map[string]interface{}, responseChan chan<- string) error {

	url := fmt.Sprintf("%schat/completions", baseURL)

	requestBody, err := json.Marshal(params)

	//requestBody, err := json.Marshal(map[string]interface{}{
	//	"model": model,
	//	"messages": []map[string]string{
	//		{
	//			"role":    "user",
	//			"content": query,
	//		},
	//	},
	//	"stream": true, // Enable streaming
	//})
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
		//if line == "" {
		//	continue
		//}

		// 如果行不为空，发送整行数据
		if line != "" {
			responseChan <- line
		}

		// 如果收到 [DONE]，结束循环但仍然发送这个标记
		if strings.TrimPrefix(line, "data: ") == "[DONE]" {
			break
		}
	}

	close(responseChan)
	return nil
}

func (c *client) SendQueryGetContent(apiKey, baseURL, model, query string, responseChan chan<- string) error {
	url := fmt.Sprintf("%schat/completions", baseURL)

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": query,
			},
		},
		"stream": true, // Enable streaming
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Accept", "text/event-stream") // Set Accept header for SSE

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	reader := bufio.NewReader(resp.Body)
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
			data := strings.TrimPrefix(line, "data: ")
			if data == "[DONE]" {
				break
			}

			var result map[string]interface{}
			if err := json.Unmarshal([]byte(data), &result); err != nil {
				return err
			}

			choices, ok := result["choices"].([]interface{})
			if !ok || len(choices) == 0 {
				continue
			}

			delta, ok := choices[0].(map[string]interface{})["delta"].(map[string]interface{})
			if !ok {
				continue
			}

			content, ok := delta["content"].(string)
			if !ok {
				continue
			}

			responseChan <- content
		}
	}

	close(responseChan)
	return nil
}

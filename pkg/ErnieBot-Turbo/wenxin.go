package wenxin

import (
	"encoding/json"
	"fmt"
	"gotalk/pkg/conf"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	Messages     []Messages `json:"messages"`
	PenaltyScore *float64   `json:"penalty_score,omitempty"`
	Stream       *bool      `json:"stream,omitempty"`
	Temperature  *float64   `json:"temperature,omitempty"`
	TopP         *float64   `json:"top_p,omitempty"`
	UserID       string     `json:"user_id"`
}

var Payload Request

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

const ()

var (
	client_secret = conf.Configs.BceAPP.ClientSecret
	client_id     = conf.Configs.BceAPP.ClientID
)

func init() {
}

var Jsmess []byte

func DoChat(content string) string {
	um := &Messages{Content: content, Role: "user"}
	Payload.Messages = append(Payload.Messages, *um)
	Jsmess, _ = json.Marshal(Payload)

	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant?access_token=" + GetAccessToken()
	payload := strings.NewReader(string(Jsmess))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	chatObj := map[string]string{}
	json.Unmarshal([]byte(body), &chatObj)
	am := &Messages{Content: chatObj["result"], Role: "assistant"}
	Payload.Messages = append(Payload.Messages, *am)
	return chatObj["result"]
}

func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", client_id, client_secret)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}

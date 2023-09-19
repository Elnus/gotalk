package wenxin

import (
	"encoding/json"
	"fmt"
	"gotalk/pkg/conf"
	"io"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	Message      []Message `json:"message"`
	PenaltyScore *float64  `json:"penalty_score,omitempty"`
	Stream       *bool     `json:"stream,omitempty"`
	Temperature  *float64  `json:"temperature,omitempty"`
	TopP         *float64  `json:"top_p,omitempty"`
	UserID       string    `json:"user_id"`
}

var Payload Request

type TokenBody struct {
	RefreshToken  string `json:"refresh_token"`
	ExpireIn      int    `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

var TokenBodys TokenBody

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

const (
	tokenApi = "https://aip.baidubce.com/oauth/2.0/token"
	chatApi  = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant"
)

// var request Request
var (
	access_token  = ""
	client_secret = conf.Configs.BceAPP.ClientSecret
	client_id     = conf.Configs.BceAPP.ClientID
	charUrl       = ""
	tokenUrl      = ""
)

func init() {
	//access_token = GetToken()
	charUrl = fmt.Sprintf("%s?access_token=%s", chatApi, access_token)
	tokenUrl = fmt.Sprintf("%s?grant_type=client_credentials&client_id=%s&client_secret=%s", tokenApi, client_id, client_secret)
}

func DoChat() []byte {

	method := "POST"

	payload := strings.NewReader(`{
		"messages": [
			{
				"role": "user",
				"content": "hello"
			}
		]
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, charUrl, payload)

	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return body
}

func GetToken() {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, tokenUrl, nil)

	if err != nil {
		log.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal(body, &TokenBodys)
}

package cryptorg_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/chenzhuoyu/base64x"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string

	key    string
	secret string
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseUrl:    "https://api2.cryptorg.net",
		key:        os.Getenv("API_KEY"),
		secret:     os.Getenv("API_SECRET"),
	}
}

func (c *Client) GetBotDetails(botId int64) (*[]BotDetail, error) {
	uri := "/bot-futures/detail"
	query := "botId=" + strconv.FormatInt(botId, 10)

	body, err := c.sendRequest(uri, query, "GET")
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	//fmt.Printf("%v+\n", string(body))

	var resp BotDetailsResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Printf("Unmarshal error: %s\n", err)
	}

	if resp.Success == false {
		return nil, fmt.Errorf("Cryptorg error: %s\n", resp.ErrorMessage)
	}

	return &resp.Data, nil
}

// sendRequest - send sendRequest to cryptorg api
//
//	uri 	- must have leader slash "/bot-futures/detail"
//	query 	- for example "botId=303017&anotherParam=123"
//	method 	- "GET" or "POST"
func (c *Client) sendRequest(uri string, query string, method string) ([]byte, error) {
	url := "https://api2.cryptorg.net" + uri + "?" + query
	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	hash := generateSignature(uri, nonce, query, c.secret)

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Error creating sendRequest:", err)
		return nil, err
	}

	request.Header.Set("CTG-API-SIGNATURE", hash)
	request.Header.Set("CTG-API-KEY", c.key)
	request.Header.Set("CTG-API-NONCE", nonce)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error sending sendRequest:", err)
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	fmt.Printf("%v+\n", string(body))

	return body, nil
}

func generateSignature(uri string, nonce string, query string, secret string) string {
	data := uri + "/" + nonce + "/" + query
	encodedData := base64x.StdEncoding.EncodeToString([]byte(data))

	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(encodedData))
	signature := h.Sum(nil)

	//fmt.Printf("%x\n", signature)

	return hex.EncodeToString(signature)
}

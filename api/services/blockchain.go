package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"fmt"
)

// Result of `validateaddress` call
type ValidateAddressInfo struct {
	IsValid   bool   `json:"isvalid"`
	Beaconed  byte   `json:"isbeaconed"`  // why isbeaconed is a number in meritd???
	Confirmed byte   `json:"isconfirmed"` // why isconfirmed is a number in meritd???
	Address   string `json:"address"`
	PublicKey string `json:"pubkey"`
	Alias     string `json:"alias"`
}

// ValidateAddress returns information about given address from the blockchain
func (c *Client) ValidateAddress(address string) (ValidateAddressInfo, error) {
	var info ValidateAddressInfo
	err := c.exec("validateaddress", &info, address)

	return info, err
}

type rawRequest struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

type rawResponse struct {
	Result json.RawMessage `json:"result"`
	Error  string          `json:"error"`
}

type logger interface {
	Printf(format string, v ...interface{})
}

// Client is a blockchain http client
type Client struct {
	url      string
	user     string
	password string
	l        logger
}

// NewClient create new blockchain http client object
func NewClient(url, user, password string, l logger) Client {
	return Client{
		url,
		user,
		password,
		l,
	}
}

func makeRequest(method string, params []interface{}) (*rawRequest, error) {
	// Method may not be empty.
	if method == "" {
		return nil, errors.New("method can not be empty")
	}

	rawParams := make([]json.RawMessage, 0, len(params))
	for _, param := range params {
		marshalledParam, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		rawMessage := json.RawMessage(marshalledParam)
		rawParams = append(rawParams, rawMessage)
	}

	rawRequest := &rawRequest{
		Jsonrpc: "1.0",
		Method:  method,
		Params:  rawParams,
	}

	return rawRequest, nil
}

func (c *Client) exec(method string, res interface{}, params ...interface{}) error {
	rawRequest, err := makeRequest(method, params)
	if err != nil {
		return err
	}

	marshalledJSON, err := json.Marshal(rawRequest)

	fmt.Printf("URL: %+v\n", c.url)

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(marshalledJSON))
	req.SetBasicAuth(c.user, c.password)

	fmt.Printf("REQ %+v\n", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Printf("RESP: %+v ERR: %+v\n", resp, err)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var out rawResponse
	if err := json.Unmarshal(bodyBytes, &out); err != nil {
		fmt.Printf("ERR: %+v\n", err);
		return err
	}
	fmt.Printf("OUT: %+v\n ERR: %+v\n", out, err);
	
	if out.Error != "" {
		return errors.New(out.Error)
	}

	if err := json.Unmarshal(out.Result, &res); err != nil {
		return err
	}

	return nil
}

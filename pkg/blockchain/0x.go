package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SwapQuoteResponse struct {
	Price            string `json:"price"`
	GuaranteedPrice  string `json:"guaranteedPrice"`
	To               string `json:"to"`
	Data             string `json:"data"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice"`
	Gas              string `json:"gas"`
	BuyTokenAddress  string `json:"buyTokenAddress"`
	SellTokenAddress string `json:"sellTokenAddress"`
	BuyAmount        string `json:"buyAmount"`
	SellAmount       string `json:"sellAmount"`
	AllowanceTarget  string `json:"allowanceTarget"`
	// Add more fields as needed
}

func GetSwapQuote(sellToken, buyToken, amount string, isSellAmount bool) (*SwapQuoteResponse, error) {
	apiKey := os.Getenv("OX_API_KEY")
	feeRecipient := os.Getenv("FEE_RECIPIENT")
	buyTokenPercentageFee := os.Getenv("BUY_TOKEN_PERCENTAGE_FEE")

	if apiKey == "" || feeRecipient == "" || buyTokenPercentageFee == "" {
		log.Fatal("Environment variables OX_API_KEY, FEE_RECIPIENT, or BUY_TOKEN_PERCENTAGE_FEE are not set.")
	}

	var url string
	if isSellAmount {
		url = fmt.Sprintf("https://api.0x.org/swap/v1/quote?sellToken=%s&buyToken=%s&sellAmount=%s&feeRecipient=%s&buyTokenPercentageFee=%s",
			sellToken, buyToken, amount, feeRecipient, buyTokenPercentageFee)
	} else {
		url = fmt.Sprintf("https://api.0x.org/swap/v1/quote?sellToken=%s&buyToken=%s&buyAmount=%s&feeRecipient=%s&buyTokenPercentageFee=%s",
			sellToken, buyToken, amount, feeRecipient, buyTokenPercentageFee)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("0x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("0x API request failed with status: %d, body: %s", resp.StatusCode, string(body))
	}

	var quoteResponse SwapQuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		return nil, err
	}

	return &quoteResponse, nil
}

package module

import (
	"conf"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

type txTransaction struct {
	From 	string  `json:"from"`
	To  	string  `json:"to"`
	Value 	string  `json:"value"`
	Data    string  `json:"data"`
}

func SendGeth(data string) string{
	client, err := rpc.Dial(conf.GETH_URL)
	if err != nil {
	   fmt.Println("rpc.Dial err", err)
	   return ""
	}

	var account[]string
	err = client.Call(&account, "eth_accounts")
	if err != nil {
	   fmt.Println("account", err)
	   return ""
	}
	a := txTransaction{
		From:account[0],
		To:account[1],
		Value:"0x9184e72a",
		Data:data,
	}
	var result interface{}
	//var result hexutil.Big
	//err = client.Call(&result, "eth_getBalance", account[0], "latest")
	err = client.Call(&result, "eth_sendTransaction", a)
	if err != nil {
	   fmt.Println("eth_sendTransaction", err)
	   return ""
	}

	if err != nil {
	   fmt.Println("client.Call err", err)
	   return ""
	}

	fmt.Printf("account[0]: %s\n result: %s\n", account[0], result.(string))
	return result.(string)
}

func GetGeth(hash string) {
	client, err := rpc.Dial(conf.GETH_URL)
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}
	var account[]string
	err = client.Call(&account, "eth_accounts")
	if err != nil {
		fmt.Println("account", err)
		return
	}
	var result interface{}
	err = client.Call(&result, "eth_getTransactionByHash",hash)
	if err != nil {
		fmt.Println("eth_getTransactionByHash", err)
		return
	}

	if err != nil {
		fmt.Println("client.Call err", err)
		return
	}

	fmt.Printf("account[0]: %s\n result[0]: %s\n", account[0], result)
}


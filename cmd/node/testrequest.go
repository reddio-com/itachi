package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"itachi/evm"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

// How To Use
// Add this line in main.go (before `app.StartupChain()`)
// go testSendTransaction(gethCfg, true)

func testEthCall(exit bool) {
	time.Sleep(5 * time.Second)
	requestBody := `{
		"jsonrpc": "2.0",
		"id": 0,
		"method": "eth_call",
		"params": [{
			"from": "0x123456789abcdef123456789abcdef123456789a",
			"to": "0x9d7bA953587B87c474a10beb65809Ea489F026bD",
			"data": "0x70a082310000000000000000000000006E0d01A76C3Cf4288372a29124A26D4353EE51BE"
		}, "latest"]
	}`
	sendRequest(requestBody)

	if exit {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
}

func testSendTransaction(gethCfg *evm.GethConfig, exit bool) {
	// A random private key. address = 0x7Bd36074b61Cfe75a53e1B9DF7678C96E6463b02
	privateKeyStr := "32e3b56c9f2763d2332e6e4188e4755815ac96441e899de121969845e343c2ff"
	nonce := uint64(0)
	to := common.HexToAddress("0x7Bd36074b61Cfe75a53e1B9DF7678C96E6463b02")
	amount := big.NewInt(10000000000000000)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(3000000000000)
	data := []byte{}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &to,
		Value:    amount,
		Data:     data,
	})

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	//signer := types.MakeSigner(gethCfg, new(big.Int).SetUint64(uint64(block.Height)), block.Timestamp)

	chainID := gethCfg.ChainConfig.ChainID
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SignedTx = %+v", signedTx)

	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	requestBody := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"id": 0,
		"method": "eth_sendRawTransaction",
		"params": ["0x%x"] 
	}`, rawTxBytes)

	sendRequest(requestBody)

	if exit {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
}

func testSendRawTransaction(exit bool) {
	time.Sleep(5 * time.Second)

	requestBody := `{
		"jsonrpc": "2.0",
		"id": 0,
		"method": "eth_sendRawTransaction",
		"params": ["0xf86c808506fc23ac00825208947bd36074b61cfe75a53e1b9df7678c96e6463b02880de0b6b3a76400008026a0b5050757a8005286d85c8ae9408a933ca1126400a6749ca64e415f77db41b439a0294f3d15727ed8231d061db0ec1014ef1bf767f1665731d62e0327464cf8ad3e"] 
	}`

	sendRequest(requestBody)

	if exit {
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
}

func sendRequest(dataString string) {
	req, err := http.NewRequest("POST", "http://localhost:9092", bytes.NewBuffer([]byte(dataString)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("curl --location 'localhost:9092' --header 'Content-Type: application/json' --data '%s'\n", dataString)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	log.Printf("Response [%v] : %v", resp.Status, string(body))
	return
}

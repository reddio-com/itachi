package l1

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"itachi/cairo/config"
	"itachi/cairo/l1/contract"
	"math/big"
	"time"
)

type MyEthClient struct {
	cfg       *config.Config
	ethClient *ethclient.Client
	client    *rpc.Client
	filterer  *contract.StarknetCoreFilterer
}

func NewEthClient(ethClientAddress string) (*ethclient.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// TODO replace with our own client once we have one.
	// Geth pulls in a lot of dependencies that we don't use.
	client, err := rpc.DialContext(ctx, ethClientAddress)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(client), nil
}

func CreateAuth(client *ethclient.Client, privateKeyHex string, address string, gasLimit uint64, chainID *big.Int) (*bind.TransactOpts, error) {
	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// Ensure private key string has 0x prefix
	if !has0xPrefix(privateKeyHex) {
		privateKeyHex = "0x" + privateKeyHex
	}

	// Decode private key
	rawPrivateKey, err := hexutil.Decode(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Convert to ECDSA private key
	privateKey, err := crypto.ToECDSA(rawPrivateKey)
	if err != nil {
		return nil, err
	}

	// Create auth object
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // 0 value for no ether transfer
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

// Helper function to check if string has 0x prefix
func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0:2] == "0x"
}

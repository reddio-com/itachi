package ethrpc

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	yucommon "github.com/yu-org/yu/common"
	yucore "github.com/yu-org/yu/core"
	"itachi/evm"
	"log"
	"math/big"
)

// Functions from goeth -> ethclient -> ethclient.go

type CallMsg struct {
	From        *common.Address // the sender of the 'transaction'
	To          *common.Address // the destination contract (nil for contract creation)
	Gas         uint64          // if 0, the call executes with near-infinite gas
	GasPrice    *big.Int        // wei <-> gas exchange ratio
	GasFeeCap   *big.Int        // EIP-1559 fee cap per gas.
	GasTipCap   *big.Int        // EIP-1559 tip per gas.
	Value       *big.Int        // amount of wei sent along with the call
	Data        string          // input data, usually an ABI-encoded contract method invocation
	BlockNumber *big.Int        // nil for latest
}

// Sample Request
// curl --location 'localhost:9092' \
// --header 'Content-Type: application/json' \
//
//	--data '{
//	   "jsonrpc": "2.0",
//	   "id": 0,
//	   "method":"eth_chainId"
//	}'
func (s *EthRPC) ChainId(ctx context.Context) (*big.Int, error) {
	log.Printf("GetChainID")
	return s.cfg.ChainConfig.ChainID, nil
}

// Sample Request
// curl --location 'localhost:9092' \
// --header 'Content-Type: application/json' \
//
//	--data '{
//	   "jsonrpc": "2.0",
//	   "id": 0,
//	   "method": "eth_call",
//	   "params": [{
//	       "from": "0x123456789abcdef123456789abcdef123456789a",
//	       "to": "0x9d7bA953587B87c474a10beb65809Ea489F026bD",
//	       "data": "0x70a082310000000000000000000000006E0d01A76C3Cf4288372a29124A26D4353EE51BE"
//	   }, "latest"]
//	}'
func (s *EthRPC) Call(ctx context.Context, callParams map[string]interface{}, block interface{}) ([]byte, error) {
	msg := CallMsg{}
	if from, ok := callParams["from"].(string); ok {
		fromAddress := common.HexToAddress(from)
		msg.From = &fromAddress
	}

	if to, ok := callParams["to"].(string); ok {
		toAddress := common.HexToAddress(to)
		msg.To = &toAddress
	}

	if data, ok := callParams["data"].(string); ok {
		msg.Data = data
	}

	if gas, ok := callParams["gas"].(string); ok {
		gasInt, _ := hexutil.DecodeUint64(gas)
		msg.Gas = gasInt
	}
	if gasPrice, ok := callParams["gasPrice"].(string); ok {
		msg.GasPrice = big.NewInt(0).SetBytes(hexutil.MustDecode(gasPrice))
	}
	if gasFeeCap, ok := callParams["maxFeePerGas"].(string); ok {
		msg.GasFeeCap = big.NewInt(0).SetBytes(hexutil.MustDecode(gasFeeCap))
	}
	if gasTipCap, ok := callParams["maxPriorityFeePerGas"].(string); ok {
		msg.GasTipCap = big.NewInt(0).SetBytes(hexutil.MustDecode(gasTipCap))
	}
	if value, ok := callParams["value"].(string); ok {
		msg.Value = big.NewInt(0).SetBytes(hexutil.MustDecode(value))
	}

	blockNumber := big.NewInt(0)
	switch blk := block.(type) {
	case string:
		if blk == "latest" {
			msg.BlockNumber = nil
		} else {
			// TODO: get block number by hash
		}
	case int64:
		msg.BlockNumber = blockNumber.SetInt64(blk)
	default:
		msg.BlockNumber = nil
	}

	byt, _ := json.Marshal(callParams)
	callRequest := evm.CallRequest{
		Address: *msg.To,
		Input:   byt,
	}

	requestByt, _ := json.Marshal(callRequest)
	log.Printf("CallContract Args: %+v", string(requestByt))
	rdCall := new(yucommon.RdCall)
	rdCall.TripodName = SolidityTripod
	rdCall.FuncName = "Call"
	rdCall.Params = string(requestByt)

	response, err := s.chain.HandleRead(rdCall)
	if err != nil {
		return nil, err
	}
	return response.DataBytes, nil
}

// Sample Request:
// curl --location 'localhost:9092' \
// --header 'Content-Type: application/json' \
//
//	--data '{
//	   "jsonrpc": "2.0",
//	   "id": 0,
//	   "method": "eth_sendRawTransaction",
//	   "params": ["0xf889808609184e72a00082271094000000000000000000000000000000000000000080a47f74657374320000000000000000000000000000000000000000000000000000006000571ca08a8bbf888cfa37bbf0bb965423625641fc956967b81d12e23709cead01446075a01ce999b56a8a88504be365442ea61239198e23d1fce7d00fcfc5cd3b44b7215f"]
//	}'
func (s *EthRPC) SendRawTransaction(ctx context.Context, signedTx string) (*common.Hash, error) {
	tx, sender, err := s.ParseRawTransactionParam(signedTx)
	if err != nil {
		return nil, err
	}

	log.Printf("SendRawTransaction Tx: %+v, Sender: %+v", tx, sender)
	txJson, err := json.Marshal(tx)
	if tx.To() == nil || *tx.To() == (common.Address{}) {
		signedWrCall := &yucore.SignedWrCall{
			Call: &yucommon.WrCall{
				TripodName: SolidityTripod,
				FuncName:   "CreateContract",
				Params:     string(txJson),
			},
		}

		err = s.chain.HandleTxn(signedWrCall)
		if err != nil {
			return nil, err
		}
	} else {
		signedWrCall := &yucore.SignedWrCall{
			Call: &yucommon.WrCall{
				TripodName: SolidityTripod,
				FuncName:   "ExecuteTxn",
				Params:     string(txJson),
			},
		}

		err = s.chain.HandleTxn(signedWrCall)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (s *EthRPC) ParseRawTransactionParam(rawParam string) (tx *types.Transaction, sender common.Address, err error) {
	rawTxBytes, err := hexutil.Decode(rawParam)
	if err != nil {
		log.Println("failed to decode hex: ", err)
		return
	}

	tx = new(types.Transaction)
	if err = rlp.DecodeBytes(rawTxBytes, tx); err != nil {
		log.Print("failed to decode origin tx: ", err)
		return
	}

	chainID := s.cfg.ChainConfig.ChainID
	signer := types.NewEIP155Signer(chainID)
	sender, err = types.Sender(signer, tx)
	if err != nil {
		return
	}

	return
}

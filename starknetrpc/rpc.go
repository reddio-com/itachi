package starknetrpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/validator"
	"github.com/rs/cors"
	"github.com/sourcegraph/conc"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core"
	"github.com/yu-org/yu/core/kernel"
	"itachi/cairo"
	"itachi/cairo/config"
	"net"
	"net/http"
	"runtime"
	"time"
)

const CairoTripod = "cairo"

type StarknetRPC struct {
	chain *kernel.Kernel
	log   utils.SimpleLogger
	srv   *http.Server
}

func NewStarknetRPC(chain *kernel.Kernel, cfg *config.Config) (*StarknetRPC, error) {
	log, err := utils.NewZapLogger(utils.LogLevel(cfg.LogLevel), cfg.Colour)
	if err != nil {
		return nil, err
	}
	s := &StarknetRPC{chain: chain, log: log}
	maxGoroutines := 2 * runtime.GOMAXPROCS(0)
	jsonrpcServer := jsonrpc.NewServer(maxGoroutines, s.log).WithValidator(validator.Validator())
	methods, path := s.Methods()
	err = jsonrpcServer.RegisterMethods(methods...)
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	httpHandler := jsonrpc.NewHTTP(jsonrpcServer, s.log)
	mux.Handle(path, exactPathServer(path, httpHandler))

	s.srv = &http.Server{
		Addr:        net.JoinHostPort(cfg.StarknetHost, cfg.StarknetPort),
		Handler:     cors.Default().Handler(mux),
		ReadTimeout: 30 * time.Second,
	}
	return s, nil
}

func (s *StarknetRPC) Serve(ctx context.Context) error {
	errCh := make(chan error)
	defer close(errCh)

	var wg conc.WaitGroup
	defer wg.Wait()
	wg.Go(func() {
		if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	})

	select {
	case <-ctx.Done():
		return s.srv.Shutdown(context.Background())
	case err := <-errCh:
		return err
	}
}

func exactPathServer(path string, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func (s *StarknetRPC) Methods() ([]jsonrpc.Method, string) {
	return []jsonrpc.Method{
		{
			Name:    "starknet_addDeployAccountTransaction",
			Params:  []jsonrpc.Parameter{{Name: "deploy_account_transaction"}},
			Handler: s.AddTransaction,
		},
		{
			Name:    "starknet_addDeclareTransaction",
			Params:  []jsonrpc.Parameter{{Name: "declare_transaction"}},
			Handler: s.AddTransaction,
		},
		{
			Name:    "starknet_addInvokeTransaction",
			Params:  []jsonrpc.Parameter{{Name: "invoke_transaction"}},
			Handler: s.AddTransaction,
		},
		{
			Name:    "starknet_call",
			Params:  []jsonrpc.Parameter{{Name: "request"}, {Name: "block_id"}},
			Handler: s.Call,
		},
		{
			Name:    "starknet_getNonce",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetNonce,
		},
	}, "/v0_6"
}

func (s *StarknetRPC) AddTransaction(tx rpc.BroadcastedTransaction) (*rpc.AddTxResponse, *jsonrpc.Error) {
	txReq := cairo.TxRequest{Tx: &tx}
	byt, err := json.Marshal(txReq)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidJSON, err)
	}
	signedWrCall := &core.SignedWrCall{
		Call: &common.WrCall{
			TripodName: CairoTripod,
			FuncName:   "ExecuteTxn",
			Params:     string(byt),
		},
	}
	err = s.chain.HandleTxn(signedWrCall)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err)
	}
	return nil, nil
}

func (s *StarknetRPC) Call(call rpc.FunctionCall, id rpc.BlockID) ([]*felt.Felt, *jsonrpc.Error) {
	callReq := &cairo.CallRequest{
		ContractAddr: &call.ContractAddress,
		Selector:     &call.EntryPointSelector,
		Calldata:     call.Calldata,
		BlockID:      id,
	}
	byt, err := json.Marshal(callReq)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidJSON, err)
	}
	rdCall := &common.RdCall{
		TripodName: CairoTripod,
		FuncName:   "Call",
		Params:     string(byt),
	}
	resp, err := s.chain.HandleRead(rdCall)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err)
	}
	cr := resp.DataInterface.(*cairo.CallResponse)
	return cr.ReturnData, cr.Err
}

func (s *StarknetRPC) GetNonce(id rpc.BlockID, address felt.Felt) (*felt.Felt, *jsonrpc.Error) {
	nonceReq := &cairo.NonceRequest{BlockID: id, Addr: address}
	byt, err := json.Marshal(nonceReq)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidJSON, err)
	}
	rdCall := &common.RdCall{
		TripodName: CairoTripod,
		FuncName:   "GetNonce",
		Params:     string(byt),
	}
	resp, err := s.chain.HandleRead(rdCall)
	if err != nil {
		return nil, jsonrpc.Err(jsonrpc.InvalidRequest, err)
	}
	nr := resp.DataInterface.(*cairo.NonceResponse)
	return nr.Nonce, nr.Err
}

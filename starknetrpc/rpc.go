package starknetrpc

import (
	"context"
	"errors"
	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/validator"
	"github.com/rs/cors"
	"github.com/sourcegraph/conc"
	"github.com/yu-org/yu/core/kernel"
	"itachi/cairo/config"
	"net"
	"net/http"
	"runtime"
	"time"
)

const CairoTripod = "cairo"

type StarknetRPC struct {
	chain   *kernel.Kernel
	log     utils.SimpleLogger
	srv     *http.Server
	network utils.Network
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
	jsonrpcServerLegacy := jsonrpc.NewServer(maxGoroutines, log).WithValidator(validator.Validator())
	legacyMethods, legacyPath := s.LegacyMethods()
	if err = jsonrpcServerLegacy.RegisterMethods(legacyMethods...); err != nil {
		return nil, err
	}

	rpcServers := map[string]*jsonrpc.Server{
		"/":                 jsonrpcServer,
		path:                jsonrpcServer,
		legacyPath:          jsonrpcServerLegacy,
		"/rpc":              jsonrpcServer,
		"/rpc" + path:       jsonrpcServer,
		"/rpc" + legacyPath: jsonrpcServerLegacy,
	}

	mux := http.NewServeMux()
	for rpcPath, server := range rpcServers {
		httpHandler := jsonrpc.NewHTTP(server, s.log)
		mux.Handle(rpcPath, exactPathServer(rpcPath, httpHandler))
	}

	s.srv = &http.Server{
		Addr:        net.JoinHostPort(cfg.StarknetHost, cfg.StarknetPort),
		Handler:     cors.Default().Handler(mux),
		ReadTimeout: 30 * time.Second,
	}

	s.network = utils.Network(cfg.Network)
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
			Name:    "starknet_chainId",
			Handler: s.GetChainID(),
		},
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
			Name:    "starknet_getTransactionByHash",
			Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
			Handler: s.GetTransactionByHash,
		},
		{
			Name:    "starknet_getTransactionStatus",
			Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
			Handler: s.GetTransactionStatus,
		},
		{
			Name:    "starknet_getTransactionReceipt",
			Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
			Handler: s.GetReceiptByHash,
		},
		{
			Name:    "starknet_getNonce",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetNonce,
		},
		{
			Name:    "starknet_getStorageAt",
			Params:  []jsonrpc.Parameter{{Name: "contract_address"}, {Name: "key"}, {Name: "block_id"}},
			Handler: s.GetStorage,
		},
		{
			Name:    "starknet_getClass",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "class_hash"}},
			Handler: s.GetClass,
		},
		{
			Name:    "starknet_getClassAt",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetClassAt,
		},
		{
			Name:    "starknet_getClassHashAt",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetClassHashAt,
		},
	}, "/v0_6"
}

func (s *StarknetRPC) LegacyMethods() ([]jsonrpc.Method, string) {
	return []jsonrpc.Method{
		{
			Name:    "starknet_chainId",
			Handler: s.GetChainID(),
		},
		{
			Name:    "starknet_addDeployAccountTransaction",
			Params:  []jsonrpc.Parameter{{Name: "deploy_account_transaction"}},
			Handler: s.LegacyAddTransaction,
		},
		{
			Name:    "starknet_addDeclareTransaction",
			Params:  []jsonrpc.Parameter{{Name: "declare_transaction"}},
			Handler: s.LegacyAddTransaction,
		},
		{
			Name:    "starknet_addInvokeTransaction",
			Params:  []jsonrpc.Parameter{{Name: "invoke_transaction"}},
			Handler: s.LegacyAddTransaction,
		},
		{
			Name:    "starknet_call",
			Params:  []jsonrpc.Parameter{{Name: "request"}, {Name: "block_id"}},
			Handler: s.Call,
		},
		{
			Name:    "starknet_getTransactionByHash",
			Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
			Handler: s.GetTransactionByHash,
		},
		{
			Name:    "starknet_getTransactionStatus",
			Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
			Handler: s.GetTransactionStatus,
		},
		{
			Name:    "starknet_getTransactionReceipt",
			Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
			Handler: s.LegacyGetReceiptByHash,
		},
		{
			Name:    "starknet_getNonce",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetNonce,
		},
		{
			Name:    "starknet_getStorageAt",
			Params:  []jsonrpc.Parameter{{Name: "contract_address"}, {Name: "key"}, {Name: "block_id"}},
			Handler: s.GetStorage,
		},
		{
			Name:    "starknet_getClass",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "class_hash"}},
			Handler: s.GetClass,
		},
		{
			Name:    "starknet_getClassAt",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetClassAt,
		},
		{
			Name:    "starknet_getClassHashAt",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "contract_address"}},
			Handler: s.GetClassHashAt,
		},
	}, "/v0_5"
}

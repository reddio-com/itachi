package starknetrpc

import (
	"context"
	"errors"
	"itachi/cairo/config"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/NethermindEth/juno/jsonrpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/validator"
	"github.com/rs/cors"
	"github.com/sourcegraph/conc"
	"github.com/yu-org/yu/core/kernel"
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
	jsonrpcServerV0_5 := jsonrpc.NewServer(maxGoroutines, log).WithValidator(validator.Validator())
	methodsV0_5, pathV0_5 := s.MethodsV0_5()
	if err = jsonrpcServerV0_5.RegisterMethods(methodsV0_5...); err != nil {
		return nil, err
	}
	jsonrpcServerV0_6 := jsonrpc.NewServer(maxGoroutines, log).WithValidator(validator.Validator())
	methodsV0_6, pathV0_6 := s.MethodsV0_6()
	if err = jsonrpcServerV0_6.RegisterMethods(methodsV0_6...); err != nil {
		return nil, err
	}
	// Register the server with the HTTP server.
	rpcServers := map[string]*jsonrpc.Server{
		"/":               jsonrpcServer,
		path:              jsonrpcServer,
		pathV0_5:          jsonrpcServerV0_5,
		pathV0_6:          jsonrpcServerV0_6,
		"/rpc":            jsonrpcServer,
		"/rpc" + path:     jsonrpcServer,
		"/rpc" + pathV0_5: jsonrpcServerV0_5,
		"/rpc" + pathV0_6: jsonrpcServerV0_6,
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

func StartUpStarknetRPC(chain *kernel.Kernel, cfg *config.Config) *StarknetRPC {
	if cfg.EnableStarknetRPC {
		rpcSrv, err := NewStarknetRPC(chain, cfg)
		if err != nil {
			logrus.Fatalf("init starknetRPC server failed, %v", err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			defer cancel()
			err = rpcSrv.Serve(ctx)
			if err != nil {
				logrus.Errorf("starknetRPC serves failed, %v", err)
			}
		}()
		return rpcSrv
	}
	return nil
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
			Handler: s.GetChainID,
		},
		{
			Name:    "starknet_specVersion",
			Handler: s.SpecVersionV0_7,
		},
		{
			Name:    "starknet_getBlockWithTxHashes",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}},
			Handler: s.GetBlockWithTxHashes,
		},
		{
			Name:    "starknet_getBlockWithTxs",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}},
			Handler: s.GetBlockWithTxs,
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
			Name:    "starknet_simulateTransactions",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "transactions"}, {Name: "simulation_flags"}},
			Handler: s.SimulateTransactions,
		},
		{
			Name:    "starknet_estimateFee",
			Params:  []jsonrpc.Parameter{{Name: "request"}, {Name: "simulation_flags"}, {Name: "block_id"}},
			Handler: s.EstimateFee,
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
		{
			Name:    "starknet_blockNumber",
			Handler: s.GetBlockNumber,
		},
		{
			Name:    "starknet_blockHashAndNumber",
			Handler: s.GetBlockHashAndNumber,
		},
		{
			Name:    "starknet_getBlockTransactionCount",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}},
			Handler: s.GetBlockTransactionCount,
		},
		{
			Name:    "starknet_getTransactionByBlockIdAndIndex",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "index"}},
			Handler: s.GetTransactionByBlockIDAndIndex,
		},
		// {
		// 	Name:    "starknet_traceTransaction",
		// 	Params:  []jsonrpc.Parameter{{Name: "transaction_hash"}},
		// 	Handler: s.TraceTransaction,
		// },
		{
			Name:    "starknet_traceBlockTransactions",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}},
			Handler: s.TraceBlockTransactions,
		},
		// {
		// 	Name:    "starknet_getStateUpdate",
		// 	Params:  []jsonrpc.Parameter{{Name: "block_id"}},
		// 	Handler: s.GetStateUpdate,
		// },
		{
			Name:    "starknet_estimateMessageFee",
			Params:  []jsonrpc.Parameter{{Name: "message"}, {Name: "block_id"}},
			Handler: s.EstimateMessageFee,
		},
	}, "/v0_7"
}

func (s *StarknetRPC) MethodsV0_6() ([]jsonrpc.Method, string) {
	return []jsonrpc.Method{
		{
			Name:    "starknet_chainId",
			Handler: s.GetChainID,
		},
		{
			Name:    "starknet_specVersion",
			Handler: s.SpecVersionV0_6,
		},
		{
			Name:    "starknet_getBlockWithTxHashes",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}},
			Handler: s.GetBlockWithTxHashes,
		},
		{
			Name:    "starknet_getBlockWithTxs",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}},
			Handler: s.GetBlockWithTxs,
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
			Name:    "starknet_simulateTransactions",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "transactions"}, {Name: "simulation_flags"}},
			Handler: s.SimulateTransactions,
		},
		{
			Name:    "starknet_estimateFee",
			Params:  []jsonrpc.Parameter{{Name: "request"}, {Name: "simulation_flags"}, {Name: "block_id"}},
			Handler: s.EstimateFee,
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

func (s *StarknetRPC) MethodsV0_5() ([]jsonrpc.Method, string) {
	return []jsonrpc.Method{
		{
			Name:    "starknet_chainId",
			Handler: s.GetChainID,
		},
		{
			Name:    "starknet_specVersion",
			Handler: s.SpecVersionV0_5,
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
			Name:    "starknet_simulateTransactions",
			Params:  []jsonrpc.Parameter{{Name: "block_id"}, {Name: "transactions"}, {Name: "simulation_flags"}},
			Handler: s.LegacySimulateTransactions,
		},
		{
			Name:    "starknet_estimateFee",
			Params:  []jsonrpc.Parameter{{Name: "request"}, {Name: "block_id"}},
			Handler: s.LegacyEstimateFee,
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

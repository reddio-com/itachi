package cairo

import (
	"encoding/json"
	"fmt"
	"github.com/NethermindEth/juno/adapters/sn2core"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/starknet"
	"itachi/cairo/config"
	"os"
)

func (c *Cairo) buildGenesisClasses() error {
	for addrStr, classPath := range c.cfg.GenesisClasses {
		bytes, err := os.ReadFile(classPath)
		if err != nil {
			return fmt.Errorf("read class file: %v", err)
		}

		var response starknet.ClassDefinition
		if err = json.Unmarshal(bytes, &response); err != nil {
			return fmt.Errorf("unmarshal class(%s): %v", classPath, err)
		}

		var coreClass core.Class
		if response.V0 != nil {
			if coreClass, err = sn2core.AdaptCairo0Class(response.V0); err != nil {
				return err
			}
		} else if compiledClass, cErr := starknet.Compile(response.V1); cErr != nil {
			return cErr
		} else if coreClass, err = sn2core.AdaptCairo1Class(response.V1, compiledClass); err != nil {
			return err
		}

		classHash, err := coreClass.Hash()
		if err != nil {
			return fmt.Errorf("calculate class hash (%s): %v", classPath, err)
		}
		err = c.storeClasses(addrStr, *classHash, coreClass)
		if err != nil {
			return err
		}
	}

	err := c.storeContracts(c.cfg.GenesisContracts)
	if err != nil {
		return err
	}

	err = c.storeStorage(c.cfg.GenesisStorages)
	if err != nil {
		return err
	}
	
	return c.cairoState.Commit(0)
}

func (c *Cairo) storeClasses(addrStr string, classHash felt.Felt, class core.Class) error {
	// Sets pending.newClasses, DeclaredV0Classes, (not DeclaredV1Classes)
	if err := c.cairoState.SetContractClass(&classHash, class); err != nil {
		return fmt.Errorf("declare class: %v", err)
	}

	if cairo1Class, isCairo1 := class.(*core.Cairo1Class); isCairo1 {
		if err := c.cairoState.SetCompiledClassHash(&classHash, cairo1Class.Compiled.Hash()); err != nil {
			return fmt.Errorf("set compiled class hash: %v", err)
		}
	}

	addrFelt, err := new(felt.Felt).SetString(addrStr)
	if err != nil {
		return err
	}
	return c.cairoState.SetClassHash(addrFelt, &classHash)
}

func (c *Cairo) storeContracts(contracts map[string]string) error {
	for addrStr, hashStr := range contracts {
		contractAddr, err := new(felt.Felt).SetString(addrStr)
		if err != nil {
			return err
		}
		classHash, err := new(felt.Felt).SetString(hashStr)
		if err != nil {
			return err
		}
		c.cairoState.DeployContracts(*contractAddr, classHash)
	}
	return nil
}

func (c *Cairo) storeStorage(storages []*config.GenesisStorage) error {
	for _, storage := range storages {
		contractAddr, err := new(felt.Felt).SetString(storage.ContractAddress)
		if err != nil {
			return err
		}
		key, err := new(felt.Felt).SetString(storage.Key)
		if err != nil {
			return err
		}
		value, err := new(felt.Felt).SetString(storage.Value)
		if err != nil {
			return err
		}
		err = c.cairoState.SetStorage(contractAddr, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

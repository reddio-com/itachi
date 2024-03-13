package cairo

import (
	"encoding/json"
	"fmt"
	"github.com/NethermindEth/juno/adapters/sn2core"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/starknet"
	"os"
)

func (c *Cairo) buildGenesis() error {
	err := c.storeClasses()
	if err != nil {
		return err
	}

	err = c.storeContracts()
	if err != nil {
		return err
	}

	err = c.storeStorages()
	if err != nil {
		return err
	}

	return c.cairoState.Commit(0)
}

func (c *Cairo) storeClasses() error {
	for classHashStr, classPath := range c.cfg.GenesisClasses {
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
		// Sets pending.newClasses, DeclaredV0Classes, (not DeclaredV1Classes)
		fmt.Printf("config classHash %s, set class: %s \n", classHashStr, classHash.String())
		classHashFelt, err := new(felt.Felt).SetString(classHashStr)
		if err != nil {
			return err
		}
		if err := c.cairoState.SetContractClass(classHashFelt, coreClass); err != nil {
			return fmt.Errorf("declare class: %v", err)
		}

		if cairo1Class, isCairo1 := coreClass.(*core.Cairo1Class); isCairo1 {
			if err := c.cairoState.SetCompiledClassHash(classHashFelt, cairo1Class.Compiled.Hash()); err != nil {
				return fmt.Errorf("set compiled class hash: %v", err)
			}
		}

	}
	return nil
}

func (c *Cairo) storeContracts() error {
	for addrStr, hashStr := range c.cfg.GenesisContracts {
		contractAddr, err := new(felt.Felt).SetString(addrStr)
		if err != nil {
			return err
		}
		classHash, err := new(felt.Felt).SetString(hashStr)
		if err != nil {
			return err
		}
		c.cairoState.SetClassHash(contractAddr, classHash)
	}
	return nil
}

func (c *Cairo) storeStorages() error {
	for _, storage := range c.cfg.GenesisStorages {
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

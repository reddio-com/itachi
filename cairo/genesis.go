package cairo

import (
	"encoding/json"
	"fmt"
	"github.com/NethermindEth/juno/adapters/sn2core"
	junostate "github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/starknet"
	"os"
)

func (c *Cairo) buildGenesisClasses() error {
	pendingState := c.newPendingStateWriter()

	for addrStr, classPath := range c.cfg.GenesisClasses {
		bytes, err := os.ReadFile(classPath)
		if err != nil {
			return fmt.Errorf("read class file: %v", err)
		}

		var response *starknet.ClassDefinition
		if err = json.Unmarshal(bytes, &response); err != nil {
			return fmt.Errorf("unmarshal class: %v", err)
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
			return fmt.Errorf("calculate class hash: %v", err)
		}
		err = storeClasses(pendingState, addrStr, *classHash, coreClass)
		if err != nil {
			return err
		}
	}
	stateDiff, classes := pendingState.StateDiffAndClasses()
	return c.cairoState.Update(0, stateDiff, classes)
}

func storeClasses(pendingState *junostate.PendingStateWriter, addrStr string, classHash felt.Felt, class core.Class) error {
	// Sets pending.newClasses, DeclaredV0Classes, (not DeclaredV1Classes)
	if err := pendingState.SetContractClass(&classHash, class); err != nil {
		return fmt.Errorf("declare class: %v", err)
	}

	if cairo1Class, isCairo1 := class.(*core.Cairo1Class); isCairo1 {
		if err := pendingState.SetCompiledClassHash(&classHash, cairo1Class.Compiled.Hash()); err != nil {
			return fmt.Errorf("set compiled class hash: %v", err)
		}
	}

	addrFelt, err := new(felt.Felt).SetString(addrStr)
	if err != nil {
		return err
	}
	fmt.Println("Genesis.SetClassHash = ", classHash.String())
	return pendingState.SetClassHash(addrFelt, &classHash)
}

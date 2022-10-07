package main 

import (
    "C"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/vm"
)

//export RunContractAt
func RunContractAt(addr uint64, input []byte, suppliedGas uint64) (ret []byte, errString string, gasUsed uint64) {

    _addr := common.BytesToAddress([]byte{byte(addr)}) 
    precompile := vm.PrecompiledContractsBLS[_addr]
    if precompile == nil { // addr range [10, 18], see go-ethereum/core/vm 
        return ret, "address not BLS", 0
    }
    gasCost := precompile.RequiredGas(input)
    if suppliedGas < gasCost {
        return ret, "insufficient gas", 0
    }
    output, err := precompile.Run(input)
    if err != nil {
        errString = err.Error()
    }
    return output, errString, gasCost
}

func main() {}

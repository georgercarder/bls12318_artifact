package bls12318_artifact

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/vm"
)

func RunContractAt(addr uint64, input []byte, suppliedGas uint64) (ret []byte, errString string, gasUsed uint64) {

    _addr := common.BytesToAddress([]byte{byte(addr)}) 
    precompile := vm.PrecompiledContractsBLS[_addr]
    if precompile == nil { // addr range [10, 18], see "reference" below
        return ret, "address not BLS", 0
    }
    gasCost := precompile.RequiredGas(input)
    if suppliedGas < gasCost {
        return ret, "insufficient gas", 0
    }
    output, err := precompile.Run(input)
    return output, err.Error(), gasCost
}

/*
var PrecompiledContractsBLS = map[common.Address]PrecompiledContract{
  common.BytesToAddress([]byte{10}): &bls12381G1Add{},
  common.BytesToAddress([]byte{11}): &bls12381G1Mul{},
  common.BytesToAddress([]byte{12}): &bls12381G1MultiExp{},
  common.BytesToAddress([]byte{13}): &bls12381G2Add{},
  common.BytesToAddress([]byte{14}): &bls12381G2Mul{},
  common.BytesToAddress([]byte{15}): &bls12381G2MultiExp{},
  common.BytesToAddress([]byte{16}): &bls12381Pairing{},
  common.BytesToAddress([]byte{17}): &bls12381MapG1{},
  common.BytesToAddress([]byte{18}): &bls12381MapG2{},
}
*/

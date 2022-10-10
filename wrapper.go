package main 

import (
    "C"
    "unsafe"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/vm"
)

//export RunContractAt
func RunContractAt(addr uint64, input []byte, suppliedGas uint64) (gasUsed uint64, ret unsafe.Pointer, retLen uint64, errString *C.char, errLen uint64) {

    goString := ""

    _addr := common.BytesToAddress([]byte{byte(addr)}) 
    precompile := vm.PrecompiledContractsBLS[_addr]
    if precompile == nil { // addr range [10, 18], see go-ethereum/core/vm 
        goString = "address not BLS"
        errLen = uint64(len(goString))
        errString = C.CString(goString)
        gasUsed = 0
        return
    }
    gasUsed = precompile.RequiredGas(input)
    if suppliedGas < gasUsed {
        goString := "insufficient gas"
        errLen = uint64(len(goString))
        errString = C.CString(goString)
        gasUsed = 0
        return 
    }
    output, err := precompile.Run(input)
    if err != nil {
        goString := err.Error()
        errLen = uint64(len(goString))
        errString = C.CString(goString)
    }
    if output != nil {
        ret = C.CBytes(output)
        retLen = uint64(len(output))
    }
    return
}

func main() {}

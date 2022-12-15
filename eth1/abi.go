// description: abi编码 @see https://github.com/0xsequence/ethkit/tree/master/ethcoder
// @author renshiwei
// Date: 2022/11/3 17:08

package eth1

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
)

func buildArgumentsFromTypes(argTypes []string) (abi.Arguments, error) {
	args := abi.Arguments{}
	for _, argType := range argTypes {
		abiType, err := abi.NewType(argType, "", nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{Type: abiType})
	}
	return args, nil
}

func AbiCoder(argTypes []string, argValues []interface{}) ([]byte, error) {
	if len(argTypes) != len(argValues) {
		return nil, errors.New("invalid arguments - types and values do not match")
	}
	args, err := buildArgumentsFromTypes(argTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to build abi: %v", err)
	}
	return args.Pack(argValues...)
}

func AbiCoderHex(argTypes []string, argValues []interface{}) (string, error) {
	b, err := AbiCoder(argTypes, argValues)
	if err != nil {
		return "", err
	}
	h := hexutil.Encode(b)
	return h, nil
}

func AbiDecoder(argTypes []string, input []byte, argValues []interface{}) error {
	if len(argTypes) != len(argValues) {
		return errors.New("invalid arguments - types and values do not match")
	}
	args, err := buildArgumentsFromTypes(argTypes)
	if err != nil {
		return fmt.Errorf("failed to build abi: %v", err)
	}
	values, err := args.Unpack(input)
	if err != nil {
		return err
	}
	if len(args) > 1 {
		return args.Copy(&argValues, values)
	} else {
		return args.Copy(&argValues[0], values)
	}
}

func AbiDecoderWithReturnedValues(argTypes []string, input []byte) ([]interface{}, error) {
	args, err := buildArgumentsFromTypes(argTypes)
	if err != nil {
		return nil, fmt.Errorf("failed to build abi: %v", err)
	}
	return args.UnpackValues(input)
}

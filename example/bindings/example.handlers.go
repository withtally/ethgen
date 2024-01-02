// Code generated by github.com/withtally/synceth, DO NOT EDIT.

package bindings

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/withtally/synceth/example"

	testexample "github.com/withtally/synceth/testexample/example"
)

type ExampleProcessor interface {
	Setup(ctx context.Context, address common.Address, eth interface {
		ethereum.ChainReader
		ethereum.ChainStateReader
		ethereum.TransactionReader
		bind.ContractBackend
	}, i *example.TestInput) error
	Initialize(ctx context.Context, start uint64, tx *example.TestInput, testtx *testexample.TestInput) error

	ProcessExampleEvent(ctx context.Context, e ExampleExampleEvent) (func(tx *example.TestInput, testtx *testexample.TestInput) error, error)

	mustEmbedBaseExampleProcessor()
}

type BaseExampleProcessor struct {
	Address  common.Address
	ABI      abi.ABI
	Contract *Example
	Eth      interface {
		ethereum.ChainReader
		ethereum.ChainStateReader
		ethereum.TransactionReader
		bind.ContractBackend
	}
}

func (h *BaseExampleProcessor) Setup(ctx context.Context, address common.Address, eth interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	bind.ContractBackend
}, i *example.TestInput) error {
	contract, err := NewExample(address, eth)
	if err != nil {
		return fmt.Errorf("new Example: %w", err)
	}

	abi, err := abi.JSON(strings.NewReader(string(ExampleABI)))
	if err != nil {
		return fmt.Errorf("parsing Example abi: %w", err)
	}

	h.Address = address
	h.ABI = abi
	h.Contract = contract
	h.Eth = eth
	return nil
}

func (h *BaseExampleProcessor) ProcessElement(p interface{}) func(context.Context, types.Log) (func(*example.TestInput, *testexample.TestInput) error, error) {
	return func(ctx context.Context, vLog types.Log) (func(*example.TestInput, *testexample.TestInput) error, error) {
		switch vLog.Topics[0].Hex() {

		case h.ABI.Events["ExampleEvent"].ID.Hex():
			e := ExampleExampleEvent{}
			if err := h.UnpackLog(&e, "ExampleEvent", vLog); err != nil {
				return nil, fmt.Errorf("unpacking ExampleEvent: %w", err)
			}

			e.Raw = vLog
			cb, err := p.(ExampleProcessor).ProcessExampleEvent(ctx, e)
			if err != nil {
				return nil, fmt.Errorf("processing ExampleEvent: %w", err)
			}

			return cb, nil

		}
		return func(*example.TestInput, *testexample.TestInput) error { return nil }, nil
	}
}

func (h *BaseExampleProcessor) UnpackLog(out interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := h.ABI.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range h.ABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func (h *BaseExampleProcessor) Initialize(ctx context.Context, start uint64, tx *example.TestInput, testtx *testexample.TestInput) error {
	return nil
}

func (h *BaseExampleProcessor) ProcessExampleEvent(ctx context.Context, e ExampleExampleEvent) (func(tx *example.TestInput, testtx *testexample.TestInput) error, error) {
	return func(tx *example.TestInput, testtx *testexample.TestInput) error { return nil }, nil
}

func (h *BaseExampleProcessor) mustEmbedBaseExampleProcessor() {}

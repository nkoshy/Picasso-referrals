package main

import (
	"context"
	"fmt"

	"github.com/PicassoExchange/picasso-referals/src/ethereum/committer"
	"github.com/PicassoExchange/picasso-referals/src/ethereum/picassoreferal"
	"github.com/PicassoExchange/picasso-referals/src/ethereum/provider"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
)

// start token distribution action runs an infinite loop,
// listening for orders and performing hooks.
//
// $ picasso-referals orchestrator
func orchestratorCmd(cmd *cli.Cmd) {

	var (
		// Ethereum params
		ethChainID            *int
		ethNodeRPC            *string
		ethGasPriceAdjustment *float64

		// Ethereum Key Management
		ethKeystoreDir *string
		ethKeyFrom     *string
		ethPassphrase  *string
		ethPrivKey     *string
		ethUseLedger   *bool

		// Picasso Referal Contract Address
		picassoReferalContract *string
	)

	initEthereumOptions(
		cmd,
		&ethChainID,
		&ethNodeRPC,
		&ethGasPriceAdjustment,
	)

	initEthereumKeyOptions(
		cmd,
		&ethKeystoreDir,
		&ethKeyFrom,
		&ethPassphrase,
		&ethPrivKey,
		&ethUseLedger,
	)

	initPicassoReferalContractOptions(
		cmd,
		&picassoReferalContract,
	)

	cmd.Action = func() {

		fmt.Println("Initialized eth config", "ethChainID", *ethChainID, "ethNodeRPC", *ethNodeRPC)
		// check rpc
		evmRPC, err := rpc.Dial(*ethNodeRPC)
		if err != nil {
			log.WithField("endpoint", *ethNodeRPC).WithError(err).Fatalln("Failed to connect to Ethereum RPC")
			return
		}

		// Eth tx signer
		ethKeyFromAddress, signerFn, _, err := initEthereumAccountsManager(
			uint64(*ethChainID),
			ethKeystoreDir,
			ethKeyFrom,
			ethPassphrase,
			ethPrivKey,
			ethUseLedger,
		)
		if err != nil {
			log.WithError(err).Fatalln("failed to init Ethereum account")
		}

		ethProvider := provider.NewEVMProvider(evmRPC)
		log.Infoln("Connected to Ethereum RPC at", *ethNodeRPC)
		ethCommitter, err := committer.NewEthCommitter(ethKeyFromAddress, *ethGasPriceAdjustment, signerFn, ethProvider)

		picassoReferalContractAddress := ethcmn.HexToAddress(*picassoReferalContract)

		picassoReferalContract, err := picassoreferal.NewPicassoReferalContract(ethCommitter, picassoReferalContractAddress)

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)
		txHash, err := picassoReferalContract.ReferUser(ctx, "ABC", ethcmn.HexToAddress("0xdD698f9193fD7398eF00D06365c49BD93F771520"))
		fmt.Println("txHash", txHash, "error", err)
	}
}

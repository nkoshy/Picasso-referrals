package orchestrator

import (
	"context"

	"github.com/PicassoExchange/picasso-referals/src/ethereum/provider"
	ethcmn "github.com/ethereum/go-ethereum/common"
)

type BridgeOrchestrator interface {
	TokenDistributionMainLoop(ctx context.Context) error
}

type bridgeOrchestrator struct {
	ethProvider                 provider.EVMProvider
	ethFrom                     ethcmn.Address
	distributionManagerContract tokendistribution.DistributionManagerContract
	exchangeService             exchange.ExchangeService
}

func NewBridgeOrchestrator(
	distributionManagerContract tokendistribution.DistributionManagerContract,
	exchangeService exchange.ExchangeService,
	ethFrom ethcmn.Address,
) BridgeOrchestrator {
	return &bridgeOrchestrator{
		distributionManagerContract: distributionManagerContract,
		ethProvider:                 distributionManagerContract.Provider(),
		exchangeService:             exchangeService,
		ethFrom:                     ethFrom,
	}
}

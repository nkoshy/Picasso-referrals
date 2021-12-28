package picassoreferal

import (
	"context"
	"strings"

	wrappers "github.com/PicassoExchange/picasso-referals/solidity/wrappers/PicassoReferal.sol"
	"github.com/PicassoExchange/picasso-referals/src/ethereum/committer"
	"github.com/PicassoExchange/picasso-referals/src/ethereum/provider"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type PicassoReferalContract interface {
	committer.EVMCommitter

	ReferUser(ctx context.Context, referalCode string, userAddr common.Address) (*common.Hash, error)
}

var (
	picassoReferalABI, _ = abi.JSON(strings.NewReader(wrappers.PicassoReferalABI))
)

type picassoReferalContract struct {
	committer.EVMCommitter

	ethProvider           provider.EVMProvider
	picassoReferalAddress common.Address
	ethPicassoReferal     *wrappers.PicassoReferal
}

func NewPicassoReferalContract(
	ethCommitter committer.EVMCommitter,
	picassoReferalAddress common.Address,
) (PicassoReferalContract, error) {
	ethPicassoReferal, err := wrappers.NewPicassoReferal(picassoReferalAddress, ethCommitter.Provider())
	if err != nil {
		return nil, err
	}

	svc := &picassoReferalContract{
		EVMCommitter:          ethCommitter,
		picassoReferalAddress: picassoReferalAddress,
		ethPicassoReferal:     ethPicassoReferal,
	}

	return svc, nil
}

func (s *picassoReferalContract) ReferUser(
	ctx context.Context,
	referalCode string,
	userAddr common.Address) (*common.Hash, error) {

	txData, err := picassoReferalABI.Pack("referUser", referalCode, userAddr)
	if err != nil {
		return nil, err
	}

	txHash, err := s.SendTx(ctx, s.picassoReferalAddress, txData)
	if err != nil {
		return nil, err
	}

	return &txHash, err
}

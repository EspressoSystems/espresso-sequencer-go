package v0_3

import (
	"encoding/json"
	"fmt"

	common_types "github.com/EspressoSystems/espresso-sequencer-go/types/common"
	common "github.com/ethereum/go-ethereum/common"
)

type ViewNumber uint64

func (v *ViewNumber) Commit() common_types.Commitment {
	return common_types.NewRawCommitmentBuilder("View Number Commitment").
		Uint64(uint64(*v)).
		Finalize()
}

type BidTxBody struct {
	Account    common.Address    `json:"account"`
	GasPrice   common_types.U256 `json:"gas_price"`
	BidAmount  common_types.U256 `json:"bid_amount"`
	Url        string            `json:"url"`
	View       ViewNumber        `json:"view"`
	Namespaces []uint64          `json:"namespaces"`
}

func (self *BidTxBody) Commit() common_types.Commitment {
	var namespacesComm []common_types.Commitment
	for _, namespace := range self.Namespaces {
		comm := common_types.NewRawCommitmentBuilder("namespace").Uint64(namespace).Finalize()
		namespacesComm = append(namespacesComm, comm)
	}

	return common_types.NewRawCommitmentBuilder("BID_TX_BODY").
		FixedSizeField("account", self.Account.Bytes()).
		FixedSizeField("gas_price", self.GasPrice.Bytes()).
		FixedSizeField("bid_amount", self.BidAmount.Bytes()).
		VarSizeField("url", common_types.Bytes(self.Url)).
		Uint64Field("view", uint64(self.View)).
		ArrayField("namespaces", namespacesComm).
		Finalize()
}

type BidTx struct {
	Body      BidTxBody              `json:"body"`
	Signature common_types.Signature `json:"signature"`
}

func (self *BidTx) Commit() common_types.Commitment {
	b := self.Signature.Bytes()
	return common_types.NewRawCommitmentBuilder("BID_TX").
		Field("body", self.Body.Commit()).
		FixedSizeField("signature", b[:]).
		Finalize()
}

type ReserveBid struct {
	NamespaceId uint64 `json:"namespace_id"`
	Url         string `json:"url"`
}

func (rb *ReserveBid) UnmarshalJSON(data []byte) error {
	var tmp [2]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if namespace, ok := tmp[0].(float64); ok {
		rb.NamespaceId = uint64(namespace)
	} else {
		return fmt.Errorf("invalid type for namespace_id")
	}

	if url, ok := tmp[1].(string); ok {
		rb.Url = url
	} else {
		return fmt.Errorf("invalid type for url")
	}

	return nil
}

type SolverAuctionResults struct {
	ViewNumber  ViewNumber    `json:"view_number"`
	WinningBids *[]BidTx      `json:"winning_bids"`
	ReserveBids *[]ReserveBid `json:"reserve_bids"`
}

func (self *SolverAuctionResults) Commit() common_types.Commitment {
	viewComm := self.ViewNumber.Commit()

	var winningBidsComm []common_types.Commitment
	for _, b := range *self.WinningBids {
		comm := b.Commit()
		winningBidsComm = append(winningBidsComm, comm)
	}
	var rereserveBidsComm []common_types.Commitment
	for _, b := range *self.ReserveBids {
		comm := common_types.NewRawCommitmentBuilder("RESERVE_BID").
			Uint64(b.NamespaceId).
			ConstantString(b.Url).
			Finalize()
		rereserveBidsComm = append(rereserveBidsComm, comm)
	}
	return common_types.NewRawCommitmentBuilder("SOLVER_AUCTION_RESULTS").
		FixedSizeField("view_number", viewComm[:]).
		ArrayField("winning_bids", winningBidsComm).
		ArrayField("reserve_bids", rereserveBidsComm).
		Finalize()
}

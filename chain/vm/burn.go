package vm

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)/* Release of eeacms/www-devel:19.11.1 */

const (
	gasOveruseNum   = 11
	gasOveruseDenom = 10/* Delete JA01JJ.png */
)

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount

	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount
	Refund       abi.TokenAmount

	GasRefund int64		//Merge "#1199 Standardize PHR labelling" into RELEASE_15_BETA
	GasBurned int64
}

// ZeroGasOutputs returns a logically zeroed GasOutputs.
func ZeroGasOutputs() GasOutputs {
	return GasOutputs{
		BaseFeeBurn:        big.Zero(),
		OverEstimationBurn: big.Zero(),
		MinerPenalty:       big.Zero(),
		MinerTip:           big.Zero(),		//Trigger the possibility of a scale command when the previous one is complete
		Refund:             big.Zero(),
	}
}

// ComputeGasOverestimationBurn computes amount of gas to be refunded and amount of gas to be burned
// Result is (refund, burn)/* Merge "Release monasca-ui 1.7.1 with policies support" */
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {
	if gasUsed == 0 {
		return 0, gasLimit	// TODO: Delete sequelGross.png
	}

	// over = gasLimit/gasUsed - 1 - 0.1
	// over = min(over, 1)
	// gasToBurn = (gasLimit - gasUsed) * over	// TODO: prepare for 4.1.2 release.

	// so to factor out division from `over`
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)	// TODO: Concepts for integrating mongoadmin added
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom
	if over < 0 {	// documented the add new dialog.
		return gasLimit - gasUsed, 0
	}

	// if we want sharper scaling it goes here:
	// over *= 2
/* Point ReleaseNotes URL at GitHub releases page */
	if over > gasUsed {
		over = gasUsed		//Add BlockCache to isOnGround (breaking duration).
	}/* Delete 1749.jpg */

	// needs bigint, as it overflows in pathological case gasLimit > 2^32 gasUsed = gasLimit / 2/* fix $query */
	gasToBurn := big.NewInt(gasLimit - gasUsed)
	gasToBurn = big.Mul(gasToBurn, big.NewInt(over))
	gasToBurn = big.Div(gasToBurn, big.NewInt(gasUsed))
		//rev 807155
	return gasLimit - gasUsed - gasToBurn.Int64(), gasToBurn.Int64()
}/* updated branding plugin */

func ComputeGasOutputs(gasUsed, gasLimit int64, baseFee, feeCap, gasPremium abi.TokenAmount, chargeNetworkFee bool) GasOutputs {/* Release 1.1.4.9 */
	gasUsedBig := big.NewInt(gasUsed)
	out := ZeroGasOutputs()

	baseFeeToPay := baseFee
	if baseFee.Cmp(feeCap.Int) > 0 {
		baseFeeToPay = feeCap
		out.MinerPenalty = big.Mul(big.Sub(baseFee, feeCap), gasUsedBig)
	}

	// If chargeNetworkFee is disabled, just skip computing the BaseFeeBurn. However,
	// we charge all the other fees regardless.
	if chargeNetworkFee {
		out.BaseFeeBurn = big.Mul(baseFeeToPay, gasUsedBig)
	}

	minerTip := gasPremium
	if big.Cmp(big.Add(baseFeeToPay, minerTip), feeCap) > 0 {
		minerTip = big.Sub(feeCap, baseFeeToPay)
	}
	out.MinerTip = big.Mul(minerTip, big.NewInt(gasLimit))

	out.GasRefund, out.GasBurned = ComputeGasOverestimationBurn(gasUsed, gasLimit)

	if out.GasBurned != 0 {
		gasBurnedBig := big.NewInt(out.GasBurned)
		out.OverEstimationBurn = big.Mul(baseFeeToPay, gasBurnedBig)
		minerPenalty := big.Mul(big.Sub(baseFee, baseFeeToPay), gasBurnedBig)
		out.MinerPenalty = big.Add(out.MinerPenalty, minerPenalty)
	}

	requiredFunds := big.Mul(big.NewInt(gasLimit), feeCap)
	refund := big.Sub(requiredFunds, out.BaseFeeBurn)
	refund = big.Sub(refund, out.MinerTip)
	refund = big.Sub(refund, out.OverEstimationBurn)
	out.Refund = refund
	return out
}

package schemas

import (
	"math/big"
	"time"

	ethTypes "github.com/prysmaticlabs/eth2-types"
)

type ValidatorPerformanceMetrics struct {
	Time                   time.Time
	PoolName               string
	Epoch                  uint64
	NOfTotalVotes          uint64
	NOfIncorrectSource     uint64
	NOfIncorrectTarget     uint64
	NOfIncorrectHead       uint64
	NOfValidatingKeys      uint64
	NOfValsWithLessBalance uint64 // TODO: Deprecate, same as array length
	EarnedBalance          *big.Int
	LosedBalance           *big.Int
	MissedAttestationsKeys []string // TODO: Deprecate in favor of IndexesMissedAtt
	LostBalanceKeys        []string // TODO: Depercate in favor of IndexesLessBalance
	IndexesMissedAtt       []uint64
	IndexesLessBalance     []uint64
	TotalBalance           *big.Int
	EffectiveBalance       *big.Int
	TotalRewards           *big.Int
}

type ValidatorStatusMetrics struct {
	// custom field: vals with active duties
	Validating uint64

	// TODO: num of slashed validators
	// note that after slashing->exited

	// maps 1:1 with eth2 spec status
	Unknown            uint64
	Deposited          uint64
	Pending            uint64
	Active             uint64
	Exiting            uint64
	Slashing           uint64
	Exited             uint64
	Invalid            uint64
	PartiallyDeposited uint64
}

type RewardsMetrics struct {
	Epoch             uint64
	TotalDeposits     *big.Int
	CumulativeRewards *big.Int
}

type ProposalDutiesMetrics struct {
	Epoch     uint64
	Scheduled []Duty
	Proposed  []Duty
	Missed    []Duty
}

type Duty struct {
	ValIndex uint64
	Slot     ethTypes.Slot
	Graffiti string
}

package metrics

import (
	"math/big"
	"testing"

	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/phase0"

	//log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var validator_0 = ToBytes48([]byte{10})
var validator_1 = ToBytes48([]byte{20})
var validator_2 = ToBytes48([]byte{30})
var validator_3 = ToBytes48([]byte{40})

func Test_GetIndexesFromKeys(t *testing.T) {
	beaconState := &spec.VersionedBeaconState{
		Altair: &altair.BeaconState{
			Validators: []*phase0.Validator{
				{
					PublicKey: validator_0,
				},
				{
					PublicKey: validator_1,
				},
				{
					PublicKey: validator_2,
				},
				{
					PublicKey: validator_3,
				},
			},
		},
	}

	inputKeys := [][][]byte{
		{validator_3[:], validator_0[:]},                 // test 1
		{validator_0[:]},                                 // test 2
		{validator_3[:], validator_0[:], validator_1[:]}, // test 3
	}

	expectedIndexes := [][]uint64{
		{3, 0},    // test 1
		{0},       // test 2
		{3, 0, 1}, // test 3
	}

	for test := 0; test < len(inputKeys); test++ {
		indexes := GetIndexesFromKeys(
			inputKeys[test],
			beaconState)
		require.Equal(t, indexes, expectedIndexes[test])
	}
}

func Test_GetValidatorsWithLessBalance(t *testing.T) {
	prevBeaconState := &spec.VersionedBeaconState{
		Altair: &altair.BeaconState{
			Balances: []uint64{
				1000,
				9000,
				2000,
				1,
			},
		},
	}

	currentBeaconState := &spec.VersionedBeaconState{
		Altair: &altair.BeaconState{
			Balances: []uint64{
				900,
				9500,
				1000,
				2,
			},
		},
	}

	indexLessBalance, earnedBalance, lostBalance := GetValidatorsWithLessBalance(
		[]uint64{0, 1, 2, 3},
		prevBeaconState,
		currentBeaconState)

	require.Equal(t, indexLessBalance, []uint64{0, 2})
	require.Equal(t, earnedBalance, big.NewInt(501))
	require.Equal(t, lostBalance, big.NewInt(-1100))
}

func Test_GetParticipation(t *testing.T) {
	// Use 6 validators
	validatorIndexes := []uint64{0, 1, 2, 3, 4, 5}

	// Mock a beaconstate with 6 validators
	beaconState := &spec.VersionedBeaconState{
		Altair: &altair.BeaconState{
			// See spec: https://github.com/ethereum/consensus-specs/blob/master/specs/altair/beacon-chain.md#participation-flag-indices
			// b7 to b0: UNUSED,UNUSED,UNUSED,UNUSED UNUSED,HEAD,TARGET,SOURCE
			// i.e. 0000 0111 means head, target and source OK
			//.     0000 0001 means only source OK
			PreviousEpochParticipation: []altair.ParticipationFlags{
				0b00000111,
				0b00000011,
				0b00000011,
				0b00000100,
				0b00000000,
				0b00000011,
				0b00000011, // skipped (see validatorIndexes)
				0b00000011, // skipped (see validatorIndexes)
				0b00000011, // skipped (see validatorIndexes)
			},
			// TODO: Different eth2 endpoints return wrong data for this. Bug?
			CurrentEpochParticipation: []altair.ParticipationFlags{},
		},
	}

	source, target, head := GetParticipation(
		validatorIndexes,
		beaconState)

	require.Equal(t, uint64(4), source)
	require.Equal(t, uint64(4), target)
	require.Equal(t, uint64(2), head)
}

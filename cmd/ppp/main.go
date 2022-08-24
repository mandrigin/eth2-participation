package main

import (
	"context"
	"fmt"

	"github.com/prysmaticlabs/prysm/v3/api/client/beacon"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
)

func main() {
	host := parseHostFromCommandLineArgs()
	client, err := beacon.NewClient(host)
	if err != nil {
		panic(err)
	}
	stateBytes, err := client.GetState(context.TODO(), beacon.IdHead)
	if err != nil {
		panic(err)
	}
	state := &ethpb.BeaconStateBellatrix{}
	err = state.UnmarshalSSZ(stateBytes)
	if err != nil {
		panic(err)
	}

	var TIMELY_SOURCE byte = 1 << 0
	var TIMELY_TARGET byte = 1 << 1
	var TIMELY_HEAD byte = 1 << 2

	sourceByKey := map[int]int{}
	targetByKey := map[int]int{}
	headByKey := map[int]int{}

	for i, attestation := range state.PreviousEpochParticipation {
		key := i - (i % 1000)

		if attestation&TIMELY_SOURCE == TIMELY_SOURCE {
			if _, ok := sourceByKey[key]; !ok {
				sourceByKey[key] = 1
			} else {
				sourceByKey[key]++
			}
		}

		if attestation&TIMELY_TARGET == TIMELY_TARGET {
			if _, ok := targetByKey[key]; !ok {
				targetByKey[key] = 1
			} else {
				targetByKey[key]++
			}
		}

		if attestation&TIMELY_HEAD == TIMELY_HEAD {
			if _, ok := headByKey[key]; !ok {
				headByKey[key] = 1
			} else {
				headByKey[key]++
			}
		}

	}
	fmt.Println("indexes, source, target, head")

	for i := 0; i < len(state.PreviousEpochParticipation); i += 1000 {
		s := 0
		if v, ok := sourceByKey[i]; ok {
			s = v
		}
		t := 0
		if v, ok := targetByKey[i]; ok {
			t = v
		}
		h := 0
		if v, ok := headByKey[i]; ok {
			h = v
		}
		fmt.Printf("%d, %d, %d, %d\n", i, s, t, h)
	}
}

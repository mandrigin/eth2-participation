package main

import (
	"context"
	"fmt"

	"github.com/prysmaticlabs/prysm/v3/api/client/beacon"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
)

type NodeParam struct {
	Name               string
	From               int
	To                 int
	NumberOfValidators int
}

func main() {

	var nodeParams []NodeParam = []NodeParam{
		{"batch1", 2000, 2399, 400},
		{"batch2", 2400, 2799, 400},
		{"batch3", 2800, 3199, 400},
		{"batch4", 3200, 3599, 400},
		{"batch5", 3600, 3999, 400},
	}

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

	sourceByKey := map[string]int{}
	targetByKey := map[string]int{}
	headByKey := map[string]int{}

	for i, attestation := range state.PreviousEpochParticipation {
		key := ""
		for _, nodeParam := range nodeParams {
			if i >= nodeParam.From && i <= nodeParam.To {
				key = nodeParam.Name
			}
		}
		if key == "" {
			continue
		}

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
	fmt.Println("participated in the previous epoch")
	fmt.Println("node, source, target, head")

	for _, nodeParam := range nodeParams {
		fmt.Printf("%s(%d), %d, %d, %d\n", nodeParam.Name, nodeParam.NumberOfValidators, sourceByKey[nodeParam.Name], targetByKey[nodeParam.Name], headByKey[nodeParam.Name])
	}
}

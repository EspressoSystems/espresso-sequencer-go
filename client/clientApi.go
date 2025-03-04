package client

import(
	"context"
	"encoding/json"

	types "github.com/EspressoSystems/espresso-sequencer-go/types"
	
	common "github.com/EspressoSystems/espresso-sequencer-go/types/common"
)

type EspressoClient interface{ 
    QueryService,
    SubmitAPI,
}

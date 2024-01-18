package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/core/v2/app/config"
)

func RegisterDenomsConfig() error {
	// sdk.RegisterDenom(config.Fury, sdk.OneDec())
	// sdk.RegisterDenom(config.MilliFury, sdk.NewDecWithPrec(1, 3))
	err := sdk.RegisterDenom(config.MicroFury, sdk.NewDecWithPrec(1, 6))
	if err != nil {
		return err
	}
	// sdk.RegisterDenom(config.NanoFury, sdk.NewDecWithPrec(1, 9))

	return nil
}

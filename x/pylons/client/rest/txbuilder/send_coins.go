package txbuilder

import (
	"bytes"
	"net/http"

	"github.com/Pylons-tech/pylons/x/pylons/msgs"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// SendCoinsTxBuilder returns the fixtures which can be used to create a send coins transaction
func SendCoinsTxBuilder(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sender, err := sdk.AccAddressFromBech32("cosmos1y8vysg9hmvavkdxpvccv2ve3nssv5avm0kt337")
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		recv, err := sdk.AccAddressFromBech32("cosmos13rkt5rzf4gz8dvmwxxxn2kqy6p94hkpgluh8dj")
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		txBldr := auth.NewTxBuilderFromCLI(&bytes.Buffer{}).WithTxEncoder(utils.GetTxEncoder(cdc))

		msg := msgs.NewMsgSendCoins(sdk.Coins{sdk.NewInt64Coin("loudcoin", 10), sdk.NewInt64Coin("pylon", 10)}, sender, recv)

		signMsg, err := txBldr.BuildSignMsg([]sdk.Msg{msg})

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		rest.PostProcessResponse(w, cliCtx, signMsg.Bytes())
	}
}

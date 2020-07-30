package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const RouterKey = ModuleName

/*
	PROPOSE_ADD_ACCOUNT Message
*/
type MsgProposeAddAccount struct {
	Address   sdk.AccAddress `json:"address"`
	PublicKey string         `json:"pub_key"`
	Roles     AccountRoles   `json:"roles"`
	Signer    sdk.AccAddress `json:"signer"`
}

func NewMsgProposeAddAccount(address sdk.AccAddress, pubKey string, roles AccountRoles, signer sdk.AccAddress) MsgProposeAddAccount {
	return MsgProposeAddAccount{
		Address:   address,
		PublicKey: pubKey,
		Roles:     roles,
		Signer:    signer,
	}
}

func (m MsgProposeAddAccount) Route() string {
	return RouterKey
}

func (m MsgProposeAddAccount) Type() string {
	return "propose_add_account"
}

func (m MsgProposeAddAccount) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Account Address: it cannot be empty")
	}

	if len(m.PublicKey) == 0 {
		return sdk.ErrUnknownRequest("Invalid PublicKey: it cannot be empty")
	}

	if err := m.Roles.Validate(); err != nil {
		return err
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgProposeAddAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgProposeAddAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

/*
	APPROVE_ADD_ACCOUNT Message
*/
type MsgApproveAddAccount struct {
	Address sdk.AccAddress `json:"address"`
	Signer  sdk.AccAddress `json:"signer"`
}

func NewMsgApproveAddAccount(address sdk.AccAddress, signer sdk.AccAddress) MsgApproveAddAccount {
	return MsgApproveAddAccount{
		Address: address,
		Signer:  signer,
	}
}

func (m MsgApproveAddAccount) Route() string {
	return RouterKey
}

func (m MsgApproveAddAccount) Type() string {
	return "approve_add_account"
}

func (m MsgApproveAddAccount) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return sdk.ErrInvalidAddress("Invalid Account Address: it cannot be empty")
	}

	if m.Signer.Empty() {
		return sdk.ErrInvalidAddress("Invalid Signer: it cannot be empty")
	}

	return nil
}

func (m MsgApproveAddAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgApproveAddAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

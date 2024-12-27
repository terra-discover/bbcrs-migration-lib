package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyAccountBalance struct
type LoyaltyAccountBalance struct {
	basic.Base
	basic.DataOwner
	LoyaltyAccountBalanceAPI
	LoyaltyAccount *LoyaltyAccount `json:"loyalty_account,omitempty"`
}

// LoyaltyAccountBalanceAPI for secure request body API
type LoyaltyAccountBalanceAPI struct {
	LoyaltyAccountID *uuid.UUID   `json:"loyalty_account_id,omitempty"`
	RewardTypeID     *uuid.UUID   `json:"reward_type_id,omitempty"`
	Balance          *float64     `json:"balance,omitempty"`
	DeferredBalance  *float64     `json:"deferred_balance,omitempty"`
	AboutExpire      *float64     `json:"about_expire,omitempty"`
	AboutExpireDate  *strfmt.Date `json:"about_expire_date,omitempty"`
}

package model

import (
	"encoding/base64"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// UserAccount struct
type UserAccount struct {
	basic.Base
	basic.DataOwner
	UserAccountAPI
	Person *Person `json:"person,omitempty"`
}

// UserAccountAPI struct
type UserAccountAPI struct {
	Login                        *string          `json:"login,omitempty" gorm:"type:varchar(256);not null;check:chk_user_account__login_lowercase,coalesce(login=lower(login),true)=true;"`
	PersonID                     *uuid.UUID       `json:"person_id,omitempty" gorm:"type:varchar(36);"`
	Email                        *string          `json:"email,omitempty" gorm:"type:varchar(256);not null;index:idx_email_deleted_at,unique;where:deleted_at is null;check:chk_user_account__email_lowercase,coalesce(email=lower(email),true)=true;" validate:"required,email"`
	Password                     *string          `json:"password,omitempty" gorm:"type:varchar(256);not null;"`
	Salt                         *string          `json:"salt,omitempty" gorm:"type:varchar(256);not null;"`
	IsPasswordSystemGenerated    *bool            `json:"is_password_system_generated,omitempty" gorm:"default:false"`
	PasswordLastChange           *strfmt.DateTime `json:"password_last_change,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	PasswordExpiration           *strfmt.DateTime `json:"password_expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	InvitedAt                    *strfmt.DateTime `json:"invited_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ActivationCode               *string          `json:"activation_code,omitempty" gorm:"type:varchar(256);"`
	IsActivated                  *bool            `json:"is_activated,omitempty" gorm:"default:false"`
	ActivatedAt                  *strfmt.DateTime `json:"activated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	PersistCode                  *string          `json:"persist_code,omitempty" gorm:"type:varchar(256);"`
	SessionPersistenceExpiration *strfmt.DateTime `json:"session_persistence_expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	ResetPasswordCode            *string          `json:"reset_password_code,omitempty" gorm:"type:varchar(256);"`
	ResetPasswordExpiration      *strfmt.DateTime `json:"reset_password_expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	AccessCode                   *string          `json:"access_code,omitempty" gorm:"type:varchar(256);"`
	OTPEnabled                   *bool            `json:"otp_enabled,omitempty" gorm:"default:false"`
	OTPCode                      *string          `json:"otp_code,omitempty" gorm:"type:varchar(256);"`
	OTPExpiration                *strfmt.DateTime `json:"otp_expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	LastLogin                    *strfmt.DateTime `json:"last_login,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	LastAccess                   *strfmt.DateTime `json:"last_access,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	LastPage                     *string          `json:"last_page,omitempty" gorm:"type:varchar(256);"`
	IsSuperuser                  *bool            `json:"is_superuser,omitempty"`
	IsAnonymous                  *bool            `json:"is_anonymous,omitempty"`
	IsAPI                        *bool            `json:"is_api,omitempty"`
}

// Seed User Accounts
func (userAccount *UserAccount) Seed(salt, aes string) *UserAccount {
	password := "Password"
	saltCipherEncrypt, _ := lib.CipherEncrypt(salt, aes)
	passwordCipherEncrypt, _ := lib.CipherEncrypt(string(saltCipherEncrypt)+password, aes)
	seed := UserAccount{UserAccountAPI: UserAccountAPI{
		Login:    lib.Strptr("admin"),
		PersonID: lib.UUIDPtr(uuid.New()),
		Email:    lib.Strptr("admin@mail.com"),
		Password: lib.Strptr(base64.StdEncoding.EncodeToString(passwordCipherEncrypt)),
		Salt:     lib.Strptr(base64.StdEncoding.EncodeToString(saltCipherEncrypt)),
	}}
	lib.Merge(seed, &userAccount)
	return userAccount
}

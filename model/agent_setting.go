package model

import (
	"encoding/json"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// AgentSetting AgentSetting
type AgentSetting struct {
	basic.Base
	basic.DataOwner
	AgentSettingAPI
	Agent             *Agent             `json:"agent,omitempty"`
	OperationSchedule *OperationSchedule `json:"operation_schedule,omitempty" gorm:"foreignKey:AfterHourOperationScheduleID;references:ID"`
}

// AgentSettingAPI AgentSetting API
type AgentSettingAPI struct {
	AgentID                                           *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;unique"`
	FromEmail                                         *string    `json:"from_email,omitempty" gorm:"type:varchar(256);check:chk_agent_setting__from_email_lowercase,coalesce(from_email=lower(from_email),true)=true;" validate:"omitempty,email"`
	FromEmailDisplay                                  *string    `json:"from_email_display,omitempty" gorm:"type:varchar(256)"`
	ToEmail                                           *string    `json:"to_email,omitempty" gorm:"type:varchar(256);check:chk_agent_setting__to_email_lowercase,coalesce(to_email=lower(to_email),true)=true;" validate:"omitempty,email"`
	IsPasswordSystemGenerated                         *bool      `json:"is_password_system_generated,omitempty"`
	MaximumSessionAge                                 *int       `json:"maximum_session_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumSessionPersistenceAge                      *int       `json:"maximum_session_persistence_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumPasswordAge                                *int       `json:"maximum_password_age,omitempty" gorm:"type:smallint;" example:"1"`
	PasswordExpiryNoticePeriod                        *int       `json:"password_expiry_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumOneTimePasswordAge                         *int       `json:"maximum_one_time_password_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumOneTimePasswordPersistenceAge              *int       `json:"maximum_one_time_password_persistence_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumEmailSignInAge                             *int       `json:"maximum_email_sign_in_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumEmailUnsubscribeAge                        *int       `json:"maximum_email_unsubscribe_age,omitempty" gorm:"type:smallint;" example:"1"`
	UserActivationMode                                *string    `json:"user_activation_mode,omitempty" gorm:"type:varchar(36);" example:"none"` // option:'none','user','administrator'
	UserLoginMode                                     *string    `json:"user_login_mode,omitempty" gorm:"type:varchar(36);" example:"none"`      // option:'none','username','member_number'
	MaximumResetPasswordAge                           *int       `json:"maximum_reset_password_age,omitempty" gorm:"type:smallint;" example:"1"`
	EnforcePasswordHistory                            *int       `json:"enforce_password_history,omitempty" gorm:"type:smallint;" example:"1"`
	MinimumLoginAttemptsBeforeRecaptcha               *int       `json:"minimum_login_attempts_before_recaptcha,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumLoginAttempts                              *int       `json:"maximum_login_attempts,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumRequestAge                                 *int       `json:"maximum_request_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumRequests                                   *int       `json:"maximum_requests,omitempty" gorm:"type:smallint;" example:"1"`
	SuspensionTime                                    *int       `json:"suspension_time,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumQueueAge                                   *int       `json:"maximum_queue_age,omitempty" gorm:"type:smallint;" example:"1"`
	CorporateContractExpiryNoticePeriod               *int       `json:"corporate_contract_expiry_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	SignUpRequirement                                 *string    `json:"sign_up_requirement,omitempty" gorm:"type:text;"`
	MaximumCreditApprovalResponseTime                 *int       `json:"maximum_credit_approval_response_time,omitempty" gorm:"type:smallint;" example:"1"`
	CreditApprovalRule                                *string    `json:"credit_approval_rule,omitempty" gorm:"type:text" example:"1"`
	TicketingTimeLimitOffset                          *int       `json:"ticketing_time_limit_offset,omitempty" gorm:"type:smallint;" example:"1"`
	TicketingTimeLimitNoticePeriod                    *int       `json:"ticketing_time_limit_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	CancellationDeadlineOffset                        *int       `json:"cancellation_deadline_offset,omitempty" gorm:"type:smallint;" example:"1"`
	CancellationDeadlineNoticePeriod                  *int       `json:"cancellation_deadline_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	PassportExpiryNoticePeriod                        *int       `json:"passport_expiry_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumRebooks                                    *int       `json:"maximum_rebooks,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumVicinityDistance                           *int       `json:"maximum_vicinity_distance,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumWeatherForecastAge                         *int       `json:"maximum_weather_forecast_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumVicinityUnitOfMeasureID                    *uuid.UUID `json:"maximum_vicinity_unit_of_measure_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	RetailProcessingFeeCategoryID                     *uuid.UUID `json:"retail_processing_fee_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	OtherProcessingFeeCategoryID                      *uuid.UUID `json:"other_processing_fee_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	AfterHourOperationScheduleID                      *uuid.UUID `json:"after_hour_operation_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CorporateContractFirstReminderScheduleID          *uuid.UUID `json:"corporate_contract_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CorporateContractRecurrenceReminderScheduleID     *uuid.UUID `json:"corporate_contract_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TravelDocumentFirstReminderScheduleID             *uuid.UUID `json:"travel_document_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TravelDocumentRecurrenceReminderScheduleID        *uuid.UUID `json:"travel_document_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	InvoiceFirstReminderScheduleID                    *uuid.UUID `json:"invoice_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	InvoiceRecurrenceReminderScheduleID               *uuid.UUID `json:"invoice_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TransactionFirstReminderScheduleID                *uuid.UUID `json:"transaction_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TransactionRecurrenceReminderScheduleID           *uuid.UUID `json:"transaction_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteTransactionFirstReminderScheduleID      *uuid.UUID `json:"last_minute_transaction_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteTransactionRecurrenceReminderScheduleID *uuid.UUID `json:"last_minute_transaction_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	HoldingFirstReminderScheduleID                    *uuid.UUID `json:"holding_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	HoldingRecurrenceReminderScheduleID               *uuid.UUID `json:"holding_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteHoldingFirstReminderScheduleID          *uuid.UUID `json:"last_minute_holding_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteHoldingRecurrenceReminderScheduleID     *uuid.UUID `json:"last_minute_holding_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	ToEmailDisplay                                    *string    `json:"to_email_display,omitempty" gorm:"type:varchar(256)"`
}

// SignUpRequirement json model
type SignUpRequirement struct {
	Type    *string         `json:"type"`    // Type of the signup form like : member | corporeate | hotel
	Enabled *bool           `json:"enabled"` // Indicates whether self signup or registration is enabled.
	Fields  *[]SignUpFields `json:"fields"`  // Contains array of objects which have the following set of fields
}

// SignUpFields json model
type SignUpFields struct {
	Name     *string `json:"name"`     // The name of the field like : first_name | last_name | login | email | password | confirm_password | user_consent
	Display  *bool   `json:"display"`  // Indicates whether to display this field or not
	Required *bool   `json:"required"` // Indicates whether this field is required or not
}

// Seed Agent Setting
func (a *AgentSetting) Seed(agentID uuid.UUID, smtpEmail, smtpSenderName string) *AgentSetting {
	// Sign Up Requirements json object define
	fields := []SignUpFields{}
	fieldValue := SignUpFields{}
	fieldValue.Name = lib.Strptr("member")
	fieldValue.Display = lib.Boolptr(false)
	fieldValue.Required = lib.Boolptr(false)
	fields = append(fields, fieldValue)
	DefineSignUpRequirement := SignUpRequirement{
		Type:    lib.Strptr("member"),
		Enabled: lib.Boolptr(true),
		Fields:  &fields,
	}
	jsonData, _ := json.Marshal(DefineSignUpRequirement)
	SignUpRequirement := lib.Strptr(string(jsonData))

	if lib.IsEmptyUUID(agentID) {
		agentID = uuid.New()
	}
	seed := AgentSetting{
		AgentSettingAPI: AgentSettingAPI{
			AgentID:                              &agentID,
			IsPasswordSystemGenerated:            lib.Boolptr(false),
			FromEmail:                            lib.Strptr(smtpEmail),
			FromEmailDisplay:                     lib.Strptr(smtpSenderName),
			ToEmail:                              lib.Strptr(smtpEmail),
			ToEmailDisplay:                       lib.Strptr(smtpSenderName),
			MaximumSessionAge:                    lib.Intptr(360), // in minutes
			MaximumSessionPersistenceAge:         lib.Intptr(2),   // in days
			MaximumPasswordAge:                   lib.Intptr(100), // in days
			PasswordExpiryNoticePeriod:           lib.Intptr(2),   // in days
			MaximumOneTimePasswordAge:            lib.Intptr(2),
			MaximumOneTimePasswordPersistenceAge: lib.Intptr(5),
			MaximumEmailSignInAge:                lib.Intptr(10),
			MaximumEmailUnsubscribeAge:           lib.Intptr(10),
			UserActivationMode:                   lib.Strptr("none"),
			UserLoginMode:                        lib.Strptr("none"),
			MaximumResetPasswordAge:              lib.Intptr(5),
			EnforcePasswordHistory:               lib.Intptr(3),
			MinimumLoginAttemptsBeforeRecaptcha:  lib.Intptr(3),
			MaximumLoginAttempts:                 lib.Intptr(3), // on trigger
			MaximumRequestAge:                    lib.Intptr(3),
			MaximumRequests:                      lib.Intptr(3),
			SuspensionTime:                       lib.Intptr(3),
			MaximumQueueAge:                      lib.Intptr(3),
			CorporateContractExpiryNoticePeriod:  lib.Intptr(3),
			SignUpRequirement:                    SignUpRequirement,
			TicketingTimeLimitOffset:             lib.Intptr(3503),
			TicketingTimeLimitNoticePeriod:       lib.Intptr(4233),
			CancellationDeadlineOffset:           lib.Intptr(5656),
			CancellationDeadlineNoticePeriod:     lib.Intptr(5676),
			PassportExpiryNoticePeriod:           lib.Intptr(1),
		},
	}
	_ = lib.Merge(seed, &a)
	return a
}

// ConfigurationBookingAPI ConfigurationBooking API
type ConfigurationBookingAPI struct {
	AgentID                                           *uuid.UUID       `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;unique"`
	FromEmail                                         *string          `json:"from_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	FromEmailDisplay                                  *string          `json:"from_email_display,omitempty" gorm:"type:varchar(256)"`
	ToEmail                                           *string          `json:"to_email,omitempty" gorm:"type:varchar(256)" validate:"omitempty,email"`
	ToEmailDisplay                                    *string          `json:"to_email_display,omitempty" gorm:"type:varchar(256)"`
	IsPasswordSystemGenerated                         *bool            `json:"is_password_system_generated,omitempty"`
	MaximumSessionAge                                 *int             `json:"maximum_session_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumSessionPersistenceAge                      *int             `json:"maximum_session_persistence_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumPasswordAge                                *int             `json:"maximum_password_age,omitempty" gorm:"type:smallint;" example:"1"`
	PasswordExpiryNoticePeriod                        *int             `json:"password_expiry_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumOneTimePasswordAge                         *int             `json:"maximum_one_time_password_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumOneTimePasswordPersistenceAge              *int             `json:"maximum_one_time_password_persistence_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumEmailSignInAge                             *int             `json:"maximum_email_sign_in_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumEmailUnsubscribeAge                        *int             `json:"maximum_email_unsubscribe_age,omitempty" gorm:"type:smallint;" example:"1"`
	UserActivationMode                                *string          `json:"user_activation_mode,omitempty" gorm:"type:varchar(36);" example:"none"` // option:'none','user','administrator'
	UserLoginMode                                     *string          `json:"user_login_mode,omitempty" gorm:"type:varchar(36);" example:"none"`      // option:'none','username','member_number'
	MaximumResetPasswordAge                           *int             `json:"maximum_reset_password_age,omitempty" gorm:"type:smallint;" example:"1"`
	EnforcePasswordHistory                            *int             `json:"enforce_password_history,omitempty" gorm:"type:smallint;" example:"1"`
	MinimumLoginAttemptsBeforeRecaptcha               *int             `json:"minimum_login_attempts_before_recaptcha,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumLoginAttempts                              *int             `json:"maximum_login_attempts,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumRequestAge                                 *int             `json:"maximum_request_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumRequests                                   *int             `json:"maximum_requests,omitempty" gorm:"type:smallint;" example:"1"`
	SuspensionTime                                    *int             `json:"suspension_time,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumQueueAge                                   *int             `json:"maximum_queue_age,omitempty" gorm:"type:smallint;" example:"1"`
	CorporateContractExpiryNoticePeriod               *int             `json:"corporate_contract_expiry_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	SignUpRequirement                                 *string          `json:"sign_up_requirement,omitempty" gorm:"type:text;"`
	MaximumCreditApprovalResponseTime                 *int             `json:"maximum_credit_approval_response_time,omitempty" gorm:"type:smallint;" example:"1"`
	CreditApprovalRule                                *int             `json:"credit_approval_rule,omitempty" gorm:"type:smallint;" example:"1"`
	TicketingTimeLimitOffset                          *int             `json:"ticketing_time_limit_offset,omitempty" gorm:"type:smallint;" example:"1"`
	TicketingTimeLimitNoticePeriod                    *int             `json:"ticketing_time_limit_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	CancellationDeadlineOffset                        *int             `json:"cancellation_deadline_offset,omitempty" gorm:"type:smallint;" example:"1"`
	CancellationDeadlineNoticePeriod                  *int             `json:"cancellation_deadline_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	PassportExpiryNoticePeriod                        *int             `json:"passport_expiry_notice_period,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumRebooks                                    *int             `json:"maximum_rebooks,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumVicinityDistance                           *int             `json:"maximum_vicinity_distance,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumWeatherForecastAge                         *int             `json:"maximum_weather_forecast_age,omitempty" gorm:"type:smallint;" example:"1"`
	MaximumVicinityUnitOfMeasureID                    *uuid.UUID       `json:"maximum_vicinity_unit_of_measure_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	RetailProcessingFeeCategoryID                     *uuid.UUID       `json:"retail_processing_fee_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	OtherProcessingFeeCategoryID                      *uuid.UUID       `json:"other_processing_fee_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	AfterHourOperationScheduleID                      *uuid.UUID       `json:"after_hour_operation_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CorporateContractFirstReminderScheduleID          *uuid.UUID       `json:"corporate_contract_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	CorporateContractRecurrenceReminderScheduleID     *uuid.UUID       `json:"corporate_contract_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TravelDocumentFirstReminderScheduleID             *uuid.UUID       `json:"travel_document_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TravelDocumentRecurrenceReminderScheduleID        *uuid.UUID       `json:"travel_document_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	InvoiceFirstReminderScheduleID                    *uuid.UUID       `json:"invoice_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	InvoiceRecurrenceReminderScheduleID               *uuid.UUID       `json:"invoice_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TransactionFirstReminderScheduleID                *uuid.UUID       `json:"transaction_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	TransactionRecurrenceReminderScheduleID           *uuid.UUID       `json:"transaction_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteTransactionFirstReminderScheduleID      *uuid.UUID       `json:"last_minute_transaction_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteTransactionRecurrenceReminderScheduleID *uuid.UUID       `json:"last_minute_transaction_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	HoldingFirstReminderScheduleID                    *uuid.UUID       `json:"holding_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	HoldingRecurrenceReminderScheduleID               *uuid.UUID       `json:"holding_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteHoldingFirstReminderScheduleID          *uuid.UUID       `json:"last_minute_holding_first_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	LastMinuteHoldingRecurrenceReminderScheduleID     *uuid.UUID       `json:"last_minute_holding_recurrence_reminder_schedule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	OperationTime                                     *[]OperationTime `json:"operation_time,omitempty"`
}

// ConfigurationsRetailAncillaryFeeDataGet struct
type ConfigurationsRetailAncillaryFeeDataGet struct {
	AgentSetting
	ConfigurationRetailAncillaryFeeFlight *ConfigurationRetailAncillaryFeeFlight `json:"configuration_retail_ancillary_fee_flight,omitempty" gorm:"-"`
	ConfigurationRetailAncillaryFeeHotel  *ConfigurationRetailAncillaryFeeHotel  `json:"configuration_retail_ancillary_fee_hotel,omitempty" gorm:"-"`
	DomesticReissueID                     *uuid.UUID                             `json:"domestic_reissue_id,omitempty" swaggerignore:"true"`
	DomesticReissue                       *ProcessingFeeRate                     `json:"-"`
	DomesticRevalidateID                  *uuid.UUID                             `json:"domestic_revalidate_id,omitempty" swaggerignore:"true"`
	DomesticRevalidate                    *ProcessingFeeRate                     `json:"-"`
	DomesticRefundID                      *uuid.UUID                             `json:"domestic_refund_id,omitempty" swaggerignore:"true"`
	DomesticRefund                        *ProcessingFeeRate                     `json:"-"`
	DomesticVoidID                        *uuid.UUID                             `json:"domestic_void_id,omitempty" swaggerignore:"true"`
	DomesticVoid                          *ProcessingFeeRate                     `json:"-"`
	DomesticRFPId                         *uuid.UUID                             `json:"domestic_rfp_id,omitempty" swaggerignore:"true"`
	DomesticRFP                           *ProcessingFeeRate                     `json:"-"`
	DomesticNonGDSId                      *uuid.UUID                             `json:"domestic_non_gds_id,omitempty" swaggerignore:"true"`
	DomesticNonGDS                        *ProcessingFeeRate                     `json:"-"`
	InternationalReissueID                *uuid.UUID                             `json:"international_reissue_id,omitempty" swaggerignore:"true"`
	InternationalReissue                  *ProcessingFeeRate                     `json:"-"`
	InternationalRevalidateID             *uuid.UUID                             `json:"international_revalidate_id,omitempty" swaggerignore:"true"`
	InternationalRevalidate               *ProcessingFeeRate                     `json:"-"`
	InternationalRefundID                 *uuid.UUID                             `json:"international_refund_id,omitempty" swaggerignore:"true"`
	InternationalRefund                   *ProcessingFeeRate                     `json:"-"`
	InternationalVoidID                   *uuid.UUID                             `json:"international_void_id,omitempty" swaggerignore:"true"`
	InternationalVoid                     *ProcessingFeeRate                     `json:"-"`
	InternationalRFPId                    *uuid.UUID                             `json:"international_rfp_id,omitempty" swaggerignore:"true"`
	InternationalRFP                      *ProcessingFeeRate                     `json:"-"`
	InternationalNonGDSId                 *uuid.UUID                             `json:"international_non_gds_id,omitempty" swaggerignore:"true"`
	InternationalNonGDS                   *ProcessingFeeRate                     `json:"-"`
	EmergencyServiceID                    *uuid.UUID                             `json:"emergency_service_id,omitempty" swaggerignore:"true"`
	EmergencyService                      *ProcessingFeeRate                     `json:"-"`
	HotelDomesticMhfID                    *uuid.UUID                             `json:"hotel_domestic_modify_fee_id,omitempty" swaggerignore:"true"`
	HotelDomesticMhf                      *ProcessingFeeRate                     `json:"-"`
	HotelDomesticRfId                     *uuid.UUID                             `json:"hotel_domestic_refund_fee_id,omitempty" swaggerignore:"true"`
	HotelDomesticRf                       *ProcessingFeeRate                     `json:"-"`
	HotelDomesticRFPId                    *uuid.UUID                             `json:"hotel_domestic_rfp_id,omitempty" swaggerignore:"true"`
	HotelDomesticRFP                      *ProcessingFeeRate                     `json:"-"`
	HotelDomesticNonGDSId                 *uuid.UUID                             `json:"hotel_domestic_non_gds_id,omitempty" swaggerignore:"true"`
	HotelDomesticNonGDS                   *ProcessingFeeRate                     `json:"-"`
	HotelInterMhfID                       *uuid.UUID                             `json:"hotel_international_modify_fee_id,omitempty" swaggerignore:"true"`
	HotelInterMhf                         *ProcessingFeeRate                     `json:"-"`
	HotelInterRfID                        *uuid.UUID                             `json:"hotel_international_refund_fee_id,omitempty" swaggerignore:"true"`
	HotelInterRf                          *ProcessingFeeRate                     `json:"-"`
	HotelInterRFPId                       *uuid.UUID                             `json:"hotel_international_rfp_id,omitempty" swaggerignore:"true"`
	HotelInterRFP                         *ProcessingFeeRate                     `json:"-"`
	HotelInterNonGDSId                    *uuid.UUID                             `json:"hotel_international_non_gds_id,omitempty" swaggerignore:"true"`
	HotelInterNonGDS                      *ProcessingFeeRate                     `json:"-"`
	HotelEmergencyServiceID               *uuid.UUID                             `json:"hotel_emergency_service_id,omitempty" swaggerignore:"true"`
	HotelEmergencyService                 *ProcessingFeeRate                     `json:"-"`
}

// ConfigurationsRetailAncillaryFeeAPI ConfigurationsRetailAncillaryFee API
type ConfigurationsRetailAncillaryFeeAPI struct {
	DomesticReissue            *ProcessingFeeRateAPI `json:"domestic_reissue,omitempty" gorm:"-"`
	DomesticRevalidate         *ProcessingFeeRateAPI `json:"domestic_revalidate,omitempty" gorm:"-"`
	DomesticRefund             *ProcessingFeeRateAPI `json:"domestic_refund,omitempty" gorm:"-"`
	DomesticVoid               *ProcessingFeeRateAPI `json:"domestic_void,omitempty" gorm:"-"`
	DomesticFrp                *ProcessingFeeRateAPI `json:"domestic_rfp,omitempty" gorm:"-"`
	DomesticNonGds             *ProcessingFeeRateAPI `json:"domestic_non_gds,omitempty" gorm:"-"`
	InternationalReissue       *ProcessingFeeRateAPI `json:"international_reissue,omitempty" gorm:"-"`
	InternationalRevalidate    *ProcessingFeeRateAPI `json:"international_revalidate,omitempty" gorm:"-"`
	InternationalRefund        *ProcessingFeeRateAPI `json:"international_refund,omitempty" gorm:"-"`
	InternationalVoid          *ProcessingFeeRateAPI `json:"international_void,omitempty" gorm:"-"`
	InternationalFrp           *ProcessingFeeRateAPI `json:"international_rfp,omitempty" gorm:"-"`
	InternationalNonGds        *ProcessingFeeRateAPI `json:"international_non_gds,omitempty" gorm:"-"`
	OtherEmergencyService      *ProcessingFeeRateAPI `json:"other_emergency_service,omitempty" gorm:"-"`
	HotelDomesticMHF           *ProcessingFeeRateAPI `json:"hotel_domestic_modify_fee,omitempty" gorm:"-"`
	HotelDomesticRF            *ProcessingFeeRateAPI `json:"hotel_domestic_refund_fee,omitempty" gorm:"-"`
	HotelDomesticFrp           *ProcessingFeeRateAPI `json:"hotel_domestic_rfp,omitempty" gorm:"-"`
	HotelDomesticNonGds        *ProcessingFeeRateAPI `json:"hotel_domestic_non_gds,omitempty" gorm:"-"`
	HotelInternationalMHF      *ProcessingFeeRateAPI `json:"hotel_international_modify_fee,omitempty" gorm:"-"`
	HotelInternationalRF       *ProcessingFeeRateAPI `json:"hotel_international_refund_fee,omitempty" gorm:"-"`
	HotelInternationalFrp      *ProcessingFeeRateAPI `json:"hotel_international_rfp,omitempty" gorm:"-"`
	HotelInternationalNonGds   *ProcessingFeeRateAPI `json:"hotel_international_non_gds,omitempty" gorm:"-"`
	HotelOtherEmergencyService *ProcessingFeeRateAPI `json:"hotel_other_emergency_service,omitempty" gorm:"-"`
}

// ConfigurationRetailAncillaryFeeFlight Configuration Retail Ancillary Fee Flight
type ConfigurationRetailAncillaryFeeFlight struct {
	DomesticReissueID         *uuid.UUID         `json:"domestic_reissue_id,omitempty"`
	DomesticReissue           *ProcessingFeeRate `json:"domestic_reissue,omitempty"`
	DomesticRevalidateID      *uuid.UUID         `json:"domestic_revalidate_id,omitempty"`
	DomesticRevalidate        *ProcessingFeeRate `json:"domestic_revalidate,omitempty"`
	DomesticRefundID          *uuid.UUID         `json:"domestic_refund_id,omitempty"`
	DomesticRefund            *ProcessingFeeRate `json:"domestic_refund,omitempty"`
	DomesticVoidID            *uuid.UUID         `json:"domestic_void_id,omitempty"`
	DomesticVoid              *ProcessingFeeRate `json:"domestic_void,omitempty"`
	DomesticRFPId             *uuid.UUID         `json:"domestic_rfp_id,omitempty"`
	DomesticRFP               *ProcessingFeeRate `json:"domestic_rfp,omitempty"`
	DomesticNonGDSId          *uuid.UUID         `json:"domestic_non_gds_id,omitempty"`
	DomesticNonGDS            *ProcessingFeeRate `json:"domestic_non_gds,omitempty"`
	InternationalReissueID    *uuid.UUID         `json:"international_reissue_id,omitempty"`
	InternationalReissue      *ProcessingFeeRate `json:"international_reissue,omitempty"`
	InternationalRevalidateID *uuid.UUID         `json:"international_revalidate_id,omitempty"`
	InternationalRevalidate   *ProcessingFeeRate `json:"international_revalidate,omitempty"`
	InternationalRefundID     *uuid.UUID         `json:"international_refund_id,omitempty"`
	InternationalRefund       *ProcessingFeeRate `json:"international_refund,omitempty"`
	InternationalVoidID       *uuid.UUID         `json:"international_void_id,omitempty"`
	InternationalVoid         *ProcessingFeeRate `json:"international_void,omitempty"`
	InternationalRFPId        *uuid.UUID         `json:"international_rfp_id,omitempty"`
	InternationalRFP          *ProcessingFeeRate `json:"international_rfp,omitempty"`
	InternationalNonGDSId     *uuid.UUID         `json:"international_non_gds_id,omitempty"`
	InternationalNonGDS       *ProcessingFeeRate `json:"international_non_gds,omitempty"`
	EmergencyServiceID        *uuid.UUID         `json:"emergency_service_id,omitempty"`
	EmergencyService          *ProcessingFeeRate `json:"emergency_service,omitempty"`
}

// ConfigurationRetailAncillaryFeeHotel Configuration Retail Ancillary Fee Hotel
type ConfigurationRetailAncillaryFeeHotel struct {
	HotelDomesticMhfID      *uuid.UUID         `json:"hotel_domestic_modify_fee_id,omitempty"`
	HotelDomesticMhf        *ProcessingFeeRate `json:"hotel_domestic_modify_fee,omitempty"`
	HotelDomesticRfId       *uuid.UUID         `json:"hotel_domestic_refund_fee_id,omitempty"`
	HotelDomesticRf         *ProcessingFeeRate `json:"hotel_domestic_refund_fee,omitempty"`
	HotelDomesticRFPId      *uuid.UUID         `json:"hotel_domestic_rfp_id,omitempty"`
	HotelDomesticRFP        *ProcessingFeeRate `json:"hotel_domestic_rfp,omitempty"`
	HotelDomesticNonGDSId   *uuid.UUID         `json:"hotel_domestic_non_gds_id,omitempty"`
	HotelDomesticNonGDS     *ProcessingFeeRate `json:"hotel_domestic_non_gds,omitempty"`
	HotelInterMhfID         *uuid.UUID         `json:"hotel_international_modify_fee_id,omitempty"`
	HotelInterMhf           *ProcessingFeeRate `json:"hotel_international_modify_fee,omitempty"`
	HotelInterRfID          *uuid.UUID         `json:"hotel_international_refund_fee_id,omitempty"`
	HotelInterRf            *ProcessingFeeRate `json:"hotel_international_refund_fee,omitempty"`
	HotelInterRFPId         *uuid.UUID         `json:"hotel_international_rfp_id,omitempty"`
	HotelInterRFP           *ProcessingFeeRate `json:"hotel_international_rfp,omitempty"`
	HotelInterNonGDSId      *uuid.UUID         `json:"hotel_international_non_gds_id,omitempty"`
	HotelInterNonGDS        *ProcessingFeeRate `json:"hotel_international_non_gds,omitempty"`
	HotelEmergencyServiceID *uuid.UUID         `json:"hotel_emergency_service_id,omitempty"`
	HotelEmergencyService   *ProcessingFeeRate `json:"hotel_emergency_service,omitempty"`
}

// AgentSettingGetData Agent Setting Get Data
type AgentSettingGetData struct {
	AgentSetting
	OperationTime []*OperationTime `json:"operation_time" gorm:"-"`
}

// AgentEmailSenderRecipientDefaultAPI API
type AgentEmailSenderRecipientDefaultAPI struct {
	FromEmail        *string `json:"from_email,omitempty" validate:"omitempty,email"`
	FromEmailDisplay *string `json:"from_email_display,omitempty" validate:"omitempty,alphanumunicodespace"`
	ToEmail          *string `json:"to_email,omitempty" validate:"omitempty,email"`
	ToEmailDisplay   *string `json:"to_email_display,omitempty" validate:"omitempty,alphanumunicodespace"`
}

// AgentEmailSenderRecipientDefaultData Agent Email Sender Recipient Default Get Data
type AgentEmailSenderRecipientDefaultData struct {
	ID               *uuid.UUID `json:"id,omitempty"`
	FromEmail        *string    `json:"from_email,omitempty" validate:"omitempty,email"`
	FromEmailDisplay *string    `json:"from_email_display,omitempty" validate:"omitempty,alphanumunicodespace"`
	ToEmail          *string    `json:"to_email,omitempty" validate:"omitempty,email"`
	ToEmailDisplay   *string    `json:"to_email_display,omitempty" validate:"omitempty,alphanumunicodespace"`
}

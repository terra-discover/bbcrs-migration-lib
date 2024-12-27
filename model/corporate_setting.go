package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateSetting Model
type CorporateSetting struct {
	basic.Base
	basic.DataOwner
	CorporateSettingAPI
	Corporate    *Corporate    `json:"corporate,omitempty"`     // Corporate
	Division     *Division     `json:"division,omitempty"`      // Division
	ServiceLevel *ServiceLevel `json:"service_level,omitempty"` // ServiceLevel
	//ApplyBundlingRate *ApplyBundlingRate `json:"apply_bundling_rate,omitempty"`
}

// CorporateSettingAPI Model
type CorporateSettingAPI struct {
	CorporateID                               *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null;index:unique,where:deleted_at is null;" swaggertype:"string" format:"uuid"` // Corporate ID
	DivisionID                                *uuid.UUID `json:"division_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                  // Division ID
	ServiceLevelID                            *uuid.UUID `json:"service_level_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                             // Service Level ID
	ApplyBundlingRateID                       *uuid.UUID `json:"apply_bundling_rate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                       // Apply Bundling Rate ID
	ProposalPrefixCode                        *string    `json:"proposal_prefix_code,omitempty" gorm:"type:varchar(8);not null;index:unique,where:deleted_at is null;"`                             // Proposal Prefix Code. Must be maximum 8 characters, alphanumeric, uppercase, unique.
	HasSelfBookingTool                        *bool      `json:"has_self_booking_tool,omitempty"`                                                                                                   // Has Self Booking Tool
	TravelerCanSelfBook                       *bool      `json:"traveler_can_self_book,omitempty"`                                                                                                  // TravelerCanSelfBook
	MaximumSelfBookingAmount                  *float64   `json:"maximum_self_booking_amount,omitempty"`                                                                                             // MaximumSelfBookingAmount
	AllowPersonalTravel                       *bool      `json:"allow_personal_travel,omitempty"`                                                                                                   // AllowPersonalTravel
	NeedTravelRequest                         *bool      `json:"need_travel_request,omitempty"`                                                                                                     // NeedTravelRequest
	SelfBookNeedTravelRequest                 *bool      `json:"self_book_need_travel_request,omitempty"`                                                                                           // SelfBookNeedTravelRequest
	MaximumOutstandingTravelRequestsPerBooker *int       `json:"maximum_outstanding_travel_requests_per_booker,omitempty"`                                                                          // MaximumOutstandingTravelRequestsPerBooker
	SelfBookCanHoldTicket                     *bool      `json:"self_book_can_hold_ticket,omitempty"`                                                                                               // SelfBookCanHoldTicket
	SelfBookCanIssueTicket                    *bool      `json:"self_book_can_issue_ticket,omitempty"`                                                                                              // SelfBookCanIssueTicket
	TravelRequestVerificationRemark           *string    `json:"travel_request_verification_remark,omitempty" gorm:"type:varchar(4000)"`                                                            // TravelRequestVerificationRemark
	NeedTravelPurpose                         *bool      `json:"need_travel_purpose,omitempty"`                                                                                                     // NeedTravelPurpose
	ApplyBundlingServiceFee                   *bool      `json:"apply_bundling_service_fee,omitempty"`                                                                                              // ApplyBundlingServiceFee
	SharingPercent                            *float64   `json:"sharing_percent,omitempty"`                                                                                                         // SharingPercent
	SelfBookSharingPercent                    *float64   `json:"self_book_sharing_percent,omitempty"`                                                                                               // SelfBookSharingPercent
	PersonalTravelSharingPercent              *float64   `json:"personal_travel_sharing_percent,omitempty"`                                                                                         // PersonalTravelSharingPercent
	NeedNonGdsFlightProcessing                *bool      `json:"need_non_gds_flight_processing,omitempty"`                                                                                          // NeedNonGdsFlightProcessing
	NeedNonGdsHotelProcessing                 *bool      `json:"need_non_gds_hotel_processing,omitempty"`                                                                                           // NeedNonGdsHotelProcessing
	InvoiceDueDays                            *int       `json:"invoice_due_days,omitempty"`                                                                                                        // InvoiceDueDays
	InvoiceDueDaysTolerance                   *int       `json:"invoice_due_days_tolerance,omitempty"`                                                                                              // InvoiceDueDaysTolerance
	InvoiceDueDateFrom                        *uuid.UUID `json:"invoice_due_date_from,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                        // InvoiceDueDateFrom
	MaximumOverdueInvoices                    *int       `json:"maximum_overdue_invoices,omitempty"`                                                                                                // MaximumOverdueInvoices
	PayInvoiceWithCreditCard                  *bool      `json:"pay_invoice_with_credit_card,omitempty"`                                                                                            // PayInvoiceWithCreditCard
	TravelBookerNeedInvoice                   *bool      `json:"travel_booker_need_invoice,omitempty"`                                                                                              // TravelBookerNeedInvoice
	TravelerNeedInvoice                       *bool      `json:"traveler_need_invoice,omitempty"`                                                                                                   // TravelerNeedInvoice
	PersonalTravelerNeedInvoice               *bool      `json:"personal_traveler_need_invoice,omitempty"`                                                                                          // PersonalTravelerNeedInvoice
	InvoiceByPostMail                         *bool      `json:"invoice_by_post_mail,omitempty"`                                                                                                    // InvoiceByPostMail
	InvoiceSeparateServiceFee                 *bool      `json:"invoice_separate_service_fee,omitempty"`                                                                                            // InvoiceSeparateServiceFee
	InvoiceSeparateProductType                *bool      `json:"invoice_separate_product_type,omitempty"`                                                                                           // InvoiceSeparateProductType
	NeedTaxInvoice                            *bool      `json:"need_tax_invoice,omitempty"`                                                                                                        // NeedTaxInvoice
	InvoiceBatches                            *int       `json:"invoice_batches,omitempty"`                                                                                                         // InvoiceBatches
	InvoiceBatchesPlusDays                    *int       `json:"invoice_batches_plus_days,omitempty"`                                                                                               // InvoiceBatchesPlusDays
	InvoiceBatchesOnWorkingDays               *bool      `json:"invoice_batches_on_working_days,omitempty"`                                                                                         // InvoiceBatchesOnWorkingDays
	ApplyStrictestTravelPolicy                *bool      `json:"apply_strictest_travel_policy,omitempty"`                                                                                           // ApplyStrictestTravelPolicy
	StorePaymentCardData                      *bool      `json:"store_payment_card_data,omitempty"`                                                                                                 // StorePaymentCardData
	Reimbursable                              *bool      `json:"reimbursable,omitempty"`                                                                                                            // Reimbursable
}

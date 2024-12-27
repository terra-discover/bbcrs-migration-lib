package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// QueueType Queue Type
type QueueType struct {
	basic.Base
	basic.DataOwner
	QueueTypeAPI
	QueueTypeTranslation *QueueTypeTranslation `json:"queue_type_translation,omitempty"` // Queue Type Translation
}

// QueueTypeAPI Queue Type API
type QueueTypeAPI struct {
	QueueTypeCode *int    `json:"queue_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`                         // Queue Type Code
	QueueTypeName *string `json:"queue_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"After Save Corporate "` // Queue Type Name
	Count         *int    `json:"count,omitempty" gorm:"type:int;default:0" example:"0"`                                                                              // Count
}

// Seed data
func (s QueueType) Seed() *[]QueueType {
	data := []QueueType{}
	items := []string{
		"After Save Corporate",
		"After Purchase",
		"After Booking",
		"After Cancellation",
		"Request for Over Credit Authorization",
		"Promotion",
		"Request for Proposal Review",
		"Reminder for Request for Proposal Review",
		"Newsletter",
		"Survey",
		"System",
		"Transactional",
		"Invoice",
		"Request for Travel Request Submission",
		"Reminder for Travel Request Submission",
		"Request for Travel Request Revision",
		"Request for Credit Card Online Payment",
		"Reminder for Credit Card Online Payment",
		"Approved to Issue",
		"Reminder for Approved to Issue",
		"Ticketing Time Limit",
		"Reminder for Ticketing Time Limit",
		"Booking Confirmation",
		"Flight Ticket",
		"Hotel Voucher",
		"Request for Payment Settlement",
		"Reminder for Payment Settlement",
		"Thank You and Survey",
		"During Stay Email",
		"Flight Booking Confirmation",
		"Hotel Booking Confirmation",
		"Flight Issuance Confirmation",
		"Internal Flight Booking Confirmation",
		"Internal Hotel Booking Confirmation",
		"Internal Flight Issuance Confirmation",
		"Cancellation Confirmation",
		"Flight Cancellation Confirmation",
		"Hotel Cancellation Confirmation",
		"Internal Flight Cancellation Confirmation",
		"Internal Hotel Cancellation Confirmation",
		"Internal Ticketing Time Limit",
		"Reminder for Internal Ticketing Time Limit",
		"General Inquiry",
		"Feedback",
		"Survey Response",
		"Internal User Activation",
		"Internal User Welcome",
		"Corporate User Activation",
		"Corporate User Welcome",
		"Member Activation",
		"Member Welcome",
		"Password Change",
		"Password Restore",
		"One Time Password",
		"Invoice Per Transaction",
		"Invoice Batch",
		"Request for Over Credit Authorization Approved",
		"Request for Over Credit Authorization Rejected",
		"Cancellation Deadline",
		"Reminder for Cancellation Deadline",
		"Internal Cancellation Deadline",
		"Reminder for Internal Cancellation Deadline",
		"Ticket Not Found",
		"Reminder for Invoice Per Transaction",
		"Reminder for Invoice Batch",
		"Reminder for Request for Over Credit Authorization",
		"Transactional - Flight",
		"Transactional - Hotel",
		"Transactional - Tour",
		"Reminder for Travel Request Revision",
		"Send to Mail",
		"Flight Reservation to TMS",
		"Hotel Reservation to TMS",
		"Flight Cancellation to TMS",
		"Hotel Cancellation to TMS",
		"Schedule Change - Change Flight Number",
		"Schedule Change - Flight Retimed",
		"Travel Request Submission",
		"Proposal Option to Book",
		"Reminder for Proposal Option to Book",
		"Invoice Per Transaction - Flight",
		"Invoice Per Transaction - Hotel",
		"Invoice Per Transaction - Other",
		"Invoice Batch - Flight",
		"Invoice Batch - Hotel",
		"Invoice Batch - Other",
		"Reminder for Invoice Per Transaction - Flight",
		"Reminder for Invoice Per Transaction - Hotel",
		"Reminder for Invoice Per Transaction - Other",
		"Reminder for Invoice Batch - Flight",
		"Reminder for Invoice Batch - Hotel",
		"Reminder for Invoice Batch - Other",
		"Passport Expiration",
		"Reminder for Passport Expiration",
		"Corporate Contract Expiration",
		"Reminder for Corporate Contract Expiration",
		"Request for Hotel Booking Acknowledgment",
		"Reminder for Hotel Booking Acknowledgment",
		"Hotel Booking Acknowledged",
		"Internal Hotel Booking Acknowledged",
		"Corporate Data Verification",
		"Request to Reissue Ticket",
		"Reminder for Request to Reissue Ticket",
		"Request to Revalidate Ticket",
		"Reminder for Request to Revalidate Ticket",
		"Request to Reissue Ticket (Name Change)",
		"Reminder for Request to Reissue Ticket (Name Change)",
		"Request to Modify Hotel Voucher",
		"Reminder for Request to Modify Hotel Voucher",
		"Request to Reissue Ticket Completed",
		"Request to Revalidate Ticket Completed",
		"Request to Reissue Ticket (Name Change) Completed",
		"Request to Modify Hotel Voucher Completed",
		"Invoice Per Transaction - Tour",
		"Invoice Per Transaction - Document",
		"Invoice Batch - Tour",
		"Invoice Batch - Document",
		"Reminder for Invoice Per Transaction - Tour",
		"Reminder for Invoice Per Transaction - Document",
		"Reminder for Invoice Batch - Tour",
		"Reminder for Invoice Batch - Document",
		"Credit Limit Almost Used Up",
		"Flight Itinerary",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, QueueType{
			QueueTypeAPI: QueueTypeAPI{
				QueueTypeCode: &code,
				QueueTypeName: &name,
			},
		})
	}

	return &data
}

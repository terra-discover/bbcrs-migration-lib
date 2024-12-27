package model

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Status Status
type Status struct {
	basic.Base
	basic.DataOwner
	StatusAPI
	StatusTranslation *StatusTranslation `json:"status_translation,omitempty"` // Status Translation
	StatusCategory    *StatusCategory    `json:"status_category,omitempty"`    // Status Category
}

// StatusAPI Status API
type StatusAPI struct {
	StatusCode         *int       `json:"status_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`                               // Status Code
	StatusName         *string    `json:"status_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Cancel/confirmed/requested "` // Status Name
	InternalStatusName *string    `json:"internal_status_name,omitempty" gorm:"type:varchar(256)"`                                                                              // Internal Status Name
	ConsumerStatusName *string    `json:"consumer_status_name,omitempty" gorm:"type:varchar(256)"`                                                                              // Consumer Status Name
	StatusCategoryID   *uuid.UUID `json:"status_category_id,omitempty" swaggertype:"string" format:"uuid"`                                                                      // Status Category ID
}

// Seed data
func (s Status) Seed() *[]Status {
	data := []Status{}
	items := []string{
		"1|Cancel/confirmed/requested||",
		"2|Cancel listing||",
		"3|Cancel only if requested segment is available||",
		"4|Cancellation recommended||",
		"5|Code to report action taken as a result of previous access||",
		"6|If holding, cancel||",
		"7|If not holding, need||",
		"8|If not holding, sold||",
		"9|List (add to waiting list)||",
		"10|List space available||",
		"11|Need||",
		"12|Need segment specified or alternative||",
		"13|Sold||",
		"14|Sold (on free sale basis)||",
		"15|Space requested||",
		"16|Cancelled||",
		"17|Confirming||",
		"18|Confirming from waitlist||",
		"19|No action taken||",
		"20|Unable - have waitlisted||",
		"21|Unable - flight does not operate||",
		"22|Unable to accept sale||",
		"23|Waitlist closed||",
		"24|Have listed||",
		"25|Have requested||",
		"26|Holds confirmed|Ready To Issue|Issue In-Progress",
		"27|Reconfirmed||",
		"28|Have sold||",
		"29|Space already requested||",
		"30|OK||",
		"31|Have cancelled||",
		"32|Pending confirmation||",
		"33|Deferred from wait list||",
		"34|Staff request||",
		"35|Active||",
		"36|Available||",
		"37|Book on board||",
		"38|Closed||",
		"39|Confirmed||",
		"40|Declined||",
		"41|Not applicable||",
		"42|Offered||",
		"43|On request||",
		"44|Pre-paid||",
		"45|Shared||",
		"46|Waitlisted||",
		"47|Guaranteed||",
		"48|Held||",
		"49|Booked||",
		"50|Open||",
		"51|Segment confirmed after schedule change||",
		"52|Segment was confirmed||",
		"53|Overbook||",
		"54|Waitlisted Priority A||",
		"55|Waitlisted Priority B||",
		"56|Waitlisted Priority C||",
		"57|Waitlisted Priority D||",
		"58|Jumpseat||",
		"59|Need group space||",
		"60|Priority group||",
		"61|Unable, group waitlisted||",
		"62|Waitlist group||",
		"63|Segment was waitlist||",
		"64|Segment was reconfirmed||",
		"65|Segment was pending||",
		"66|Booked outside system||",
		"67|Pending cancel||",
		"68|Mandatory need||",
		"69|Need protection||",
		"70|Pending request||",
		"71|Pending cancellation||",
		"72|Staff listing (in SSR)||",
		"73|Link sold (history only)||",
		"74|Codeshare sold (history only)||",
		"75|Super-guaranteed sell (history only)||",
		"76|Confirming/new schedule||",
		"77|On request/new schedule||",
		"78|Placing on waitlist/new schedule||",
		"79|Standby boarded||",
		"80|Flown||",
		"81|No record available (NOREC)||",
		"82|Go-show||",
		"83|No-show||",
		"84|Offload firm booked||",
		"85|Offload of go-show||",
		"86|Downgrade to economy class||",
		"87|Down/upgrade to business class||",
		"88|Upgrade to first class||",
		"89|Offload of no record passenger||",
		"90|Holds confirmed, conditional EMD required||",
		"91|Holds confirmed, EMD required||",
		"92|Holds confirmed, EMD issued||",
		"93|Holds confirmed/new schedule, EMD already issued||",
		"151|Pending proposal|New RFP|Waiting for Proposal",
		"152|Proposal ready|Proposal Sent|Proposal Ready/ Waiting Selection",
		"153|Proposal accepted|Ready to Book|Waiting for Booking",
		"154|Proposal rejected||",
		"155|Proposal accepted, Over Credit Authorization required|Credit Being Reviewed|Credit Being Reviewed",
		"156|Proposal accepted, Travel Request required|Ready to Book|Require TR/PO",
		"157|Holds confirmed, Over Credit Authorization required|Credit Being Reviewed|Credit Being Reviewed",
		"158|Holds confirmed, Travel Request required|Ready to Issue|Require TR/PO",
		"159|Holds confirmed, Issuance Approval required|Waiting for Issuance Approval|Require Issuance Approval",
		"160|Confirmed, Over Credit Authorization required|Credit Being Reviewed|Credit Being Reviewed",
		"161|Confirmed, Travel Request required|Require TR/PO|Require TR/PO",
		"162|Confirmed, Payment required|Waiting For Payment|Waiting For Payment",
		"163|Holds cancelled|Cancelled by time limit|Cancelled by time limit",
		"164|Proposal holds confirmed|Ready to Issue|Issue In-Progress",
		"181|Credit OK|Credit OK|Credit OK",
		"182|Credit Not OK|Credit Not OK|Credit Not OK",
		"183|Credit Being Reviewed, Over Credit Authorization required|Credit Being Reviewed|Credit Being Reviewed",
		"184|Credit Approved, Over Credit Authorization approved|Credit Approved|Credit Approved",
		"185|Credit Declined, Over Credit Authorization declined|Credit Declined|Credit Declined",
		"191|Required, Travel Request required|Travel Request required|Travel Request required",
		"192|Not Required, Travel Request not required|Travel Request not required|Travel Request not required",
		"193|Submitted, Travel Request submitted|Travel Request submitted|Travel Request submitted",
		"194|Approved, Travel Request verified and approved|Travel Request verified|Travel Request verified",
		"195|Declined, Travel Request declined|Travel Request declined|Travel Request declined",
		"201|Message accepted||",
		"202|Message being delivered||",
		"203|Message delivered||",
		"204|Message failed||",
		"205|Message opened||",
		"206|Message clicked||",
		"207|Message unsubscribed||",
		"208|Message complained||",
		"221|Payment created||",
		"222|Payment being paid||",
		"223|Payment authorized||",
		"224|Payment captured||",
		"225|Payment completed||",
		"226|Payment denied||",
		"227|Payment expired||",
		"228|Payment voided||",
		"229|Payment refunded||",
		"230|Payment failed||",
		"251|Reward created||",
		"252|Reward deferred||",
		"253|Reward issued||",
		"254|Reward redeemed||",
		"255|Reward voided||",
		"256|Reward refunded||",
		"257|Reward expired||",
		"258|Reward failed||",
		"271|Invoice draft||",
		"272|Invoice open||",
		"273|Invoice paid||",
		"274|Invoice uncollectible||",
		"275|Invoice void||",
		"276|Invoice deleted||",
		"291|Survey started||",
		"292|Survey completed||",
		"311|Voucher created||",
		"312|Voucher deferred||",
		"313|Voucher issued||",
		"314|Voucher authorized||",
		"315|Voucher denied||",
		"316|Voucher redeemed||",
		"317|Voucher voided||",
		"318|Voucher refunded||",
		"319|Voucher expired||",
		"320|Voucher failed||",
		"331|Service started||",
		"332|Service restarted||",
		"333|Service stopped||",
		"334|Service succeed||",
		"335|Service failed||",
		"351|Ticket open|Ticket open, unused|",
		"352|Ticket used|Ticket used, lifted/ boarded|",
		"353|Ticket void|Ticket void, transaction voided|",
		"354|Ticket printed|Ticket printed, flight coupons printed|",
		"355|Ticket exchanged|Ticket exchanged, exchanged/ reissued|",
		"356|Ticket refunded|Ticket refunded|",
		"357|Ticket checked-in|Ticket checked-in|",
		"358|Ticket under airport control|Ticket under airport control|",
		"359|Ticket suspended|Ticket suspended by carrier|",
		"360|Ticket OK|Ticket OK, okay for travel|",
		"361|Ticket reactivated|Ticket reactivated, open, when void an exchange|",
		"371|Modification requested|New Request to Modify|",
		"372|Pending modification|Modification In-Progress|",
		"373|Modification ready|Modification Ready/ Waiting Acceptance|",
		"374|Modification accepted|Modification Confirmed|",
		"375|Modification rejected||",
		"391|Cancellation requested|New Request to Cancel|",
		"392|Pending cancellation|Cancellation In-Progress|",
		"393|Cancellation ready|Cancellation Ready/ Waiting Acceptance|",
		"394|Cancellation accepted|Cancellation Confirmed|",
		"395|Cancellation rejected||",
		"411|Booking acknowledgment requested|Request to Acknowledge Booking|",
		"412|Booking acknowledged||",
		"413|Booking rejected||",
		"431|Payload received|Payload Received|",
		"432|Payload accepted|Payload Valid|",
		"433|Payload approved|Payload Approved|",
		"434|Payload rejected|Payload Rejected|",
		"435|Payload failed|Payload Invalid|",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code, _ := strconv.Atoi(c[0])
		name := c[1]
		itnl := ""
		cust := ""
		if len(c) > 2 {
			itnl = c[1]
		}
		if len(c) > 3 {
			cust = c[2]
		}

		data = append(data, Status{
			StatusAPI: StatusAPI{
				StatusCode:         &code,
				StatusName:         &name,
				InternalStatusName: &itnl,
				ConsumerStatusName: &cust,
			},
		})
	}

	return &data
}

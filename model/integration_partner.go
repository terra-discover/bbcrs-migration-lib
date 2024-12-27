package model

import (
	"strconv"
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IntegrationPartner Integration Partner
type IntegrationPartner struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerAPI
	SupplierType           *SupplierType           `json:"supplier_type,omitempty" gorm:"foreignKey:SupplierTypeID;references:ID"`                      // Supplier Type
	BusinessEntity         *BusinessEntity         `json:"business_entity,omitempty" gorm:"foreignKey:BusinessEntityID;references:ID"`                  // Business Entity
	IntegrationPartnerType *IntegrationPartnerType `json:"integration_partner_type,omitempty" gorm:"foreignKey:IntegrationPartnerTypeID;references:ID"` // Integration Partner Type
}

// IntegrationPartnerAPI Integration Partner API
type IntegrationPartnerAPI struct {
	IntegrationPartnerCode                      *int       `json:"integration_partner_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null;"`    // Integration Partner Code
	IntegrationPartnerName                      *string    `json:"integration_partner_name,omitempty" gorm:"type:varchar(64);not null;index:,unique,where:deleted_at is null;"` // Integration Partner Name
	IntegrationPartnerTypeID                    *uuid.UUID `json:"integration_partner_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Integration Partner Type ID
	SupplierTypeID                              *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // Supplier Type ID
	BusinessEntityID                            *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                     // Business Entity ID
	PartnerURL                                  *string    `json:"partner_url,omitempty" gorm:"type:varchar(256)"`                                                              // The URL to access the partner's web service.
	PartnerUsername                             *string    `json:"partner_username,omitempty" gorm:"type:varchar(256)"`                                                         // The username element to authenticate to access the partner's web service.
	PartnerPassword                             *string    `json:"partner_password,omitempty" gorm:"type:varchar(256)"`                                                         // The password element (encrypted) associated with the username to authenticate to access the partner's web service.
	CrsUsername                                 *string    `json:"crs_username,omitempty" gorm:"type:varchar(256)"`                                                             // The username element to authenticate to access the CRS' web service.
	CrsPassword                                 *string    `json:"crs_password,omitempty" gorm:"type:varchar(256)"`                                                             // The password element (encrypted) associated with the username to authenticate to access the CRS' web service.
	TransactionPerformanceEnabled               *bool      `json:"transaction_performance_enabled,omitempty"`                                                                   // Transaction Performance Enabled
	AgentPerformanceEnabled                     *bool      `json:"agent_performance_enabled,omitempty"`                                                                         // Agent Performance Enabled
	CorporateInformationEnabled                 *bool      `json:"corporate_information_enabled,omitempty"`                                                                     // Corporate Information Enabled
	CorporatePerformanceEnabled                 *bool      `json:"corporate_performance_enabled,omitempty"`                                                                     // Corporate Performance Enabled
	CreditLimitEnabled                          *bool      `json:"credit_limit_enabled,omitempty"`                                                                              // Credit Limit Enabled
	EmployeePerformanceEnabled                  *bool      `json:"employee_performance_enabled,omitempty"`                                                                      // Employee Performance Enabled
	TravelerPerformanceEnabled                  *bool      `json:"traveler_performance_enabled,omitempty"`                                                                      // Traveler Performance Enabled
	InvoiceEnabled                              *bool      `json:"invoice_enabled,omitempty"`                                                                                   // Invoice Enabled
	InvoluntaryChangeBookingNotificationEnabled *bool      `json:"involuntary_change_booking_notification_enabled,omitempty"`                                                   // Involuntary Change Booking Notification Enabled
	Description                                 *string    `json:"description,omitempty" gorm:"type:text"`                                                                      // Description
}

// Seed data
func (s IntegrationPartner) Seed() *[]IntegrationPartner {
	data := []IntegrationPartner{}
	items := []string{
		"1|Sabre",
		"2|iTank",
		"3|Via.com",
		"4|TMS",
		"5|Open Weather Map",
		"6|Google Map",
		"7|SQ NDC",
		"8|Sabre NDC",
		"9|Garuda",
		"10|Expedia",
		"11|Midtrans",
		"12|Google Directions",
		"13|Google Distance Matrix",
		"14|Google Geocode",
		"15|Google Places",
		"16|Opsigo",
		"17|Foreign Exchange Rate",
		"18|Timatic",
		"19|Triptease",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		code, _ := strconv.Atoi(contents[0])
		businessEntityID, _ := uuid.Parse("4db13874-2519-49a2-b93d-646eaeb86077")

		var name string = contents[1]
		data = append(data, IntegrationPartner{
			IntegrationPartnerAPI: IntegrationPartnerAPI{
				BusinessEntityID:       &businessEntityID,
				IntegrationPartnerCode: &code,
				IntegrationPartnerName: &name,
			},
		})
	}

	return &data
}

// AfterCreate IntegrationPartner
func (s *IntegrationPartner) AfterCreate(tx *gorm.DB) error {
	// Create agent integration partner
	agent := Agent{}
	err := tx.First(&agent).Error
	if err != nil {
		return err
	}

	agentID := *agent.ID
	var agentIntegrationPartner AgentIntegrationPartner
	result := tx.Model(&agentIntegrationPartner).
		Where(tx.Where(AgentIntegrationPartner{
			AgentIntegrationPartnerAPI: AgentIntegrationPartnerAPI{
				AgentID: &agentID,
			},
		})).
		First(&agentIntegrationPartner)
	if result.RowsAffected < 1 {
		agentIntegrationPartner := AgentIntegrationPartner{}
		agentIntegrationPartner.AgentID = &agentID
		agentIntegrationPartner.IntegrationPartnerID = s.ID
		tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&agentIntegrationPartner)
	}

	// Create integration partner setting
	var integrationPartnerSetting IntegrationPartnerSetting
	result = tx.Model(&integrationPartnerSetting).
		Where(tx.Where(IntegrationPartnerSetting{
			IntegrationPartnerSettingAPI: IntegrationPartnerSettingAPI{
				IntegrationPartnerID: s.ID,
			},
		})).
		First(&integrationPartnerSetting)
	if result.RowsAffected < 1 {
		integrationPartnerSetting := IntegrationPartnerSetting{}
		integrationPartnerSetting.IntegrationPartnerID = s.ID
		integrationPartnerSetting.PartnerEventEnabled = lib.Boolptr(true)
		tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&integrationPartnerSetting)
	}

	return nil
}

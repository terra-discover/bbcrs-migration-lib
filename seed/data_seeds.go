package migration

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/model"
)

type DataSeedsRequest struct {
	AgentID        uuid.UUID
	UserID         uuid.UUID
	SmtpEmail      string
	SmtpSenderName string
	Salt           string
	Aes            string
}

var (
	menusSeeder                model.Menus
	modulePackages             model.ModulePackage
	modules                    model.Module
	genders                    model.Gender
	agentCorporate             model.AgentCorporate
	industriesSeeder           model.Industries
	documentTypeSeeder         model.DocumentType
	accreditationTypes         model.AccreditationTypes
	urlTypes                   model.URLTypes
	currencies                 model.Currency
	addressTypes               model.AddressTypes
	phoneLocationTypes         model.PhoneLocationType
	contactTypes               model.ContactType
	formTypes                  model.FormTypes
	messageCategories          model.MessageCategories
	messageTypes               model.MessageTypes
	operatingSystems           model.OperatingSystems
	menuLinkTypes              model.MenuLinkTypes
	eventTypes                 model.EventType
	eventLevels                model.EventLevel
	markets                    model.Market
	statusCategories           model.StatusCategory
	statuses                   model.Status
	securityFeatures           model.SecurityFeature
	classifications            model.RoomClassification
	segmentCategories          model.SegmentCategory
	uniqueIDTypes              model.UniqueIDType
	guestRoomInfoType          model.GuestRoomInfoType
	fareTypes                  model.FareType
	chargeTypes                model.ChargeType
	taskTypes                  model.TaskType
	queueTypes                 model.QueueType
	distributionTypes          model.DistributionType
	transportationModes        model.TransportationMode
	builtInTypes               model.BuiltInType
	recipientTypes             model.RecipientType
	proximities                model.Proximity
	metrics                    model.Metrics
	rateTimeUnits              model.RateTimeUnit
	rpointCategories           model.ReferencePointCategory
	recurrenceTypes            model.RecurrenceType
	basisTypes                 model.BasisType
	messageVariables           model.MessageVariable
	offsetTimeUnits            model.OffsetTimeUnit
	offsetDropTimes            model.OffsetDropTime
	capabilities               model.Capability
	messagePriorities          model.MessagePriorities
	rewardTypes                model.RewardType
	agents                     model.Agent
	languages                  model.Language
	corp                       model.Corporate
	corpRatingType             model.CorporateRatingType
	corpRatingTypeLevel        model.CorporateRatingTypeLevel
	productType                model.ProductType
	event                      model.Event
	phoneTechType              model.PhoneTechnologyType
	emailAddressType           model.EmailAddressType
	roomAmenityCat             model.RoomAmenityCategory
	passengerType              model.PassengerType
	feeTaxType                 model.FeeTaxType
	integrationPartnerStatuses model.IntegrationPartnerStatuses
	integrationPartner         model.IntegrationPartner
	agentIdentityRule          model.AgentIdentityRule
	serviceLevel               model.ServiceLevel
	additionalDetailType       model.AdditionalDetailType
	informationType            model.InformationType
	airlineCategory            model.AirlineCategory
	emailCategory              model.EmailCategory
	agentEmailCategory         model.AgentEmailCategory
	emailTemplate              model.EmailTemplate
	agentEmailTemplate         model.AgentEmailTemplate
	emailTemplateRecipientType model.EmailTemplateRecipientType
	reservation                model.Reservation
	proposal                   model.Proposal
	agentSetting               model.AgentSetting
	profileType                model.ProfileType
	sites                      model.Site
	siteType                   model.SiteType
	userTypes                  model.UserType
	profileStatus              model.ProfileStatus
	corporateType              model.CorporateType
	agentCorporateType         model.AgentCorporateType
	unitOfMeasure              model.UnitOfMeasure
	destinationGroup           model.DestinationGroup
	emailTemplateTranslations  model.EmailTemplateTranslation
)

// DataSeeds data to seeds
func DataSeeds(req DataSeedsRequest) []interface{} {
	return []interface{}{
		menusSeeder.Seed(),
		modulePackages.Seed(),
		modules.Seed(),
		industriesSeeder.Seed(),
		documentTypeSeeder.Seed(),
		accreditationTypes.Seed(),
		roomAmenityCat.Seed(),
		urlTypes.Seed(),
		currencies.Seed(),
		addressTypes.Seed(),
		phoneLocationTypes.Seed(),
		contactTypes.Seed(),
		formTypes.Seed(),
		messageCategories.Seed(),
		messageTypes.Seed(),
		operatingSystems.Seed(),
		menuLinkTypes.Seed(),
		eventTypes.Seed(),
		eventLevels.Seed(),
		markets.Seed(),
		statusCategories.Seed(),
		statuses.Seed(),
		securityFeatures.Seed(),
		classifications.Seed(),
		segmentCategories.Seed(),
		uniqueIDTypes.Seed(),
		guestRoomInfoType.Seed(),
		fareTypes.Seed(),
		chargeTypes.Seed(),
		taskTypes.Seed(),
		queueTypes.Seed(),
		distributionTypes.Seed(),
		transportationModes.Seed(),
		builtInTypes.Seed(),
		recipientTypes.Seed(),
		proximities.Seed(),
		metrics.Seed(),
		rateTimeUnits.Seed(),
		rpointCategories.Seed(),
		recurrenceTypes.Seed(),
		basisTypes.Seed(),
		messageVariables.Seed(),
		offsetTimeUnits.Seed(),
		offsetDropTimes.Seed(),
		capabilities.Seed(),
		messagePriorities.Seed(),
		rewardTypes.Seed(),
		agents.Seed(req.AgentID),
		languages.Seed(),
		corpRatingType.Seed(),
		corpRatingTypeLevel.Seed(),
		corp.Seed(req.AgentID),
		agentCorporate.Seed(req.AgentID),
		productType.Seed(),
		event.Seed(),
		phoneTechType.Seed(),
		emailAddressType.Seed(),
		passengerType.Seed(),
		feeTaxType.Seed(),
		integrationPartnerStatuses.Seed(),
		integrationPartner.Seed(),
		agentIdentityRule.Seed(req.AgentID),
		serviceLevel.Seed(),
		additionalDetailType.Seed(),
		informationType.Seed(),
		airlineCategory.Seed(),
		genders.Seed(),
		emailCategory.Seed(),
		agentEmailCategory.Seed(),
		emailTemplate.Seed(),
		agentEmailTemplate.Seed(),
		emailTemplateRecipientType.Seed(),
		agentSetting.Seed(req.AgentID, req.SmtpEmail, req.SmtpSenderName),
		profileType.Seed(),
		sites.Seed(),
		siteType.Seed(),
		userTypes.Seed(),
		profileStatus.Seed(),
		corporateType.Seed(),
		agentCorporateType.Seed(req.AgentID),
		unitOfMeasure.Seed(),
		destinationGroup.Seed(),
		emailTemplateTranslations.Seed(),
	}
}

package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// SecurityFeature Security Feature
type SecurityFeature struct {
	basic.Base
	basic.DataOwner
	SecurityFeatureAPI
	SecurityFeatureTranslation *SecurityFeatureTranslation `json:"security_feature_translation,omitempty"` // Security Feature Translation
}

// SecurityFeatureAPI Security Feature API
type SecurityFeatureAPI struct {
	SecurityFeatureCode *int    `json:"security_feature_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`                           // Security Feature Code
	SecurityFeatureName *string `json:"security_feature_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"2nd lock on guest doors"` // Security Feature Name
}

// Seed data
func (s SecurityFeature) Seed() *[]SecurityFeature {
	data := []SecurityFeature{}
	items := []string{
		"2nd lock on guest doors",
		"Address of nearest police station",
		"Alarms continuously monitored",
		"Audible alarm smoke detectors in guest rooms",
		"Auto link to fire station",
		"Automatic fire doors",
		"Building meets all current local, state and country building codes",
		"Complies with Hotel/Motel Safety Act",
		"Complies with Local/State/Federal fire laws",
		"Dead bolts on guest room doors",
		"Distance to nearest police station",
		"Doctor on call",
		"Electronic room key",
		"Elevator auto recall",
		"Emergency back-up generators",
		"Emergency evacuation plan",
		"Emergency exit maps",
		"Emergency info in room",
		"Emergency lighting",
		"Fire detectors in guest rooms",
		"Fire detectors in hallways",
		"Fire detectors in public areas",
		"Fire extinguishers",
		"Fire resistant guest room doors",
		"Fire resistant hall room doors",
		"Frequency of evacuation drills",
		"Guest room doors self-closing",
		"Hard wired smoke detectors",
		"If no 24 hour security, what are the hours?",
		"Lighted parking area",
		"Lighted walkways",
		"Multiple exits on each floor",
		"Parking area attendants",
		"Patrolled parking area",
		"Phone number of nearest police station",
		"Property has elevators",
		"Public address system",
		"Response time (minutes) from fire/police department",
		"Restricted public access",
		"Room access through exterior corridor",
		"Room access through interior corridor",
		"Room accessible through balcony sliding glass doors",
		"Room windows open",
		"Safety chains on guest doors",
		"Secondary locks on room windows",
		"Secondary locks on sliding glass doors",
		"Security 24 hours/day",
		"Smoke detector in guest rooms",
		"Smoke detector in hallways",
		"Smoke detector in public areas",
		"Some guest rooms have a balcony",
		"Sprinklers in guest rooms",
		"Sprinklers in hallways",
		"Sprinklers in public areas",
		"Staff trained in CPR",
		"Staff trained in duplicate key issue",
		"Staff trained in first aid",
		"Uniformed security",
		"Ventilated stair wells",
		"Video cameras at entrance",
		"Video cameras in hallways",
		"Video cameras in public areas",
		"Viewports in guest room doors",
		"Visual alarm smoke detectors in guest rooms",
		"Well lighted exit signs",
		"Which floors have exterior entrances",
		"Which floors have interior entrances",
		"Balcony accessibly by adjoining rooms",
		"Double glazed windows",
		"Fire extinguishers in hallways",
		"Security",
		"Audible alarms in hallways",
		"Audible alarms in public areas",
		"Automated External Defibrillator (AED) on-site",
		"Basic medical equipment on-site",
		"Camera monitoring parking area 24 hrs",
		"Camera recording parking area 24 hrs",
		"Controlled access to parking",
		"Exterior doors (except lobby entrance) require key access at night",
		"Staff Red Cross certified in CPR",
		"Staff trained in Automated External Defibrillator (AED) usage",
		"Video surveillance of parking",
		"Visual alarms in hallways",
		"Visual alarms in public areas",
		"Emergency exits on each floor",
		"Video surveillance recorded 24 hrs a day",
		"Video surveillance monitored 24 hrs a day",
		"Carbon monoxide detector",
		"Fire extinguishers in public areas",
		"First aid available",
		"Private security available",
		"Secured floors",
		"U.S. Fire Safety compliant",
		"Evacuation drills",
		"Hotel has fire safety measures in place but does not meet a national fire safety standard",
		"FEMA approved",
		"Hours security",
		"Emergency svc response time (minutes)",
		"Security escorts available on request",
		"Swingbolt lock",
		"AAA security standards",
		"Connecting doors have deadbolts",
		"Emergency call button on phone",
		"Audible alarms",
		"A floor only accessible via a guest room key",
		"Health club facilities (pool/gym) require key access for entrance",
		"VIP security",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, SecurityFeature{
			SecurityFeatureAPI: SecurityFeatureAPI{
				SecurityFeatureCode: &code,
				SecurityFeatureName: &name,
			},
		})
	}

	return &data
}

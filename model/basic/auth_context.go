package basic

import (
	"github.com/google/uuid"
)

// AuthContext will hold current request context user-specific ids
type AuthContext struct {
	UserID      *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	AgentID     *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	CorporateID *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
	LanguageID  *uuid.UUID `json:"-" swaggertype:"string" format:"uuid"`
}

// Dummy will generate a random dummy auth context for unit testing purpose
func (a *AuthContext) Dummy() *AuthContext {
	corporateID, _ := uuid.NewRandom()
	agentID, _ := uuid.NewRandom()
	userID, _ := uuid.NewRandom()
	languageID, _ := uuid.NewRandom()

	a.CorporateID = &corporateID
	a.AgentID = &agentID
	a.UserID = &userID
	a.LanguageID = &languageID

	return a
}

package model

import (
	"log"
	"time"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FlightCachingRequestStatus model will store background process flight search status
type FlightCachingRequestStatus struct {
	basic.Base
	basic.DataOwner
	SessionID        *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);not null"`
	Status           *int       `json:"status,omitempty" gorm:"type:integer;not null'"`      // status (0=pending, 1=done, 2=error, 3=expired)
	Description      *string    `json:"description,omitempty" gorm:"type:varchar(255);null"` // additional description when there is error(2) status
	MeasureTime      *int64     `json:"measure_time,omitempty" gorm:"type:integer"`          // measure time request milliseconds unit
	IsDomestic       *bool      `json:"is_domestic,omitempty"`                               // flag this flight is domestic or international
	IsCachedResponse *bool      `json:"is_cached_response,omitempty"`
	ChunkSize        *int       `json:"chunk_size"`     // total of flight search chunk
	FinishedChunk    *int       `json:"finished_chunk"` // total of flight search chunk resolved. if same as ChunkSize, status=DONE, else=PARTIAL
}

var FLIGHT_SEARCH_BACKGROUND_STATUS_PENDING = 0
var FLIGHT_SEARCH_BACKGROUND_STATUS_DONE = 1
var FLIGHT_SEARCH_BACKGROUND_STATUS_ERROR = 2
var FLIGHT_SEARCH_BACKGROUND_STATUS_EXPIRED = 3
var FLIGHT_SEARCH_BACKGROUND_STATUS_INFO = 4
var FLIGHT_SEARCH_BACKGROUND_STATUS_WARNING = 5
var FLIGHT_SEARCH_BACKGROUND_STATUS_PARTIAL = 9

// FlightSearchBackgroundStatus
var FlightSearchBackgroundStatus = map[int]string{
	0: "PENDING",
	1: "DONE",
	2: "ERROR",
	3: "EXPIRED",
	4: "INFO",
	5: "WARNING",
	9: "PARTIAL",
}

// IsDomesticFlight by query filter
func (data *FlightCachingRequestStatus) IsDomesticFlight(tx *gorm.DB, sessionID *uuid.UUID) bool {
	var total int64
	tx.Model(&FlightCachingRequestStatus{}).Where(tx.Where(FlightCachingRequestStatus{
		IsDomestic: lib.Boolptr(true),
		SessionID:  sessionID,
	})).Count(&total)
	if total == 0 {
		return total != 0
	}
	return true
}

func (data *FlightCachingRequestStatus) SetTotalChunk(tx *gorm.DB, chunkSize int) {
	data.ChunkSize = &chunkSize
	tx.Save(data)
}

func (data *FlightCachingRequestStatus) SetChunkIsFinished(tx *gorm.DB, measureTime ...time.Time) {
	var totalChunkSize int
	var finishedChunk int

	if data.ChunkSize != nil {
		totalChunkSize = *data.ChunkSize
	}
	if data.FinishedChunk != nil {
		finishedChunk = *data.FinishedChunk
	}

	finishedChunk++
	data.FinishedChunk = &finishedChunk
	if finishedChunk >= totalChunkSize {
		data.Status = &FLIGHT_SEARCH_BACKGROUND_STATUS_DONE
	} else {
		data.Status = &FLIGHT_SEARCH_BACKGROUND_STATUS_PARTIAL
	}

	if len(measureTime) > 0 {
		data.MeasureTime = lib.Int64ptr(time.Since(measureTime[0]).Milliseconds())
	}

	log.Printf("SetChunkIsFinished %d/%d", finishedChunk, totalChunkSize)

	tx.Save(data)
}

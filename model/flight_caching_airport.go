package model

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/fs"

	static "github.com/terra-discover/bbcrs-migration-lib/data"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
)

// FlightCachingAirport model
type FlightCachingAirport struct {
	basic.Base
	basic.DataOwner
	IATA        *string  `json:"iata,omitempty" gorm:"type:varchar(3);index:idx_flight_caching_airport_iata,unique;not null" example:"CGK"` // International Air Transport Association airport code
	Timezone    *string  `json:"timezone,omitempty" example:"Asia/Jakarta"`                                                                 // Timezone of the airport, e.g., "Asia/Jakarta"
	GMTOffset   *string  `json:"gmt_offset,omitempty" example:"GMT +07:00"`                                                                 // GMT offset for the airport, e.g., "GMT +07:00"
	OffsetHours *int     `json:"offset_hours,omitempty" example:"+7"`                                                                       // Offset in hours from GMT, e.g., +7
	Name        *string  `json:"name,omitempty" example:"Soekarno-Hatta International Airport"`                                             // Name of the airport
	CityCode    *string  `json:"city_code,omitempty" example:"JKT"`                                                                         // City code where the airport is located, e.g., "JKT"
	CountryID   *string  `json:"country_id,omitempty" example:"ID"`                                                                         // Country code of the airport, e.g., "ID"
	Latitude    *float64 `json:"latitude,omitempty" example:"106.65658968644006"`                                                           // Latitude of the airport, e.g., 106.65658968644006
	Longitude   *float64 `json:"longitude,omitempty" example:"-6.12389155"`                                                                 // Longitude of the airport, e.g., -6.12389155
	Elevation   *int     `json:"elevation,omitempty" example:"29"`                                                                          // Elevation of the airport, e.g., 29
	ICAO        *string  `json:"icao,omitempty" example:"WIII"`                                                                             // ICAO code for the airport, e.g., "WIII"
}

// Seed data
func (p FlightCachingAirport) Seed() *[]FlightCachingAirport {
	var err error
	var fileContent []byte
	var file io.Reader
	var lines [][]string
	var data []FlightCachingAirport

	fileContent, err = fs.ReadFile(static.FS, "airports.csv")
	file = bytes.NewReader(fileContent)

	// file, _ = os.Open("data/airports.csv")
	// defer file.Close()

	if err == nil {
		reader := csv.NewReader(file)
		lines, err = reader.ReadAll()
	}

	if err == nil && len(lines) > 1 {
		for key, line := range lines {
			if key == 0 {
				continue
			}
			airport := FlightCachingAirport{
				IATA:        lib.Strptr(line[0]),
				Timezone:    lib.Strptr(line[1]),
				GMTOffset:   lib.Strptr(line[2]),
				OffsetHours: lib.Intptr(lib.StrToInt(line[3])),
				Name:        lib.Strptr(line[4]),
				CityCode:    lib.Strptr(line[5]),
				CountryID:   lib.Strptr(line[6]),
				Latitude:    lib.Float64ptr(lib.StrToFloat(line[7])),
				Longitude:   lib.Float64ptr(lib.StrToFloat(line[8])),
				Elevation:   lib.Intptr(lib.StrToInt(line[9])),
				ICAO:        lib.Strptr(line[10]),
			}
			data = append(data, airport)
		}
	}

	return &data
}

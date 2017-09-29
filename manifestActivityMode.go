package goBungieNet

import (
	"fmt"
	"errors"
	"encoding/json"
)

type DestinyActivityModeCategory int32

const (
	DestinyActivityModeCategoryNone DestinyActivityModeCategory = 0
	DestinyActivityModeCategoryPvE = 1
	DestinyActivityModeCategoryPvP = 2
)

func (atype *DestinyActivityModeCategory)String() string {
	switch (*atype) {
		case DestinyActivityModeCategoryNone: return "None"
		case DestinyActivityModeCategoryPvE: return "PvE"
		case DestinyActivityModeCategoryPvP: return "PvP"
	}

	return "Unknown"
}

type DestinyActivityModeDefinition struct {
	ActivityModeCategory DestinyActivityModeCategory
	ActivityModeMappings map[uint32]DestinyActivityMode
	Display bool
	DisplayProperties DestinyDisplayProperties
	FriendlyName string
	IsTeamBased bool
	IsAggregateMode bool
	Index int32
	ModeType DestinyActivityMode
	Order int32
	ParentHashes []uint32
	PGCRImage string
	Redacted bool
}

func manifestActivityModeDefinition(languageCode string, hash uint32) (*DestinyActivityModeDefinition, error) {
	dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
	db, err := manifestOpenData(dataFile)
	if err != nil { return nil, err }
  
	rows, err := db.Query("SELECT json FROM DestinyActivityModeDefinition WHERE id = $1", int32(hash))
	if err != nil { return nil, err }
	defer rows.Close()
	if !rows.Next() { return nil, errors.New("No Results!") }
	var jsonData string
	err = rows.Scan(&jsonData)
	if err != nil { return nil, err }
  
	result := DestinyActivityModeDefinition{}
	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil { return nil, err }
  
	return &result, nil
}

func (base *CharacterActivities)ActivityModeDefinition(languageCode string) (*DestinyActivityModeDefinition, error) {
	return manifestActivityModeDefinition(languageCode, base.CurrentActivityModeHash)
}
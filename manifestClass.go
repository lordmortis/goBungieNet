package goBungieNet

import (
	"fmt"
	"errors"
	"encoding/json"
)

type DestinyClassDefinition struct {
	DisplayProperties DestinyDisplayProperties
	Index int32
	GenderedClassNames map[string]string
	Redacted bool
}

func manifestClassDefinition(languageCode string, hash uint32) (*DestinyClassDefinition, error) {
	dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
	db, err := manifestOpenData(dataFile)
	if err != nil { return nil, err }
  
	rows, err := db.Query("SELECT json FROM DestinyClassDefinition WHERE id = $1", int32(hash))
	if err != nil { return nil, err }
	defer rows.Close()
	if !rows.Next() { return nil, errors.New("No Results!") }
	var jsonData string
	err = rows.Scan(&jsonData)
	if err != nil { return nil, err }
  
	result := DestinyClassDefinition{}
	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil { return nil, err }
  
	return &result, nil
  }

  func (base *Character)Class(languageCode string) (*DestinyClassDefinition, error) {
	  return manifestClassDefinition(languageCode, base.ClassHash)
  }
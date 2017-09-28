package goBungieNet

import (
  "fmt"
  "encoding/json"
)

type DestinyActivityModifierDefinition struct {
  DisplayProperties DestinyDisplayProperties
  Index int32
  Redacted bool
}

func (base *DestinyMilestoneActivity)Modifiers(languageCode string) (*[]DestinyActivityModifierDefinition, error) {
  dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
  db, err := manifestOpenData(dataFile)
  if err != nil { return nil, err }

  queryStr := "SELECT json FROM DestinyActivityModifierDefinition WHERE id IN ("
  params := make([]interface{}, len(base.ModifierHashes))
  for index, value := range(base.ModifierHashes) {
    params[index] = int32(value)
    if index + 1 < len(base.ModifierHashes) {
      queryStr += "?,"
    } else {
      queryStr += "?"
    }
  }
  queryStr += ")"

  rows, err := db.Query(queryStr, params...)
  if err != nil { return nil, err }
  defer rows.Close()

  var jsonData string
  modifiers := make([]DestinyActivityModifierDefinition, 0, 5)
  for rows.Next() {
    err = rows.Scan(&jsonData)
    if err != nil { return nil, err }

    modifier := DestinyActivityModifierDefinition{}
    err = json.Unmarshal([]byte(jsonData), &modifier)
    if err != nil { return nil, err }
    modifiers = append(modifiers, modifier)
  }

  return &modifiers, nil
}

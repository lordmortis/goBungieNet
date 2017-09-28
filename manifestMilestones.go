package goBungieNet

import (
  "fmt"
  "errors"
//  "time"
  "encoding/json"
)

type DestinyMilestoneType int

const (
  DestinyMilestoneTypeUnknown DestinyMilestoneType = 0
  DestinyMilestoneTypeTutorial = 1
  DestinyMilestoneTypeOneTime = 2
  DestinyMilestoneTypeWeekly = 3
  DestinyMilestoneTypeDaily = 4
  DestinyMilestoneTypeSpecial = 5
)

func (atype DestinyMilestoneType)String() string {
  switch(atype) {
    case DestinyMilestoneTypeUnknown: return "Unknown"
    case DestinyMilestoneTypeTutorial: return "Tutorial"
    case DestinyMilestoneTypeOneTime: return "One Time"
    case DestinyMilestoneTypeWeekly: return "Weekly"
    case DestinyMilestoneTypeDaily: return "Daily"
    case DestinyMilestoneTypeSpecial: return "Special"
  }

  return "Unknown"
}

type DestinyMilestoneQuestDefinition struct {

}

type DestinyMilestoneRewardCategory struct {

}

type DestinyMilestoneVendor struct {

}

type DestinyMilestoneValue struct {

}

type DestinyMilestoneDefinition struct {
  FriendlyName string
  HasPredictableDates bool
  Image string
  IsInGameMilestone bool
  MilestoneType DestinyMilestoneType
  Quests map[uint32]DestinyMilestoneQuestDefinition
  Recruitable bool
  Rewards map[uint32]DestinyMilestoneRewardCategory
  ShowInExplorer bool
  Values map[string]DestinyMilestoneValue
  Vendors []DestinyMilestoneVendor
  Redacted bool
}

func ManifestMilestoneDefinition(languageCode string, milestoneHash uint32) (*DestinyMilestoneDefinition, error) {
  dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
  db, err := manifestOpenData(dataFile)
  if err != nil { return nil, err }

  rows, err := db.Query("SELECT json FROM DestinyMilestoneDefinition WHERE id = $1", int32(milestoneHash))
  if err != nil { return nil, err }
  if !rows.Next() { return nil, errors.New("No Results!") }
  var jsonData string
  err = rows.Scan(&jsonData)
  if err != nil { return nil, err }

  result := DestinyMilestoneDefinition{}
  err = json.Unmarshal([]byte(jsonData), &result)
  if err != nil { return nil, err }

  return &result, nil
}

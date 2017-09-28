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

type DestinyMilestoneQuestRewardItem struct {
  ItemInstanceID int64
  VendorHash uint32
  VendorItemIndex int32
  Quantity int32
}

type DestinyMilestoneQuestRewards struct {
  Items []DestinyMilestoneQuestRewardItem
}

type DestinyMilestoneActivityDefinition struct {
  BaseActivityHash uint32 `json:"conceptualActivityHash"`
  Variants map[uint32]DestinyMilestoneActivityVariantDefinition
}

type DestinyMilestoneActivityVariantDefinition struct {
  ActivityHash uint32
  Order int32
}

type DestinyMilestoneQuestDefinition struct {
  Activities map[uint32]DestinyMilestoneActivityDefinition
  DisplayProperties DestinyDisplayProperties
  QuestItemHash uint32
  OverrideImage string
  QuestRewards DestinyMilestoneQuestRewards
}

type DestinyMilestoneRewardEntry struct {
  DisplayProperties DestinyDisplayProperties
  Identifier string `json:"rewardEntryIdentifier"`
  Items []DestinyItemQuantity
  VendorHash uint32
  Order int32
}

type DestinyMilestoneRewardCategory struct {
  CategoryIdentifier string
  DisplayProperties DestinyDisplayProperties
  RewardEntries map[uint32]DestinyMilestoneRewardEntry
  Order int32
}

type DestinyMilestoneVendor struct {
  VendorHash uint32
}

type DestinyMilestoneValue struct {
  DisplayProperties DestinyDisplayProperties
  Key string
}

type DestinyMilestoneDefinition struct {
  DisplayProperties DestinyDisplayProperties
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

func milestoneDefinition(languageCode string, milestoneHash uint32) (*DestinyMilestoneDefinition, error) {
  dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
  db, err := manifestOpenData(dataFile)
  if err != nil { return nil, err }

  rows, err := db.Query("SELECT json FROM DestinyMilestoneDefinition WHERE id = $1", int32(milestoneHash))
  if err != nil { return nil, err }
  defer rows.Close()
  if !rows.Next() { return nil, errors.New("No Results!") }
  var jsonData string
  err = rows.Scan(&jsonData)
  if err != nil { return nil, err }

  result := DestinyMilestoneDefinition{}
  err = json.Unmarshal([]byte(jsonData), &result)
  if err != nil { return nil, err }

  return &result, nil
}

func (base *DestinyMilestone)Definition(languageCode string) (*DestinyMilestoneDefinition, error) {
  return milestoneDefinition(languageCode, base.MilestoneHash)
}


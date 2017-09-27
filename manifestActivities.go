package goBungieNet

import (
  "fmt"
  "errors"
  "time"
  "encoding/json"
)

type DestinyActivityReward struct {
  Text string `json:rewardText`
  Items []DestinyItemQuantity `json:rewardItems`
}

type DestinyActivityModifier struct {
  ActivityModifierHash uint32
}

type DestinyActivityChallenges struct {
  ObjectiveHash uint32
}

type DestinyActivityUnlockString struct {
  DisplayString string
}

type DestinyActivityPlaylistItem struct {
  ActivityHash uint32
  DirectActivityModeHash uint32
  DirectActivityModeType DestinyActivityMode
  ActivityModeHashes []uint32
  ActivityModeTypes []DestinyActivityMode
}

type DestinyActivityMatchmaking struct {
  IsMatchmade bool
  MinParty int32
  MaxParty int32
  MaxPlayers int32
  RequiresGuardianOath bool
}

type DestinyActivityGuided struct {
  MaxLobbySize int32 `json:"guidedMaxLobbySize"`
  MinLobbySize int32 `json:"guidedMinLobbySize"`
  DisbandCount int32 `json:"guidedDisbandCount"`
}

type DestinyActivity struct {
  ActivityLevel int32
  ActivityLightLevel int32
  ActivityTypeHash uint32
  ActivityModeHashes []uint32
  ActivityModeTypes []DestinyActivityMode
  Challenges []DestinyActivityChallenges
  DisplayProperties DestinyDisplayProperties
  DestinationHash uint32
  DirectActivityModeHash uint32
  DirectActivityModeType DestinyActivityMode
  GuidedGame DestinyActivityGuided
  Matchmaking DestinyActivityMatchmaking
  Modifiers []DestinyActivityModifier
  OptionalUnlockStrings []DestinyActivityUnlockString
  PCGRImage string
  PlaceHash uint32
  IsPlaylist bool
  IsPVP bool
  PlaylistItems []DestinyActivityPlaylistItem
  ReleaseIcon string
  ReleaseTimeEpoch int32 `json:"releaseTime"`
  ReleaseTime time.Time `json:"-"`
  Rewards []DestinyActivityReward
  Redacted bool
  Tier int32
}

func (data *CharacterActivities) CurrentActivity(languageCode string) error {
  return nil
}

func ManifestActivityDefinition(languageCode string, activityHash uint32) (*DestinyActivity, error) {
  dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
  db, err := manifestOpenData(dataFile)
  if err != nil { return nil, err }

  rows, err := db.Query("SELECT json FROM DestinyActivityDefinition WHERE id = $1", int32(activityHash))
  if err != nil { return nil, err }
  if !rows.Next() { return nil, errors.New("No Results!") }
  var jsonData string
  err = rows.Scan(&jsonData)
  if err != nil { return nil, err }

  result := DestinyActivity{}
  result.ReleaseTimeEpoch = 10
  err = json.Unmarshal([]byte(jsonData), &result)
  if err != nil { return nil, err }
  result.ReleaseTime = time.Unix(int64(result.ReleaseTimeEpoch), 0)

  return &result, nil
}

package goBungieNet

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type DestinyActivityReward struct {
	Text  string                `json:rewardText`
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
	ActivityHash           uint32
	DirectActivityModeHash uint32
	DirectActivityModeType DestinyActivityMode
	ActivityModeHashes     []uint32
	ActivityModeTypes      []DestinyActivityMode
}

type DestinyActivityMatchmaking struct {
	IsMatchmade          bool
	MinParty             int32
	MaxParty             int32
	MaxPlayers           int32
	RequiresGuardianOath bool
}

type DestinyActivityGuided struct {
	MaxLobbySize int32 `json:"guidedMaxLobbySize"`
	MinLobbySize int32 `json:"guidedMinLobbySize"`
	DisbandCount int32 `json:"guidedDisbandCount"`
}

type DestinyActivity struct {
	ActivityLevel          int32
	ActivityLightLevel     int32
	ActivityTypeHash       uint32
	ActivityModeHashes     []uint32
	ActivityModeTypes      []DestinyActivityMode
	Challenges             []DestinyActivityChallenges
	DisplayProperties      DestinyDisplayProperties
	DestinationHash        uint32
	DirectActivityModeHash uint32
	DirectActivityModeType DestinyActivityMode
	GuidedGame             DestinyActivityGuided
	Matchmaking            DestinyActivityMatchmaking
	Modifiers              []DestinyActivityModifier
	OptionalUnlockStrings  []DestinyActivityUnlockString
	PCGRImage              string
	PlaceHash              uint32
	IsPlaylist             bool
	IsPVP                  bool
	PlaylistItems          []DestinyActivityPlaylistItem
	ReleaseIcon            string
	ReleaseTimeEpoch       int32     `json:"releaseTime"`
	ReleaseTime            time.Time `json:"-"`
	Rewards                []DestinyActivityReward
	Redacted               bool
	Tier                   int32
}

func manifestActivityDefinition(languageCode string, activityHash uint32) (*DestinyActivity, error) {
	dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
	db, err := manifestOpenData(dataFile)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT json FROM DestinyActivityDefinition WHERE id = $1", int32(activityHash))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.New("No Results!")
	}
	var jsonData string
	err = rows.Scan(&jsonData)
	if err != nil {
		return nil, err
	}

	result := DestinyActivity{}
	err = json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		return nil, err
	}
	result.ReleaseTime = time.Unix(int64(result.ReleaseTimeEpoch), 0)

	return &result, nil
}

func (base *DestinyMilestoneActivity) Definition(languageCode string) (*DestinyActivity, error) {
	return manifestActivityDefinition(languageCode, base.ActivityHash)
}

func (base *CharacterActivities) ActivityDefinition(languageCode string) (*DestinyActivity, error) {
	return manifestActivityDefinition(languageCode, base.CurrentActivityHash)
}

func ActivityDefinitions(languageCode string, hashes []uint32) (*[]DestinyActivity, error) {
	dataFile := fmt.Sprintf("%s-%s", manifestWorld, languageCode)
	db, err := manifestOpenData(dataFile)
	if err != nil {
		return nil, err
	}

	params := make([]interface{}, len(hashes))
	queryStr := "SELECT json FROM DestinyActivityDefinition WHERE id IN ("
	for index, hash := range hashes {
		params[index] = int32(hash)
		if index+1 < len(hashes) {
			queryStr += "?,"
		} else {
			queryStr += "?"
		}
	}
	queryStr += ")"

	rows, err := db.Query(queryStr, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var defns []DestinyActivity

	for rows.Next() {
		var jsonData string
		err = rows.Scan(&jsonData)
		if err != nil {
			return &defns, err
		}

		result := DestinyActivity{}
		err = json.Unmarshal([]byte(jsonData), &result)
		if err != nil {
			return &defns, err
		}
		result.ReleaseTime = time.Unix(int64(result.ReleaseTimeEpoch), 0)

		defns = append(defns, result)
	}

	return &defns, nil
}

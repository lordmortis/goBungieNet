package goBungieNet

import (
  "fmt"
  "errors"
//  "reflect"
  "time"
)

type GetProfileResponse struct {
  Characters interface{}
  CharacterActivities SingleActivities
  CharacterEquipment interface{}
  CharacterInventories interface{}
  CharacterKiosks interface{}
  CharacterProgressions interface{}
  CharacterRenderData interface{}
  ItemComponents interface{}
  ProfileInventory interface{}
  ProfileCurrencies interface{}
  Profile SingleProfile
  ProfileKiosks interface{}
  VendorReciepts interface{}
}

type SingleProfile struct {
  Data DestinyProfile
  Privacy PrivacySetting
}

type SingleActivities struct {
  Data map[int64]CharacterActivities
  Privacy PrivacySetting
}

type CharacterActivities struct {
  AvailableActivities []Activity
  DateActivityStarted time.Time
  CurrentActivityHash uint32
  CurrentActivityModeHash uint32
  CurrentActivityModeType ActivityMode
  CurrentActivityModeHashes []uint32
  CurrentActivityModeTypes []ActivityMode
  CurrentPlaylistActibityHash uint32
  LastCompletedStoryHash uint32
}

type Activity struct {

}

type DestinyProfile struct {
  CharacterIDs []int64
  UserInfo UserInfo
  DateLastPlayed time.Time
  VersionsOwned DestinyVersion
}

func GetProfile(membershipID int64, networkType BungieMembershipType, components []DestinyComponentType) (*GetProfileResponse, error) {
  componentString := ""

  for index, component := range(components) {
    if (index + 1 == len(components)) {
      componentString += fmt.Sprintf("%d", component)
    } else {
      componentString += fmt.Sprintf("%s,", component)
    }
  }

  url := fmt.Sprintf("/Destiny2/%d/Profile/%d/?components=%s", networkType, membershipID, componentString)
  response, err := get(url)
  if (err != nil) {
    return nil, errors.New("Get Error: " + err.Error())
  }

  if response.ErrorCode.isError() {
    return nil, errors.New("Response Error: " + string(response.ErrorCode))
  }

  var profileResponse GetProfileResponse

  err = decode(response.Response, &profileResponse)
  if err != nil {
    return nil, errors.New("Could not decode response: " + err.Error())
  }

  return &profileResponse, nil
}

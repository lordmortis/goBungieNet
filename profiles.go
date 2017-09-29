package goBungieNet

import (
  "fmt"
  "errors"
  "time"
)

type GetProfileResponse struct {
  Characters SingleCharacters
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

type SingleCharacters struct {
  Data map[int64]Character
  Privacy PrivacySetting
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
      componentString += fmt.Sprintf("%d,", component)
    }
  }

  url := fmt.Sprintf("/Destiny2/%d/Profile/%d/?components=%s", networkType, membershipID, componentString)
  response, err := get(url)
  if (err != nil) {
    return nil, errors.New("Get Error: " + err.Error())
  }

  if response.ErrorCode.isError() {
    errorString := fmt.Sprintf("Response Error: %d", response.ErrorCode)
    return nil, errors.New(errorString)
  }

  var profileResponse GetProfileResponse

  err = decode(response.Response, &profileResponse)
  if err != nil {
    return nil, errors.New("Could not decode response: " + err.Error())
  }

  return &profileResponse, nil
}

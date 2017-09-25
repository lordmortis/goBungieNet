package goBungieNet

import (
  "fmt"
  "errors"
//  "reflect"
  "time"
)

type GetProfileResponse struct {
  Characters interface{}
  CharacterActivities interface{}
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

type DestinyProfile struct {
  CharacterIds []int64
  UserInfo UserInfo
  DateLastPlayed time.Time
  VersionsOwned DestinyVersion
}

func GetProfile(membershipID string, networkType BungieMembershipType, components []DestinyComponentType) (*GetProfileResponse, error) {
  componentString := ""

  for index, component := range(components) {
    if (index + 1 == len(components)) {
      componentString += fmt.Sprintf("%d", component)
    } else {
      componentString += fmt.Sprintf("%s,", component)
    }
  }

  url := fmt.Sprintf("/Destiny2/%d/Profile/%s/?components=%s", networkType, membershipID, componentString)
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

  return nil, nil
}

package goBungieNet

import (
  "fmt"
  "errors"
  "strings"
)

type BungieMembershipType int

const (
  NetworkNone BungieMembershipType = 0
  NetworkXbox = 1
  NetworkPsn = 2
  NetworkBlizzard = 4
  NetworkDemon = 10
  NetworkBungieNext = 254
  NetworkAll = -1
)

func (atype BungieMembershipType)String() string {
  switch(atype) {
    case NetworkNone: return "None"
    case NetworkXbox: return "Xbox Live"
    case NetworkPsn: return "Playstation Network"
    case NetworkBlizzard: return "Battle.Net"
    case NetworkAll: return "All Networks"
  }

  return "Unknown"
}

func StringToBungieMembershipType(text string) BungieMembershipType {
  accountTypeString := strings.ToLower(text)

  if strings.HasPrefix(accountTypeString, "ps") {
    return NetworkPsn
  } else if strings.HasPrefix(accountTypeString, "xbox") {
    return NetworkXbox
  } else if strings.HasPrefix(accountTypeString, "battle") {
    return NetworkBlizzard
  } else if strings.HasPrefix(accountTypeString, "all") {
    return NetworkAll
  }

  return NetworkNone
}

type UserInfo struct {
  IconPath string
  SupplementalDisplayName string
  MembershipType BungieMembershipType
  MembershipID int64
  DisplayName string
}

func FindDestinyUser(displayName string, networkType BungieMembershipType) (*[]UserInfo, error) {
  url := fmt.Sprintf("/Destiny2/SearchDestinyPlayer/%d/%s/", networkType, displayName)
  response, err := get(url)
  if (err != nil) {
    return nil, errors.New("Get Error: " + err.Error())
  }

  if response.ErrorCode.isError() {
    return nil, errors.New("Response Error: " + string(response.ErrorCode))
  }

  var info []UserInfo

  err = decode(response.Response, &info)
  if err != nil {
    return nil, errors.New("Could not decode response: " + err.Error())
  }
  return &info, err
}


package goBungieNet

import (
  "fmt"
  "errors"
//  "reflect"
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

type SingleActivities struct {
  Data map[int64]CharacterActivities
  Privacy PrivacySetting
}

type SingleCharacters struct {
  Data map[int64]Character
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

type Character struct {
  CharacterID int64
  MembershipID int64
  MembershipType BungieMembershipType
  DateLastPlayed time.Time
  MinutesPlayedThisSession int64
  MinutesPlayedTotal int64
  Light int32
  Stats map[string]int32
  RaceHash uint32
  GenderHash uint32
  ClassHash uint32
  EmblemPath string
  EmblemBakgroundPath string
  EmblemHash uint32
  LevelProgression DestinyProgression
  BaseCharacterLevel int32
  PercentToNextLevel float32
}

type Activity struct {

}

type DestinyProgression struct {
  /** The hash identifier of the Progression in question. Use it to
   * look up the DestinyProgressionDefinition in static data. */
  ProgressionHash uint32
  /** The amount of progress earned today for this progression. */
  DailyProgress int32
  /** If this progression has a daily limit, this is that limit. */
  DailyLimit int32
  /** The amount of progress earned toward this progression in the
   * current week. */
  WeeklyProgress int32
  /** If this progression has a weekly limit, this is that limit. */
  WeeklyLimit int32
  /** This is the total amount of progress obtained overall for this
   * progression (for instance, the total amount of Character Level
   * experience earned) */
  CurrentProgress int32
  /** This is the level of the progression (for instance, the Character Level). */
  Level int32
  /** This is the maximum possible level you can achieve for this
   * progression (for example, the maximum character level obtainable) */
  LevelCap int32
  /** Progressions define their levels in "steps". Since the last step
   * may be repeatable, the user may be at a higher level than the
   * actual Step achieved in the progression. Not necessarily useful,
   * but potentially interesting for those cruising the API. Relate
   * this to the "steps" property of the DestinyProgression to see
   * which step the user is on, if you care about that. (Note that
    * this is Content Version dependent since it refers to indexes.) */
  StepIndex int32
  /** The amount of progression (i.e. "Experience") needed to reach
   * the next level of this Progression. Jeez, progression is such
   * an overloaded word. */
  ProgressToNextLevel int32
  /** The total amount of progression (i.e. "Experience") needed in
   * order to reach the next level. */
  NextLevelAt int32
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

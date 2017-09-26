package goBungieNet

import (
  "time"
)

type SingleActivities struct {
  Data map[int64]CharacterActivities
  Privacy PrivacySetting
}

type CharacterActivities struct {
  AvailableActivities []Activity
  DateActivityStarted time.Time
  CurrentActivityHash uint32
  CurrentActivityModeHash uint32
  CurrentActivityModeType DestinyActivityMode
  CurrentActivityModeHashes []uint32
  CurrentActivityModeTypes []DestinyActivityMode
  CurrentPlaylistActibityHash uint32
  LastCompletedStoryHash uint32
}

func (data *SingleActivities)MostRecentCharacter() int64 {
  activityTime := time.Unix(0,0)
  var charID int64 = -1

  for id, activities := range(data.Data) {
    if (activityTime.Before(activities.DateActivityStarted)) {
      activityTime = activities.DateActivityStarted
      charID = id
    }
  }

  return charID
}

package goBungieNet

import (
  "time"
  "errors"
//  "fmt"
//  "errors"
//    "strings"

//  "github.com/davecgh/go-spew/spew"
)

type DestinyMilestoneActivityVariant struct {
  ActivityHash uint32
}

type DestinyMilestoneActivity struct {
  ActivityHash uint32
  ModifierHashes []uint32
  Variants []DestinyMilestoneActivityVariant
}

type DestinyMilestoneChallenge struct {
  ObjectiveHash uint32
  ActivityHash uint32
}

type DestinyMilestoneQuest struct {
  QuestItemHash uint32
  Activity DestinyMilestoneActivity
  Challenges []DestinyMilestoneChallenge
}

type DestinyMilestone struct {
  MilestoneHash uint32
  AvailableQuests []DestinyMilestoneQuest
  VendorHashes []uint32
  StartDate time.Time
  EndDate time.Time
}

func GetMilestones() (*map[uint32]DestinyMilestone, error) {
  response, err := get("/Destiny2/Milestones/")
  if err != nil { return nil, err }

  var milestones map[uint32]DestinyMilestone

  err = decode(response.Response, &milestones)
  if err != nil {
    return nil, errors.New("Could not decode response: " + err.Error())
  }

  return &milestones, nil
}

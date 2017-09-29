package goBungieNet

type DestinyActivityMode int

const (
  DestinyActivityModeNone DestinyActivityMode = 0
  DestinyActivityModeStory = 2
  DestinyActivityModeStrike = 3
  DestinyActivityModeRaid = 4
  DestinyActivityModeAllPvP = 5
  DestinyActivityModePatrol = 6
  DestinyActivityModeAllPvE = 7
  DestinyActivityModeReserved9 = 9
  DestinyActivityModeControl = 10
  DestinyActivityModeReserved11 = 11
  DestinyActivityModeClash = 12
  DestinyActivityModeReserved13 = 13
  DestinyActivityModeReserved15 = 15
  DestinyActivityModeNightfall = 16
  DestinyActivityModeHeroicNightfall = 17
  DestinyActivityModeAllStrikes = 18
  DestinyActivityModeIronBanner = 19
  DestinyActivityModeReserved20 = 20
  DestinyActivityModeReserved21 = 21
  DestinyActivityModeReserved22 = 22
  DestinyActivityModeReserved24 = 24
  DestinyActivityModeReserved25 = 25
  DestinyActivityModeReserved26 = 26
  DestinyActivityModeReserved27 = 27
  DestinyActivityModeReserved28 = 28
  DestinyActivityModeReserved29 = 29
  DestinyActivityModeReserved30 = 30
  DestinyActivityModeSupremacy = 31
  DestinyActivityModeReserved32 = 32
  DestinyActivityModeSurvival = 37
  DestinyActivityModeCountdown = 38
  DestinyActivityModeTrialsOfTheNine = 39
  DestinyActivityModeSocial = 40
)

func (mode DestinyActivityMode)String() string {
  switch(mode) {
    case DestinyActivityModeNone: return "None"
    case DestinyActivityModeStory: return "Story"
    case DestinyActivityModeStrike: return "Strike"
    case DestinyActivityModeRaid: return "Raid"
    case DestinyActivityModeAllPvP: return "PvP"
    case DestinyActivityModePatrol: return "Patrol"
    case DestinyActivityModeAllPvE: return "PvE"
    case DestinyActivityModeControl: return "Control"
    case DestinyActivityModeClash: return "Clash"
    case DestinyActivityModeNightfall: return "Nightfall"
    case DestinyActivityModeHeroicNightfall: return "Nightfall Heroic"
    case DestinyActivityModeAllStrikes: return "Strikes"
    case DestinyActivityModeIronBanner: return "Iron Banner"
    case DestinyActivityModeSupremacy: return "Supremacy"
    case DestinyActivityModeSurvival: return "Survival"
    case DestinyActivityModeCountdown: return "Countdown"
    case DestinyActivityModeTrialsOfTheNine: return "Trials of the Nine"
    case DestinyActivityModeSocial: return "Social"
  }

  return "Unknown"
}


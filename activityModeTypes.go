package goBungieNet

type DestinyActivityMode int

const (
  ActivityModeNone DestinyActivityMode = 0
  ActivityModeStory = 2
  ActivityModeStrike = 3
  ActivityModeRaid = 4
  ActivityModeAllPvP = 5
  ActivityModePatrol = 6
  ActivityModeAllPvE = 7
  ActivityModeReserved9 = 9
  ActivityModeControl = 10
  ActivityModeReserved11 = 11
  ActivityModeClash = 12
  ActivityModeReserved13 = 13
  ActivityModeReserved15 = 15
  ActivityModeNightfall = 16
  ActivityModeHeroicNightfall = 17
  ActivityModeAllStrikes = 18
  ActivityModeIronBanner = 19
  ActivityModeReserved20 = 20
  ActivityModeReserved21 = 21
  ActivityModeReserved22 = 22
  ActivityModeReserved24 = 24
  ActivityModeReserved25 = 25
  ActivityModeReserved26 = 26
  ActivityModeReserved27 = 27
  ActivityModeReserved28 = 28
  ActivityModeReserved29 = 29
  ActivityModeReserved30 = 30
  ActivityModeSupremacy = 31
  ActivityModeReserved32 = 32
  ActivityModeSurvival = 37
  ActivityModeCountdown = 38
  ActivityModeTrialsOfTheNine = 39
  ActivityModeSocial = 40
)

func (mode DestinyActivityMode)String() string {
  switch(mode) {
    case ActivityModeNone: return "None"
    case ActivityModeStory: return "Story"
    case ActivityModeStrike: return "Strike"
    case ActivityModeRaid: return "Raid"
    case ActivityModeAllPvP: return "PvP"
    case ActivityModePatrol: return "Patrol"
    case ActivityModeAllPvE: return "PvE"
    case ActivityModeControl: return "Control"
    case ActivityModeClash: return "Clash"
    case ActivityModeNightfall: return "Nightfall"
    case ActivityModeHeroicNightfall: return "Nightfall Heroic"
    case ActivityModeAllStrikes: return "Strikes"
    case ActivityModeIronBanner: return "Iron Banner"
    case ActivityModeSupremacy: return "Supremacy"
    case ActivityModeSurvival: return "Survival"
    case ActivityModeCountdown: return "Countdown"
    case ActivityModeTrialsOfTheNine: return "Trials of the Nine"
    case ActivityModeSocial: return "Social"
  }

  return "Unknown"
}


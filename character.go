package goBungieNet

import (
	"time"
)

type DestinyGender int32

const (
	DestinyGenderMale DestinyGender = 0
	DestinyGenderFemale = 1
	DestinyGenderUnknown = 2
)

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
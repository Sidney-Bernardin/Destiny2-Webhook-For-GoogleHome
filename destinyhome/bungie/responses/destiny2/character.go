package destiny2

import (
	"time"
)

type Character struct {
	Characters         CharacterComponent
	CharacterEquipment EquipmentComponent
}

func (c *Character) Get() *Character {
	return c
}

type CharacterComponent struct {
	Character struct {
		Data struct {
			MembershipID             string    `json:"membershipId"`
			MembershipType           int       `json:"membershipType"`
			CharacterID              string    `json:"characterId"`
			DateLastPlayed           time.Time `json:"dateLastPlayed"`
			MinutesPlayedThisSession string    `json:"minutesPlayedThisSession"`
			MinutesPlayedTotal       string    `json:"minutesPlayedTotal"`
			Light                    int       `json:"light"`
			Stats                    struct {
				Num144602215  int `json:"144602215"`
				Num392767087  int `json:"392767087"`
				Num1735777505 int `json:"1735777505"`
				Num1935470627 int `json:"1935470627"`
				Num1943323491 int `json:"1943323491"`
				Num2996146975 int `json:"2996146975"`
				Num4244567218 int `json:"4244567218"`
			} `json:"stats"`
			RaceHash             int    `json:"raceHash"`
			GenderHash           int64  `json:"genderHash"`
			ClassHash            int64  `json:"classHash"`
			RaceType             int    `json:"raceType"`
			ClassType            int    `json:"classType"`
			GenderType           int    `json:"genderType"`
			EmblemPath           string `json:"emblemPath"`
			EmblemBackgroundPath string `json:"emblemBackgroundPath"`
			EmblemHash           int    `json:"emblemHash"`
			EmblemColor          struct {
				Red   int `json:"red"`
				Green int `json:"green"`
				Blue  int `json:"blue"`
				Alpha int `json:"alpha"`
			} `json:"emblemColor"`
			LevelProgression struct {
				ProgressionHash     int `json:"progressionHash"`
				DailyProgress       int `json:"dailyProgress"`
				DailyLimit          int `json:"dailyLimit"`
				WeeklyProgress      int `json:"weeklyProgress"`
				WeeklyLimit         int `json:"weeklyLimit"`
				CurrentProgress     int `json:"currentProgress"`
				Level               int `json:"level"`
				LevelCap            int `json:"levelCap"`
				StepIndex           int `json:"stepIndex"`
				ProgressToNextLevel int `json:"progressToNextLevel"`
				NextLevelAt         int `json:"nextLevelAt"`
			} `json:"levelProgression"`
			BaseCharacterLevel float64 `json:"baseCharacterLevel"`
			PercentToNextLevel float64 `json:"percentToNextLevel"`
		} `json:"data"`
		Privacy int `json:"privacy"`
	} `json:"character"`
}

type EquipmentComponent struct {
	Equipment struct {
		Data struct {
			Items []struct {
				ItemHash              int    `json:"itemHash"`
				ItemInstanceID        string `json:"itemInstanceId"`
				Quantity              int    `json:"quantity"`
				BindStatus            int    `json:"bindStatus"`
				Location              int    `json:"location"`
				BucketHash            int    `json:"bucketHash"`
				TransferStatus        int    `json:"transferStatus"`
				Lockable              bool   `json:"lockable"`
				State                 int    `json:"state"`
				DismantlePermission   int    `json:"dismantlePermission"`
				IsWrapper             bool   `json:"isWrapper"`
				VersionNumber         int    `json:"versionNumber,omitempty"`
				OverrideStyleItemHash int    `json:"overrideStyleItemHash,omitempty"`
				MetricHash            int    `json:"metricHash,omitempty"`
				MetricObjective       struct {
					ObjectiveHash   int64 `json:"objectiveHash"`
					Progress        int   `json:"progress"`
					CompletionValue int   `json:"completionValue"`
					Complete        bool  `json:"complete"`
					Visible         bool  `json:"visible"`
				} `json:"metricObjective,omitempty"`
			} `json:"items"`
		} `json:"data"`
		Privacy int `json:"privacy"`
	} `json:"equipment"`
}

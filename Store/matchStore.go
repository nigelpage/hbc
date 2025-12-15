package store

type MatchStore struct {
	Matches []struct {
		Competition struct {
			Name        string `json:"name"`
			BowlslinkID string `json:"bowlslinkId"`
		} `json:"competition"`
		DutySelector struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
		} `json:"dutySelector"`
		Round struct {
			Number   int    `json:"number"`
			PlayedAt string `json:"playedAt"`
			PlayedOn struct {
				Title string `json:"title"`
				Details string `json:"details"`
			} `json:"playedOn"`
			Venue    string `json:"venue"`
			Opponent string `json:"opponent"`
		} `json:"round"`
		Side struct {
			Updated string `json:"updated"`
			Name    string `json:"name"`
			Manager string `json:"manager"`
			Teams 	[]Team	`json:"teams"`
		} `json:"side"`
	} `json:"matches"`
}

type Team struct {
	Duty string `json:"duty,omitempty"`
	Bowlers []Bowler `json:"team"`
    Shots struct {
        For     int `json:"for"`
        Against int `json:"against"`
    } `json:"shots"`
}

type Bowler struct {
	Position string `json:"position"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

func HasResults(teams []Team) bool {
	if CalculateSidePointsFor(teams) > 0 || CalculateSidePointsAgainst(teams) > 0 {
		return true
	}
	
	return false
}

func CalculateSidePointsFor(teams []Team) int {
	totalFor := 0
	
	for _, team := range teams {
		totalFor += team.Shots.For
		}

	return totalFor
}

func CalculateSidePointsAgainst(teams []Team) int {
	totalAgainst := 0
	
	for _, team := range teams {
		totalAgainst += team.Shots.Against
		}

	return totalAgainst
}

func CalculateSidePoints(teams []Team) int {
	teamWins := 0
	teamLosses := 0
	teamDraws := 0
	totalPointsFor := 0
	totalPointsAgainst := 0
	points := 0

	for _, team := range teams {
		pointsFor := team.Shots.For
		pointsAgainst := team.Shots.Against

		totalPointsFor += pointsFor
		totalPointsAgainst += pointsAgainst

		if pointsFor > pointsAgainst {
			teamWins++
		} else if pointsFor < pointsAgainst {
			teamLosses++
		} else {
			teamDraws++
		}
	}

	points = (teamWins * 2) + teamDraws
	if totalPointsFor > totalPointsAgainst {
		points += 10
	} else if totalPointsFor == totalPointsAgainst {
		points +=5
	}

	return points
}

func GetSideWinDrawLoss(teams []Team) string {
	totalPointsFor := CalculateSidePointsFor(teams)
	totalPointsAgainst := CalculateSidePointsAgainst(teams)
	data := ""

	if totalPointsFor > totalPointsAgainst {
		data = "up"
	} else if totalPointsFor < totalPointsAgainst {
		data = "down"
	} else {
		data = "draw"
	}

	return "data-SideResult=" + data
}

func GetTeamWinDrawLoss(team Team) string {
data := ""

	if team.Shots.For > team.Shots.Against {
		data = "up"
	} else if team.Shots.For < team.Shots.Against {
		data = "down"
	} else {
		data = "draw"
	}

	return "data-teamResult=" + data
}

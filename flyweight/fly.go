package flyweight

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[string]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamID string) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

// getTeamFactory could be use factory interface
func getTeamFactory(team string) Team {
	switch team {
	case "TEAM_B":
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

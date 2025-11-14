package repositories

import "fmt"

type TeamRepository interface {
	SaveTeam(filePath string) error
}

type Team struct {
	TeamName   string   `json:"team_name"`
	TeamMember []string `json:"team_member"`
}

func NewTeamRepository() *Team {
	return &Team{}
}

func (repo *Team) SaveTeam(filePath string) error {
	fmt.Printf("Saving team %+v\n", repo)
	// TODO Realize saving
	return nil
}

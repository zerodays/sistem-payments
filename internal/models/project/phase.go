package project

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/sistem-payments/internal/config"
	"net/http"
	"time"
)

type Phase struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	DateCreated time.Time `json:"date_created"`
}

type Phases []*Phase

func PhasesForProject(id int) (Phases, error) {
	url := fmt.Sprintf("%s/projects/%d/phases", config.Microservices.ProjectsUrl(), id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	phases := make(Phases, 0)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&phases)

	return phases, err
}

func (p *Project) Phases() (Phases, error) {
	return PhasesForProject(p.ID)
}

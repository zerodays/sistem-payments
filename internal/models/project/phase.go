package project

import (
	"encoding/json"
	"fmt"
	"github.com/zerodays/sistem-payments/internal/config"
	"github.com/zerodays/sistem-payments/internal/logger"
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

var phasesCache = make(map[int]Phases)

func PhasesForProject(id int) (Phases, error) {
	ph, err := cb.Execute(func() (interface{}, error) {
		url := fmt.Sprintf("%s/projects/%d/phases", config.Microservices.ProjectsUrl(), id)
		cli := &http.Client{}
		cli.Timeout = 2 * time.Second
		resp, err := cli.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		phases := make(Phases, 0)
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&phases)

		if err != nil {
			return nil, err
		}

		return phases, err
	})

	if err != nil {
		phases, ok := phasesCache[id]
		if ok {
			logger.Log.Warn().Err(err).Msg("using fallback for project phases")
			return phases, nil
		} else {
			return nil, err
		}
	} else {
		phases := ph.(Phases)
		phasesCache[id] = phases

		return phases, nil
	}
}

func (p *Project) Phases() (Phases, error) {
	return PhasesForProject(p.ID)
}

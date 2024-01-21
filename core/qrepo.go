package core

import (
	"encoding/json"
	"os"
)

type Qrepo struct {
	Name    string              `json:"name"`
	Author  string              `json:"author"`
	Git     string              `json:"git"`
	Scripts map[string][]string `json:"scripts"`
}

func readCurrentQrepo() ([]byte, error) {
	dat, err := os.ReadFile("qrepo.json")
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func hasQrepoJsonAlready() bool {
	_, err := os.Stat("qrepo.json")
	return !os.IsNotExist(err)
}

func extractQrepoInfos() (Qrepo, error) {
	var repo Qrepo
	repoBytes, err := readCurrentQrepo()

	if err != nil {
		return repo, err
	}

	if err := json.Unmarshal(repoBytes, &repo); err != nil {
		return repo, err
	}

	return repo, nil
}

func (infos Qrepo) getMapKeys() []string {
	inputMap := infos.Scripts
	keys := make([]string, 0, len(inputMap))

	for key := range inputMap {
		keys = append(keys, key)
	}

	return keys
}

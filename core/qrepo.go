/*
 * This file is part of the Qrepo.
 * Copyright (c) 2024 Nathanne Isip
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
package core

import (
	"encoding/json"
	"os"
)

type Qrepo struct {
	Name        string              `json:"name"`
	Author      string              `json:"author"`
	Version     string              `json:"version"`
	Description string              `json:"description"`
	Git         string              `json:"git"`
	Scripts     map[string][]string `json:"scripts"`
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

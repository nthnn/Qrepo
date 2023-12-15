package main

type Qrepo struct {
	Name    string              `json:"name"`
	Author  string              `json:"author"`
	Git     string              `json:"git"`
	Scripts map[string][]string `json:"scripts"`
}

package main

import (
	"errors"
	"log"
	"os"
)

// GetEnvVar checks if a variable exists.
// It returns an error if the variable is not found
func GetEnvVar(varName string) (string, error) {
	varValue, varExists := os.LookupEnv(varName)
	if !varExists {
		return "", errors.New("Environment variable" + varName + " is not set.")
	}
	return varValue, nil
}

// getDataFromEnv gathers player data from the environment variables
// The logic of which vars are not found is not explicit or well-documented yet.
// For now, Name and Environment vars are chosen arbitrarily to be required.
// It returns error if the required variables are not found.
func (p *Player) setDataFromEnv() error {
	log.Printf("Gathering data from Environment Variables\n")
	playerEnv := "unset"
	playerName := "unset"
	// XXX: This default should be moved to a default_color, setup through NewPlayer(...)
	playerColor := "green"
	err := errors.New("Placeholder error")
	if playerEnv, err = GetEnvVar("ENV_NAME"); err != nil {
		return err
	}
	p.environment = playerEnv
	if playerName, err = GetEnvVar("PLAYER_NAME"); err != nil {
		return err
	}
	p.name = playerName
	if playerColor, err = GetEnvVar("COLOR"); err == nil {
		p.color = playerColor
	}
	return nil
}

// GatherData sets player data from different sources
// It attempts to fill data from the environment first.
// It could read the data from a file, etcd or redis in the future.
// It returns error if it cannot get the data by the available means.
func (p *Player) GatherData() error {
	return p.setDataFromEnv()
}

package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

//Command is a configuration structure which consists of a name (e.g., print_hello) and a command (e.g., "echo hello")
type Command struct {
	Name    string
	Command string
}

//Environment is a specific environment which can be entered with "sanic env"
type Environment struct {
	Commands []Command
}

//SanicConfig is the global structure of entries in sanic.yaml
type SanicConfig struct {
	Environments map[string]Environment
}

//ReadFromPath returns a new SanicConfig from the given filesystem path to a yaml file
func ReadFromPath(configPath string) (*SanicConfig, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.New("configuration file could not be read: " + err.Error())
	}

	ret := new(SanicConfig)
	err = yaml.UnmarshalStrict(data, ret)
	if err != nil {
		return nil, errors.New("configuration file error: " + err.Error())
	}

	return ret, nil
}

//Read returns a new SanicConfig, given that the environment (e.g., sanic env) has one configured
func Read() (*SanicConfig, error) {
	configPath := os.Getenv("SANIC_CONFIG")
	if configPath == "" {
		return nil, errors.New("enter an environment with 'sanic env'")
	}

	return ReadFromPath(configPath)
}

//CurrentEnvironment returns an Environment struct corresponding to the environment the user is in.
//Fails if the user is not in an environment
func CurrentEnvironment(cfg *SanicConfig) (*Environment, error) {
	sanicEnv := os.Getenv("SANIC_ENV")
	if sanicEnv == "" {
		return nil, errors.New("enter an environment with 'sanic env'")
	}
	if ret, exists := cfg.Environments[sanicEnv]; exists {
		return &ret, nil
	}
	return nil, errors.New("the environment you are you does not exist, ensure it was not removed from the configuration")
}

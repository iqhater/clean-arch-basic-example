package internal

import (
	"os"
	"testing"
)

// declare cfg as global variable
var cfg *Config

// Good alternative for init() function
func TestMain(m *testing.M) {

	// init .env config
	InitEnv("../.env.test")

	// init config environments
	cfg = NewConfig()

	// run tests
	os.Exit(m.Run())
}

func TestNewConfigNotEmptyData(t *testing.T) {

	if cfg.HTTP_PORT == "" || cfg.DB_FILENAME == "" {
		t.Errorf("Config struct should not have an empty values: got %v", cfg)
	}
}

func TestNewConfigEmptyData(t *testing.T) {

	envs := []string{"HTTP_PORT", "DB_FILENAME"}
	envsBuffer := make(map[string]string)

	// clear environments variables
	for _, env := range envs {

		// save to buffer for later restore envs
		envsBuffer[env] = os.Getenv(env)

		// clear env
		os.Setenv(env, "")
	}

	if os.Getenv(cfg.HTTP_PORT) != "" || os.Getenv(cfg.DB_FILENAME) != "" {
		t.Errorf("Config struct should be an empty values: got %v", cfg)
	}

	// restore envs for another tests
	// loads values from .env into the system
	for k, v := range envsBuffer {
		os.Setenv(k, v)
	}
}

func TestNewConfigEnvsNotExist(t *testing.T) {

	envs := []string{"HTTP_PORT", "DB_FILENAME"}
	envsBuffer := make(map[string]string)

	// clear environments variables
	for _, env := range envs {

		// save to buffer for later restore envs
		envsBuffer[env] = os.Getenv(env)

		// clear env
		os.Setenv(env, "")
	}

	// flush all environments
	// os.Clearenv()

	for _, env := range envs {

		// clear env
		os.Setenv(env, "")

		value := getEnv(env)
		if value != "" {
			t.Errorf("Env variable %s should not exist!", env)
		}
	}

	// restore envs for another tests
	// loads values from .env into the system
	for k, v := range envsBuffer {
		os.Setenv(k, v)
	}
}

func TestGetEnvExist(t *testing.T) {

	envs := []string{"HTTP_PORT", "DB_FILENAME"}

	for _, env := range envs {

		value := getEnv(env)
		if value == "" {
			t.Errorf("Env variable %s does not exist!", env)
		}
	}
}

func TestGetEnvNotExist(t *testing.T) {

	envs := []string{"FAKE_HTTP_PORT", "FAKE_DB_FILENAME"}

	for _, env := range envs {

		// clear env
		os.Setenv(env, "")

		value := getEnv(env)
		if value != "" {
			t.Errorf("Env variable %s should not exist!", env)
		}
	}
}

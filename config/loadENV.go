package config

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)
var buildInfo *debug.BuildInfo
var logger zerolog.Logger

// TRACE (-1): for tracing the code execution path.
// DEBUG (0): messages useful for troubleshooting the program.
// INFO (1): messages describing the normal operation of an application.
// WARNING (2): for logging events that need may need to be checked later.
// ERROR (3): error messages for a specific operation.
// FATAL (4): severe errors where the application cannot recover. os.Exit(1) is called after the message is logged.
// PANIC (5): similar to FATAL, but panic() is called instead.

func init() {

	buildInfo, _ = debug.ReadBuildInfo()

	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
}

func LoadEnvVars() {
	box := packr.New("myBox", "../")

	s, err := box.FindString(".env")
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	}
	//   LOAD ENVIRONMENT VARIABLES
	myEnv, err := godotenv.Unmarshal(s)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	}
	for i, j := range myEnv {
		os.Setenv(i, j)
	}
	logger.Info().Msg("📢 Info message :" + "Environment variables loaded")
}

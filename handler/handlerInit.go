package handler

import (
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
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

type DatabaseCollections struct {
	Mongo *mongo.Database
}

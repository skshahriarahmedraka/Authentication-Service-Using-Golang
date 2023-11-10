package monodb

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"os"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func MongodbConnection() *mongo.Database {

	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/?maxPoolSize=%v&w=majority", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_IP"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_MAXPOOLSIZE"))
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Context to set timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	fmt.Println("🚀 ~ file: connection.go ~ line 51 ~ funcMongodbConnection ~ err : ", err)

	if err != nil {

		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {

		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	}
	Mydb := client.Database("userdb")
	if err == nil {
		logger.Info().Msg("📢 Info message : ⚡😍 sucessfully connected to database")
	}
	return Mydb

}

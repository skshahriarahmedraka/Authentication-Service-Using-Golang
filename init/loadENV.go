package init

import (
	"log"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	box := packr.New("myBox", "../")

	s, err := box.FindString(".env")
    // fmt.Println("🚀 ~ file: loadEnv.go ~ line 14 ~ funcLoadEnvironmentVar ~ s : ", s)
	if err != nil {
		log.Println("❌ ~ file: loadEnv.go ~ line 13 ~ funcLoadEnvironmentVar ~ err : ", err)
	}
	log.Println(s)
	//   LOAD ENVIRONMENT VARIABLES

	myEnv, e := godotenv.Unmarshal(s)
	if e != nil {
		log.Println("❌ ~ file: loadEnv.go ~ line 19 ~ funcLoadEnvironmentVar ~ e : ", e)
	}
	for i, j := range myEnv {
		os.Setenv(i, j)
	}
}

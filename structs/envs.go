package structs

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type yamlVars struct {
	Clear struct {
		IgnoreExtensions []string `yaml:"ignore_extensions"`
		IgnoreFiles      []string `yaml:"ingore_files"`
	} `yaml:"clear"`
}

type envsVars struct {
	ElasticUrl    string
	BulkSizeMB    int
	BulkWorkers   int
	EventsPerFile int
	EventsPerPage int
	FoxesConfig   yamlVars
}

var Variables *envsVars = &envsVars{}

func getValueDefault(key string, default_value string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return default_value
	}

	return value
}

func (i *envsVars) FillEnvs() {
	_ = godotenv.Load()

	Variables.ElasticUrl = getValueDefault("ELASTIC_URL", "http://127.0.0.1:9200")
	Variables.BulkSizeMB, _ = strconv.Atoi(getValueDefault("BULK_SIZE_MB", "5"))
	Variables.BulkWorkers, _ = strconv.Atoi(getValueDefault("BULK_WORKERS", "4"))
	Variables.EventsPerFile, _ = strconv.Atoi(getValueDefault("BACKUP_EVENTS_PER_FILE", "1000000"))
	Variables.EventsPerPage, _ = strconv.Atoi(getValueDefault("EVENTS_PER_PAGE", "5000"))

	fillByYml()
}

func fillByYml() {
	file, err := ioutil.ReadFile("foxes.yaml")

	if err != nil {
		log.Fatalln("foxes.yaml not found")
	}

	var data yamlVars

	err = yaml.Unmarshal(file, &data)

	if err != nil {
		log.Fatalln("foxes.yaml bad formated")
	}

	Variables.FoxesConfig = data
}

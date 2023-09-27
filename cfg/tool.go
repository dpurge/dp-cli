package cfg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func ReadConfig() {

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(filepath.Join(homedir, ".dp-cli"))
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()
}

func GetTool(application string, name string) (string, error) {
	toolPath := viper.GetString(fmt.Sprintf("%s.%s", application, name))
	return toolPath, nil
}

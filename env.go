package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Configure :
func Configure() {

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("ENV : Failed to get working directory")
		return
	}

	//check for environment variable in both cases(lower and upper)
	env := os.Getenv("ENVIRONMENT")
	if len(env) == 0 {
		env = os.Getenv("environment")
	}

	if len(env) == 0 {
		fmt.Println("** Enviroment is not define **")
		return
	}

	filePath := fmt.Sprintf("%v/%v.env", workingDir, env)
	variables := readFile(filePath)
	setEnv(variables)
}

func setEnv(variables map[string]string) {

	for key, value := range variables {
		os.Setenv(key, value)
	}
}

func readFile(filePath string) map[string]string {

	variablesMap := make(map[string]string)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ENV : Failed to get file for current environment")
		return variablesMap
	}
	values := strings.Split(string(data), "\n")

	for _, value := range values {
		keyValue := strings.Split(value, "=")
		if len(keyValue) >= 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.Join(keyValue[1:], "=")
			variablesMap[key] = strings.TrimSpace(value)
		}
	}
	return variablesMap
}

// IsProduction :
func IsProduction() bool {

	env := os.Getenv("ENVIRONMENT")
	if len(env) == 0 {
		env = os.Getenv("environment")
	}

	if len(env) == 0 {
		return false
	}
	return (env == "production" || env == "PRODUCTION")
}

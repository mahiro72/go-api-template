package env

import (
	"fmt"
	"log"
	"os"
)

var CONFIG config

type config struct {
	ENVIRONMENT string
	MYSQL       mySQL
}

type mySQL struct {
	database     string
	user         string
	password     string
	rootPassword string
	host         string
	DSN          string
}

func getEnv(envPath string) (string, error) {
	v := os.Getenv(envPath)
	if v == "" {
		return "", fmt.Errorf("%s is empty", envPath)
	}

	return v, nil
}

func initEnvironmentConfig() error {
	var err error

	if CONFIG.ENVIRONMENT, err = getEnv("ENVIRONMENT"); err != nil {
		return err
	}

	return nil
}

func initMYSQLConfig() error {
	var err error

	if CONFIG.ENVIRONMENT == "prod" {
		// prod環境の場合MYSQL_DSNからdsnを取得する
		if CONFIG.MYSQL.DSN, err = getEnv("MYSQL_DSN"); err != nil {
			return err
		}
		return nil
	} else if CONFIG.ENVIRONMENT == "dev" {
		log.Println("environment is dev")
	} else {
		return fmt.Errorf("config.environment is invalid")
	}

	if CONFIG.MYSQL.database, err = getEnv("MYSQL_DATABASE"); err != nil {
		return err
	}

	if CONFIG.MYSQL.user, err = getEnv("MYSQL_USER"); err != nil {
		return err
	}

	if CONFIG.MYSQL.password, err = getEnv("MYSQL_PASSWORD"); err != nil {
		return err
	}

	if CONFIG.MYSQL.rootPassword, err = getEnv("MYSQL_ROOT_PASSWORD"); err != nil {
		return err
	}

	if CONFIG.MYSQL.host, err = getEnv("MYSQL_HOST"); err != nil {
		return err
	}

	CONFIG.MYSQL.DSN = fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true",
		CONFIG.MYSQL.user,
		CONFIG.MYSQL.password,
		CONFIG.MYSQL.host,
		CONFIG.MYSQL.database,
	)
	return nil
}

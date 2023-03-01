package env

import "fmt"

func init() {
	if err := initEnvironmentConfig(); err != nil {
		panic(
			fmt.Sprintf("error: config.initEnvironmentConfig failed: %v", err),
		)
	}
	if err := initMYSQLConfig(); err != nil {
		panic(
			fmt.Sprintf("error: config.initMYSQLConfig failed: %v", err),
		)
	}
}

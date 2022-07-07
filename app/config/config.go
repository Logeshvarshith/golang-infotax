package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppPort                       string
	DriverName                    string
	DatabaseSourceName            string
	RefreshSecret                 string
	UserLoginDetailBaseUrl        string
	EmployeeOfficialDetailBaseUrl string
	IdTokenExp                    int64
	RefreshTokenExp               int64
	HandlerTimeout                int64
}

func ParseConfig() *Config {

	idTokenExp, _ := strconv.Atoi(os.Getenv("ID_TOKEN_EXP"))
	refreshTokenExp, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXP"))
	handlerTimeout, _ := strconv.Atoi(os.Getenv("HANDLER_TIMEOUT"))

	return &Config{
		AppPort:                       os.Getenv("APP_PORT"),
		DriverName:                    os.Getenv("DRIVER_NAME"),
		DatabaseSourceName:            os.Getenv("DATABASE_SOURCE_NAME"),
		RefreshSecret:                 os.Getenv("REFRESH_SECRET"),
		UserLoginDetailBaseUrl:        os.Getenv("USER_LOGIN_DETAIL_BASE_URL"),
		EmployeeOfficialDetailBaseUrl: os.Getenv("EMPLOYEE_OFFICIAL_DETAIL_BASE_URL"),
		IdTokenExp:                    int64(idTokenExp),
		RefreshTokenExp:               int64(refreshTokenExp),
		HandlerTimeout:                int64(handlerTimeout),
	}

}

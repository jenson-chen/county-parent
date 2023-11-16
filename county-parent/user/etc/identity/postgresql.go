package identity

import "fmt"

type PostgreSQL struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

func (p PostgreSQL) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", p.Host, p.User, p.Password, p.DBName, p.Port)
}

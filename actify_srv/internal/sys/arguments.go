package sys

import (
	"flag"
	"fmt"
)

type Arguments struct {
	Host     string
	Port     uint
	DBName   string
	User     string
	Password string
	SSLMode  bool
	HelpReq  bool
}

func NewActifyArgs() Arguments {
	return Arguments{
		Host:     "localhost",
		Port:     5432,
		User:     "",
		DBName:   "",
		Password: "",
		SSLMode:  false,
		HelpReq:  false,
	}
}

func (args Arguments) PostgresConnStr() string {
	sslMode := "disable"
	if args.SSLMode {
		sslMode = "require"
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		args.Host, args.Port, args.User, args.Password, args.DBName, sslMode,
	)
}


func (args *Arguments) Parse(input []string) error {
	fs := flag.NewFlagSet("dbflags", flag.ContinueOnError)

	fs.BoolVar(&args.HelpReq, "h", false, "show help")
	fs.BoolVar(&args.HelpReq, "help", false, "show help")
	fs.StringVar(&args.Host, "db_host", "localhost", "database host")
	fs.UintVar(&args.Port, "db_port", 5432, "database port")
	fs.StringVar(&args.DBName, "db_name", "postgres", "Database name")
	fs.BoolVar(&args.SSLMode, "db_ssl", false, "database use ssl check")
	fs.StringVar(&args.Password, "db_pw", "", "Database connection password")

	fs.Usage = func() {
		fmt.Println("Custom Usage:")
		fs.PrintDefaults()
	}

	err := fs.Parse(input)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	if args.HelpReq {
		fs.Usage()
		return nil
	}

	if args.DBName == "" || args.Password == "" || args.User == "" {
		return nil
	}

	return nil
}



package flag

import (
	"AfterEnd/global"
	"flag"
)

type Options struct {
	DB         bool
	Encrypt    bool
	UserCreate string
}

// Parse 解析命令参数，并对不同的命令行参数的值来执行不同的操作
func Parse() {
	dbFlag := flag.Bool("db", false, "auto migrate database")
	EncryptFlag := flag.Bool("encrypt", false, "encrypt profile")
	userFlag := flag.String("u", "", "create user through terminal")
	flag.Parse()
	var option = Options{
		DB:         *dbFlag,
		Encrypt:    *EncryptFlag,
		UserCreate: *userFlag,
	}
	Execute(option)
}

func Execute(options Options) {
	if options.DB {
		AutoMigration()
		return
	}
	if options.Encrypt {
		EncryptProfile()
	}

	// 输入非法
	if options.UserCreate != "" && options.UserCreate != "user" && options.UserCreate != "admin" {
		global.Log.Error("Invalid user type. Please use \"user\" or \"admin\".")
		return
	}
	switch options.UserCreate {
	case "user":
		CreateUser("user")
	case "admin":
		CreateUser("admin")
	}
}

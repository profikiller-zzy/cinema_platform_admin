package flag

import (
	"flag"
)

type Options struct {
	DB      bool
	Encrypt bool
}

// Parse 解析命令参数，并对不同的命令行参数的值来执行不同的操作
func Parse() {
	dbFlag := flag.Bool("db", false, "auto migrate database")
	EncryptFlag := flag.Bool("encrypt", false, "encrypt profile")
	flag.Parse()
	var option = Options{
		DB:      *dbFlag,
		Encrypt: *EncryptFlag,
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

}

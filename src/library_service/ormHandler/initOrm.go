package ormHandler

import (
	"os"
	"sync"

	"github.com/astaxie/beego/orm"
)

var instantiated orm.Ormer = nil
var once sync.Once

func New() *orm.Ormer {
	once.Do(func() {
		orm.RegisterDataBase("default", os.Getenv("DBDriverName"),
			os.Getenv("DBUserName")+":"+
				os.Getenv("DBPassword")+"@tcp("+os.Getenv("DBUrl")+":"+os.Getenv("DBPort")+")/"+
				os.Getenv("DBName")+"?charset=utf8")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		instantiated = orm.NewOrm()
		// fmt.Println(instantiated)
	})
	return &instantiated
}

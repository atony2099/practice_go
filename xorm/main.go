package main

import (
	"fmt"
	//"log"

	//"test/models"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine



var eg *xorm.EngineGroup

func main() {
	conns := []string{
		"postgres://postgres:root@localhost:5432/test?sslmode=disable;", // 第一个默认是master
		"postgres://postgres:root@localhost:5432/test1?sslmode=disable;", // 第二个开始都是slave
		"postgres://postgres:root@localhost:5432/test2?sslmode=disable",
	}


	xorm.RandomPolicy()
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

    var err error
	eg, err = xorm.NewEngineGroup("postgres", conns)

	fmt.Println(err,eg)
}


// 
// func main() {

	




	// var err error
	// engine, err = xorm.NewEngine("mysql", "root:123456@/xorm_t?charset=utf8")

	// // INSERT INTO user (name) values (?)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var user = models.MediaAppAdsenseShield{}
	// affected, err := engine.Insert(&user)

	// fmt.Println(user, err, affected)


	// engine.Delete(user)

	// user := &models.MediaAppAdsenseShield{Id: 3}
	// has, err := engine.Unscoped().Get(user)

	// fmt.Println(has, err)
	// fmt.Println(affected, err)

// }

// Package postgres
// @Description:
// @CreateTime: 2022-01-11 15:27:17
// @author: guirf
package postgres
import (
	"fmt"
	. "github.com/aiotbox/edgebox/internal/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)
//
//  NewClient
//  @time: 2022-01-11 15:27:21
//  @Description:
//  @param conf
//  @return *gorm.DB
//  @return error
//
func NewClient(conf Configuration)(*gorm.DB,error){
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai",conf.Host,conf.Port,conf.Username,conf.DatabaseName,conf.Password)
	var c gorm.Config
	if conf.BatchSize > 0{
		c.CreateBatchSize = conf.BatchSize
	}
	db, err :=gorm.Open(postgres.Open(dsn), &c)
	return db,err
}

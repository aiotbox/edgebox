// Package mysql
// @Description:
// @CreateTime: 2022-01-11 13:39:25
// @author: guirf
package mysql
import (
	"fmt"
	. "github.com/aiotbox/edgebox/internal/pkg/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//
//  NewClient
//  @time: 2022-01-11 15:27:27
//  @Description:
//  @param conf
//  @return *gorm.DB
//  @return error
//
func NewClient(conf Configuration)(*gorm.DB,error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",conf.Username,conf.Password,conf.Host,conf.Port,conf.DatabaseName,conf.Timeout)
	var c gorm.Config
	if conf.BatchSize > 0{
		c.CreateBatchSize = conf.BatchSize
	}
	db, err :=gorm.Open(mysql.Open(dsn), &c)
	return db,err
}

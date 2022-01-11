// Package sqlite
// @Description:
// @CreateTime: 2022-01-11 15:27:08
// @author: guirf
package sqlite
import (
	. "github.com/aiotbox/edgebox/internal/pkg/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
//
//  NewClient
//  @time: 2022-01-11 15:27:11
//  @Description:
//  @param conf
//  @return *gorm.DB
//  @return error
//
func NewClient(conf Configuration)(*gorm.DB,error){
	dsn := conf.DatabaseName
	var c gorm.Config
	if conf.BatchSize > 0{
		c.CreateBatchSize = conf.BatchSize
	}
	db, err :=gorm.Open(sqlite.Open(dsn), &c)
	return db,err
}


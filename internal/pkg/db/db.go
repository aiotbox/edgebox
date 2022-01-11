// Package db
// @Description:
// @CreateTime: 2022-01-11 15:27:35
// @author: guirf
package db
import (
	"errors"
	"github.com/aiotbox/edgebox/internal/pkg/db/mysql"
	"github.com/aiotbox/edgebox/internal/pkg/db/postgres"
	"github.com/aiotbox/edgebox/internal/pkg/db/sqlite"
	"gorm.io/gorm"
)
var (
	ErrNotFound            = errors.New("Item not found")
	ErrUnsupportedDatabase = errors.New("Unsupported database type")
	ErrInvalidObjectId     = errors.New("Invalid object ID")
	ErrNotUnique           = errors.New("Resource already exists")
	ErrCommandStillInUse   = errors.New("Command is still in use by device profiles")
	ErrSlugEmpty           = errors.New("Slug is nil or empty")
	ErrNameEmpty           = errors.New("Name is required")
)
//
//  Configuration
//  @Description:
//
type Configuration struct {
	DbType       string
	Host         string
	Port         int
	Timeout      int
	DatabaseName string
	Username     string
	Password     string
	BatchSize    int
}
var sqldb *gorm.DB
//
//  GetDB
//  @time: 2022-01-11 15:27:48
//  @Description:
//  @return *gorm.DB
//
func GetDB()*gorm.DB{
	return sqldb
}
//
//  Register
//  @time: 2022-01-11 15:27:44
//  @Description:
//  @param dbname
//  @param conf
//  @return err
//
func Register(dbname string,conf Configuration)(err error){
	switch dbname{
	case "sqlite":
		sqldb,err = sqlite.NewClient(conf)
		return err
		break
	case "mysql":
		sqldb,err = mysql.NewClient(conf)
		return err
		break
	case "postgres":
		sqldb,err = postgres.NewClient(conf)
		return err
		break
	default:
		sqldb = nil
		return ErrUnsupportedDatabase
		break
	}
	return err
}


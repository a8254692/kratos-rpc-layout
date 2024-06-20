package data

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"

	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data/ent"
)

var EntClient *ent.Client
var RdbClient *redis.Client

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData() (*Data, func(), error) {
	var entClient *ent.Client
	if EntClient == nil {
		drv, err := getNewEntDrv()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get ent client, error:%s", err.Error()))
			return nil, nil, err
		}

		entClient = drv
	} else {
		entClient = EntClient
	}

	var rdbClient *redis.Client
	if RdbClient == nil {
		rdb, err := getNewRedisCli()
		if err != nil || rdb == nil {
			panic(fmt.Sprintf("failed to connect redis error:%s", err.Error()))
			return nil, nil, err
		}

		rdbClient = rdb
	} else {
		rdbClient = RdbClient
	}

	d := &Data{
		db:  entClient,
		rdb: rdbClient,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func getNewEntDrv() (*ent.Client, error) {
	prefix := fmt.Sprintf("data.database")
	dialect := viper.GetString(prefix + ".dialect")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetString(prefix + ".port")
	maxIdle := viper.GetInt(prefix + ".maxIdle")
	maxOpen := viper.GetInt(prefix + ".maxOpen")
	maxLifetime := viper.GetInt(prefix + ".maxLifetime")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	db, err := sql.Open(dialect, url)
	if err != nil {
		log.Errorf("failed opening connection to: %s path:%s err:%s", dialect, host, err.Error())
		return nil, err
	}
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)
	db.DB().SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	drv := sql.OpenDB(dialect, db.DB())
	entDrv := ent.NewClient(ent.Driver(drv))
	if entDrv == nil {
		log.Errorf("failed opening ent")
		return nil, err
	}

	EntClient = entDrv
	return entDrv, nil
}

func getNewRedisCli() (*redis.Client, error) {
	prefix := fmt.Sprintf("data.redis")
	addr := viper.GetString(prefix + ".addr")
	password := viper.GetString(prefix + ".password")
	database := viper.GetInt(prefix + ".database")
	maxActive := viper.GetInt(prefix + ".maxActive")
	//idleTimeout := time.Duration(viper.GetInt(prefix+".idleTimeout")) * time.Second

	rdb := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         database,
		MaxRetries: 3,
		PoolSize:   maxActive,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("failed opening rdb path: %s err: %s", addr, err.Error())
		return nil, err
	}

	RdbClient = rdb

	return rdb, nil
}

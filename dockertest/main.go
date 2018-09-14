package main

import (
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/gin-gonic/gin.v1"

	"fmt"

	"database/sql"
	"io/ioutil"
	"log"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		ul := UserList{}
		Orm.Get(&ul)

		c.JSON(http.StatusOK, &ul)
	})
	router.Run(":8000")
}

type UserList struct {
	Id       string `xorm:"not null pk VARCHAR(16)"`
	Username string `xorm:"not null default '' comment('用户名') unique VARCHAR(30)"`
	Password string `xorm:"default '' comment('密码') VARCHAR(1024)"`
	WebCode  string `xorm:"default 'tbet' comment('网站编码') VARCHAR(30)"`
}

var Orm = &xorm.Engine{}

func init() {

	Config.GetConf()
	Mysql_connStr = Config.Mysqluser + ":" + Config.Mysqlpass + "@tcp(" + Config.Mysqlhost + ":" + Config.Mysqlport + ")/" + Config.Mysqldb + "?charset=" + Config.Mysqlcharset
	SALT = Config.JwtSecret
	var err error

	Orm, err = xorm.NewEngine("mysql", Mysql_connStr)

	Orm.DB().SetConnMaxLifetime(30)

	Orm.DB().SetMaxOpenConns(5000)

	fmt.Printf(Mysql_connStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Orm.ShowSQL(true)

}

type Conf struct {
	Ginurl       string `yaml:"ginurl"`
	Ginport      string `yaml:"ginport"`
	Logurl       string `yaml:"logurl"`
	Logport      string `yaml:"logport"`
	Tokenurl     string `yaml:"tokenurl"`
	Tokenport    string `yaml:"tokenport"`
	Grpcurl      string `yaml:"grpcurl"`
	Grpcport     string `yaml:"grpcport"`
	Mysqlhost    string `yaml:"mysqlhost"`
	Mysqlport    string `yaml:"mysqlport"`
	Mysqluser    string `yaml:"mysqluser"`
	Mysqlpass    string `yaml:"mysqlpass"`
	Mysqldb      string `yaml:"mysqldb"`
	Mysqlcharset string `yaml:"mysqlcharset"`
	Redisstatus  string `yaml:"redisstatus"`
	Redisnetwork string `yaml:"redisnetwork"`
	Redisaddr    string `yaml:"redisaddr"`
	Redisport    string `yaml:"redisport"`
	Redisprefix  string `yaml:"redisprefix"`
	Redispwd     string `yaml:"redispwd"`
	Redisdb      int    `yaml:"redisdb"`
	Redisbug     string `yaml:"redisbug"`
	Encodeurl    string `yaml:"encodeurl"`
	Decodeurl    string `yaml:"decodeurl"`
	Webcode      string `yaml:"webcode"`
	JwtSecret    string `yaml:"jwtsecret"`
}

func (c *Conf) GetConf() *Conf {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println(dir)
	yamlFile, err := ioutil.ReadFile("/config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

var Config Conf
var db *sql.DB //数据库句柄指针
var Mysql_connStr = ""

// SALT 密钥
var SALT = ""

package conf

import (
  "encoding/json"
  "fmt"
  "log"
  "os"
  "xProcessBackend/model"
)

type Configer struct {
  Database         Database `json:"database"`
  Server           string   `json:"server"`
  FilePath         string   `json:"file_path"`
  ExportPath       string   `json:"export_path"`
  UsvName          string   `json:"usv_name"`
  ClientName       string   `json:"client_name"`
  ActivityAliasNum int      `json:"activity_alias_num"`
}
type Database struct {
  MysqlDsn string `json:"mysql_dsn"`
  PgDsn    string `json:"pg_dsn"`
  RedisUrl string `json:"redis_url"`
  RedisPsd string `json:"redis_psd"`
}

// Init 初始化配置项,返回端口号
func Init() {
  conf := GetConfig()
  // 连接数据库
  log.Print(conf)
  model.Database(conf.Database.PgDsn)
  model.RedisPool(conf.Database.RedisUrl, conf.Database.RedisPsd)

  // 返回配置文件中服务端口号

}

func GetConfig() Configer {
  var config Configer
  file, _ := os.Open("config.json")
  defer file.Close()
  decoder := json.NewDecoder(file)
  err := decoder.Decode(&config)
  if err != nil {
    fmt.Println("Get Config Error :", err)
  }
  return config
}

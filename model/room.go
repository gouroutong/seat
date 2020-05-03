package model

import (
  "encoding/json"
  "fmt"
  "github.com/garyburd/redigo/redis"
  "log"
)

type Room struct {
  Id       int64  `json:"id"`
  Name     string `json:"name"`
  X        int    `json:"x"`
  Y        int    `json:"y"`
  SeatList string `json:"seat_list"`
  Floor    int8   `json:"floor"`
}

func (room *Room) GetRoom() error {
  return DB.Where("id =?", room.Id).Find(&room).Error
}

func (room *Room) NewRoom() error {
  var err error
  if room.Id == 0 {
    err = DB.Create(room).Error
  } else {
    err = DB.Model(room).Where("id =?", room.Id).Update(room).Error
  }
  room.FmtSeatList(err)
  return err
}
func (room *Room) FmtSeatList(err error) {
  var (
    seatId int64
    bytes  []byte
  )
  var seatList [][]struct {
    Type   int    `json:"type"`
    SeatId int64  `json:"seat_id"`
    Id     string `json:"id"`
  }
  err = json.Unmarshal([]byte(room.SeatList), &seatList)
  if err != nil {
    return
  }
  seatId, err = GetId("seat_id")
  if err != nil {
    err = SetId("seat_id", 0)
    if err != nil {
      return
    }
    seatId, err = GetId("seat_id")
    if err != nil {
      return
    }
  }
  seatId++
  log.Print(seatId)

  var tableCel = "id,room_id"
  var sql = fmt.Sprintf("INSERT INTO seat (%s) VALUES", tableCel)
  //var sql = fmt.Sprintf("INSERT INTO seat (%s) VALUES", tableCel)
  for i := 0; i < len(seatList); i++ {
    seat := seatList[i]
    for i2 := 0; i2 < len(seat); i2++ {
      seat2 := &seat[i2]
      if seat2.Type == 0 {
        seat2.SeatId = seatId
        seatId++
        sql += fmt.Sprintf("('%d','%d'),", seat2.SeatId, room.Id)
      }
    }
  }
  log.Println(sql[:len(sql)-1] + ";")
  err = DB.Exec(sql[:len(sql)-1] + ";").Error
  if err != nil {
    return
  }
  //log.Print(seatList)
  err = SetId("seat_id", seatId)
  if err != nil {
    return
  }
  bytes, err = json.Marshal(seatList)
  if err != nil {
    return
  }
  room.SeatList = string(bytes)
  err = DB.Model(room).Where("id =?", room.Id).Update(room).Error
  return
}

// 设置最近浏览的 id
func SetId(k string, v int64) error {
  c := RS.Get()
  defer c.Close()
  _, err := c.Do("SET", k, v)
  if err != nil {
    fmt.Println("redis set failed:", err)
    return err
  }
  return nil
}

// 获取最近浏览记录
func GetId(k string) (int64, error) {
  c := RS.Get()
  defer c.Close()
  res, err := redis.Int64(c.Do("GET", k))
  if err != nil {
    fmt.Println("redis get failed:", err)
    return 0, err
  }
  return res, nil //return res, err
}

package model

type Seat struct {
  Id      int64  `json:"id"`
  Name    string `json:"name"`
  FloorId int64  `json:"floor_id"`
  RoomId  int64  `json:"room_id"`
}

package model

type Order struct {
  Id          int64 `json:"id"`
  OrderUserId int64 `json:"order_user_id"`
  StartTime   int64 `json:"start_time"`
  EndTime     int64 `json:"end_time"`
  SeatId      int64 `json:"seat_id"`
}

func (order *Order) NewOrder() error {
  var err error
  if order.Id == 0 {
    err = DB.Create(order).Error
  } else {
    err = DB.Model(order).Where("id =?", order.Id).Update(order).Error
  }
  return err
}

func OrderList(seatId int64, list *[]Order) error {
  return DB.Model(&Order{}).Select("*").Where("seat_id =?", seatId).Scan(list).Error
}

package model

func migration() {
  DB.AutoMigrate(&User{}).
    AutoMigrate(&Order{}).
    AutoMigrate(&Room{}).
    AutoMigrate(&Seat{})

}

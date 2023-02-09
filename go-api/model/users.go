package main

import {
  "sync"
}

type Users struct {
  id string `json:"id"`
  name string `josn:"name"`
  email string `gorm:"primaryKey" json:"email"`
  password string `json:"password"`
  created_at datetime `json:"created_at"`
  updated_at datetime `json:"updated_at"`
}

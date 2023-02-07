package submarine_war

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	Title             string     `json:"title"`
	TitleFontSize     int        `json:"titleFontSize"`
	FontSize          int        `json:"fontSize"`
	SmallFontSize     int        `json:"smallFontSize"`
	ScreenWidth       int        `json:"screenWidth"`
	ScreenHeight      int        `json:"screenHeight"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BgColor           color.RGBA `json:"bgColor"`
	BulletWidth       int        `json:"bulletWidth"`
	BulletHeight      int        `json:"bulletHeight"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	BulletInterval    int64      `json:"bulletInterval"`
	MaxBulletNum      int        `json:"maxBulletNum"`
	BulletColor       color.RGBA `json:"bulletColor"`
	AlienSpeedFactor  float64    `json:"alienSpeedFactor"`
}

func loadConfig() *Config {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}
	defer f.Close()
	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}
	return &cfg
}

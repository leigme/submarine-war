package submarine_war

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Input struct {
	lastBulletTime time.Time
}

func (i *Input) Update(game *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		game.ship.x -= game.cfg.ShipSpeedFactor
		maxLeft := -float64(game.ship.width) / 2
		if game.ship.x < maxLeft {
			game.ship.x = maxLeft
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		game.ship.x += game.cfg.ShipSpeedFactor
		maxRight := float64(game.cfg.ScreenWidth) - float64(game.ship.width)/2
		if game.ship.x > maxRight {
			game.ship.x = maxRight
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if len(game.bullets) < game.cfg.MaxBulletNum &&
			time.Now().Sub(i.lastBulletTime).Milliseconds() > game.cfg.BulletInterval {
			bullet := NewBullet(game.cfg, game.ship)
			game.addBullet(bullet)
		}
	}
}

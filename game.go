package submarine_war

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeOVer
)

type Game struct {
	cfg     *Config
	ship    *Ship
	bullets map[*Bullet]struct{}
	aliens  map[*Alien]struct{}
	input   *Input
	mode    Mode
}

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

func (g *Game) init() {
	g.CreateAliens()
	g.CrateFonts()
}

func (g *Game) CrateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	g := &Game{
		cfg:     cfg,
		ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		bullets: make(map[*Bullet]struct{}),
		aliens:  make(map[*Alien]struct{}),
		input:   &Input{},
	}
	g.init()
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.cfg.BgColor)
	g.ship.Draw(screen)
	for bullet := range g.bullets {
		bullet.Draw(screen)
	}
	for alien := range g.aliens {
		alien.Draw(screen)
	}
}

func (g *Game) Update() error {
	g.input.Update(g)
	for bullet := range g.bullets {
		if bullet.outOfScreen() {
			delete(g.bullets, bullet)
		}
		bullet.y -= bullet.speedFactor
	}
	for alien := range g.aliens {
		alien.y += alien.speedFactor
	}
	g.CheckCollision()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

func (g *Game) CreateAliens() {
	alien := NewAlien(g.cfg)
	availableSpaceX := g.cfg.ScreenWidth - 2*alien.width
	numAliens := availableSpaceX / (2 * alien.width)
	for row := 0; row < 2; row++ {
		for i := 0; i < numAliens; i++ {
			alien = NewAlien(g.cfg)
			alien.x = float64(alien.width + 2*alien.width*i)
			alien.y = float64(alien.height*row) * 1.5
			g.addAlien(alien)
		}
	}
}

func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(bullet, alien) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

func CheckCollision(bullet *Bullet, alien *Alien) bool {
	alienTop, alienLeft := alien.y, alien.x
	alienBottom, alienRight := alien.y+float64(alien.height), alien.x+float64(alien.width)
	// 左上角
	x, y := bullet.x, bullet.y
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	// 右上角
	x, y = bullet.x+float64(bullet.width), bullet.y
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	// 左下角
	x, y = bullet.x, bullet.y+float64(bullet.height)
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	// 右下角
	x, y = bullet.x+float64(bullet.width), bullet.y+float64(bullet.height)
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}
	return false
}

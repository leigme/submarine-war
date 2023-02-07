package submarine_war

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x      float64
	y      float64
}

func NewShip(screenWidth, screenHeight int) *Ship {
	imgPath := "C:\\Users\\leig\\Developer\\github\\submarine-war\\20230207\\submarine-war\\images\\ship.png"
	img, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}
	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}

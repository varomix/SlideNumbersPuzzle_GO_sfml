package main

import (
	"runtime"
	"time"

	sf "github.com/manyminds/gosfml"
	"github.com/varomix/SlideNumbersPuzzle_GO_sfml/mix"
)

func init() {
	runtime.LockOSThread()
}

type MainMenu struct {
	rect *sf.RectangleShape
	text *sf.Text
	font *sf.Font
}

func main() {
	ticker := time.NewTicker(time.Second / 60)

	var (
		gameWidth  uint = 544
		gameHeight uint = 650
	)

	// main window
	renderWindow := sf.NewRenderWindow(sf.VideoMode{gameWidth, gameHeight, 32}, "Slide Numbers Puzzle", sf.StyleClose, sf.DefaultContextSettings())

	// bg texture
	bgTex, _ := sf.NewTextureFromFile("assets/images/bg.png", nil)
	bg, _ := sf.NewSprite(bgTex)

	// Load font
	font, _ := sf.NewFontFromFile("assets/fonts/LuckiestGuy.ttf")

	// Title
	testTxt, _ := sf.NewText(font)
	testTxt.SetString("Test text")
	testTxt.SetCharacterSize(36)

	//btn, _ := sf.NewRectangleShape()
	//btn.SetSize(sf.Vector2f{64, 32})
	//btn.SetOutlineThickness(3)
	//btn.SetOutlineColor(sf.ColorBlack())
	//btn.SetFillColor(sf.ColorGreen())
	//btn.Move(sf.Vector2f{200, 300})

	btn2 := mix.NewButton()
	btn2.SetText("1")
	btn2.SetSize(128, 128)
	btn2.Move(300, 15)
	btn2.SetTextSize(48)

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			// poll events
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch ev := event.(type) {
				case sf.EventKeyPressed:
					if ev.Code == sf.KeyEscape {
						renderWindow.Close()
					}
				case sf.EventClosed:
					renderWindow.Close()

				default:
					btn2.Events(ev) // send click to button

				}
			}

		}

		// clear window
		renderWindow.Clear(sf.ColorCyan())

		if 200-btn2.GetPos().X > 0 {
			btn2.Move(5, 0)
		} else if -200-btn2.GetPos().X < 0 {
			btn2.Move(-5, 0)
		}

		// Draw here
		renderWindow.Draw(bg, sf.DefaultRenderStates())
		//renderWindow.Draw(btn, sf.DefaultRenderStates())
		renderWindow.Draw(btn2, sf.DefaultRenderStates())
		//renderWindow.Draw(testTxt, sf.DefaultRenderStates())

		// and display it
		renderWindow.Display()

	}

}

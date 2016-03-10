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

func main() {
	ticker := time.NewTicker(time.Second / 60)

	var (
		gameWidth  uint = 544
		gameHeight uint = 650
	)

	// main window
	renderWindow := sf.NewRenderWindow(sf.VideoMode{Width: gameWidth, Height: gameHeight, BitsPerPixel: 32}, "Slide Numbers Puzzle", sf.StyleClose, sf.DefaultContextSettings())

	// bg texture
	bgTex, _ := sf.NewTextureFromFile("assets/images/bg.png", nil)
	bg, _ := sf.NewSprite(bgTex)

	// Load font
	font, _ := sf.NewFontFromFile("assets/fonts/LuckiestGuy.ttf")

	// Title
	testTxt, _ := sf.NewText(font)
	testTxt.SetString("Slide Numbers Puzzle")
	testTxt.SetCharacterSize(36)

	btn1 := mix.NewButton()
	btn1.SetText("1")
	//btn1.SetSize(128, 128)
	btn1.Move(0, 0)
	btn1.SetTextSize(48)
	//btn1.SetTextSize(48)

	btn2 := mix.NewButton()
	btn2.SetText("2")
	// btn2.SetSize(128, 128)
	btn2.Move(0, 128)
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
					// send click to button
					btn1.Events(ev)
					btn2.Events(ev)

				}
			}

			// clear window
			renderWindow.Clear(sf.ColorCyan())

			// mix.Mover(btn1, 0, 0, 5, "left")

			// Draw here
			renderWindow.Draw(bg, sf.DefaultRenderStates())
			//renderWindow.Draw(btn, sf.DefaultRenderStates())
			renderWindow.Draw(btn1, sf.DefaultRenderStates())
			renderWindow.Draw(btn2, sf.DefaultRenderStates())
			renderWindow.Draw(testTxt, sf.DefaultRenderStates())

			// and display it
			renderWindow.Display()

		}

	}

}

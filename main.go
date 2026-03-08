package main

import (
	_ "embed"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
)

//go:embed assets/icon.ico
var trayIcon []byte

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(trayIcon)
	systray.SetTitle("Blink")
	systray.SetTooltip("Protect your eyes")
	mQuit := systray.AddMenuItem("Quit", "Quit Blink")
	// mQuit.SetIcon(exitIcon)

	go func() {
		tickerBreak := time.NewTicker(10 * time.Minute)
		tickerBlink := time.NewTicker(15 * time.Second)
		defer tickerBreak.Stop()
		defer tickerBlink.Stop()

		for {
			select {
			case <-tickerBreak.C:
				_ = beeep.Notify("Blink", "Close your eyes for 5 seconds", trayIcon)
				// drain the blink ticker during the break
				select {
				case <-tickerBlink.C:
				default:
				}
			case <-tickerBlink.C:
				_ = beeep.Notify("Blink", "Slow full blink!", trayIcon)
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

}

func onExit() {
	// clean up here
}

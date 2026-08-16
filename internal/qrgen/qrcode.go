package qrgen

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

func QrGen(uri string) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(uri, config)
}

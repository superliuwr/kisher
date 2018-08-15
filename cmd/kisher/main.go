package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/spin"
	"github.com/superliuwr/kisher/lib/kisher"
	"github.com/urfave/cli"
)

var version = "V1.0.0"

func main() {
	app := cli.NewApp()
	app.Name = "kisher"
	app.Version = version
	app.Author = "Marvin Liu (superliuwr@gmail.com)"
	app.Usage = "AWS Kinesis Data Stream Helper"
	app.Action = func(c *cli.Context) error {
		spin := spin.New("\033[36m %s Working...\033[m")
		spin.Start()
		err := kisher.Run()
		spin.Stop()
		if err != nil {
			return err
		}
		fmt.Println("Done")
		return nil
	}
	_ = app.Run(os.Args)
}

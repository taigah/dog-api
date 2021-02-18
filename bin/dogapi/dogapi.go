package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/taigah/dog-api/pkg/image"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dogapi",
		Usage: "get dog images",
		Commands: []*cli.Command{{
			Name:  "random",
			Usage: "get a random dog image",
			Action: func(cli *cli.Context) error {
				rep := image.NewRepository(http.DefaultClient)
				image, err := rep.GetRandom()
				if err != nil {
					return err
				}
				fmt.Println(image)
				return nil
			},
		}},
	}
	app.Run(os.Args)
}

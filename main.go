package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"syscall"

	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	app := cli.NewApp()
	app.Version = "3.99"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:      "config,c",
			Usage:     "load configuration from `FILE`",
			Value:     "/etc/insights-client/insights-client.conf",
			TakesFile: true,
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "update",
			Usage: "update to latest core",
			Action: func(c *cli.Context) error {
				cfg := defaultConfig()
				return update(cfg)
			},
		},
		{
			Name:  "collect",
			Usage: "run data collection",
			Action: func(c *cli.Context) error {
				return collect(c.Bool("verbose"), c.String("core"))
			},
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "verbose",
					Usage: "increase output",
				},
				&cli.StringFlag{
					Name:      "core,c",
					Usage:     "Use `FILE` when collecting",
					Value:     corePath,
					TakesFile: true,
				},
			},
		},
		{
			Name:  "upload",
			Usage: "upload an archive",
			Action: func(c *cli.Context) error {
				archivePath := c.Args().First()

				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Username: ")
				username, _ := reader.ReadString('\n')

				fmt.Print("Password: ")
				password, _ := terminal.ReadPassword(int(syscall.Stdin))
				fmt.Print("\n")

				if archivePath == "" {
					return fmt.Errorf("error: path to archive is required")
				}
				cfg := defaultConfig()
				cfg.Username = username
				cfg.Password = string(password)
				return upload(cfg, archivePath)
			},
		},
		{
			Name:  "show",
			Usage: "show insights about this machine",
			Action: func(c *cli.Context) error {
				cfg := defaultConfig()
				return show(cfg)
			},
		},
		{
			Name: "status",
		},
		{
			Name: "unregister",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

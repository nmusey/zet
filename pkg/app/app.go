package app

import (
	"fmt"

	"github.com/nmusey/zet/pkg/notes"
	"github.com/urfave/cli"
)

func BuildApp() *cli.App {
	app := cli.NewApp()

	app.Name = "zet"
	app.Usage = "Manage notes easily from anywhere."
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		buildWriteCommand(),
		buildReadCommand(),
		buildListCommand(),
		buildDeleteCommand(),
	}

	app.CustomAppHelpTemplate = `
{{.Name}} - {{.Usage}}

Ensure the ZET_NOTES_DIR is set to the directory where your notes are stored.

-----

USAGE:
{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
{{if len .Authors}}
AUTHOR:
{{range .Authors}}{{ . }}{{end}}
{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
{{range .VisibleFlags}}{{.}}
{{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
{{.Copyright}}
{{end}}{{if .Version}}
VERSION:
{{.Version}}
{{end}}`

	return app
}

func buildWriteCommand() cli.Command {
	return cli.Command{
		Name:    "write",
		Aliases: []string{"new", "n", "w"},
		Usage:   "Create a new note or append to an existing one",
		Action: func(ctx *cli.Context) error {
			filename := ctx.Args().Get(0)
			if filename == "" {
				return fmt.Errorf("Please provide a filename")
			}

			contents := ctx.Args().Get(1)
			err := notes.WriteNote(filename, contents)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func buildReadCommand() cli.Command {
	return cli.Command{
		Name:    "read",
		Aliases: []string{"r"},
		Usage:   "Read a note",
		Action: func(ctx *cli.Context) error {
			filename := ctx.Args().Get(0)
			if filename == "" {
				return fmt.Errorf("Please provide a filename")
			}

			contents, err := notes.ReadNote(filename)
			if err != nil {
				return err
			}

			fmt.Println(contents)
			return nil
		},
	}
}

func buildListCommand() cli.Command {
	return cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List files in the notes directory",
		Action: func(ctx *cli.Context) error {
			return notes.ListNotes()
		},
	}
}

func buildDeleteCommand() cli.Command {
	return cli.Command{
		Name:    "delete",
		Aliases: []string{"d", "remove", "rm"},
		Usage:   "Delete a note in the notes directory",
		Action: func(ctx *cli.Context) error {
			filename := ctx.Args().Get(0)
			if filename == "" {
				return fmt.Errorf("Please provide a filename")
			}

			return notes.DeleteNote(filename)
		},
	}
}

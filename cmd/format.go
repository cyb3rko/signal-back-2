package cmd

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cyb3rko/signal-back-2/signal"
	"github.com/cyb3rko/signal-back-2/types"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// Format fulfils the `format` subcommand.
var Format = cli.Command{
	Name:               "format",
	Usage:              "Read and format the backup file",
	UsageText:          "Parse and transform the backup file into other formats.\nValid formats include: CSV, RAW.",
	CustomHelpTemplate: SubcommandHelp,
	Flags: append([]cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Usage: "output the backup as `FORMAT`",
			Value: "csv",
		},
		cli.StringFlag{
			Name:  "message, m",
			Usage: "format `TYPE` messages",
			Value: "message",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "write decrypted format to `FILE`",
		},
	}, coreFlags...),
	Action: func(c *cli.Context) error {
		bf, err := setup(c)
		if err != nil {
			return err
		}

		var out io.Writer
		if c.String("output") != "" {
			var file *os.File
			file, err = os.OpenFile(c.String("output"), os.O_CREATE|os.O_WRONLY, 0644)
			out = io.Writer(file)
			if err != nil {
				return errors.Wrap(err, "unable to open output file")
			}
			defer func() {
				if file.Close() != nil {
					log.Fatalf("unable to close output file: %s", err.Error())
				}
			}()
		} else {
			out = os.Stdout
		}

		switch strings.ToLower(c.String("format")) {
		case "csv":
			err = CSV(bf, strings.ToLower(c.String("message")), out)
		case "raw":
			err = Raw(bf, out)
		default:
			return errors.Errorf("format %s not recognised", c.String("format"))
		}
		if err != nil {
			return errors.Wrap(err, "failed to format output")
		}

		return nil
	},
}

// CSV dumps the raw backup data into a comma-separated value format.
func CSV(bf *types.BackupFile, message string, out io.Writer) error {
	ss := make([][]string, 0)

	fns := types.ConsumeFuncs{
		StatementFunc: func(s *signal.SqlStatement) error {
			if strings.HasPrefix(*s.Statement, "INSERT INTO "+message) {
				ss = append(ss, types.StatementToStringArray(s))
			}
			return nil
		},
	}

	if err := bf.Consume(fns); err != nil {
		return err
	}

	w := csv.NewWriter(out)
	var headers []string
	switch strings.ToLower(message) {
	case "message":
		headers = types.MessageCSVHeaders
	case "recipient":
		headers = types.RecipientCSVHeaders
	case "thread":
		headers = types.ThreadCSVHeaders
	case "call":
		headers = types.CallCSVHeaders
	case "groups":
		headers = types.GroupsCSVHeaders
	case "group_membership":
		headers = types.GroupMembershipCSVHeaders
	case "group_receipts":
		headers = types.GroupReceiptsCSVHeaders
	default:
		headers[0] = "Headers unknown"
	}

	if err := w.Write(headers); err != nil {
		return errors.Wrap(err, "unable to write CSV headers")
	}

	for _, sms := range ss {
		if err := w.Write(sms); err != nil {
			return errors.Wrap(err, "unable to format CSV")
		}
	}

	w.Flush()

	return errors.WithMessage(w.Error(), "unable to end CSV writer or something")
}

// Raw performs an ever plainer dump than CSV, and is largely unusable for any purpose outside
// debugging.
func Raw(bf *types.BackupFile, out io.Writer) error {
	fns := types.ConsumeFuncs{
		StatementFunc: func(s *signal.SqlStatement) error {
			_, err := out.Write(append([]byte(s.String()), '\n'))
			return err
		},
	}

	return errors.WithMessage(bf.Consume(fns), "failed to write raw")
}

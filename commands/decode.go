package commands

import (
	"errors"
	"github.com/urfave/cli/v2"
)

var Decode = &cli.Command{
	Name:    "decode",
	Aliases: []string{"d"},
	Usage:   "decode hashed ID",
	UsageText: `hashid d [hashedIDs...]

hashid decode 25yodre1
hashid decode 25yodre1 re7og60y
`,
	ArgsUsage: `decode [hashedIDs...]`,
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return errors.New("none hash id present")
		}

		for _, v := range c.Args().Slice() {
			id, err := decode(v)
			if err != nil {
				return err
			}

			print(id)
		}
		return nil
	},
}

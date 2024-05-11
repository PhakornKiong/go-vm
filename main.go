package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/PhakornKiong/go-vm/compiler"
	"github.com/PhakornKiong/go-vm/vm"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "Input file containing bytecode instructions",
			},
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Usage:   "Type of the return value (int64, uint64, string)",
			},
		},
		Action: func(c *cli.Context) error {
			inputFilePath := c.String("input")
			if inputFilePath == "" {
				return errors.New("input file is required")
			}

			fileContent, err := os.ReadFile(inputFilePath)
			if err != nil {
				return err
			}

			a := vm.NewVM()
			compiler := compiler.NewCompiler()
			compiler.Compile(string(fileContent))
			a.LoadBytecode(compiler.Output())
			res, err := a.Execute()
			if err != nil {
				return err
			}

			// Handle representation of the data type
			retType := c.String("type")
			switch retType {
			case "int64":
				var val int64
				if len(res) > 8 {
					return fmt.Errorf("invalid result length: %d, expected between 0 and 8", len(res))
				}

				if len(res) == 0 {
					val = 0
				} else {
					padded := make([]byte, 8)
					copy(padded[8-len(res):], res)

					// Sign extension for negative numbers
					// Check if MSB is 1 for negative number
					if res[0]&0x80 == 0x80 {
						for i := 0; i < 8-len(res); i++ {
							padded[i] = 0xFF
						}
					}

					val = int64(binary.BigEndian.Uint64(padded))
				}

				fmt.Println("Result in int64 format")
				fmt.Println(val)
			case "uint64":
				var val uint64
				if len(res) > 8 {
					return fmt.Errorf("invalid result length: %d, expected between 0 and 8", len(res))
				}

				if len(res) == 0 {
					val = 0
				} else {
					padded := make([]byte, 8)
					copy(padded[8-len(res):], res)
					val = binary.BigEndian.Uint64(padded)
				}

				fmt.Println("Result in uint64 format")
				fmt.Println(val)
			case "string":
				fmt.Println("Result in string format")
				fmt.Println(string(res))
			default:
				fmt.Println("Result in byte array format")
				fmt.Println(res)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

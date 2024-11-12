package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Config struct {
	InputFile   string
	Delimiter   string
	FieldsIndex []int /*I don't see another reason to increase not processing that huge of a file anyway*/
	Scanner     *bufio.Scanner
}

var (
	config = Config{
		InputFile:   "", /*in case of empty should be stdin*/
		Delimiter:   "\t",
		FieldsIndex: []int{},
	}

	// split on whitespace or comma
	fieldRgx = regexp.MustCompile(`[\s,]+`)
)

func cccut() {
	var scanner *bufio.Scanner

	if config.InputFile == "" || config.InputFile == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(config.InputFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		// The reason I'm keeping the scanner here is for the close func
		// I can also pass a closer
		defer func() {
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		line := scanner.Text()

		var values []string

		fields := strings.Split(line, config.Delimiter)
		for _, fieldsIndex := range config.FieldsIndex {
			if fieldsIndex-1 < len(fields) {
				values = append(values, fields[fieldsIndex-1])
			}
		}

		fmt.Println(strings.Join(values, config.Delimiter))
	}
}

func handleFieldArg(cmd *cobra.Command) error {
	fieldValue, err := cmd.Flags().GetString("field")
	if err != nil {
		return err
	}

	// to enable the following -f5,6,4 or -f"6 5,4"
	fieldValues := fieldRgx.Split(fieldValue, -1)
	var fieldsList []int
	for _, item := range fieldValues {
		field, err := strconv.Atoi(item)
		if err != nil {
			return err
		}

		if field < 1 {
			return fmt.Errorf("field number should be greater than 0")
		}

		fieldsList = append(fieldsList, field)
	}

	// ensure that the output is consistent
	slices.Sort(fieldsList)

	config.FieldsIndex = fieldsList
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "cccut",
	Short: "cccut cuts out the selected portions from each line in a file.",
	Long:  `cccut cuts out the selected portions from each line in a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			config.InputFile = args[0]
		}

		var err error
		delimValue, err := cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatal(err)
			return
		}

		if delimValue != "" {
			config.Delimiter = delimValue
		}

		if err = handleFieldArg(cmd); err != nil {
			log.Fatal(err)
			return
		}

		cccut()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("delimiter", "d", "", "delimiter")
	rootCmd.PersistentFlags().StringP("field", "f", "", "field number separated by delimiter")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

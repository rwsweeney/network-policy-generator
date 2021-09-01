package cmd

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generate)

}

type Inventory struct {
	Name      string
	Namespace string
	Port      uint
}

var generate = &cobra.Command{
	Use:   "generate",
	Short: "Generates a network policy",
	Long:  `Generates a network policy`,

	Run: func(cmd *cobra.Command, args []string) {
		netpol_values := Inventory{"Rachel's Cool network policy", "kube-system-lol", 80}

		// I think this line defines the template file at file.txt
		tmpl, err := template.ParseFiles("file.txt")

		if err != nil {
			panic(err)
		}

		// and this line executes the tmp template with the values stored in sweater and prints it to Stdout
		err = tmpl.Execute(os.Stdout, netpol_values)
		if err != nil {
			panic(err)
		}
	},
}

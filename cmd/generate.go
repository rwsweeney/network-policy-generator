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
	Name      						string
	Namespace 						string
	SpecMatchLabelRole				string
	IngressCidr						string 
	IngressIpBlockExcept			string
	// Should the below be called Project?
	IngressNamespaceMatchLabel		string
	// Should the below be called Role?
	IngressPodSelectorMatchLabel	string
	IngressPort    					string
	EgressCidr						string
	EgressPort						string

}

var generate = &cobra.Command{
	Use:   "generate",
	Short: "Generates a network policy",
	Long:  `Generates a network policy`,

	Run: func(cmd *cobra.Command, args []string) {
		netpol_values := Inventory{"Rachel's Cool network policy", "kube-system-lol", "web", "192.168.0.1/24", "192.168.1.1/32", "myProject", "frontend", "80", "172.168.0.1/24", "8080"}

		// I think this line defines the template file at netpol-template.yaml
		tmpl, err := template.ParseFiles("netpol-template.yaml")

		if err != nil {
			panic(err)
		}

		// and this line executes the tmp template with the values stored in netpol_values and prints it to Stdout
		err = tmpl.Execute(os.Stdout, netpol_values)
		if err != nil {
			panic(err)
		}
	},
}

package cmd

import (
	"context"
	"lambda-runtime-checker/pkg/config"
	"lambda-runtime-checker/pkg/table"
	"lambda-runtime-checker/pkg/utils"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/spf13/cobra"
)

var runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		table := table.SetTable()
		client := lambda.NewFromConfig(cfg)
		resp, err := client.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{
			MaxItems: aws.Int32(100),
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, v := range resp.Functions {
			if utils.Contains(args, string(*&v.Runtime)) == true {
				table.Append([]string{*v.FunctionName, string(*&v.Runtime)})
			}
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(runtimeCmd)
}

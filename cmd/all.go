package cmd

import (
	"context"
	"lambda-runtime-checker/pkg/config"
	"lambda-runtime-checker/pkg/table"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Show runtimes for all Lambda functions",
	Long:  `Show runtimes for all Lambda functions.`,
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
			table.Append([]string{*v.FunctionName, string(*&v.Runtime)})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}

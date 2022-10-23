package cmd

import (
	"context"
	"fmt"
	"lambda-runtime-checker/pkg/config"
	"lambda-runtime-checker/pkg/table"
	"lambda-runtime-checker/pkg/utils"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/spf13/cobra"
)

// deprecatedCmd represents the deprecated command
var deprecatedCmd = &cobra.Command{
	Use:   "deprecated",
	Short: "Search for Lambda functions using deprecated runtimes",
	Long:  `Search for Lambda functions using deprecated runtimes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
		deprecated_runtimes := []string{
			"nodejs",
			"nodejs4.3-edge",
			"nodejs6.10",
			"nodejs8.10",
			"nodejs10.x",
			"nodejs12.x",
			"python2.7",
			"python3.6",
			"dotnetcore1.0",
			"dotnetcore2.0",
			"dotnetcore2.1",
			"ruby2.5",
		}

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
			if utils.Contains(deprecated_runtimes, string(*&v.Runtime)) {
				table.Append([]string{*v.FunctionName, string(*&v.Runtime)})
			}
		}
		fmt.Println("References->https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html")
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(deprecatedCmd)
}

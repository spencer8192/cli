package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Variable to store the output filename flag
var outputFilename string

func init() {
	rootCmd.AddCommand(promCmd)

	// Add a flag to customize the output filename
	// Usage: ./cli prom --output mydata.prom
	promCmd.Flags().StringVarP(&outputFilename, "output", "o", "metrics.prom", "The filename to generate")
}

var promCmd = &cobra.Command{
	Use:   "prom",
	Short: "Generate a sample Prometheus metrics file",
	Long:  `Creates a .prom file containing sample data in the standard Prometheus text exposition format.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Sample data in Prometheus Exposition Format
		// Includes Type and Help metadata
		data := `# HELP http_requests_total The total number of HTTP requests.
# TYPE http_requests_total counter
http_requests_total{method="post",code="200"} 1024
http_requests_total{method="get",code="200"} 500
http_requests_total{method="post",code="400"} 5

# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 145.2

# HELP node_memory_usage_bytes Current memory usage in bytes.
# TYPE node_memory_usage_bytes gauge
node_memory_usage_bytes{type="ram"} 4294967296
node_memory_usage_bytes{type="swap"} 1073741824
`

		// Write the data to the file
		err := os.WriteFile(outputFilename, []byte(data), 0644)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully generated Prometheus metrics in: %s\n", outputFilename)
	},
}

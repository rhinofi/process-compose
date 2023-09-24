package cmd

import (
	"github.com/f1bonacc1/process-compose/src/api"
	"github.com/spf13/cobra"
)

// runCmd represents the up command
var runCmd = &cobra.Command{
	Use:   "run [PROCESS]",
	Short: "Run PROCESS in the foreground, and it's dependencies in the background",
	Long: `Run selected process with std(in|out|err) attached, while other processes run in the background.
Additional command line arguments are passed to the PROCESS.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		*pcFlags.Headless = false
		processName := args[0]

		runner := getProjectRunner(args, *pcFlags.NoDependencies, processName)
		api.StartHttpServer(false, *pcFlags.PortNum, runner)
		runProject(runner)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(pcFlags.NoDependencies, "no-deps", "", *pcFlags.NoDependencies, "don't start dependent processes")
	runCmd.Flags().AddFlag(rootCmd.Flags().Lookup("config"))

}

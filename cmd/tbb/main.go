package main

import (
	"fmt"
	"os"

	"github.com/erikrios/my-blockchain-bar/fs"
	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"
const flagPort = "port"

func main() {
	var tbbCmd = &cobra.Command{
		Use:   "tbb",
		Short: "The Blockchain Bar CLI",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	tbbCmd.AddCommand(migrateCmd())
	tbbCmd.AddCommand(versionCmd)
	tbbCmd.AddCommand(balancesCmd())
	tbbCmd.AddCommand(txCmd())
	tbbCmd.AddCommand(runCmd())

	err := tbbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)
	return fs.ExpandPath(dataDir)
}

func incorrentUsageErr() error {
	return fmt.Errorf("incorrect usage")
}

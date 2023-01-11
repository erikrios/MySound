package main

import (
	"fmt"
	"os"
	"time"

	"github.com/erikrios/my-blockchain-bar/database"
	"github.com/spf13/cobra"
)

var migrateCmd = func() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrates the blockchain database according to new business rules.",
		Run: func(cmd *cobra.Command, args []string) {

			state, err := database.NewStateFromDisk(getDataDirFromCmd(cmd))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			block0 := database.NewBlock(
				database.Hash{},
				uint64(time.Now().Unix()),
				0,
				[]database.Tx{
					database.NewTx("erikrios", "erikrios", 3, ""),
					database.NewTx("erikrios", "erikrios", 700, "reward"),
				},
			)

			state.AddBlock(block0)
			block0Hash, _ := state.Persist()

			block1 := database.NewBlock(
				block0Hash,
				1,
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("erikrios", "babayaga", 2000, ""),
					database.NewTx("erikrios", "erikrios", 100, "reward"),
					database.NewTx("babayaga", "erikrios", 1, ""),
					database.NewTx("babayaga", "caesar", 1000, ""),
					database.NewTx("babayaga", "erikrios", 50, ""),
					database.NewTx("erikrios", "erikrios", 600, "reward"),
				},
			)

			state.AddBlock(block1)
			state.Persist()
		},
	}

	addDefaultRequiredFlags(migrateCmd)

	return migrateCmd
}

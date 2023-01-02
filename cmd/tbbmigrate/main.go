package main

import (
	"fmt"
	"os"
	"time"

	"github.com/erikrios/my-blockchain-bar/database"
)

func main() {
	cwd, _ := os.Getwd()
	state, err := database.NewStateFromDisk(cwd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("erikrios", "erikrios", 3, ""),
			database.NewTx("erikrios", "erikrios", 700, "reward"),
		},
	)

	state.AddBlock(block0)
	block0Hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0Hash,
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
}

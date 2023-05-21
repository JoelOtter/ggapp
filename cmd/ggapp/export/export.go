package export

import (
	"encoding/csv"
	"fmt"
	"github.com/JoelOtter/ggapp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func NewCommand(logger *logrus.Logger) *cobra.Command {
	return &cobra.Command{
		Use:  "export user-id",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.WithField("command", "export")
			userId, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("failed to parse user ID %s: %w", args[0], err)
			}
			log.Infof("Exporting games for user %d", userId)
			games, err := ggapp.ListGamesForStatuses(log, ggapp.ListGamesForStatusesVariables{
				StatusIds: ggapp.AllPlayStatuses,
				UserId:    userId,
			})
			if err != nil {
				return fmt.Errorf("failed to fetch games: %w", err)
			}
			log.Infof("Fetched %d games", len(games))

			writer := csv.NewWriter(os.Stdout)
			if err := writer.Write([]string{"Title", "Status"}); err != nil {
				return fmt.Errorf("failed to write CSV header: %w", err)
			}
			for _, game := range games {
				if err := writer.Write([]string{game.Game.Name, game.PlayStatus.Title}); err != nil {
					return fmt.Errorf("failed to write CSV entry: %w", err)
				}
			}
			writer.Flush()
			if err := writer.Error(); err != nil {
				return fmt.Errorf("CSV writer failed: %w", err)
			}
			return nil
		},
	}
}

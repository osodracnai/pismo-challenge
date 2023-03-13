package cmd

import (
	"fmt"
	"github.com/osodracnai/pismo-challenge/cmd/config"
	"github.com/osodracnai/pismo-challenge/pkg/database"
	"github.com/osodracnai/pismo-challenge/pkg/server"
	"github.com/osodracnai/pismo-challenge/pkg/server/accounts"
	"github.com/osodracnai/pismo-challenge/pkg/server/transactions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
)

func ServerCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "server",
		Short:         "Start the server process",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			var configCmd config.Global
			if err = viper.Unmarshal(&configCmd, DecoderConfigOptions); err != nil {
				return fmt.Errorf("parse config: %v", err)
			}
			cassandra, err := database.NewCassandra(configCmd.Cassandra)
			if err != nil {
				return fmt.Errorf("creating cassandra: %v", err)
			}
			closer, err := configCmd.Jaeger.Config()
			if err != nil {
				return fmt.Errorf("creating jaeger: %v", err)
			}
			defer closer.Close()

			acc := accounts.New(cassandra)
			trans := transactions.New(cassandra)
			s, err := server.New(acc, trans)
			if err != nil {
				return err
			}
			listenURL, err := url.Parse(configCmd.Server.Listen)
			if err != nil {
				logrus.Fatal(err)
			}

			r := s.NewEngine(true)
			logrus.Infof("Running server on %s", listenURL.Host)
			err = r.Run(listenURL.Host)
			if err != nil {
				return err
			}
			return nil
		},
	}
	f := command.Flags()
	serverFlags(f)
	cassandraFlags(f)
	jaegerFlags(f)
	return command
}

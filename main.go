package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/app/template"
	"github.com/adrianliechti/devkit/app/utility/cloc"
	"github.com/adrianliechti/devkit/app/utility/code"
	"github.com/adrianliechti/devkit/app/utility/git"
	"github.com/adrianliechti/devkit/app/utility/image"
	"github.com/adrianliechti/devkit/app/utility/proxy"
	"github.com/adrianliechti/devkit/app/utility/server"

	"github.com/adrianliechti/devkit/pkg/catalog/activemq"
	"github.com/adrianliechti/devkit/pkg/catalog/azurite"
	"github.com/adrianliechti/devkit/pkg/catalog/cassandra"
	"github.com/adrianliechti/devkit/pkg/catalog/cockroachdb"
	"github.com/adrianliechti/devkit/pkg/catalog/db2"
	"github.com/adrianliechti/devkit/pkg/catalog/directus"
	"github.com/adrianliechti/devkit/pkg/catalog/elasticsearch"
	"github.com/adrianliechti/devkit/pkg/catalog/etcd"
	"github.com/adrianliechti/devkit/pkg/catalog/ghost"
	"github.com/adrianliechti/devkit/pkg/catalog/grafana"
	"github.com/adrianliechti/devkit/pkg/catalog/immudb"
	"github.com/adrianliechti/devkit/pkg/catalog/influxdb"
	"github.com/adrianliechti/devkit/pkg/catalog/jaeger"
	"github.com/adrianliechti/devkit/pkg/catalog/jenkins"
	"github.com/adrianliechti/devkit/pkg/catalog/jupyter"
	"github.com/adrianliechti/devkit/pkg/catalog/kafka"
	"github.com/adrianliechti/devkit/pkg/catalog/mailtrap"
	"github.com/adrianliechti/devkit/pkg/catalog/mariadb"
	"github.com/adrianliechti/devkit/pkg/catalog/minio"
	"github.com/adrianliechti/devkit/pkg/catalog/mongodb"
	"github.com/adrianliechti/devkit/pkg/catalog/mosquitto"
	"github.com/adrianliechti/devkit/pkg/catalog/mssql"
	"github.com/adrianliechti/devkit/pkg/catalog/mysql"
	"github.com/adrianliechti/devkit/pkg/catalog/nats"
	"github.com/adrianliechti/devkit/pkg/catalog/nexus"
	"github.com/adrianliechti/devkit/pkg/catalog/oracle"
	"github.com/adrianliechti/devkit/pkg/catalog/postgres"
	"github.com/adrianliechti/devkit/pkg/catalog/prometheus"
	"github.com/adrianliechti/devkit/pkg/catalog/rabbitmq"
	"github.com/adrianliechti/devkit/pkg/catalog/ravendb"
	"github.com/adrianliechti/devkit/pkg/catalog/redis"
	"github.com/adrianliechti/devkit/pkg/catalog/sonarqube"
	"github.com/adrianliechti/devkit/pkg/catalog/tooljet"
	"github.com/adrianliechti/devkit/pkg/catalog/unleash"
	"github.com/adrianliechti/devkit/pkg/catalog/vault"

	"github.com/adrianliechti/devkit/pkg/cli"
)

var version string

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	app := initApp()

	if err := app.RunContext(ctx, os.Args); err != nil {
		cli.Fatal(err)
	}
}

func initApp() cli.App {
	return cli.App{
		Usage: "DevKit",

		Suggest: true,
		Version: version,

		HideHelpCommand: true,

		Commands: []*cli.Command{
			catalog.Command(&activemq.Manager{}),
			catalog.Command(&azurite.Manager{}),
			catalog.Command(&cassandra.Manager{}),
			catalog.Command(&cockroachdb.Manager{}),
			catalog.Command(&db2.Manager{}),
			catalog.Command(&directus.Manager{}),
			catalog.Command(&elasticsearch.Manager{}),
			catalog.Command(&etcd.Manager{}),
			catalog.Command(&ghost.Manager{}),
			catalog.Command(&grafana.Manager{}),
			catalog.Command(&immudb.Manager{}),
			catalog.Command(&influxdb.Manager{}),
			catalog.Command(&jaeger.Manager{}),
			catalog.Command(&jenkins.Manager{}),
			catalog.Command(&jupyter.Manager{}),
			catalog.Command(&kafka.Manager{}),
			catalog.Command(&mailtrap.Manager{}),
			catalog.Command(&mariadb.Manager{}),
			catalog.Command(&minio.Manager{}),
			catalog.Command(&mongodb.Manager{}),
			catalog.Command(&mosquitto.Manager{}),
			catalog.Command(&mssql.Manager{}),
			catalog.Command(&mysql.Manager{}),
			catalog.Command(&nats.Manager{}),
			catalog.Command(&nexus.Manager{}),
			catalog.Command(&oracle.Manager{}),
			catalog.Command(&postgres.Manager{}),
			catalog.Command(&prometheus.Manager{}),
			catalog.Command(&rabbitmq.Manager{}),
			catalog.Command(&ravendb.Manager{}),
			catalog.Command(&redis.Manager{}),
			catalog.Command(&sonarqube.Manager{}),
			catalog.Command(&tooljet.Manager{}),
			catalog.Command(&unleash.Manager{}),
			catalog.Command(&vault.Manager{}),

			template.Command,

			git.Command,
			image.Command,

			cloc.Command,

			code.Command,
			server.Command,
			proxy.Command,
		},
	}
}

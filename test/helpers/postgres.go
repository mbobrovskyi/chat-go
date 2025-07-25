// Copyright 2025 Mykhailo Bobrovskyi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/testcontainers/testcontainers-go"
	testcontainerspostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"chat-go/internal/infrastructure/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	PostgresImage = "postgres:17-alpine"
	DbName        = "test-db"
	DbUser        = "user"
	DbPass        = "password"
)

func PostgresURIForContainer(ctx context.Context, postgresContainer *testcontainerspostgres.PostgresContainer) (string, error) {
	port, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return "", err
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		return "", err
	}

	return BuildPostgresURI(host, port.Port()), nil
}

func Migrate(ctx context.Context, postgresContainer *testcontainerspostgres.PostgresContainer) error {
	postgresURL, err := PostgresURIForContainer(ctx, postgresContainer)
	if err != nil {
		return err
	}

	db, err := postgres.NewPostgres(ctx, postgresURL)
	if err != nil {
		return err
	}
	defer db.Close()

	driver, err := migratepostgres.WithInstance(db, &migratepostgres.Config{})
	if err != nil {
		return err
	}

	migrationsDir, err := GetMigrationsDir()
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationsDir), DbName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}

func RunPostgresContainer(ctx context.Context) (*testcontainerspostgres.PostgresContainer, error) {
	return testcontainerspostgres.Run(
		ctx,
		PostgresImage,
		testcontainerspostgres.WithDatabase(DbName),
		testcontainerspostgres.WithUsername(DbUser),
		testcontainerspostgres.WithPassword(DbPass),
		testcontainers.WithWaitStrategy(wait.ForListeningPort("5432/tcp")),
	)
}

func BuildPostgresURI(host, port string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", DbUser, DbPass, host, port, DbName)
}

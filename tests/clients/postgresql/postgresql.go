//go:build integrations

package postgresql

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/sariya23/manage_pr_service/internal/config"
	"github.com/sariya23/manage_pr_service/internal/storage/database"
)

type TestDB struct {
	DB *database.Database
}

func NewTestDB() *TestDB {
	cfg := config.MustLoadByPath(filepath.Join("..", "..", "..", "..", "config", "test.env"))
	DB, err := database.NewConnection(
		context.Background(),
		database.GenerateDBUrl(
			cfg.PostgresUsername,
			cfg.PostgresPassword,
			cfg.PostgresOuterHost,
			strconv.Itoa(cfg.PostgresPort),
			cfg.PostgresDB,
			cfg.SSLMode,
		),
	)
	if err != nil {
		panic(err)
	}
	return &TestDB{DB: DB}
}

func (d *TestDB) SetUp(ctx context.Context, t *testing.T, tablenames ...string) {
	t.Helper()
	d.Truncate(ctx, tablenames...)
}

func (d *TestDB) TearDown(t *testing.T) {
	t.Helper()
}

func (d *TestDB) Truncate(ctx context.Context, tables ...string) {
	q := fmt.Sprintf("truncate %s cascade", strings.Join(tables, ","))
	if _, err := d.DB.GetPool().Exec(ctx, q); err != nil {
		panic(err)
	}
}

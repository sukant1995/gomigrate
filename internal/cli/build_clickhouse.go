//go:build clickhouse
// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/sukant1995/gomigrate/database/clickhouse"
)

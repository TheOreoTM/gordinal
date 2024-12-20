package friemon

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/disgoorg/snowflake/v2"
	"github.com/pelletier/go-toml/v2"
	"github.com/theoreotm/friemon/friemon/db"
)

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}

	var cfg Config
	if err = toml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Config struct {
	Log      LogConfig `toml:"log"`
	Bot      BotConfig `toml:"bot"`
	Database db.Config `toml:"database"`
}

type BotConfig struct {
	DevGuilds    []snowflake.ID `toml:"dev_guilds"`
	Token        string         `toml:"token"`
	SyncCommands bool           `toml:"sync_commands"`
	DevMode      bool           `toml:"dev_mode"`
	Version      string         `toml:"version"`
}

type LogConfig struct {
	Level     slog.Level `toml:"level"`
	Format    string     `toml:"format"`
	AddSource bool       `toml:"add_source"`
}

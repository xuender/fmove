package move

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/expr-lang/expr"
	"github.com/pelletier/go-toml/v2"
	"github.com/xuender/fmeta/meta"
	"github.com/xuender/kit/oss"
)

const _minDatetime = 18

func Move(path string, simulate bool) error {
	cfg, err := readCfg(path)
	if err != nil {
		return err
	}

	slog.Debug("move", "path", path, "cfg", cfg.GetTarget())

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Name() == "fm.toml" {
			continue
		}

		moveEntry(path, entry.Name(), cfg.GetTarget(), simulate)
	}

	return nil
}

func moveEntry(path, entry string, targets []*Target, simulate bool) {
	file := filepath.Join(path, entry)

	fileMeta, err := meta.FileMeta(file)
	if err != nil {
		slog.Error("FileMeta", "err", err)

		return
	}

	env := metaEnv(fileMeta)

	for _, target := range targets {
		targetPath, err := oss.Abs(target.GetPath())
		if err != nil {
			slog.Error("path", "path", target.GetPath(), "err", err)

			continue
		}

		if check(env, target) {
			if target.GetSubDir() == "" {
				targetPath = filepath.Join(targetPath, entry)
			} else {
				targetPath = filepath.Join(targetPath, subDir(env, target), entry)
			}

			if simulate {
				slog.Info("move", "name", file, "target", targetPath)

				continue
			}

			move(file, targetPath)
		}
	}
}

func readCfg(path string) (*Config, error) {
	file, err := os.Open(filepath.Join(path, "fm.toml"))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	cfg := &Config{}
	if err := toml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func move(file, target string) {
	dir := filepath.Dir(target)
	if err := os.MkdirAll(dir, oss.DefaultDirFileMod); err != nil {
		slog.Error("mkdir", "path", dir, "err", err)

		return
	}

	if err := os.Rename(file, target); err == nil {
		slog.Info("move", "name", file, "target", target)
	} else {
		slog.Error("move", "err", err)
	}
}

func check(env map[string]any, target *Target) bool {
	return run[bool](env, target.GetCondition())
}

func subDir(env map[string]any, target *Target) string {
	return run[string](env, target.GetSubDir())
}

func metaEnv(env *meta.Meta) map[string]any {
	values := map[string]any{}

	for key, val := range meta.MetaType_value {
		values[strings.ToLower(key)] = meta.MetaType(val)
	}

	values["type"] = env.GetType()
	values["sub"] = env.GetSubtype()
	values["subtype"] = env.GetSubtype()
	values["ext"] = env.GetExtension()
	values["extension"] = env.GetExtension()
	values["size"] = env.GetSize()
	values["width"] = env.GetWidth()
	values["height"] = env.GetHeight()
	values["dur"] = env.GetDuration()
	values["duration"] = env.GetDuration()
	values["chan"] = env.GetChannels()
	values["channels"] = env.GetChannels()

	if len(env.GetDatetime()) > _minDatetime {
		if date, err := time.Parse(time.DateTime, env.GetDatetime()); err == nil {
			values["yyyy"] = date.Format("2006")
			values["yy"] = date.Format("06")
			values["MM"] = date.Format("01")
			values["M"] = date.Format("1")
			values["dd"] = date.Format("02")
			values["d"] = date.Format("2")
			values["HH"] = date.Format("15")
			values["H"] = strings.TrimLeft(date.Format("15"), "0")
			values["hh"] = date.Format("03")
			values["h"] = date.Format("3")
			values["mm"] = date.Format("04")
			values["m"] = date.Format("4")
			values["ss"] = date.Format("05")
			values["s"] = date.Format("5")
		}
	}

	return values
}

func run[T any](env map[string]any, code string) T {
	var zero T

	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		slog.Error("condition", "code", code, "env", env, "err", err)

		return zero
	}

	output, err := expr.Run(program, env)
	if err != nil {
		slog.Error("run", "code", code, "env", env, "err", err)

		return zero
	}

	if ret, ok := output.(T); ok {
		return ret
	}

	return zero
}

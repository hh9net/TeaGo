package Tea

import (
	"os"
	"path/filepath"
	"github.com/iwind/TeaGo/logs"
	"github.com/iwind/TeaGo/utils/string"
	"strings"
)

const (
	EnvDev  = "dev"
	EnvTest = "test"
	EnvProd = "prod"
)

var Env = EnvDev
var DS = string(os.PathSeparator)

var publicDir string
var viewsDir string
var configDir string

var Root string

// 初始化
func init() {
	Root = findRoot()
}

func Is(env ... string) bool {
	if len(env) == 0 {
		return false
	}
	for _, envItem := range env {
		if envItem == Env {
			return true
		}
	}
	return false
}

func TmpDir() string {
	return Root + DS + "tmp"
}

func TmpFile(file string) string {
	return TmpDir() + DS + file
}

func LogDir() string {
	return Root + DS + "logs"
}

func LogFile(file string) string {
	return LogDir() + DS + file
}

func BinDir() string {
	return Root + DS + "bin"
}

func PublicDir() string {
	if len(publicDir) > 0 {
		return publicDir
	}

	publicDir = findLatestDir(Root, "public")
	return publicDir
}

func PublicFile(file string) string {
	return PublicDir() + DS + file
}

func ViewsDir() string {
	if len(viewsDir) > 0 {
		return viewsDir
	}

	viewsDir = findLatestDir(Root, "views")
	return viewsDir
}

func ConfigDir() string {
	if len(configDir) > 0 {
		return configDir
	}

	configDir = findLatestDir(Root, "configs")
	return configDir
}

func ConfigFile(file string) string {
	return ConfigDir() + DS + file
}

func findRoot() string {
	// TEAROOT变量
	root := strings.TrimSpace(os.Getenv("TEAROOT"))
	if len(root) > 0 {
		abs, err := filepath.Abs(root)
		if err != nil {
			logs.Errorf("invalid GOPATH '%s'", root)
			return root
		}
		return abs
	}

	// GOPATH变量
	root = strings.TrimSpace(os.Getenv("GOPATH"))
	if len(root) > 0 {
		abs, err := filepath.Abs(root)
		if err != nil {
			logs.Errorf("invalid GOPATH '%s'", root)
			return root + DS + "src" + DS + "main"
		}
		return abs + DS + "src" + DS + "main"
	}

	// 当前执行的目录
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}
	return "./"
}

func findLatestDir(parent string, name string) string {
	matches, err := filepath.Glob(parent + DS + name + ".*")
	if err != nil {
		logs.Errorf("%s", err.Error())
		return parent + DS + name
	}

	if len(matches) == 0 {
		return parent + DS + name
	}

	var lastVersion = ""
	var resultDir = ""

	for _, match := range matches {
		dirname := match
		stat, err := os.Stat(dirname)
		if err != nil || !stat.IsDir() {
			continue
		}

		version := filepath.Base(match)[len(name)+1:]

		if len(lastVersion) == 0 {
			lastVersion = version
			resultDir = dirname
			continue
		}

		if stringutil.VersionCompare(lastVersion, version) < 0 {
			lastVersion = version
			resultDir = dirname
			continue
		}
	}

	if len(resultDir) == 0 {
		return parent + DS + name
	}

	return resultDir
}
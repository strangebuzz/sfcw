package symfony

import (
	"fmt"
	"github.com/strangebuzz/cc/structs"
	"github.com/strangebuzz/cc/tools"
	"os"
	"path/filepath"
)

/**
 * Get all the files of the Symfony application which have to be watched.
 *
 * @todo handle nesting level? Add a parameter for that?
 */
func getFilesToWatch(config structs.Config) ([]string, error) {
	var filesToWatch []string

	// 1) ".env" files at the root of the project
	filesToWatch = append(filesToWatch, getFilesFromPath(config, ".env*")...)

	// 2) Yaml files in "config/"...
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*.yaml", config.SymfonyConfigDir))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*.yaml", config.SymfonyConfigDir))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*/*.yaml", config.SymfonyConfigDir))...)

	// 3) Yaml files in "translations/"...
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*.yaml", config.SymfonyTranslationsDir))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*.yaml", config.SymfonyTranslationsDir))...)

	// 4) Twig files in "templates/"...
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*.twig", config.SymfonyTemplatesDir))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*.twig", config.SymfonyTemplatesDir))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*/*.twig", config.SymfonyTemplatesDir))...)
	filesToWatch = append(filesToWatch, getFilesFromPath(config, fmt.Sprintf("%s/*/*/*/*.twig", config.SymfonyTemplatesDir))...)

	return filesToWatch, nil
}

/**
 * Get all files corresponding to a glob pattern. To simplify the code we don't
 * return an error. The process will stop if an error is encountered.
 */
func getFilesFromPath(config structs.Config, glob string) []string {
	files, err := filepath.Glob(fmt.Sprintf("%s/%s", config.SymfonyProjectDir, glob))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return files
}

/**
 * Get the watched files as a simple map: "path" => "datetime"
 *
 * @example "/Users/coil/Sites/strangebuzz.com/config/services.yaml": "2020-04-28 16:33:44.73756727 +0200 CEST"
 */
func GetWatchMap(config structs.Config) (map[string]string, error) {
	watchMap := map[string]string{}
	filesTowatch, _ := getFilesToWatch(config)

	for _, file := range filesTowatch {
		stats, err := os.Stat(file)
		if err != nil {
			tools.PrintError(fmt.Errorf("Can't get stats for the \"%s\" file, check the project permissions.", err))
			os.Exit(1)
		}
		watchMap[file] = fmt.Sprintf("%s", stats.ModTime())
	}

	return watchMap, nil
}

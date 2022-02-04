package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/longdoan7421/rename-files/utils"
)

var supportedCaseTypes = []string{"title", "pascal", "camel", "snake", "kebab", "pascal-snake", "pascal-kebab"}
var smallWords = []string{"a", "an", "the", "and", "as", "but", "for", "if", "nor", "or", "so", "yet", "at", "by", "for", "in", "of", "off", "on", "per", "to", "up", "via"}

/* Flag Usages */
var pathUsage = `The path to file or directory which has files need to be renamed. (required)

If path is a directory, every files, even files included in sub directory, will also be renamed. Use "-depth" flag to limit the depth.`
var depthUsage = `If path is a directory, all files which is within the depth will be renamed. Depth is count from 1.`
var caseTypeUsage = `The case type which file will be renamed to. (required)

Support types:
	* title: This is an Example
	* pascal: ThisIsAnExample
	* camel: thisIsAnExample
	* snake: this_is_an_example
	* kebab: this-is-an-example
	* pascal-snake: This_Is_An_Example
	* pascal-kebab: This-Is-An-Example`
var dryRunUsage = `Show filename after renaming, but no files would be renamed.`
var keepUpperUsage = `The original word which is uppercase will be preserved, e.g. "GO is fun" -> "GO-is-fun".`

/* End Flag Usages */

/* Colors */
var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

/* End Colors */

func main() {

	var pathFlag *string = flag.String("path", "", pathUsage)
	var depthFlag *int = flag.Int("depth", 10, depthUsage)
	var caseTypeFlag *string = flag.String("case", "", caseTypeUsage)
	var dryRunFlag *bool = flag.Bool("dry-run", false, dryRunUsage)
	var keepUpperPartFlag *bool = flag.Bool("keep-upper", false, keepUpperUsage)
	flag.Parse()

	validateFlags(*pathFlag, *caseTypeFlag, *depthFlag)

	if *dryRunFlag {
		log.Println("Running in dry mode")
	}

	var err error
	*pathFlag, err = filepath.Abs(*pathFlag)
	if err != nil {
		log.Fatal(err)
	}

	isDir, err := utils.IsDirectory(*pathFlag)
	if err != nil {
		log.Fatal(err)
	}

	if isDir {
		log.Printf(`Reading files in "%s"`, colorYellow+(*pathFlag)+colorReset)
		err := filepath.Walk(*pathFlag,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					if depth := calculateDepthOfChildDirectory(*pathFlag, path); depth > *depthFlag {
						return filepath.SkipDir
					}

					if isHidden, err := utils.IsHiddenFile(info.Name()); isHidden || err != nil {
						return filepath.SkipDir
					}

					return nil
				}

				_, filename := filepath.Split(path)
				if isHidden, err := utils.IsHiddenFile(filename); isHidden || err != nil {
					return nil
				}

				renameFile(path, *caseTypeFlag, *dryRunFlag, *keepUpperPartFlag)
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	} else {
		renameFile(*pathFlag, *caseTypeFlag, *dryRunFlag, *keepUpperPartFlag)
	}
}

func validateFlags(path string, caseType string, depth int) {
	path = strings.TrimSpace(path)
	caseType = strings.TrimSpace(caseType)

	switch {
	case path == "":
		log.Fatal("Invalid path")
	case !utils.Contains(supportedCaseTypes, caseType):
		log.Fatal("Invalid case type")
	case depth < 1:
		log.Fatal("Depth needs to be positive number")
	}
}

func calculateDepthOfChildDirectory(root string, dir string) int {
	rootParts := strings.Split(root, string(os.PathSeparator))
	dirParts := strings.Split(dir, string(os.PathSeparator))
	depth := len(dirParts) - len(rootParts) + 1

	return depth
}

func renameFile(path string, caseType string, isDryRun bool, keepUpper bool) {
	dir, file := filepath.Split(path)
	ext := filepath.Ext(path)
	oldName := strings.TrimSuffix(file, ext)
	nameParts := splitNameParts(oldName, keepUpper)

	var newName string
	// NOTE: Make sure case types are consistent with supported case types above
	switch caseType {
	case "title":
		newName = toTitleCaseString(nameParts)
	case "pascal":
		newName = toPascalCaseString(nameParts)
	case "camel":
		newName = toCamelCaseString(nameParts)
	case "snake":
		newName = toSnakeCaseString(nameParts)
	case "kebab":
		newName = toKebabCaseString(nameParts)
	case "pascal-snake":
		newName = toPascalSnakeCaseString(nameParts)
	case "pascal-kebab":
		newName = toPascalKebabCaseString(nameParts)
	}

	if !isDryRun {
		os.Rename(path, dir+newName+ext)
		log.Printf(`Renamed: "{%s}%s" -> "%s"`, colorYellow+dir+colorReset, colorRed+oldName+ext+colorReset, colorGreen+newName+ext+colorReset)
	} else {
		log.Printf(`Would rename: "{%s}%s" -> "%s"`, colorYellow+dir+colorReset, colorRed+oldName+ext+colorReset, colorGreen+newName+ext+colorReset)
	}
}

func splitNameParts(oldName string, keepUpper bool) []string {
	var chunks []string
	// split path by dash "-"
	chunks = utils.Map(strings.Split(oldName, "-"), func(_ int, part string) string {
		return strings.TrimSpace(part)
	})

	// split path by underscore "_"
	oldName = strings.Join(chunks, "_")
	chunks = utils.Map(strings.Split(oldName, "_"), func(_ int, part string) string {
		return strings.TrimSpace(part)
	})

	// split path by comma ","
	oldName = strings.Join(chunks, ",")
	chunks = utils.Map(strings.Split(oldName, ","), func(_ int, part string) string {
		return strings.TrimSpace(part)
	})

	// NOTE: split by space should always be last
	// split path by space " "
	oldName = strings.Join(chunks, " ")
	chunks = utils.Map(strings.Split(oldName, " "), func(_ int, part string) string {
		return strings.TrimSpace(part)
	})

	// convert all parts to lowercase
	chunks = utils.Map(chunks, func(_ int, part string) string {
		if keepUpper && part == strings.ToUpper(part) {
			return part
		}

		return strings.ToLower(part)
	})

	return chunks
}

func toTitleCaseString(parts []string) string {
	newParts := utils.Map(parts, func(index int, value string) string {
		if index != 0 && utils.Contains(smallWords, value) {
			return value
		} else {
			return strings.Title(value)
		}
	})

	return strings.Join(newParts, " ")
}

func toPascalCaseString(parts []string) string {
	newParts := utils.Map(parts, func(index int, value string) string {
		return strings.ToUpper(value[0:1]) + value[1:]
	})

	return strings.Join(newParts, "")
}

func toCamelCaseString(parts []string) string {
	newParts := utils.Map(parts, func(index int, value string) string {
		if index == 0 {
			return value
		}

		return strings.ToUpper(value[0:1]) + value[1:]
	})

	return strings.Join(newParts, "")
}

func toSnakeCaseString(parts []string) string {
	return strings.Join(parts, "_")
}

func toKebabCaseString(parts []string) string {
	return strings.Join(parts, "-")
}

func toPascalSnakeCaseString(parts []string) string {
	newParts := utils.Map(parts, func(index int, value string) string {
		return strings.ToUpper(value[0:1]) + value[1:]
	})

	return strings.Join(newParts, "_")
}

func toPascalKebabCaseString(parts []string) string {
	newParts := utils.Map(parts, func(index int, value string) string {
		return strings.ToUpper(value[0:1]) + value[1:]
	})

	return strings.Join(newParts, "-")
}

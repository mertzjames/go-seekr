/*
Copyright © 2025 James Mertz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

// TODO: Slated for a later release
// type VERBOSITY_LEVELS int

// const (
// 	QUIET VERBOSITY_LEVELS = iota
// 	DEFAULT
// 	VERBOSE
// 	DEBUG
// )

var FLAG_SCAN_PATH string
var FLAG_INCLUDE_BINARY bool
var FLAG_INCLUDE_ALL bool
var FLAG_LANGUAGE string
var FLAG_USR_VARS string
var FLAG_USR_REGEX string
var FLAG_USR_VARS_FILE string
var FLAG_USR_REGEX_FILE string
var FLAG_IGNORE_DEFAULT_VARS bool

// TODO: Slated for a later release
// var FLAG_VERBOSITY VERBOSITY_LEVELS = DEFAULT
var FLAG_CASE_INSENSITIVE bool

//go:embed variables.txt
var VARIABLES_LIST string

var LANG_EXT = map[string][]string{
	// Top 20 languages according to Google Gemini 2.5 Pro
	"javascript": {".js", ".mjs", ".cjs"},
	"python":     {".py", ".pyw", ".pyi"},
	"java":       {".java", ".class", ".jar"},
	"typescript": {".ts", ".tsx"},
	"csharp":     {".cs", ".csx"},
	"cpp":        {".cpp", ".hpp", ".cc", ".h"},
	"php":        {".php", ".phtml"},
	"go":         {".go"},
	"swift":      {".swift"},
	"ruby":       {".rb", ".rbw"},
	"kotlin":     {".kt", ".kts"},
	"rust":       {".rs"},
	"sql":        {".sql"},
	"r":          {".r", ".R"},
	"perl":       {".pl", ".pm"},
	"lua":        {".lua"},
	"objc":       {".m", ".h"},
	"dart":       {".dart"},
	"scala":      {".scala", ".sc"},
	"matlab":     {".m"},

	// Additional Languages/file types Supported
	"c":          {".c", ".h"},
	"shell":      {".sh", ".bash", ".zsh"},
	"powershell": {".ps1", ".psm1"},
	"yaml":       {".yaml", ".yml"},
	"xml":        {".xml"},
	"json":       {".json"},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seekr",
	Short: "Seekr - Vulnerable Variable Scanner",
	Long: `
███████╗███████╗███████╗██╗  ██╗██████╗
██╔════╝██╔════╝██╔════╝██║ ██╔╝██╔══██╗
███████╗███████╗███████╗█████╔╝ ██████╔╝
╚════██║██╚════║██╚════║██╔═██╗ ██╔══██═╗
███████║███████║███████║██║  ██╗██╚═══██║
╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═══════╝
Seekr - Vulnerable Variable Scanner

a comprehensive security scanner designed to protect your codebase from 
accidentally leaked secrets and sensitive environment variables. Built with Go for cross-platform 
compatibility, it detects over 400 types of potentially dangerous exposures including API keys, 
tokens, passwords, database credentials, and SSH private keys across 20+ programming languages 
and configuration formats.

The tool features intelligent pattern matching with support for custom variables (--vars) and 
user-defined regex patterns (--regex_str), making it adaptable to organization-specific security 
requirements. With language-specific filtering, binary file scanning capabilities, and precise 
line number reporting, go-seekr integrates seamlessly into development workflows, CI/CD pipelines, 
and security audit processes.

Quick Examples:
  ./seekr                                          # Scan current directory
  ./seekr --path /project --language python,go     # Scan specific languages
  ./seekr --vars "MY_SECRET,CUSTOM_KEY"            # Include custom variables
  ./seekr --regex_str "secret[_-]?key.*=.*"        # Use custom regex patterns

Detects: AWS credentials, GitHub tokens, API keys (Stripe, Twilio, etc.), database passwords, 
OAuth secrets, SSH private keys, Docker secrets, CI/CD variables, and much more.`,

	Run: func(cmd *cobra.Command, args []string) {
		var varsToCheck []string
		checkParams()

		selectedExtensions := []string{}
		selectedLanguages := strings.Split(FLAG_LANGUAGE, ",")

		// fmt.Println("[DEBG]    Selected languages for scanning:", selectedLanguages)

		// If nothing or all is provided, even with a list of other languages, just
		// include all supported languages.  Otherwise add all extensions that were
		// selected by the user
		if FLAG_LANGUAGE == "" || slices.Contains(selectedLanguages, "all") {
			for _, ext := range LANG_EXT {
				selectedExtensions = append(selectedExtensions, ext...)
			}
		} else {
			for _, lang := range selectedLanguages {
				if extensions, ok := LANG_EXT[lang]; ok {
					selectedExtensions = append(selectedExtensions, extensions...)
				} else {
					log.Printf("[WARN]    Language '%s' is not supported or has no associated file extensions. Skipping.", lang)
				}
			}
		}

		if !FLAG_IGNORE_DEFAULT_VARS {
			defaultVars := strings.Join(extractVars(VARIABLES_LIST, DelimeterNewline), "|")
			varsToCheck = append(varsToCheck, defaultVars)
		}

		if FLAG_USR_VARS_FILE != "" {
			// Read user-defined variables from a file
			fileContent, err := os.ReadFile(FLAG_USR_VARS_FILE)
			if err != nil {
				log.Fatalf("[FATAL] Unable to read user-defined variables file: %v", err)
			}
			usrVars := strings.Join(extractVars(string(fileContent), DelimeterNewline), "|")
			varsToCheck = append(varsToCheck, usrVars)
		}

		if FLAG_USR_VARS != "" {
			varsToCheck = append(varsToCheck, strings.Join(extractVars(FLAG_USR_VARS, DelimeterComma), "|"))
		}

		if FLAG_USR_REGEX_FILE != "" {
			// Read user-defined regular expressions from a file
			fileContent, err := os.ReadFile(FLAG_USR_REGEX_FILE)
			if err != nil {
				log.Fatalf("[FATAL] Unable to read user-defined regex file: %v", err)
			}
			usrRegex := strings.Join(extractVars(string(fileContent), DelimeterNewline), "|")
			varsToCheck = append(varsToCheck, usrRegex)
		}
		if FLAG_USR_REGEX != "" {
			usrRegex := strings.Join(extractVars(FLAG_USR_REGEX, DelimeterComma), "|")
			varsToCheck = append(varsToCheck, usrRegex)
		}

		// fmt.Println("[DEBG]    Selected file extensions for scanning:", selectedExtensions)

		fileInfo, err := os.Stat(FLAG_SCAN_PATH)
		if err != nil {
			log.Fatalf("[FATAL] Unable to open scan path (does it exist?): %v", err)
		}

		if !fileInfo.IsDir() {
			FLAG_INCLUDE_ALL = true
			if !checkIfText(FLAG_SCAN_PATH) && !FLAG_INCLUDE_BINARY {
				log.Fatal("[FATAL] A binary file was passed without setting the --binary_check flag.")
			}
		}

		processBinary := FLAG_INCLUDE_BINARY && FLAG_INCLUDE_ALL

		err = filepath.WalkDir(FLAG_SCAN_PATH, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("[FATAL] accessing path %q: %v", path, err)
			}
			if !d.IsDir() {
				// fmt.Println("[DEBG]    Processing file:", path)
				ext := filepath.Ext(path)
				if slices.Contains(selectedExtensions, ext) || FLAG_INCLUDE_ALL {

					fmt.Println("[INFO] Scanning file:", path)
					vulns, err := processFile(path, varsToCheck, !FLAG_CASE_INSENSITIVE, processBinary)
					if err != nil {
						log.Printf("[ERROR]    Error processing file %q: %v", path, err)
					}
					for _, vuln := range vulns {
						fmt.Printf("[VULN]    Found potentially leaked secret (Line %03d): %s\n", vuln.LineNum, vuln.VarName)
						fmt.Printf("[VULN]      With Value: '%s'\n", vuln.VarContent)
					}
				} else {
					fmt.Println("[INFO]    Scanning directory:", path)
				}
			}
			// printing empty line to cleanly separate output
			fmt.Println("")
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func checkParams() {
	if FLAG_INCLUDE_BINARY && !FLAG_INCLUDE_ALL {
		fmt.Println("[WARN]    The --binary_check flag is only effective when the --all_files flag is also set.")
		fmt.Println("[WARN]      It will be ignored.")
	} else if FLAG_INCLUDE_BINARY && FLAG_INCLUDE_ALL {
		fmt.Println("[WARN]    Scanning binary files only scans for embedded text based secrets and can take a long")
		fmt.Println("[WARN]      time to process files.  Scanning binaries may also result in system instability.")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seekr.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&FLAG_SCAN_PATH, "path", "p", ".", "The path to the file or directory to scan.")
	rootCmd.Flags().BoolVarP(&FLAG_INCLUDE_BINARY, "binary_check", "b", false, "Include binary files in the scan.")
	rootCmd.Flags().StringVarP(&FLAG_LANGUAGE, "language", "l", "", "The programming language to scan for secrets. If not specified, or 'all' is provided then all supported languages will be scanned.")
	rootCmd.Flags().BoolVarP(&FLAG_INCLUDE_ALL, "all_files", "a", false, "Include all files in the scan, regardless of file extension.  This overrides the language flag.")
	rootCmd.Flags().StringVarP(&FLAG_USR_VARS, "vars", "v", "", "Comma-separated list of additional variables to include in the scan.")
	rootCmd.Flags().StringVarP(&FLAG_USR_REGEX, "regex_str", "r", "", "User-defined regular expression for matching custom/unsupported secrets.")
	rootCmd.Flags().StringVarP(&FLAG_USR_VARS_FILE, "vars_file", "V", "", "Path to a file containing additional variables to include in the scan, one per line.")
	rootCmd.Flags().StringVarP(&FLAG_USR_REGEX_FILE, "regex_file", "R", "", "Path to a file containing user-defined regular expressions for matching custom/unsupported secrets, one per line.")
	rootCmd.Flags().BoolVarP(&FLAG_IGNORE_DEFAULT_VARS, "ignore_default", "i", false, "Ignore the default set of vulnerable variables and only use user-defined variables and regex patterns.")
	rootCmd.Flags().BoolVarP(&FLAG_CASE_INSENSITIVE, "case_insensitive", "C", false, "Perform case-inensitive checks for variables and regex patterns.  By default, all checks are case-sensitive.  Note: This may produce significantly more false positives.")

}

type DelimeterOptions string

const (
	DelimeterNewline DelimeterOptions = "\n"
	DelimeterComma   DelimeterOptions = ","
)

func extractVars(content string, delimeter DelimeterOptions) []string {
	var extractedVars []string

	// Remove leading and trailing whitespace
	cleaned := strings.TrimSpace(content)

	// Remove comments (lines starting with #)
	lines := strings.Split(cleaned, string(delimeter))
	for _, line := range lines {
		// Skip the line if it's empty or a comment
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		extractedVars = append(extractedVars, line)
	}
	return extractedVars
}

// unique removes duplicate elements from a slice of any comparable type.
func unique[T comparable](slice []T) []T {
	// Create a map to store keys we've seen.
	// The empty struct struct{} is used because it takes up zero memory.
	seen := make(map[T]struct{})

	// Create a new slice to store the unique results.
	// Pre-allocating with a capacity can be a small optimization.
	result := make([]T, 0, len(slice))

	// Iterate over the input slice.
	for _, item := range slice {
		// If the item has not been seen before...
		if _, ok := seen[item]; !ok {
			// Mark it as seen.
			seen[item] = struct{}{}
			// Append it to the result slice.
			result = append(result, item)
		}
	}

	return result
}

// uniqueSlices removes duplicate inner slices from a slice of slices.
func uniqueSlices[T comparable](sliceOfSlices [][]T) [][]T {
	// A map to store string representations of slices we have already seen.
	seen := make(map[string]struct{})

	// The slice to store the unique results.
	result := make([][]T, 0, len(sliceOfSlices))

	for _, innerSlice := range sliceOfSlices {
		// Convert the inner slice to a string to use as a map key.
		// fmt.Sprint() creates a consistent string like "[0 5]".
		key := fmt.Sprint(innerSlice)

		// If we haven't seen this string key before...
		if _, ok := seen[key]; !ok {
			// Mark it as seen.
			seen[key] = struct{}{}
			// Append the original inner slice to our result.
			result = append(result, innerSlice)
		}
	}

	return result
}

// processFile: processes a file based on its type (binary or text) and the selected extensions.
func processFile(filePath string, varsToCheck []string, case_sensitive bool, processBinary bool) ([]VulnVars, error) {
	var contentStr string
	var err error
	var vulns []VulnVars

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Unable to read file: %s, %v", filePath, err)
	}

	is_text := checkIfText(filePath)
	if is_text {
		contentStr = string(content)
	} else if processBinary {
		contentStr, _ = binToStrings(content)
	}
	for _, varToCheck := range varsToCheck {
		vulns = append(vulns, checkForVulnVars(contentStr, varToCheck, case_sensitive)...)
	}
	vulns = unique(vulns)
	return vulns, nil
}

// checkIfText: checks if a file is a text file based on its content.
func checkIfText(filePath string) bool {
	file, err := os.Open(filePath)

	// Skip files that cannot be opened but alert the user
	if err != nil {
		fmt.Println("[ERROR] Unable to open file:", filePath, err)
		return false
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, err := file.Read(buf[:])

	// Skip files that cannot be read but alert the user
	if err != nil {
		fmt.Println("[ERROR] Unable to read file:", filePath, err)
		return false
	}

	// Use the http module to automatically detect the file type
	contentType := http.DetectContentType(buf[:n])

	// We assume that all of these are text based.  Anything other than that is
	// assumed to be binary files and will be processed as such
	assumedTextTypes := []string{"text/", "application/json", "application/xml", "application/yaml", "application/sql"}
	is_text := false
	for _, t := range assumedTextTypes {
		if strings.HasPrefix(contentType, t) {
			is_text = true
			break
		}
	}

	return is_text
}

type VulnVars struct {
	LineNum    int
	VarName    string
	VarContent string
}

// checkForVulnVars: checks for vulnerable variables in the given content.
func checkForVulnVars(content string, vars_str string, case_sensitive bool) []VulnVars {
	var foundVulnVars []VulnVars
	var matches [][]int

	var vuln_reg string
	if case_sensitive {
		vuln_reg = "(" + vars_str + ")(.*)"
	} else {
		vuln_reg = "(?i)(" + vars_str + ")(.*)"
	}
	re, err := regexp.Compile(vuln_reg)
	if err != nil {
		log.Fatal(err)
	}
	matches = re.FindAllStringSubmatchIndex(content, -1)

	// Because we may have duplicate matches from user provided vars/regex
	// make sure to remove them
	matches = uniqueSlices(matches)

	if len(matches) == 0 {
		return nil
	} else {
		for _, matchIndices := range matches {
			startIndex := matchIndices[0]
			varName := content[matchIndices[2]:matchIndices[3]]
			varContent := content[matchIndices[len(matchIndices)-2]:matchIndices[len(matchIndices)-1]]
			lineNum := strings.Count(content[:startIndex], "\n") + 1
			foundVulnVars = append(foundVulnVars, VulnVars{LineNum: lineNum, VarName: varName, VarContent: varContent})
		}

		return foundVulnVars
	}
}

// isPrintable: checks if a byte is printable including new line chars.
func isPrintable(b byte) bool {
	return (b >= 32 && b <= 126) || b == '\n' || b == '\r' || b == '\t'
}

// binToStrings: converts binary content to a slice of strings.
//
//	effectively the same as using the linux command `strings` on a binary file
func binToStrings(content []byte) (string, error) {
	const minLen = 4

	if len(content) < minLen {
		return "", fmt.Errorf("binary content too short")
	}

	var currentString bytes.Buffer
	var foundStrings []string
	for _, b := range content {
		if isPrintable(b) {
			currentString.WriteByte(b)
		} else {
			if currentString.Len() >= minLen {
				foundStrings = append(foundStrings, currentString.String())
			}
			currentString.Reset()
		}
	}

	if currentString.Len() >= minLen {
		foundStrings = append(foundStrings, currentString.String())
	}

	return strings.Join(foundStrings, ""), nil
}

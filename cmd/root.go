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
	"strings"

	"github.com/spf13/cobra"
)

var FLAG_SCAN_PATH string
var FLAG_INCLUDE_BINARY bool
var FLAG_INCLUDE_ALL bool
var FLAG_LANGUAGE string
var FLAG_USR_VARS string
var FLAG_USR_REGEX string
var FLAG_USR_VARS_FILE string
var FLAG_USR_REGEX_FILE string
var FLAG_IGNORE_DEFAULT_VARS bool

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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if FLAG_INCLUDE_BINARY && !FLAG_INCLUDE_ALL {
			fmt.Println("[WARN]    The --binary_check flag is only effective when the --all_files flag is also set.")
			fmt.Println("[WARN]      It will be ignored.")
		} else if FLAG_INCLUDE_BINARY && FLAG_INCLUDE_ALL {
			fmt.Println("[WARN]    Scanning binary files only scans for embedded text based secrets and can take a long")
			fmt.Println("[WARN]      time to process files.  Scanning binaries may also result in system instability.")
		}

		selectedExtensions := []string{}
		selectedLanguages := strings.Split(FLAG_LANGUAGE, ",")

		// fmt.Println("[DEBG]    Selected languages for scanning:", selectedLanguages)

		// If nothing or all is provided, even with a list of other languages, just
		// include all supported languages.  Otherwise add all extensions that were
		// selected by the user
		if FLAG_LANGUAGE == "" || contains(selectedLanguages, "all") {
			for _, ext := range LANG_EXT {
				selectedExtensions = append(selectedExtensions, ext...)
			}
		} else {
			for _, lang := range selectedLanguages {
				if extensions, ok := LANG_EXT[lang]; ok {
					selectedExtensions = append(selectedExtensions, extensions...)
				} else if lang == "all" {
					for _, ext := range LANG_EXT {
						selectedExtensions = append(selectedExtensions, ext...)
					}
				}
			}
		}

		if FLAG_USR_VARS_FILE != "" {
			// Read user-defined variables from a file
			fileContent, err := os.ReadFile(FLAG_USR_VARS_FILE)
			if err != nil {
				log.Fatalf("[FATAL] Unable to read user-defined variables file: %v", err)
			}
			variablesList := extractContent(string(fileContent))
			if FLAG_USR_VARS != "" {
				FLAG_USR_VARS += "," + strings.Join(variablesList, ",")
			} else {
				FLAG_USR_VARS = strings.Join(variablesList, ",")
			}
		}

		if FLAG_USR_REGEX_FILE != "" {
			// Read user-defined regular expressions from a file
			fileContent, err := os.ReadFile(FLAG_USR_REGEX_FILE)
			if err != nil {
				log.Fatalf("[FATAL] Unable to read user-defined regex file: %v", err)
			}
			regexList := extractContent(string(fileContent))
			if FLAG_USR_REGEX != "" {
				FLAG_USR_REGEX += "|" + strings.Join(regexList, "|")
			} else {
				FLAG_USR_REGEX = strings.Join(regexList, "|")
			}
		}

		// fmt.Println("[DEBG]    Selected file extensions for scanning:", selectedExtensions)

		// TODO: Consider the case of processing a single file but isn't directly supported/is a binary file
		//  should we force the processing or alert the user that it won't do anything?
		fileInfo, err := os.Stat(FLAG_SCAN_PATH)
		if err != nil {
			log.Fatalf("[FATAL] Unable to stat scan path: %v", err)
		}

		if !fileInfo.IsDir() {
			FLAG_INCLUDE_ALL = true
			if !checkIfText(FLAG_SCAN_PATH) && !FLAG_INCLUDE_BINARY {
				log.Fatal("[FATAL] A binary file was passed without setting the --binary_check flag.")
			}
		}

		err = filepath.WalkDir(FLAG_SCAN_PATH, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("[FATAL] accessing path %q: %v", path, err)
			}
			if !d.IsDir() {
				// fmt.Println("[DEBG]    Processing file:", path)
				processFile(path, selectedExtensions, FLAG_INCLUDE_ALL, FLAG_INCLUDE_BINARY, FLAG_USR_VARS, FLAG_USR_REGEX, FLAG_IGNORE_DEFAULT_VARS)
			} else {
				fmt.Println("[INFO]    Scanning directory:", path)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
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
}

func extractContent(content string) []string {
	// Remove leading and trailing whitespace
	cleaned := strings.TrimSpace(content)

	// Remove comments (lines starting with #)
	lines := strings.Split(cleaned, "\n")
	var result []string
	for _, line := range lines {
		// Skip the line if it's empty or a comment
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		result = append(result, line)
	}
	return result
}

// contains: checks if a value exists in a slice of any comparable type.
func contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
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
func processFile(filePath string, selectedExtensions []string, check_all bool, check_binary bool, user_vars_str string, user_regex_str string, ignore_default bool) {
	is_text := checkIfText(filePath)
	if is_text {
		// process file only if with selected extensions OR if all_files is set
		ext := filepath.Ext(filePath)
		if contains(selectedExtensions, ext) || check_all {
			fmt.Println("[INFO]    Scanning file:", filePath)
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("[ERROR] Unable to read file:", filePath, err)
				return
			}
			checkForVulnVars(string(content), user_vars_str, user_regex_str, ignore_default)
		} else {
			// fmt.Println("[DEBG]    Skipping file due to unselected/unsupported extension:", ext)
		}

		// only scan binary files if we have the "all_files" and binary_check
		// flags set
	} else if check_all && check_binary {
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("[ERROR] Unable to read file:", filePath, err)
			return
		}
		fmt.Println("[INFO]    Scanning binary file:", filePath)
		checkForVulnVarsBinary(content, user_vars_str, user_regex_str, ignore_default)
	}

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

// checkForVulnVars: checks for vulnerable variables in the given content.
func checkForVulnVars(content string, user_vars_str string, user_regex_str string, ignore_default bool) bool {

	var matches [][]int
	if !ignore_default {
		variablesList := extractContent(VARIABLES_LIST)
		vuln_vars := strings.Join(variablesList, "|")
		vuln_reg := "(?i)(" + vuln_vars + ")(.*)"
		// for all checks we do case insensitive checks as variables may use different casing
		re, err := regexp.Compile(vuln_reg)
		if err != nil {
			log.Fatal(err)
		}
		matches = re.FindAllStringSubmatchIndex(content, -1)
	}

	// find matches according to the user defined variables and append them
	// to the matches array
	if user_vars_str != "" {
		usr_vars := strings.Join(strings.Split(user_vars_str, ","), "|")
		usr_vars_re, err := regexp.Compile("(?i)(" + usr_vars + ")(.*)")
		if err != nil {
			log.Fatal(err)
		}
		matches = append(matches, usr_vars_re.FindAllStringSubmatchIndex(content, -1)...)
	}

	if user_regex_str != "" {
		usr_regex_re, err := regexp.Compile("(?i)(" + user_regex_str + ")(.*)")
		if err != nil {
			log.Fatal(err)
		}
		matches = append(matches, usr_regex_re.FindAllStringSubmatchIndex(content, -1)...)
	}

	// Because we may have duplicate matches from user provided vars/regex
	// make sure to remove them
	matches = uniqueSlices(matches)

	if len(matches) == 0 {
		fmt.Println("[INFO]    No vulnerable variables found.")
		return false
	} else {
		fmt.Println("[VULN]    Found the following potential vulnerable variables:")
		for _, matchIndices := range matches {
			startIndex := matchIndices[0]
			endIndex := matchIndices[1]

			match := content[startIndex:endIndex]
			lineNum := strings.Count(content[:startIndex], "\n") + 1
			fmt.Printf("[VULN]      - %03d: %s\n", lineNum, match)
		}

		return true
	}
}

// isPrintable: checks if a byte is printable.
func isPrintable(b byte) bool {
	return (b >= 32 && b <= 126) || b == '\n' || b == '\r' || b == '\t'
}

// checkForVulnVarsBinary: checks for vulnerable variables in binary content.
//
//	effectively the same as using the linux command `strings` on a binary file
func checkForVulnVarsBinary(content []byte, user_vars_str string, user_regex_str string, ignore_default bool) bool {
	const minLen = 4

	if len(content) < minLen {
		fmt.Println("[INFO]    Skipping binary file (too short)")
		return false
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

	return checkForVulnVars(strings.Join(foundStrings, ""), user_vars_str, user_regex_str, ignore_default)
}

// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [CHRISTOPHER OKOH]
// Squad:  [THE iNTERFACE]

/*
// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [The Interface]
// ───────────────────────────────────────────
// Input line types:
//   [
        Lines in ALL CAPS
        Lines in all lowercase
        Lines with extra leading/trailing spaces
        Lines that are only dashes or blanks
       ]
//
// Transformation rules (in order):
//   1. [Convert ALL CAPS lines to Title Case]
//   2. [Convert all lowercase lines to uppercase]
//   3. [Trim all leading and trailing whitespace]
//   4. [Reverse the words in any line that contains the word REVERSE]
//   5. [Remove lines that are only dashes or blanks]
//
// Output format:
//   [Header: yes]
//   [Line numbering format: "1"]
//   [Summary block: yes]
//
// Terminal summary fields:
//   [Lines read    : [number]]
//   [Lines written : [number]]
//   [Lines removed : [number]]
//   [Rules applied : [list your 5 rules]]
// ═══════════════════════════════════════════
═════════════════════════════════════════
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func CapAll(line string) string {
	if line == strings.ToUpper(line) && line != "" {
		words := strings.Fields(line)
		for i, w := range words {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
		return strings.Join(words, " ")
	}
	return line
}

func lowerToUpper(line string) string {
	if line == strings.ToLower(line) && line != "" {
		return strings.ToUpper(line)
	}
	return line
}

func TRIM(line string) string {
	return strings.TrimSpace(line)
}
func REVERSE(line string) string {
	if strings.Contains(line, "REVERSE") {
		words := strings.Fields(line)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	return line
}

func IsBlankOrDash(line string) bool {
	trimmed := strings.TrimSpace(line)
	return trimmed == "" || trimmed == "─────────────────────────"
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if input == output {
		fmt.Println("Input and output cannot be the same file.")
		return
	}

	file, err := os.Open(input)
	if err != nil {
		fmt.Println("File not found:", input)
		return
	}
	defer file.Close()

	if info, err := os.Stat(output); err == nil && info.IsDir() {
		fmt.Println("Cannot write to output: path is a directory, not a file.")
		return
	}

	scanner := bufio.NewScanner(file)

	var processed []string
	linesRead := 0
	linesRemoved := 0

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		line = CapAll(line)
		line = lowerToUpper(line)
		line = TRIM(line)
		line = REVERSE(line)
		if IsBlankOrDash(line) {
			linesRemoved++
			continue
		}

		processed = append(processed, line)
	}

	if linesRead == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
	}

	out, err := os.Create(output)
	if err != nil {
		fmt.Println("Failed to write output file.")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	writer.WriteString("SENTINEL INTERFACE REPORT\n")
	writer.WriteString("─────────────────────────\n")

	for i, line := range processed {
		writer.WriteString(fmt.Sprintln(i+1, line))
	}

	writer.WriteString("─────────────────────────\n")
	writer.WriteString(fmt.Sprintln("Lines read : ", linesRead))
	writer.WriteString(fmt.Sprintln("Lines written :", len(processed)))
	writer.WriteString(fmt.Sprintln("Lines removed :", linesRemoved))
	writer.WriteString("Lines removed : 0\n")

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	writer.WriteString(fmt.Sprintln("Processed at :", timestamp))

	writer.Flush()
	fmt.Println("Lines read :", linesRead)
	fmt.Println("Lines written :", len(processed))
	fmt.Println("Lines removed :", linesRemoved)
	fmt.Println("Rules applied : CAPS→Title, lower→UPPER, Trim, Reverse(REVERSE), Remove Dashes/Blanks")
	fmt.Println("Processed at :", timestamp)
}

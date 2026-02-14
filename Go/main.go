package main
import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"flag"
	"io"
	"archive/tar"
	"compress/gzip"

)
func main() {

	// Application Header
	fmt.Println("======= Log Archive Tool =======")

	// Parse Command-Line Argument
	LogDir := flag.String("logdir", "", "Directory containing log files")
	DaysToKeep := flag.Int("days", 30, "Number of days to keep logs before archiving and deleting")
	flag.Parse()

	// Validate Log Directory (ask user if not provided)
	if *LogDir == "" {
		fmt.Print("No log directory provided. Enter log directory: ")
		fmt.Scanln(LogDir)
		if *LogDir == "" {
			*LogDir = "var/logs"
			fmt.Println("Using default:", *LogDir)
		}
	}
	
	// Create Archive Name with Timestamp
	timestamp := time.Now().Format("20060102_150405") 
	archiveName := fmt.Sprintf("logs_archive_%s.tar.gz", timestamp) 
	archiveDir := filepath.Join(*LogDir, "archives")
	archivePath := filepath.Join(archiveDir, archiveName)

	// Ensure Archive Directory Exists
	err := os.MkdirAll(archiveDir, 0755)
	if err != nil {
		fmt.Printf("Error: Archive directory could not be created: %v\n", err)
		return
	}

	// Compress Logs into Archive
	err = compressLogs(*LogDir, archivePath, archiveName)

	if err != nil {
		noPermErr := fmt.Sprintf("open %s/%s: permission denied", archivePath, archiveName)
		errstr := fmt.Sprintf("%s", err)
		if errstr == noPermErr {
			fmt.Println("Required permissions were not found, please run as root so the files can be created")
		} else {
			fmt.Printf("Error: Logs could not be compressed: %v\n", err)
			return
		}

	}

	// the process of logging in details
	logFilePath := filepath.Join(*LogDir, "archive_log.txt")
	logEntry := fmt.Sprintf("%s: Archived logs to %s\n", timestamp, archivePath)
	err = appendToFile(logFilePath, logEntry)
	if err != nil {
		fmt.Printf("Error: Could not write to log file: %v\n", err)
		return
	}
	fmt.Printf("Archiving completed: %s\n", archivePath)

	// Delete logs older than specified days
	cutoff := time.Now().AddDate(0, 0, -*DaysToKeep)
	err = filepath.Walk(*LogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if filepath.Base(path) == "archives" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Base(path) == "archive_log.txt" {
			return nil
		}
		if info.ModTime().Before(cutoff) {
			return os.Remove(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: Could not delete old logs: %v\n", err)
		return
	}
	fmt.Printf("Deleted logs older than %d days\n", *DaysToKeep)



func compressLogs(LogDir, archivePath, archiveName string) error {
	file, err := os.Create(archivePath)
	if err != nil {
		return err
}
	defer file.Close()
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	
	err = filepath.Walk(LogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if filepath.Base(path) == "archives" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Base(path) == "archive_log.txt" {
			return nil
		}
		return addFileToTarGz(tw, path, LogDir)
	} )
	return err}
func addFileToTarGz(tw *tar.Writer, filePath, baseDir string) error {
		file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	relPath, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		return err
	}
	header.Name = relPath

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	return err
}

func appendToFile(filePath, text string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}
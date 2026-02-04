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
	flag.Parse()

	// Validate Log Directory (default to "var/logs" if not provided)
	if *LogDir == "" {
		*LogDir = "var/logs"
		fmt.Println(" No log directory provided. Using default:", *LogDir)
		
	}
	
	// Create Archive Name with Timestamp
	timestamp := time.Now().Format("20060102_150405") 
	archiveName := fmt.Sprintf("logs_archive_%s.tar.gz", timestamp) 
	archivePath := filepath.Join(*LogDir, archiveName) 

	// Ensure Archive Directory Exists
	err := os.MkdirAll(archivePath, 0755)
	if err != nil {
		fmt.Printf("Error: Archive directory could not be created: %v\n", err)
		return
	}

	// Compress Logs into Archive
	archiveFullPath := filepath.Join(archivePath, archiveName)
	err = compressLogs(*logDir, archiveFullPath, archiveName)

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
	logFilePath := filepath.Join(*logDir, "archive_log.txt")
	logEntry := fmt.Sprintf("%s: Archived logs to %s\n", timestamp, archiveName)
	err = appendToFile(logFilePath, logEntry)
	if err != nil {
		fmt.Printf("Error: Could not write to log file: %v\n", err)
		return
	}
	fmt.Printf("Archiving completed: %s\n", archiveFullPath)
}



func compressLogs(logDir, archivePath, archiveName string) error {}
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
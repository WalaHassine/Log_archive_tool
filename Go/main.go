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
	fmt.Println("======= Log Archive Tool =======")
	LogDir := flag.String("logdir", "", "Directory containing log files")
	flag.Parse()

	if *LogDir == "" {
		*LogDir = "var/logs"
		fmt.Println(" No log directory provided. Using default:", *LogDir)
		
	}
	
	timestamp := time.Now().Format("20060102_150405")
	archiveName := fmt.Sprintf("logs_%s.tar.gz", timestamp)
	archivePath := filepath.Join(*LogDir, archiveName)
}
func compressLogs(sourceDir, destDir string) error {}
func addFileToTarGz(tw *tar.Writer, filePath, baseDir string) error {}
func appendToFile(filePath, text string) error {}
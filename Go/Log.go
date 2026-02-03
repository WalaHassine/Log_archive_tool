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
func main() {}
func compressLogs(sourceDir, destDir string) error {}
func addFileToTarGz(tw *tar.Writer, filePath, baseDir string) error {}
func appendToFile(filePath, text string) error {}
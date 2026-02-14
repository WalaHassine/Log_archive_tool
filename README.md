# Log_archive_tool
------------------------------------------------
# Description
This CLI tool compresses and archives log files of a Linux server.
This tool is created using different approachs and solutions :
# 1. GoLang (Go/main.go)
------------------------------------------------
**Overview**  
A command-line tool written in Go that archives log files into a compressed tar.gz format with timestamped naming, using Go's standard library for compression and file handling.

**Features**  
- Command-line flag support for specifying log directory  
- Interactive prompt for log directory if it's not specified via command-line flag  
- Configurable retention period for logs (days to keep before deletion)  
- Automatic timestamped archive naming  
- Compression using tar.gz format  
- Logging of archiving actions to a file  
- Deletion of original logs after archiving (based on retention period)  
- Error handling for permissions and directory creation  

**Architecture**  
- **main()**: Parses command-line flags, validates input, creates archive paths, calls compression function, and logs results.  
- **compressLogs()**: Handles the creation of tar.gz archives (implementation incomplete in current code).  
- **addFileToTarGz()**: Adds individual files to the tar archive with relative paths.  
- **appendToFile()**: Appends log entries to a text file.  

**Execution Flow**  
1. Parse `--logdir` and `--days` flags (prompt for logdir if not provided).  
2. Generate timestamp and create archive name (e.g., logs_archive_20260207_120000.tar.gz).  
3. Ensure archive directory exists (create if needed).  
4. Compress all log files into the tar.gz archive.  
5. Log the archiving action to "archive_log.txt".  
6. Delete original log files older than the specified number of days.  
7. Print completion messages.

**Configuration Variables**  
- `logDir`: Directory containing log files (command-line flag `--logdir`, prompted if not provided).  
- `daysToKeep`: Number of days to retain logs before archiving and deleting (command-line flag `--days`, default: 30).  

**Output**  
- Console: Status messages, errors (e.g., permission issues), and completion confirmation.  
- Files:  
  - Archive: `logs_archive_<timestamp>.tar.gz` in the specified log directory.  
  - Log: `archive_log.txt` with entries like "timestamp: Archived logs to archive_name".  

# 2. Shell Script (Shell/Log.sh)
------------------------------------------------
**Overview**  
An interactive bash script that provides a menu-driven interface for configuring log archiving parameters and executing the process, with support for automated scheduling via cron jobs.

**Features**  
- Interactive menu for setting log directory, retention periods, and running archiving.  
- Archives logs older than specified days using `find` and `tar` commands.  
- Deletes archived logs from the source directory.  
- Cleans up old archive files based on retention settings.  
- Optional cron job setup for daily automated execution.  

**Architecture**  
- **input_prompt()**: Function for prompting user input with default values.  
- **Main loop**: Case-based menu system handling user selections.  
- **Setup_cron_job()**: Function to configure a cron job for automation.  

**Execution Flow**  
1. Display interactive menu with options (1-5).  
2. Based on user choice:  
   - Set log directory (validate existence).  
   - Set days to keep logs.  
   - Set days to keep backup archives.  
   - Run archiving: Create archive directory, find old logs, compress with tar.gz, delete originals, clean old archives.  
   - Exit the script.  
3. After menu loop, prompt for cron job setup (runs script daily at 2 AM if confirmed).  

**Configuration Variables**  
- `log_directory`: Path to log files (default: "/var/log").  
- `days_to_keep`: Number of days to retain logs before archiving (default: 30).  
- `backup_days_to_keep`: Number of days to retain archive files (default: 90).  

**Output**  
- Console: Menu prompts, user confirmations, status messages, and error notifications.  
- Files:  
  - Archive: `logs_<timestamp>.tar.gz` in the archive subdirectory.  
  - Log: Appends to `archive_log.txt` in the archive directory with archiving details.  
  - Cron: Updates user's crontab for automation (if selected).  

# 3. Python (Python/Log.py)
------------------------------------------------
**Overview**  
(Not implemented ).  

**Features**  
(Not implemented).  

**Architecture**  
(Not implemented).  

**Execution Flow**  
(Not implemented).  

**Configuration Variables**  
(Not implemented).  

**Output**  
(Not implemented).  

# URL of the project idea (Roadmap)
https://roadmap.sh/projects/log-archive-tool



# for dynamic inputs
input_prompt(){
    read -r -p "$1 [$2]: " input
    echo "${input:-$2}"
}

# the interactive loop

while true; do
echo "1. Specify the log directory"
echo "2. Specify number of days to keep logs"
echo "3. Specify number of days to keep backup archives"
echo "4. Run the log archiving process"
echo "5. Exit"
echo ""

read -r -p "Please select an option (1-5): " option

case $option in
    1)
        log_directory=$(input_prompt "Enter the log directory" "/var/log")
        if [ ! -d "$log_directory" ]; then
                echo "Error: Log directory does not exist."
                log_directory=""
        else
                echo "Log directory set to $log_directory"
        fi
        ;;
    2)
        days_to_keep=$(input_prompt "Enter the number of days to keep logs" "30")
        echo "Number of days to keep logs set to: $days_to_keep"
        ;;
    3)
        backup_days_to_keep=$(input_prompt "Enter the number of days to keep backup archives" "90")
        echo "Number of days to keep backup archives set to: $backup_days_to_keep"
        ;;
    4)
        echo "Running the log archiving process..."
        # Here you would call your log archiving function or script
        if [ -z $log_directory ]; then
            echo "Error: Log directory is not set. Please specify it first."
        else

        #archive directory
        archive_directory="$log_directory/archives"
        mkdir -p "$archive_directory"

        #archive name with timestamp
        timestamp =$(date +"%Y%m%d%H%M%S")
        archive_file="$archive_directory/logs_$timestamp.tar.gz"

        #find and compress logs older than specified days (days to keep)
        find "$log_directory" -type f -mtime +$days_to_keep print0 | tar -czvf "$archive_file" --null -T -
        echo "Logs older than $days_to_keep days have been archived to $archive_file" >> "$archive_directory/archive_log.txt"
        #find and delete logs older than the specified days (days to keep)
        find "$log_directory" -type f -mtime +$days_to_keep -exec rm {} \;
        echo "Logs older than $days_to_keep days have been deleted from $log_directory"
        echo "Log archiving process completed. Archive file: $archive_file"
        #find and delete backup archives older than the specified days (backup days to keep)
        find "$archive_directory"  -type f -name "*.tar.gz" -mtime +$backup_days_to_keep -exec rm {} \;
        echo "Backup archives older than $backup_days_to_keep days have been deleted from $archive_directory"
        fi
        ;;
    5)
        echo "Exiting..."
        break
        ;;
    *)
        echo "Invalid option. Please select a number between 1 and 5."
        ;;
esac
done

Setup_cron_job(){
    read -r -p " Do you want to set up a cron job to run this script automatically? (y/n): " setup_cron
    if [[ "$setup_cron" == "y" || "$setup_cron" == "Y" ]]; then
        echo "Setting up a cron job..."
        crontab -l > mycron
        echo "0 2 * * * /usr/local/bin/Log.sh" >> mycron
        crontab mycron
        rm mycron
        echo "Cron job set up successfully."
    else
        echo "Cron job setup skipped."
    fi
}

Setup_cron_job
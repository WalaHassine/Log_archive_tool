# present options to the user , archive logs, days to keep, bakcup days etc ..
# take the log directory as argument
# compress logs in a tar.gz file and timestamp it and store it in the archive log directory

import datetime
import os
import tarfile
import glob


log_directory = None
days_to_keep = None
backup_days = None


def print_menu():
    print("\n Log archiving script ")
    print("---------------------------------------")
    print("Please select an option (1-5): ")
    print("1. Specify the log directory")
    print("2. Specify number of days to keep logs")
    print("3. Specify number of days to keep backup archives")
    print("4. Run the log archiving process")
    print("5. Exit")
    return input("Enter your choice: ")


def configure_log_directory():
    global log_directory
    log_directory = input("Enter the log directory path: ")
    print(f"Log directory set to: {log_directory}")


def configure_days_to_keep():
    global days_to_keep
    days_to_keep = int(input("Enter the number of days to keep logs: "))
    print(f"Number of days to keep logs set to: {days_to_keep}")


def configure_backup_days():
    global backup_days
    backup_days = int(input("Enter the number of days to keep backup archives: "))
    print(f"Number of days to keep backup archives set to: {backup_days}")


def run_archiving_process():
    global log_directory, days_to_keep, backup_days
    
    print("Running the log archiving process...")
    
    if not log_directory or not days_to_keep or not backup_days:
        print("Error: Please configure all settings first (options 1, 2, and 3).")
        return
    
    timestamp = datetime.datetime.now().strftime("%Y%m%d_%H%M%S")
    archive_name = f"logs_archive_{timestamp}.tar.gz"
    
    if not os.path.exists(log_directory):
        os.makedirs(log_directory)
    
    cutoff_date = datetime.datetime.now() - datetime.timedelta(days=days_to_keep)
    
    log_files = glob.glob(os.path.join(log_directory, "*.log"))
    
    files_to_archive = []
    for log_file in log_files:
        file_mtime = datetime.datetime.fromtimestamp(os.path.getmtime(log_file))
        if file_mtime < cutoff_date:
            files_to_archive.append(log_file)
    
    if files_to_archive:
        try:
            with tarfile.open(archive_name, "w:gz") as tar:
                for file in files_to_archive:
                    tar.add(file, arcname=os.path.basename(file))
            
            print(f"Successfully created archive: {archive_name}")
            print(f"Archived {len(files_to_archive)} log files:")
            for f in files_to_archive:
                print(f"  - {os.path.basename(f)}")
            
            for file in files_to_archive:
                os.remove(file)
                print(f"Removed original file: {os.path.basename(file)}")
        
        except Exception as e:
            print(f"Error creating archive: {e}")
    
    else:
        print("No log files older than the specified days found.")
    
    archive_dir = log_directory
    archive_files = glob.glob(os.path.join(archive_dir, "logs_archive_*.tar.gz"))
    cutoff_backup_date = datetime.datetime.now() - datetime.timedelta(days=backup_days)
    
    removed_count = 0
    for archive_file in archive_files:
        file_mtime = datetime.datetime.fromtimestamp(os.path.getmtime(archive_file))
        if file_mtime < cutoff_backup_date:
            os.remove(archive_file)
            print(f"Removed old backup: {os.path.basename(archive_file)}")
            removed_count += 1
    
    if removed_count > 0:
        print(f"Cleaned up {removed_count} old backup archives.")
    else:
        print("No old backup archives to clean up.")
    
    print(f"Archiving logs from {log_directory} that are older than {days_to_keep} days and keeping backups for {backup_days} days.")


def show_current_settings():
    print("\n--- Current Settings ---")
    print(f"Log directory: {log_directory if log_directory else 'Not set'}")
    print(f"Days to keep logs: {days_to_keep if days_to_keep else 'Not set'}")
    print(f"Backup days: {backup_days if backup_days else 'Not set'}")
    print("------------------------\n")


while True:
    show_current_settings()
    choice = print_menu()
    
    if choice == "1":
        configure_log_directory()
    
    elif choice == "2":
        configure_days_to_keep()
    
    elif choice == "3":
        configure_backup_days()
    
    elif choice == "4":
        run_archiving_process()
    
    elif choice == "5":
        print("Exiting the script. Goodbye!")
        break
    
    else:
        print("Invalid choice. Please select a valid option (1-5).")



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
        echo "Log directory set to: $log_directory"
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
        ;;
    5)
        echo "Exiting..."
        break
        ;;
    *)
        echo "Invalid option. Please select a number between 1 and 5."
        ;;
esac
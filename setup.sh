#!/bin/bash
# ~~ Error handling setup ~~
set -e # Exit right away if a command returns non-zero
set -u # Errors if you try to reference an undeclared variable
set -o pipefail # Ensures failures during pipe execution aren't hidden by later successes

# ~~ Logging Stuff ~~

# ~ Color Definitions ~
# Check if stdout is a terminal
if [ -t 1 ]; then
  RCol='\033[0;31m' # Red
  GCol='\033[0;32m' # Green
  YCol='\033[0;33m' # Yellow
  BCol='\033[0;34m' # Blue
  NCol='\033[0m'    # No Color
else
  RCol=''
  GCol=''
  YCol=''
  BCol=''
  NCol=''
fi

# ~ Logging Function ~
# Usage: log_message <LEVEL> "message"
# LEVEL can be INFO, WARN, ERROR, DEBUG
log_message() {
  local level="$1"
  local message="$2"
  local color="${NCol}"
  local P_level=""

  case "$level" in
    PASS)  color="${GCol}"; P_level="PASS "; ;;
    WARN)  color="${YCol}"; P_level="WARN "; ;;
    ERROR) color="${RCol}"; P_level="ERROR"; ;;
    INFO)  color="${BCol}"; P_level="INFO"; ;;
    *)     P_level="LOG  "; ;;
  esac

  # Get current timestamp e.g., 2023-10-27 15:30:00
  timestamp=$(date '+%Y-%m-%d %H:%M:%S')
  echo -e "${color}[${timestamp} ${P_level}] ${message}${NCol}" >&2 # Log to stderr
}

# ~~ Variable setup ~~

BASE_URL="setup.dalton.dog"
BINARY_FILE="bootstrap.x86"
CONFIG_FILE="config.yaml"

BINARY_URL="${BASE_URL}/${BINARY_FILE}"
CONFIG_URL="${BASE_URL}/${CONFIG_FILE}"

# ~~ Cleanup Function and Trap ~~

ORIGINAL_PWD=$(pwd) # Save the directory where the script was invoked
TEMP_DIR=""         # Initialize TEMP_DIR

cleanup() {
  local exit_status=$? # Capture the exit status of the last command
  # Only attempt to cd if ORIGINAL_PWD is set and is a directory
  if [ -n "${ORIGINAL_PWD:-}" ] && [ -d "${ORIGINAL_PWD}" ]; then
    cd "${ORIGINAL_PWD}"
  fi

  if [ -n "${TEMP_DIR:-}" ] && [ -d "${TEMP_DIR}" ]; then
    log_message INFO "Cleaning up temporary directory: ${TEMP_DIR}"
    rm -rf "${TEMP_DIR}"
  else
    # Use simple echo if log_message infrastructure might be part of what failed
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] INFO : No temporary directory to clean or TEMP_DIR (${TEMP_DIR:-unset}) not valid." >&2
  fi

  if [ ${exit_status} -ne 0 ]; then
    log_message ERROR "Bootstrap script finished with an error (Exit Status: ${exit_status})."
  fi
  # The trap will ensure this function is called, then the script will exit with the original status
}
trap cleanup EXIT HUP INT TERM

# ~~ Main Execution ~~

log_message INFO "Beginning bootstrap process..."
log_message INFO "Current directory: ${ORIGINAL_PWD}"

log_message INFO "Creating temp directory..."
TEMP_DIR=$(mktemp -d)

if [ -z "$TEMP_DIR" ] || ! [ -d "$TEMP_DIR" ]; then
    log_message ERROR "Failed to create temporary directory."
    exit 1
fi
log_message PASS "Temporary directory created: ${TEMP_DIR}"

cd "$TEMP_DIR"
log_message INFO "Changed into ${TEMP_DIR}"

log_message INFO "Downloading ${BINARY_FILE} from ${BINARY_URL}..."
if ! wget --quiet -O "${BINARY_FILE}" "${BINARY_URL}"; then
    local wget_exit_code=$?
    log_message ERROR "Failed to download ${BINARY_FILE} from ${BINARY_URL}. wget exit code: ${wget_exit_code}"
    exit 1 # Trap will handle cleanup
fi

if ! [ -s "${BINARY_FILE}" ]; then
    log_message ERROR "Downloaded ${BINARY_FILE} is empty or does not exist!"
    exit 1
fi

log_message PASS "${BINARY_FILE} successfully downloaded."

log_message INFO "Downloading ${CONFIG_FILE} from ${CONFIG_URL}..."
if ! wget --quiet -O "${CONFIG_FILE}" "${CONFIG_URL}"; then
    local wget_exit_code=$?
    log_message ERROR "Failed to download ${CONFIG_FILE} from ${CONFIG_URL}. wget exit code: ${wget_exit_code}"
    exit 1 # Trap will handle cleanup
fi
if ! [ -s "${CONFIG_FILE}" ]; then # Check if file exists and is not empty
    log_message ERROR "Downloaded ${CONFIG_FILE} is empty or does not exist!"
    exit 1
fi
log_message PASS "${CONFIG_FILE} downloaded successfully."

chmod +x "${BINARY_FILE}"
log_message PASS "Made binary executable."

log_message INFO "Starting bootstrap program: ./${BINARY_FILE} ${CONFIG_FILE}"
(./${BINARY_FILE} ${CONFIG_FILE})
exit_code=$?

if [ ${exit_code} -ne 0]; then
	log_message ERROR "Error occurred in bootstrap program. Failed with exit code ${ exit_code }."
	exit $exit_code
fi

log_message PASS "Bootstrap complete!"

#!/bin/bash
# ~~ Error handling setup ~~
set -e # Exit right away if a command returns non-zero
set -u # Errors if you try to reference an undeclared variable
set -o pipefail # Ensures failures during pipe execution aren't hidden by later successes

# ~~ Variable setup ~~
echo "Beginning bootstrap"
echo "Setting up configuration"

BASE_URL="bootstrap.dalton.dog"
BINARY_FILE="bootstrap.x86"
CONFIG_FILE="config.yaml"

BINARY_URL="${BASE_URL}/${BINARY_FILE}"
CONFIG_URL="${BASE_URL}/${CONFIG_FILE}"

# ~~ Main Execution ~~
SCRIPT_DIR=$(dirname $(readlink -f $0))
TEMP_DIR=$(mktemp -d)

cd $TEMP_DIR

echo "Downloading ${BINARY_FILE} from ${BINARY_URL} into ${TEMP_DIR}"
wget "$BINARY_URL"

if ! [ -f "${BINARY_FILE}" ]; then
	echo "Unable to download binary file!"
	cd ..
	rm -rf ${TEMP_DIR}
	exit 1
fi

echo "Downloading ${CONFIG_FILE} from ${CONFIG_URL} into ${TEMP_DIR}"
wget "$CONFIG_URL"

if ! [ -f "${CONFIG_FILE}" ]; then
	echo "Unable to download config file!"
	cd ..
	rm -rf ${TEMP_DIR}
	exit 1
fi

echo "Files successfully downloaded."

chmod +x "${BINARY_FILE}"
echo "Made binary executable. Starting bootstrap program..."

(./${BINARY_FILE} ${CONFIG_FILE})
local exit_code=$?

cd ${SCRIPT_DIR}
echo "Cleaning up temp data"
rm -rf ${TEMP_DIR}

if [ $exit_code -ne 0]; then
	echo "Error occurred in bootstrap program. Failed with exit code $exit_code."
	exit $exit_code
fi

echo "Bootstrap complete!"

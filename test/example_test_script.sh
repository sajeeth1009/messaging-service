export GLOBAL_DB_CONNECTION_STR="<connection string without prefix and auth data>"
export GLOBAL_DB_CONNECTION_PREFIX="<connection string prefix (e.g. +srv for atlas)>"
export GLOBAL_DB_USERNAME="<db username>"
export GLOBAL_DB_PASSWORD="<db password>"

export MESSAGING_DB_CONNECTION_STR="<connection string without prefix and auth data>"
export MESSAGING_DB_CONNECTION_PREFIX="<connection string prefix (e.g. +srv for atlas)>"
export MESSAGING_DB_USERNAME="<db username>"
export MESSAGING_DB_PASSWORD="<db password>"

export DB_TIMEOUT=30 # seconds until connection times out
export DB_IDLE_CONN_TIMEOUT=45 # terminate idle connection after seconds
export DB_MAX_POOL_SIZE=8
export DB_DB_NAME_PREFIX="<db name prefix used in the test>" # DB names will be then > <DB_PREFIX>+"hard-coded-db-name-as-we-need-it"

export MESSAGING_SERVICE_LISTEN_PORT=5004
export USER_MANAGEMENT_LISTEN_PORT=5002
export STUDY_SERVICE_LISTEN_PORT=5003
export EMAIL_CLIENT_SERVICE_LISTEN_PORT=5005

export MESSAGING_CONFIG_FOLDER=$PWD

IP_ADDR="localhost"
export ADDR_USER_MANAGEMENT_SERVICE=$IP_ADDR:$USER_MANAGEMENT_LISTEN_PORT
export ADDR_STUDY_SERVICE=$IP_ADDR:$STUDY_SERVICE_LISTEN_PORT
export ADDR_EMAIL_CLIENT_SERVICE=$IP_ADDR:$EMAIL_CLIENT_SERVICE_LISTEN_PORT

# use -v for verbose for the script
go test ./... $1

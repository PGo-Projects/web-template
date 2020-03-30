#!/usr/bin/env bash

read -p 'What is the path of this module: ' MODULE_PATH
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SCRIPT_DIR="$(cygpath -m $SCRIPT_DIR)"
TEMPLATE_MODULE_PATH="github.com/PGo-Projects/web-template"

perl -pi -e s,$TEMPLATE_MODULE_PATH,$MODULE_PATH,g $SCRIPT_DIR/main.go
perl -pi -e s,$TEMPLATE_MODULE_PATH,$MODULE_PATH,g $SCRIPT_DIR/internal/security/security.go
perl -pi -e s,$TEMPLATE_MODULE_PATH,$MODULE_PATH,g $SCRIPT_DIR/internal/securitydb/mongo.go
perl -pi -e s,$TEMPLATE_MODULE_PATH,$MODULE_PATH,g $SCRIPT_DIR/internal/server/server.go

cd "$SCRIPT_DIR"
rm "$SCRIPT_DIR/init_windows.sh"
rm "$SCRIPT_DIR/init_unix.sh"

go mod init $MODULE_PATH


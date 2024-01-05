#!/usr/bin/env sh

SCRIPTPATH="$(
  cd "$(dirname "$0")"
  pwd -P
)"

CURRENT_DIR=$SCRIPTPATH
ROOT_DIR="$(dirname $CURRENT_DIR)"

INFRA_LOCAL_COMPOSE_FILE=$ROOT_DIR/build/docker-compose.dev.yaml

function local_infra() {
  docker-compose -f $INFRA_LOCAL_COMPOSE_FILE $@
}

function init() {
    cd $CURRENT_DIR/..
    goimports -w ./..
    go fmt ./...
}

function infra() {
  case $1 in
  up)
    local_infra up ${@:2}
    ;;
  down)
    local_infra down ${@:2}
    ;;
  build)
    local_infra build ${@:2}
    ;;
  *)
    echo "up|down|build [docker-compose command arguments]"
    ;;
  esac
}

function api_start() {
  echo "Starting infrastructure..."
  infra up -d
  setup_env_variables
  echo "Start api app config file: $CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/cmd/service/main.go"
  go run $ENTRY_FILE --config-file=$CONFIG_FILE
}

function worker_start() {
  echo "Starting infrastructure..."
  infra up -d
  setup_env_variables
  echo "Start api app config file: $CONFIG_FILE"
  ENTRY_FILE="$ROOT_DIR/cmd/worker/main.go"
  go run $ENTRY_FILE --config-file=$CONFIG_FILE
}

function setup_env_variables() {
    set -a
    export $(grep -v '^#' "$ROOT_DIR/build/.base.env" | xargs -0) >/dev/null 2>&1
    . $ROOT_DIR/build/.base.env
    set +a
    export CONFIG_FILE=$ROOT_DIR/build/app.yaml
}

function api() {
    case $1 in
    start)
        api_start
        ;;
    worker_start)
        worker_start
        ;;
    *)
        echo "[test|start|worker_start|docs_gen|migrate|gqlgen|benchmark]"
        ;;
    esac
}

case $1 in
init)
    init
    ;;
infra)
    infra ${@:2}
    ;;
api)
    api ${@:2}
    ;;
*)
    echo "./scripts/bin.sh [infra|api|lint|add_version|test]"
    ;;
esac
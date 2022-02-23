#!/bin/bash

TIMEOUT=15
QUIET=0

ADDRS=()

echoerr() {
  if [ "$QUIET" -ne 1 ]; then printf "%s\n" "$*" 1>&2; fi
}

usage() {
  exitcode="$1"
  cat << USAGE >&2
client:
  $cmdname host:port [host:port] [host:port] [-t timeout] [-- command args]
  -q | --quiet                        Do not output any status messages
  -t TIMEOUT | --timeout=timeout      Timeout in seconds, zero for no timeout
  -- COMMAND ARGS                     Execute command with args after the test finishes
USAGE
  exit "$exitcode"
}

wait_for() {
  results=()
  for addr in ${ADDRS[@]}
  do
    HOST=$(printf "%s\n" "$addr"| cut -d : -f 1)
    PORT=$(printf "%s\n" "$addr"| cut -d : -f 2)
    result=1
    for i in `seq $TIMEOUT` ; do
      nc -z "$HOST" "$PORT" > /dev/null 2>&1
      result=$?
      if [ $result -ne 0 ] ; then
        sleep 1
	continue
      fi
      break
    done
    results=(${results[@]} $result)
  done
  num=${#results[@]}
  for result in ${results[@]}
  do
    if [ $result -eq 0 ] ; then
	    num=`expr $num - 1`
    fi
  done
  if [ $num -eq 0 ] ; then
    if [ $# -gt 0 ] ; then
      exec "$@"
    fi
    exit 0
  fi
  echo "Operation timed out" >&2
  exit 1
}

while [ $# -gt 0 ]
do
  case "$1" in
    *:* )
    ADDRS=(${ADDRS[@]} $1)
    shift 1
    ;;
    -q | --quiet)
    QUIET=1
    shift 1
    ;;
    -t)
    TIMEOUT="$2"
    if [ "$TIMEOUT" = "" ]; then break; fi
    shift 2
    ;;
    --timeout=*)
    TIMEOUT="${1#*=}"
    shift 1
    ;;
    --)
    shift
    break
    ;;
    --help)
    usage 0
    ;;
    *)
    echoerr "Unknown argument: $1"
    usage 1
    ;;
  esac
done

if [ "${#ADDRS[@]}" -eq 0 ]; then
  echoerr "Error: you need to provide a host and port to test."
  usage 2
fi

wait_for "$@"
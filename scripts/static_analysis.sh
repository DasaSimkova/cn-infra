#!/bin/bash

WHITELIST_CONTENT="^// DO NOT EDIT|^// File generated by|^// Automatically generated"
WHITELIST_ERRORS="should not use dot imports"

function static_analysis() {
  local TOOL="${@}"
  local PWD=$(pwd)

  local FILES=$(find "${PWD}" -mount -name "*.go" -type f -not -path "${PWD}/vendor/*" -exec grep -LE "${WHITELIST_CONTENT}"  {} +)

  local DB=$(${TOOL} "${PWD}/db${SELECTOR}")
  local LOGGING=$(${TOOL} "${PWD}/logging${SELECTOR}")
  local UTILS=$(${TOOL} "${PWD}/utils${SELECTOR}")
  local MESSAGING=$(${TOOL} "${PWD}/messaging${SELECTOR}")

  local ALL="$DB
$LOGGING
$UTILS
$MESSAGING
"

  local OUT=$(echo "${ALL}" | grep -F "${FILES}" | grep -v "${WHITELIST_ERRORS}")
  if [[ ! -z $OUT ]] ; then
    echo "${OUT}" 1>&2
    exit 1
  fi
}

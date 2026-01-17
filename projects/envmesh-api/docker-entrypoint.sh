#!/bin/sh

echo Starting envmesh-api...
exec ./main || {
  exitCode=$?
  echo "Failed to start usermanager service."
  echo "Exiting with error code $exitCode"
  exit $exitCode
}
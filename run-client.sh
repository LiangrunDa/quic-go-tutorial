#!/bin/bash

SCRIPTDIR=`dirname "$(readlink -f "$0")"`
# export ENABLEPROFILE=1

if [ "$TESTCASE" == "optimize" ] ; then
    ${SCRIPTDIR}/optimize_client
else
    ${SCRIPTDIR}/client $TESTCASE
fi

retVal=$?
if [ $retVal -eq 0 ]; then
    echo "client exited with code 0"
fi

if [ $retVal -eq 127 ]; then
    echo "exited with code 127"
fi

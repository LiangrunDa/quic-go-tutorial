#!/bin/bash

SCRIPTDIR=`dirname "$(readlink -f "$0")"`

if [ "$TESTCASE" == "optimize" ] ; then
    ${SCRIPTDIR}/optimize_server
else
    ${SCRIPTDIR}/server $TESTCASE
fi

retVal=$?
if [ $retVal -eq 127 ]; then
    echo "exited with code 127"
fi

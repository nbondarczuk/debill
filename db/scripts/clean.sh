#!/bin/bash

if test -d ${PGDATA} -a -d ${PGLOG}
then
	rm -Rf ${PGDATA} ${PGLOG}
	echo Cleaned ${PGDATABASE} in path ${PGDATA}
else
	echo Skip cleanup as ${PGDATA} or ${PGLOG} paths do not exist
	exit 1
fi

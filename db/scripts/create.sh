#!/bin/bash

if test -d ${PGDATA} -o -d ${PGLOG}
then
	echo Cleanup must be done on ${PGDATA|} ${PGLOG}
	exit 1
fi

mkdir -p ${PGDATA} ${PGLOG}
initdb ${PGDATA} -U ${PGUSER}
createdb -U ${PGUSER} ${PGDATABASE}
echo Created ${PGDATABASE} for ${PGUSER} in path ${PGDATA}

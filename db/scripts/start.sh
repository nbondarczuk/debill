#!/bin/bash

DT=$(date '+%Y%m%d%H%M%S')
pg_ctl -D ${PGDATA} -l ${PGLOG}/${PGDATABASE}-${DT}.log start
echo Started ${PGDATABASE} in path ${PGDATA}

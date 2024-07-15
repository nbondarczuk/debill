#!/bin/bash

DT=$(date '+%Y%m%d%H%M%S')
pg_ctl -D ${PGDATA} -l ${PGLOG}/${PGNAME}-${DT}.log start

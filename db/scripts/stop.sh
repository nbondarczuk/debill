#!/bin/bash

pg_ctl -D ${PGDATA} stop
echo Stopped ${PGDATABASE} in path ${PGDATA}

init:
	psql -U ${PGUSER} -d postgres -f init/init.sql -f init/schema.sql -f init/tables.sql

show:
	psql -U ${PGUSER} -d postgres -c '\l'
	psql -U ${PGUSER} -d postgres -c '\dn'
	psql -U ${PGUSER} -d postgres -c '\dt debill.*'

gen:
	gentool -c ./gen/gen.tool

drop:
	psql -U ${PGUSER} -d postgres -c 'DROP SCHEMA debill CASCADE;'

.PHONY: init gen sql

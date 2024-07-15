pg/create:
	make -C ./db create

pg/start:
	make -C ./db start

pg/psql:
	make -C ./db psql

pg/stop:
	make -C ./db stop

pg/status:
	make -C ./db status

pg/drop:
	make -C ./db drop

pg/clean:
	make -C ./db clean

pg/help:
	@echo
	@echo 'Postgres utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make pg/create'
	@echo '    make pg/start'
	@echo '    make pg/psql'
	@echo '    make pg/stop'
	@echo '    make pg/status'
	@echo '    make pg/drop'
	@echo '    make pg/clean'


pg/create:
	make -C ./db create

pg/start:
	make -C ./db start

pg/stop:
	make -C ./db stop

pg/status:
	make -C ./db status

pg/clean: db/stop
	make -C ./db clean

pg/help:
	@echo
	@echo 'Postgres utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make pg/create'


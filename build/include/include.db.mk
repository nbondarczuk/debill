db/create:
	make -C ./db create

db/start:
	make -C ./db start

db/psql:
	make -C ./db psql

db/stop:
	make -C ./db stop

db/status:
	make -C ./db status

db/drop:
	make -C ./db drop

db/clean:
	make -C ./db clean

db/help:
	@echo
	@echo 'Postgres utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make db/create'
	@echo '    make db/start'
	@echo '    make db/psql'
	@echo '    make db/stop'
	@echo '    make db/status'
	@echo '    make db/drop'
	@echo '    make db/clean'


include .env

db/create:
	make -C ./db create

db/start:
	make -C ./db start

db/stop:
	make -C ./db stop

db/status:
	make -C ./db status

db/clean: db/stop
	make -C ./db clean

schema:
	make -C ./schema init

install:
	go install gorm.io/gen/tools/gentool@latest

gen:
	gentool -c ./schema/gen/gen.tool

clean:
	find . -name *~ -exec rm {} \;
	find . -name .DS_Store -exec rm {} \;
	find . -name *.log -exec rm {} \;

.PHONY:	schema install gen clean


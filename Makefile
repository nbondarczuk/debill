schema:
	make -C ./schema/sql

install:
	go install gorm.io/gen/tools/gentool@latest

gen:
	gentool -c ./schema/gen/gen.tool

clean:
	find . -name *~ -exec rm {} \;
	find . -name .DS_Store -exec rm {} \;

.PHONY:	schema install gen clean

-include .env

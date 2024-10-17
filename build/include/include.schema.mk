schema/init:
	make -C ./schema init

schema/show:
	make -C ./schema show

schema/gen:
	make -C ./schema gen

schema/drop:
	make -C ./schema drop

schema/help:
	@echo
	@echo 'Schema utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make schema/init'
	@echo '    make schema/show'
	@echo '    make schema/gen'
	@echo '    make schema/drop'

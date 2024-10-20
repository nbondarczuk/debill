swagger/pull:
	@docker pull quay.io/goswagger/swagger

swagger/version:
	@$(SWAGGER_TOOL_PATH)/swagger.sh version

swagger/validate:
	@$(SWAGGER_TOOL_PATH)/swagger.sh validate

swagger/generate/model:
	@$(SWAGGER_TOOL_PATH)/swagger.sh generate model --skip-validation  -t ../gen/model -A debill-api

swagger/generate/server:
	@$(SWAGGER_TOOL_PATH)/swagger.sh generate server --skip-validation -t ../gen/server -A debill-api

swagger/generate/client:
	@$(SWAGGER_TOOL_PATH)/swagger.sh generate client --skip-validation  -t ../gen/client -A debill-api

swagger/serve:
	@$(SWAGGER_TOOL_PATH)/swagger.sh serve

swagger/help:
	@echo
	@echo 'Swagger utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make swagger/pull'
	@echo '    make swagger/version'
	@echo '    make swagger/validate'
	@echo '    make swagger/generate/model'
	@echo '    make swagger/generate/server'
	@echo '    make swagger/generate/client'
	@echo '    make swagger/serve'

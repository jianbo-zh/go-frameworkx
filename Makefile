#.RECIPEPREFIX = >

app_version := $(app_version)

remote_host := $(remote_host)
remote_port := $(remote_port)
remote_user := $(remote_user)
remote_key := $(remote_key)

dockerfile := $(dockerfile)

deploy_dir := /data/project
upload_dir := /home/server/jianbo/uploads

local_run_dir=./runlocal

#----------------------- docker deploy ---------------------------
checkout.%:
	@./scripts/checkout-version.sh ${*} ${app_version}

upload.%: checkout.%
	@./scripts/copy-to-server.sh ${remote_host} ${remote_port} ${remote_user} ${remote_key} ${upload_dir} ${*} ${app_version}

deploy.%: upload.%
	@ssh -i ${remote_key} -p ${remote_port} ${remote_user}@${remote_host} 'bash -s' < ./scripts/deploy-container.sh ${upload_dir} ${deploy_dir} ${*} ${app_version} ${dockerfile}

restart.%:
	@ssh -i ${remote_key} -p ${remote_port} ${remote_user}@${remote_host} 'bash -s' < ./scripts/restart-container.sh ${upload_dir} ${deploy_dir} ${*} ${app_version}


#----------------------- run local ---------------------------
swagger.%:
	@for module in $$(ls ./cmd/${*}/module/); do \
		if [ -e "./cmd/${*}/module/$$module/swag.go" ]; then \
            swag init -d ./cmd/${*}/module/$$module -g swag.go -o ./cmd/${*}/module/$$module/doc/  --instanceName $$module; \
		fi \
	done

clean.%:
	@rm -rf ${local_run_dir}/${*}/{server,server.log} 

build.%: clean.%
	@mkdir -p ${local_run_dir}/${*}
	@go build -v -o ${local_run_dir}/${*}/server ./cmd/${*}
	@echo "build go binary at: ./cmd/${*}"

trans.%:
	@i18n -extract.dir=./cmd/${*} -output.dir=./cmd/${*}/translations extract
	@i18n -output.dir=./cmd/${*}/translations generator

run.%: trans.% document.% build.%
	@cd ${local_run_dir}/${*}; ./server --conf.type env --conf.file .env --log.level debug

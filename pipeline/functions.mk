define curl
	docker run --rm -v `pwd`:/work appropriate/curl $(1)
endef

define target_env_to_environment
	$(call split_target_env, $(1), 1)
endef

define target_env_to_location
	$(call split_target_env, $(1), 2)
endef

define split_target_env
	$(call word, $(2), $(subst -, ,$(1)) )
endef

define deployer
	$(call curl,sh -c "cd /work; curl --insecure -o $(GPG_KEY_NAME)-private.asc $(DEPLOY_PRIVATE_KEY_URL)"); \
	$(call curl,sh -c "cd /work; curl --insecure -o admin.enc $(DEPLOY_KUBE_CONFIG_URL)"); \
	ls -l $(GPG_KEY_NAME)-private.asc; \
	ls -l admin.enc;

	docker run \
	--name $(APP_NAME)-$(APP_VERSION)-PIPELINE_DEPLOY \
	--rm \
	-e GO_PIPELINE_LABEL=$(GO_PIPELINE_LABEL) \
	-e GPG_SECRET_KEY=/secrets/private-key.asc \
	-e GPG_KEY_NAME=$(GPG_KEY_NAME) \
	-v `pwd`/$(GPG_KEY_NAME)-private.asc:/secrets/private-key.asc:ro \
	-v `pwd`/admin.enc:/secrets/admin.enc:ro \
	$(1)
endef


#Usage: $1 = make target
define pipeline_build
	docker network create test-network || true
	docker volume rm $(BUILD_DIR_VOLUME) || true
	docker volume create -d local-persist -o mountpoint=$(BUILD_DIR) --name=$(BUILD_DIR_VOLUME)
	docker run \
	--name $(APP_NAME)-$(VERSION)-PIPELINE_BUILD \
	--rm \
	--net=test-network \
	-v $(BUILD_DIR):/build:rw \
	-v ~/.gradle:/root/.gradle:rw \
	repo.sns.sky.com:8186/dost/pipeline-build:$(PIPELINE_BUILD_IMAGE_VERSION) \
	$(1)
endef

#Usage: $1 = gpg key name, $2 = app version, $3 = target env, $4 = make target
define pipeline_deploy
	$(call curl,sh -c "cd /work && curl --insecure -o $(1)-private.asc $(DEPLOY_PRIVATE_KEY_URL)")
	$(call curl,sh -c "cd /work && curl --insecure -o admin.enc $(DEPLOY_KUBE_CONFIG_URL)")
	ls -l $(1)-private.asc
	ls -l admin.enc
	docker run \
	--name $(APP_NAME)-$(VERSION)-PIPELINE_DEPLOY \
	--rm \
	-e GPG_KEY_NAME=$(1) \
	-e GO_PIPELINE_LABEL=$(2) \
	-e TARGET_ENV=$(3) \
	-v `pwd`:/deploy:rw \
	-v `pwd`/$(1)-private.asc:/secrets/private-key.asc:ro \
	-v `pwd`/admin.enc:/secrets/admin.enc:ro \
	repo.sns.sky.com:8186/dost/pipeline-deploy:$(PIPELINE_DEPLOY_IMAGE_VERSION) \
	make $(4)
endef

#Usage: $1 = command
define curl
	docker run --rm -v `pwd`:/work appropriate/curl $(1)
endef

#Usage: $1 = input file, $2 output file
define decrypt_secrets
	$(call curl,sh -c "cd /work && curl --insecure -o $(GPG_KEY_NAME)-private.asc $(DEPLOY_PRIVATE_KEY_URL)")
	ls -l $(GPG_KEY_NAME)-private.asc
	docker run \
	--rm \
	-v `pwd`:/work:rw \
	-v `pwd`/$(GPG_KEY_NAME)-private.asc:/tmp/secret.asc \
	-e INPUT_FILE=/work/$(1) \
	-e OUTPUT_FILE=/work/$(2) \
	-e SECRET_KEY=/tmp/secret.asc \
	repo.sns.sky.com:8186/dost/gpg-yaml-decryptor:$(GPG_YAML_DECRYPTOR_VERSION)
endef
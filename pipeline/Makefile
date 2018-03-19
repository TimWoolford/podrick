include functions.mk

APP_NAME=podrick
NAMESPACE=monitoring

TARGET_ENV?=unknown
ARTIFACTORY_USERNAME?=docker-build
ARTIFACTORY_PASSWORD?=unknown
DOCKER_REGISTRY=repo.sns.sky.com:8185/dost

PIPELINE_BUILD_IMAGE_VERSION=0.1.2

GO_PIPELINE_LABEL?=DEV-SNAPSHOT

VERSION=$(GO_PIPELINE_LABEL)

GPG_KEY_NAME?=unknown

DEPLOY_PRIVATE_KEY_URL=https://git.sns.sky.com/dostadmin/dost-private-keys/raw/master/$(GPG_KEY_NAME)-private.asc?private_token=$(GIT_ACCESS_TOKEN)
DEPLOY_KUBE_CONFIG_URL=https://git.sns.sky.com/dost/vdc-kubeconfig/raw/master/$(TARGET_ENV)/admin.enc?private_token=$(GIT_ACCESS_TOKEN)


deploy:
	$(eval ENVIRONMENT = $(call target_env_to_environment, $(TARGET_ENV)))
	$(eval LOCATION = $(call target_env_to_location, $(TARGET_ENV)))
	helm upgrade --install ${APP_NAME} ./${APP_NAME} --version ${VERSION} --namespace ${NAMESPACE} --set "cluster.environment=$(ENVIRONMENT),cluster.location=$(LOCATION)"

quality-gate:
	$(call curl,sh -c "curl --insecure http://podrick.$(TARGET_ENV).bmtapps.bskyb.com/status/$(NAMESPACE)/$(APP_NAME) | grep \"status\" | grep OK")

rollback:
	$(call pipeline_deploy,$(GPG_KEY_NAME),$(GO_PIPELINE_LABEL),$(TARGET_ENV),describe-deployment helm-rollback)

helm-upgrade:
	helm upgrade --install --wait --timeout=300 $(APP_NAME) ./$(APP_NAME)-${VERSION}.tgz --values $(APP_NAME)/_$(TARGET_ENV)-values.yaml -f $(APP_NAME)/_$(CLUSTER)-values-secrets.dec.yaml

package-chart:
	helm package --save=false --version $(VERSION) -d ./ $(APP_NAME)

helm-list:
	helm list

describe-deployment:
	kubectl describe deployment $(APP_NAME) -n utils

helm-rollback:
	helm rollback $(APP_NAME) 0 --wait --timeout=300
	helm history $(APP_NAME)

promote-image: CURL_COMMAND=sh -c "curl -v --silent --output /dev/stderr --write-out \"%{http_code}\" -H \"Content-Type: application/json\" -X POST -d '{\"targetRepo\":\"docker-prod\",\"dockerRepository\":\"dost/$(APP_NAME)\",\"tag\":\"$(VERSION)\",\"copy\":true}' '$(ARTIFACTORY_URL)/api/docker/docker-local/v2/promote' | grep 200"
promote-image:
	$(call curl,$(CURL_COMMAND))

promote-artifacts: CURL_COMMAND=sh -c "curl -v --silent --output /dev/stderr --write-out \"%{http_code}\" -H \"Content-Type: application/json\" -X POST -d '{ \"ciUser\": \"$(ARTIFACTORY_USERNAME)\", \"dryRun\" : false, \"targetRepo\" : \"production-releases-local\" , \"copy\": true , \"artifacts\" : true , \"dependencies\" : false , \"scopes\" : [ \"compile\", \"runtime\" ] , \"failFast\": true }' -u $(ARTIFACTORY_USERNAME):$(ARTIFACTORY_PASSWORD) '$(ARTIFACTORY_URL)/api/build/promote/$(APP_NAME)/$(VERSION)' | grep 200"
promote-artifacts:
	$(call curl,$(CURL_COMMAND))
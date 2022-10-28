# ====================================================================================
# Setup Project

PLATFORMS := linux_amd64 linux_arm64
include ../../../build/makelib/common.mk

# ====================================================================================
#  Options
IMAGE = $(BUILD_REGISTRY)/provider-tencentcloud-controller-$(ARCH)
include ../../../build/makelib/imagelight.mk

# OSBASEIMAGE = scratch
# OSBASEIMAGE = gcr.io/distroless/static@sha256:1f580b0a1922c3e54ae15b0758b5747b260bd99d39d40c2edb3e7f6e2452298b

ifeq ($(origin OSBASEIMAGE),undefined)
OSBASE ?= alpine:3.13
ifeq ($(TARGETARCH),$(filter $(TARGETARCH),amd64 ppc64le))
OSBASEIMAGE = $(OSBASE)
else ifeq ($(TARGETARCH),arm64)
OSBASEIMAGE = arm64v8/$(OSBASE)
else
$(error unsupported architecture $(TARGETARCH))
endif
endif
# ====================================================================================
# Targets

img.build:
	@$(INFO) docker build $(IMAGE)
	@$(MAKE) BUILD_ARGS="--load" img.build.shared
	@$(OK) docker build $(IMAGE)

img.publish:
	@$(INFO) Skipping image publish for $(IMAGE)
	@echo Publish is deferred to xpkg machinery
	@$(OK) Image publish skipped for $(IMAGE)

img.build.shared:
	@$(INFO) docker build $(IMAGE)
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp terraformrc.hcl $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp $(OUTPUT_DIR)/bin/$(OS)_$(ARCH)/provider $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && $(SED_CMD) 's|BASEIMAGE|$(OSBASEIMAGE)|g' Dockerfile || $(FAIL)
	@docker buildx build $(BUILD_ARGS) \
		--platform $(IMAGE_PLATFORMS) \
		--build-arg ARCH=$(ARCH) \
		--build-arg TINI_VERSION=$(TINI_VERSION) \
		--build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) \
		--build-arg TERRAFORM_PROVIDER_SOURCE=$(TERRAFORM_PROVIDER_SOURCE) \
		--build-arg TERRAFORM_PROVIDER_VERSION=$(TERRAFORM_PROVIDER_VERSION) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME=$(TERRAFORM_PROVIDER_DOWNLOAD_NAME) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=$(TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX) \
		--build-arg TERRAFORM_NATIVE_PROVIDER_BINARY=$(TERRAFORM_NATIVE_PROVIDER_BINARY) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)
	@$(OK) docker build $(IMAGE)

img.promote:
	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
	@echo Promote is deferred to xpkg machinery
	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)

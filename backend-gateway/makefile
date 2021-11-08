# -------------------------------------------------------------
# This makefile defines the following targets
#
#   - all (default) - builds all targets and runs all tests/checks
#   - checks - runs all tests/checks
#   - linter - runs all code checks                                                                                                                                                           
#   - clean - cleans the build area
# -------------------------------------------------------------

#local:		本地环境
#joindebug:	联调环境
#dev:		内部测试环境
#beta:		外部测试环境(灰度环境)
#prod:		生产环境
ENV_LOCAL               := local
ENV_JOIN_DEBUG          := joindebug
ENV_DEV                 := dev
ENV_BETA                := beta
ENV_PROD                := prod

# 传递参数定义
RUN_ENV                 :=
export RUN_ENV
PKG_NAME                  := backend-gateway
APP_NAME                  := backend
APP_NAME_CRON_SERVER      := $(APP_NAME)-corn
APP_NAME_GATEWAY          := $(APP_NAME)-gateway
APP_FILE_NAME_GATEWAY     := $(APP_NAME_GATEWAY)
APP_FILE_NAME_CRON_SERVER := $(APP_NAME_CRON_SERVER)
APP_BUILD_TARGET_DIR      := build/bin
# 以下两项编译时需要根据实际情况进行修正
APP_VERSION	            := 1.0.0
APP_ENV	                := $(ENV_LOCAL)

# windows环境时间串有一些差异处理
ifeq ($(OS),Windows_NT)
APP_BUILD                   := $(shell git rev-parse --short HEAD)
APP_BUILD_BUILD_SHORT_TIME  := $(shell echo %date:~5,2%%date:~8,2%%time:~0,2%%time:~3,2%)
APP_BUILD_BUILD_TIME        := $(shell echo %date:~0,4%%date:~5,2%%date:~8,2%%time:~0,2%%time:~3,2%%time:~6,2%)
else
APP_BUILD                   := $(shell git rev-parse --short HEAD)
APP_BUILD_BUILD_SHORT_TIME  := $(shell date "+%m%d%H%M")
APP_BUILD_BUILD_TIME        := $(shell date "+%F %T")
endif

APP_SOURCE			    := ./*.go
APP_TARGET_DIR		    := /usr/local/bin/${APP_NAME}
# 附加参数 用于指定是否输出gdb调试信息
BUILD_EXTRA		        :=

ifeq ($(RUN_ENV), $(ENV_LOCAL))
    APP_ENV     := $(ENV_LOCAL)
else ifeq ($(RUN_ENV),$(ENV_JOIN_DEBUG))
    APP_ENV     := $(ENV_JOIN_DEBUG)
else ifeq ($(RUN_ENV),$(ENV_DEV))
    APP_ENV     := $(ENV_DEV)
else ifeq ($(RUN_ENV),$(ENV_BETA))
    APP_ENV     := $(ENV_BETA)
else ifeq ($(RUN_ENV),$(ENV_PROD))
    APP_ENV     := $(ENV_PROD)
endif

# 处理各类配置名
ifeq ($(ENV_LOCAL),$(APP_ENV))
    APP_FILE_NAME_SUFFIX := -$(APP_ENV)
else ifeq ($(ENV_JOIN_DEBUG),$(APP_ENV))
    APP_FILE_NAME_SUFFIX := -$(APP_ENV)-build-$(APP_BUILD)-$(APP_BUILD_BUILD_SHORT_TIME)
else ifeq ($(ENV_PROD),$(APP_ENV))
    APP_FILE_NAME_SUFFIX := -$(APP_VERSION)-build-$(APP_BUILD)-$(APP_BUILD_BUILD_SHORT_TIME)
    BUILD_EXTRA := -w -s
else ifeq ($(ENV_BETA),$(APP_ENV))
    APP_FILE_NAME_SUFFIX := -$(APP_ENV)-$(APP_VERSION)-build-$(APP_BUILD)-$(APP_BUILD_BUILD_SHORT_TIME)
else ifeq ($(ENV_DEV),$(APP_ENV))
    APP_FILE_NAME_SUFFIX := -$(APP_ENV)-$(APP_VERSION)-build-$(APP_BUILD)-$(APP_BUILD_BUILD_SHORT_TIME)
endif

# windows平台需要在输出文件名后加上exe
ifeq ($(OS),Windows_NT)
    APP_FILE_NAME := $(APP_FILE_NAME).exe
    CLEAN_FILE_NAME := $(CLEAN_FILE_NAME).exe
    RM_CMD := del
else
    RM_CMD := rm
endif

check: linter

linter: 
	@echo "LINT: Running code checks.."
	@echo "Running go vet ..."
	go vet ./server/...
	go vet ./service/...
	go vet ./dao/...
	go vet ./model/...
	go vet ./utils/...
	# go vet ./coin/...
	go vet ./library/...
	@echo "Running gofmt ..."
	@./scripts/gofmt.sh

all: export GOPROXY=https://goproxy.cn
all: check server

server: gateway cron

build/bin/%:
	go build -mod=vendor -ldflags \
	" \
	${BUILD_EXTRA}    \
	-X 'main.appletVersion=${APP_VERSION}' \
	-X 'main.appletBuild=${APP_BUILD}' \
	-X 'main.appletBuildTime=${APP_BUILD_BUILD_TIME}' \
	-X 'main.appletRunEnv=${APP_ENV}' \
	-X  main.appletName=$(@F)-${APP_ENV}   \
	" \
	-o $@${APP_FILE_NAME_SUFFIX} ${SOURCE} \
	${PKG_NAME}/server/$(@F)

.PHONY: gateway
gateway: build/bin/${APP_NAME}-gateway

.PHONY: cron
cron: build/bin/${APP_NAME}-cron

clean:
	$(RM_CMD) -rf ${APP_BUILD_TARGET_DIR}

install:
	mkdir -p ${APP_TARGET_DIR}
	cp ${APP_BUILD_TARGET_DIR}${APP_FILE_NAME} ${APP_TARGET_DIR} -f
	cp ./*.xml ${APP_TARGET_DIR} -f

.PHONY : all clean install ${APP_FILE_NAME}

.PHONY: mockgen
mockgen:
# 在更新和修改dao层函数参数、定义与删除函数时，需执行此命令更新mock
	mockgen -source=./dao/interface.go -destination=./dao/mocks/interface_mock.go -package=mocks
	@echo "重建数据库mocks成功"

.PHONY: apidoc
apidoc:
#	@npm install --registry=https://registry.npm.taobao.org
#	@npm config set strict-ssl false
#	@echo "generating api docs"
	apidoc -i ./server -o ./apidoc/ -f .go --private false

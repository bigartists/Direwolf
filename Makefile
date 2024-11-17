


# 相关文件路径


WEB_SRC_DIR = web

.PHONY: check-nvm
check-nvm:
	@if [ ! -f "$(HOME)/.nvm/nvm.sh" ]; then \
		echo "NVM not found. Please install NVM first." >&2; \
		exit 1; \
	fi



.PHONY: web-release
web-release: check-nvm
	cd $(WEB_SRC_DIR) && \
		. $(HOME)/.nvm/nvm.sh && \
		sleep 1;\
		nvm use v22.11.0 && \
		sleep 1;\
		npm run release



.PHONY: web
web: check-nvm
	cd $(WEB_SRC_DIR) && \
		. $(HOME)/.nvm/nvm.sh && \
		sleep 1;\
		nvm use v22.11.0 && \
		sleep 1;\
		npm run dev

# server run

.PHONY: run
run:
	go run cmd/direwolf/main.go


# 开发环境同时启动前后端
.PHONY: dev
dev:
	@echo "Starting both backend and frontend services..."
	@trap 'kill 0' EXIT; \
	( make run & \
	  sleep 2; \
	  make web & \
	  wait )





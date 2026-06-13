.PHONY: help fmt-check test-basic test-race run-basic-hello run-task-manager run-log-analyzer run-user-center run-blog-api run-compute-server run-compute-client ecommerce-up ecommerce-down

help:
	@echo "Targets:"
	@echo "  fmt-check           Check Go formatting"
	@echo "  test-basic          Run tests that do not require external services"
	@echo "  test-race           Run selected race detector tests"
	@echo "  run-basic-hello     Run module01 hello"
	@echo "  run-task-manager    Run module01 task manager"
	@echo "  run-log-analyzer    Run module02 log analyzer"
	@echo "  run-user-center     Run Gin user center"
	@echo "  run-blog-api        Run GORM blog API"
	@echo "  run-compute-server  Run gRPC compute server"
	@echo "  run-compute-client  Run gRPC compute client"
	@echo "  ecommerce-up        Start ecommerce infra"
	@echo "  ecommerce-down      Stop ecommerce infra"

fmt-check:
	@files="$$(gofmt -l $$(git ls-files '*.go'))"; \
	if [ -n "$$files" ]; then \
		echo "Go files need formatting:"; \
		echo "$$files"; \
		exit 1; \
	fi

test-basic:
	go test ./module01_basics/... ./module02_advanced/... ./module03_web_gin/01_net_basics ./module03_web_gin/07_testing_httptest ./module03_web_gin/project_user_center/internal/service ./module03_web_gin/project_user_center/internal/handler ./module04_gorm/project_blog_api/internal/service

test-race:
	go test -race ./module02_advanced/project_log_analyzer

run-basic-hello:
	go run ./module01_basics/01_hello

run-task-manager:
	go run ./module01_basics/project_task_manager

run-log-analyzer:
	go run ./module02_advanced/project_log_analyzer

run-user-center:
	go run ./module03_web_gin/project_user_center

run-blog-api:
	go run ./module04_gorm/project_blog_api

run-compute-server:
	go run ./module05_grpc/project_distributed_compute/server

run-compute-client:
	go run ./module05_grpc/project_distributed_compute/client

ecommerce-up:
	cd module06_gozero/project_ecommerce && docker compose up -d

ecommerce-down:
	cd module06_gozero/project_ecommerce && docker compose down

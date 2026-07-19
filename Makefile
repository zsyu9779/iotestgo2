.PHONY: help fmt-check test-basic test-race run-basic-hello run-task-manager run-log-analyzer run-user-center run-blog-api run-compute-server run-compute-client ecommerce-up ecommerce-down module01-verify module01-demo-contracts module01-lab-01 module01-lab-02 module01-lab-03 module01-lab-04 module01-integrated-lab module01-homework-solution

help:
	@echo "Targets: fmt-check test-basic test-race module01-verify"

fmt-check:
	@files="$$(gofmt -l $$(git ls-files '*.go'))"; \
	if [ -n "$$files" ]; then echo "Go files need formatting:"; echo "$$files"; exit 1; fi

test-basic:
	go test ./module01_basics/... ./module02_advanced/... ./module03_web_gin/01_net_basics ./module03_web_gin/07_testing_httptest ./module03_web_gin/project_user_center/internal/service ./module03_web_gin/project_user_center/internal/handler ./module04_gorm/project_blog_api/internal/service

test-race:
	go test -race ./module02_advanced/project_log_analyzer

run-basic-hello:
	go run ./module01_basics/blocks/01_go_basics/demo/01_hello

run-task-manager:
	go run ./module01_basics/homework/task_manager/teacher/solution/cmd/taskmanager

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

module01-lab-01:
	go test ./module01_basics/blocks/01_go_basics/lab/solution

module01-lab-02:
	go test ./module01_basics/blocks/02_collections/lab/solution

module01-lab-03:
	go test ./module01_basics/blocks/03_modeling/lab/solution

module01-lab-04:
	go test ./module01_basics/blocks/04_functions_testing/lab/solution

module01-integrated-lab:
	go test ./module01_basics/integrated_lab/scorebook/solution

module01-homework-solution:
	cd module01_basics/homework/task_manager/teacher/solution && ../../student_pack/scripts/grade.sh

module01-verify:
	@test -z "$$(find module01_basics -type f -name '*.go' -exec gofmt -l {} \;)" || (echo "Module 01 存在未格式化 Go 文件" && exit 1)
	go vet ./module01_basics/...
	go test ./module01_basics/...
	$(MAKE) module01-homework-solution

module01-demo-contracts:
	bash module01_basics/instructor/scripts/verify_demo_contracts.sh

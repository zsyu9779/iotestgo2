.PHONY: help fmt-check test-basic test-race run-basic-hello run-task-manager run-log-analyzer run-user-center run-blog-api run-compute-server run-compute-client ecommerce-up ecommerce-down module01-verify module01-demo-contracts module01-teaching-failures module01-audit module01-lab-01 module01-lab-02 module01-lab-03 module01-lab-04 module01-integrated-lab module01-homework-solution module02-lab-01 module02-lab-03 module02-lab-04 module02-integrated-lab module02-homework-solution module02-verify module02-demo-contracts module02-teaching-failures module02-audit module04-env-check module04-verify module04-demo-contracts module04-integration module04-lab module04-audit

help:
	@echo "Targets: fmt-check test-basic test-race module01-verify module02-verify module04-verify module04-audit"

fmt-check:
	@formatted_files="$$(mktemp)"; \
	trap 'rm -f "$$formatted_files"' EXIT; \
	if ! git ls-files -z -- '*.go' ':(exclude)module01_basics/teaching_failures/testdata/**' | xargs -0 gofmt -l > "$$formatted_files"; then \
		echo "gofmt failed while checking tracked Go files" >&2; \
		exit 1; \
	fi; \
	if [ -s "$$formatted_files" ]; then echo "Go files need formatting:"; cat "$$formatted_files"; exit 1; fi

test-basic:
	go test ./module01_basics/... ./module02_advanced/... ./module03_web_gin/01_net_basics ./module03_web_gin/07_testing_httptest ./module03_web_gin/project_user_center/internal/service ./module03_web_gin/project_user_center/internal/handler ./module04_gorm/integrated_lab/blog_api/solution/internal/service

test-race:
	go test -race ./module02_advanced/blocks/... ./module02_advanced/integrated_lab/log_analyzer/solution

run-basic-hello:
	go run ./module01_basics/blocks/01_go_basics/demo/01_hello

run-task-manager:
	go run ./module01_basics/homework/task_manager/teacher/solution/cmd/taskmanager

run-log-analyzer:
	go run ./module02_advanced/integrated_lab/log_analyzer/solution

run-user-center:
	go run ./module03_web_gin/project_user_center

run-blog-api:
	go run ./module04_gorm/integrated_lab/blog_api/solution

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
	@test -z "$$(find module01_basics -path 'module01_basics/teaching_failures/testdata' -prune -o -type f -name '*.go' -exec gofmt -l {} \;)" || (echo "Module 01 存在未格式化 Go 文件" && exit 1)
	go vet ./module01_basics/...
	go test ./module01_basics/...
	$(MAKE) module01-homework-solution

module01-demo-contracts:
	bash module01_basics/instructor/scripts/verify_demo_contracts.sh

module01-teaching-failures:
	bash module01_basics/teaching_failures/verify.sh

module01-audit: module01-verify module01-demo-contracts module01-teaching-failures

module02-lab-01:
	go test ./module02_advanced/blocks/01_interfaces_errors/lab/solution

module02-lab-03:
	go test -race ./module02_advanced/blocks/03_context_concurrency/lab/solution

module02-lab-04:
	go test ./module02_advanced/blocks/04_testing_reflection/lab/solution

module02-integrated-lab:
	go test -race ./module02_advanced/integrated_lab/log_analyzer/solution

module02-homework-solution:
	cd module02_advanced/homework/file_scanner/teacher/solution && ../../student_pack/scripts/grade.sh

module02-demo-contracts:
	bash module02_advanced/scripts/demo_contracts.sh

module02-verify:
	bash module02_advanced/scripts/grade.sh

module02-teaching-failures:
	bash module02_advanced/teaching_failures/verify.sh

module02-audit: module02-verify module02-demo-contracts module02-teaching-failures

module04-env-check:
	go run ./module04_gorm/cmd/envcheck

module04-verify:
	bash module04_gorm/scripts/grade.sh

module04-demo-contracts:
	bash module04_gorm/scripts/demo_contracts.sh

module04-integration:
	go test -tags=integration ./module04_gorm/integration ./module04_gorm/07_testing_mysql ./module04_gorm/integrated_lab/blog_api/solution

module04-lab:
	go test ./module04_gorm/integrated_lab/blog_api/solution/...

module04-audit: module04-verify module04-lab module04-env-check module04-demo-contracts module04-integration

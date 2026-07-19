.PHONY: module01-verify module01-lab-01 module01-lab-02 module01-lab-03 module01-lab-04 module01-integrated-lab module01-homework-solution

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

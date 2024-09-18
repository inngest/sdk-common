.PHONY: check-gen
check-gen:
	@sh ./scripts/check_gen.sh

.PHONY: gen-enums
gen-enums:
	@go run ./scripts/gen_enums

.PHONY: gen
gen:
	@make gen-enums

.PHONY: protobuf
protobuf:
	rm gen/go/v0/*.pb.go
	buf generate

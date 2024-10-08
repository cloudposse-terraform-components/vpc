-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

all: init readme

readme: readme/build

test::
	@echo "🚀 Starting tests..."
	cd test && go test -v ./... -count=1 -timeout=10m
	@echo "✅ All tests passed."

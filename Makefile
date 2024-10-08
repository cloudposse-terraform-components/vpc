-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

all: init readme

readme: readme/build

test::
	@echo "🚀 Starting tests..."
	./test/run.sh
	@echo "✅ All tests passed."

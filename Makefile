.PHONY: coverage test clean

test:
	go test ./... -v

coverage:
	chmod +x test/run_coverage.sh
	./test/run_coverage.sh
	# Convert coverage to lcov format (requires gcov2lcov installed)
	gcov2lcov -infile test/coverage/coverage.out -outfile test/coverage/coverage.lcov
	# Generate HTML report (requires lcov/genhtml installed)
	genhtml test/coverage/coverage.lcov --output-dir coverage-html
	@echo "Coverage report generated in coverage-html/index.html"

clean:
	rm -rf test/coverage
	rm -rf coverage-html
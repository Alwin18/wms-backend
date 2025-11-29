#!/bin/bash

# WMS Backend - Test Coverage Script
# This script runs all tests and generates HTML coverage report

set -e

echo "========================================="
echo "WMS Backend - Running Tests with Coverage"
echo "========================================="
echo ""

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Create coverage directory if not exists
mkdir -p test/coverage

# Run tests with coverage
echo -e "${BLUE}Running tests...${NC}"
go test ./... -coverprofile=test/coverage/coverage.out -covermode=atomic -v

# Check if tests passed
if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}✓ All tests completed!${NC}"
    echo ""
    
    # Check if coverage file exists and has content
    if [ -s test/coverage/coverage.out ]; then
        # Generate HTML coverage report
        echo -e "${BLUE}Generating HTML coverage report...${NC}"
        go tool cover -html=test/coverage/coverage.out -o test/coverage/coverage.html
        
        # Generate coverage summary
        echo -e "${BLUE}Coverage Summary:${NC}"
        go tool cover -func=test/coverage/coverage.out | tail -n 1
        
        echo ""
        echo -e "${GREEN}✓ Coverage report generated!${NC}"
        echo -e "  - Text report: ${YELLOW}test/coverage/coverage.out${NC}"
        echo -e "  - HTML report: ${YELLOW}test/coverage/coverage.html${NC}"
        echo ""
        echo -e "Open HTML report with:"
        echo -e "  ${BLUE}xdg-open test/coverage/coverage.html${NC}"
    else
        echo -e "${YELLOW}⚠ No coverage data generated (tests may have been skipped)${NC}"
    fi
    echo ""
else
    echo -e "${RED}✗ Tests failed!${NC}"
    exit 1
fi

echo "========================================="
echo "Test coverage completed!"
echo "========================================="

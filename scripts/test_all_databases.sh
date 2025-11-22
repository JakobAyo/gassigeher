#!/bin/bash
# Test All Databases Script
# Runs the complete test suite against SQLite, MySQL, and PostgreSQL

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Helper functions
header() {
    echo -e "\n${CYAN}========================================${NC}"
    echo -e "${CYAN}$1${NC}"
    echo -e "${CYAN}========================================${NC}\n"
}

success() {
    echo -e "${GREEN}[OK]${NC} $1"
}

error() {
    echo -e "${RED}[FAIL]${NC} $1"
}

info() {
    echo -e "${CYAN}[INFO]${NC} $1"
}

step() {
    echo -e "${YELLOW}==>${NC} $1"
}

# Parse arguments
SKIP_DOCKER=false
MYSQL_ONLY=false
POSTGRES_ONLY=false
SQLITE_ONLY=false

while [[ $# -gt 0 ]]; do
    case $1 in
        --skip-docker) SKIP_DOCKER=true; shift ;;
        --mysql-only) MYSQL_ONLY=true; shift ;;
        --postgres-only) POSTGRES_ONLY=true; shift ;;
        --sqlite-only) SQLITE_ONLY=true; shift ;;
        *) echo "Unknown option: $1"; exit 1 ;;
    esac
done

header "Gassigeher Multi-Database Test Suite"

# Track results
SQLITE_RESULT=""
MYSQL_RESULT=""
POSTGRES_RESULT=""

# ============================================
# Phase 1: SQLite Tests (Default)
# ============================================

if [ "$MYSQL_ONLY" = false ] && [ "$POSTGRES_ONLY" = false ]; then
    header "Phase 1: Testing with SQLite"
    step "Running all tests with SQLite (in-memory)..."

    export DB_TYPE="sqlite"
    if go test ./... -v > test_results_sqlite.txt 2>&1; then
        SQLITE_RESULT="PASS"
        success "SQLite tests passed"
        SQLITE_COUNT=$(grep -c "^--- PASS:" test_results_sqlite.txt || echo "0")
        info "SQLite: $SQLITE_COUNT individual tests passed"
    else
        SQLITE_RESULT="FAIL"
        error "SQLite tests failed"
        info "See test_results_sqlite.txt for details"
    fi
    unset DB_TYPE
fi

# ============================================
# Phase 2: MySQL Tests (Requires Docker)
# ============================================

if [ "$SQLITE_ONLY" = false ] && [ "$POSTGRES_ONLY" = false ]; then
    header "Phase 2: Testing with MySQL"

    if [ "$SKIP_DOCKER" = false ]; then
        step "Starting MySQL test database via Docker..."

        if docker-compose -f docker-compose.test.yml up -d mysql-test; then
            success "MySQL container started"

            step "Waiting for MySQL to be ready (max 30 seconds)..."
            TIMEOUT=30
            ELAPSED=0
            READY=false

            while [ $ELAPSED -lt $TIMEOUT ]; do
                if docker-compose -f docker-compose.test.yml ps mysql-test | grep -q "healthy"; then
                    READY=true
                    break
                fi
                sleep 2
                ELAPSED=$((ELAPSED + 2))
                echo -n "."
            done
            echo ""

            if [ "$READY" = true ]; then
                success "MySQL is ready"

                export DB_TEST_MYSQL="gassigeher_test:testpass@tcp(localhost:3307)/gassigeher_test?parseTime=true&charset=utf8mb4"

                step "Running all tests with MySQL..."
                if go test ./... -v > test_results_mysql.txt 2>&1; then
                    MYSQL_RESULT="PASS"
                    success "MySQL tests passed"
                    MYSQL_COUNT=$(grep -c "^--- PASS:" test_results_mysql.txt || echo "0")
                    info "MySQL: $MYSQL_COUNT individual tests passed"
                else
                    MYSQL_RESULT="FAIL"
                    error "MySQL tests failed"
                    info "See test_results_mysql.txt for details"
                fi

                unset DB_TEST_MYSQL
            else
                error "MySQL failed to become ready within $TIMEOUT seconds"
                MYSQL_RESULT="FAIL"
            fi
        else
            error "Failed to start MySQL container"
            MYSQL_RESULT="FAIL"
        fi
    else
        info "Skipping MySQL tests (--skip-docker specified)"
    fi
fi

# ============================================
# Phase 3: PostgreSQL Tests (Requires Docker)
# ============================================

if [ "$SQLITE_ONLY" = false ] && [ "$MYSQL_ONLY" = false ]; then
    header "Phase 3: Testing with PostgreSQL"

    if [ "$SKIP_DOCKER" = false ]; then
        step "Starting PostgreSQL test database via Docker..."

        if docker-compose -f docker-compose.test.yml up -d postgres-test; then
            success "PostgreSQL container started"

            step "Waiting for PostgreSQL to be ready (max 20 seconds)..."
            TIMEOUT=20
            ELAPSED=0
            READY=false

            while [ $ELAPSED -lt $TIMEOUT ]; do
                if docker-compose -f docker-compose.test.yml ps postgres-test | grep -q "healthy"; then
                    READY=true
                    break
                fi
                sleep 2
                ELAPSED=$((ELAPSED + 2))
                echo -n "."
            done
            echo ""

            if [ "$READY" = true ]; then
                success "PostgreSQL is ready"

                export DB_TEST_POSTGRES="postgres://gassigeher_test:testpass@localhost:5433/gassigeher_test?sslmode=disable"

                step "Running all tests with PostgreSQL..."
                if go test ./... -v > test_results_postgres.txt 2>&1; then
                    POSTGRES_RESULT="PASS"
                    success "PostgreSQL tests passed"
                    POSTGRES_COUNT=$(grep -c "^--- PASS:" test_results_postgres.txt || echo "0")
                    info "PostgreSQL: $POSTGRES_COUNT individual tests passed"
                else
                    POSTGRES_RESULT="FAIL"
                    error "PostgreSQL tests failed"
                    info "See test_results_postgres.txt for details"
                fi

                unset DB_TEST_POSTGRES
            else
                error "PostgreSQL failed to become ready within $TIMEOUT seconds"
                POSTGRES_RESULT="FAIL"
            fi
        else
            error "Failed to start PostgreSQL container"
            POSTGRES_RESULT="FAIL"
        fi
    else
        info "Skipping PostgreSQL tests (--skip-docker specified)"
    fi
fi

# ============================================
# Summary
# ============================================

header "Test Results Summary"

TOTAL=0
PASSED=0
FAILED=0

if [ -n "$SQLITE_RESULT" ]; then
    TOTAL=$((TOTAL + 1))
    if [ "$SQLITE_RESULT" = "PASS" ]; then
        success "SQLite: PASSED"
        PASSED=$((PASSED + 1))
    else
        error "SQLite: FAILED"
        FAILED=$((FAILED + 1))
    fi
fi

if [ -n "$MYSQL_RESULT" ]; then
    TOTAL=$((TOTAL + 1))
    if [ "$MYSQL_RESULT" = "PASS" ]; then
        success "MySQL: PASSED"
        PASSED=$((PASSED + 1))
    else
        error "MySQL: FAILED"
        FAILED=$((FAILED + 1))
    fi
fi

if [ -n "$POSTGRES_RESULT" ]; then
    TOTAL=$((TOTAL + 1))
    if [ "$POSTGRES_RESULT" = "PASS" ]; then
        success "PostgreSQL: PASSED"
        PASSED=$((PASSED + 1))
    else
        error "PostgreSQL: FAILED"
        FAILED=$((FAILED + 1))
    fi
fi

echo ""
echo -e "${CYAN}Total Databases Tested: $TOTAL${NC}"
echo -e "${GREEN}Passed: $PASSED${NC}"
echo -e "${RED}Failed: $FAILED${NC}"

if [ "$SKIP_DOCKER" = false ]; then
    echo ""
    step "To stop test databases:"
    info "  docker-compose -f docker-compose.test.yml down"
    step "To clean up (remove volumes):"
    info "  docker-compose -f docker-compose.test.yml down -v"
fi

echo ""

# Exit with appropriate code
if [ $FAILED -gt 0 ]; then
    header "Some tests failed - review output above"
    exit 1
else
    header "All database tests passed! ✅"
    exit 0
fi

# Test All Databases Script
# Runs the complete test suite against SQLite, MySQL, and PostgreSQL

param(
    [switch]$SkipDocker = $false,
    [switch]$MySQLOnly = $false,
    [switch]$PostgreSQLOnly = $false,
    [switch]$SQLiteOnly = $false
)

$ErrorActionPreference = "Continue"

# Color output helpers
function Write-Header { param($msg) Write-Host "`n========================================" -ForegroundColor Cyan; Write-Host $msg -ForegroundColor Cyan; Write-Host "========================================`n" -ForegroundColor Cyan }
function Write-Step { param($msg) Write-Host "==> $msg" -ForegroundColor Yellow }
function Write-Success { param($msg) Write-Host "[OK] $msg" -ForegroundColor Green }
function Write-Error { param($msg) Write-Host "[FAIL] $msg" -ForegroundColor Red }
function Write-Info { param($msg) Write-Host "[INFO] $msg" -ForegroundColor Cyan }

Write-Header "Gassigeher Multi-Database Test Suite"

# Test results tracking
$sqliteResult = $null
$mysqlResult = $null
$postgresResult = $null

# ============================================
# Phase 1: SQLite Tests (Default)
# ============================================

if (-not $MySQLOnly -and -not $PostgreSQLOnly) {
    Write-Header "Phase 1: Testing with SQLite"
    Write-Step "Running all tests with SQLite (in-memory)..."

    # SQLite tests (default, no environment variables needed)
    $env:DB_TYPE = "sqlite"
    go test ./... -v > test_results_sqlite.txt 2>&1
    $sqliteResult = $LASTEXITCODE

    if ($sqliteResult -eq 0) {
        Write-Success "SQLite tests passed"
        $sqliteCount = (Select-String -Path "test_results_sqlite.txt" -Pattern "^--- PASS:").Count
        Write-Info "SQLite: $sqliteCount individual tests passed"
    } else {
        Write-Error "SQLite tests failed"
        Write-Info "See test_results_sqlite.txt for details"
    }
}

# ============================================
# Phase 2: MySQL Tests (Requires Docker)
# ============================================

if (-not $SQLiteOnly -and -not $PostgreSQLOnly) {
    Write-Header "Phase 2: Testing with MySQL"

    if (-not $SkipDocker) {
        Write-Step "Starting MySQL test database via Docker..."

        # Start MySQL container
        docker-compose -f docker-compose.test.yml up -d mysql-test
        if ($LASTEXITCODE -ne 0) {
            Write-Error "Failed to start MySQL container"
            $mysqlResult = 1
        } else {
            Write-Success "MySQL container started"

            # Wait for MySQL to be ready
            Write-Step "Waiting for MySQL to be ready (max 30 seconds)..."
            $timeout = 30
            $elapsed = 0
            $ready = $false

            while ($elapsed -lt $timeout) {
                $healthCheck = docker-compose -f docker-compose.test.yml ps mysql-test --format json | ConvertFrom-Json
                if ($healthCheck.Health -eq "healthy") {
                    $ready = $true
                    break
                }
                Start-Sleep -Seconds 2
                $elapsed += 2
                Write-Host "." -NoNewline
            }
            Write-Host ""

            if ($ready) {
                Write-Success "MySQL is ready"

                # Set connection string for tests
                $env:DB_TEST_MYSQL = "gassigeher_test:testpass@tcp(localhost:3307)/gassigeher_test?parseTime=true&charset=utf8mb4"

                Write-Step "Running all tests with MySQL..."
                go test ./... -v > test_results_mysql.txt 2>&1
                $mysqlResult = $LASTEXITCODE

                if ($mysqlResult -eq 0) {
                    Write-Success "MySQL tests passed"
                    $mysqlCount = (Select-String -Path "test_results_mysql.txt" -Pattern "^--- PASS:").Count
                    Write-Info "MySQL: $mysqlCount individual tests passed"
                } else {
                    Write-Error "MySQL tests failed"
                    Write-Info "See test_results_mysql.txt for details"
                }

                # Clean up env var
                Remove-Item Env:DB_TEST_MYSQL -ErrorAction SilentlyContinue
            } else {
                Write-Error "MySQL failed to become ready within $timeout seconds"
                $mysqlResult = 1
            }
        }
    } else {
        Write-Info "Skipping MySQL tests (--SkipDocker specified)"
    }
}

# ============================================
# Phase 3: PostgreSQL Tests (Requires Docker)
# ============================================

if (-not $SQLiteOnly -and -not $MySQLOnly) {
    Write-Header "Phase 3: Testing with PostgreSQL"

    if (-not $SkipDocker) {
        Write-Step "Starting PostgreSQL test database via Docker..."

        # Start PostgreSQL container
        docker-compose -f docker-compose.test.yml up -d postgres-test
        if ($LASTEXITCODE -ne 0) {
            Write-Error "Failed to start PostgreSQL container"
            $postgresResult = 1
        } else {
            Write-Success "PostgreSQL container started"

            # Wait for PostgreSQL to be ready
            Write-Step "Waiting for PostgreSQL to be ready (max 20 seconds)..."
            $timeout = 20
            $elapsed = 0
            $ready = $false

            while ($elapsed -lt $timeout) {
                $healthCheck = docker-compose -f docker-compose.test.yml ps postgres-test --format json | ConvertFrom-Json
                if ($healthCheck.Health -eq "healthy") {
                    $ready = $true
                    break
                }
                Start-Sleep -Seconds 2
                $elapsed += 2
                Write-Host "." -NoNewline
            }
            Write-Host ""

            if ($ready) {
                Write-Success "PostgreSQL is ready"

                # Set connection string for tests
                $env:DB_TEST_POSTGRES = "postgres://gassigeher_test:testpass@localhost:5433/gassigeher_test?sslmode=disable"

                Write-Step "Running all tests with PostgreSQL..."
                go test ./... -v > test_results_postgres.txt 2>&1
                $postgresResult = $LASTEXITCODE

                if ($postgresResult -eq 0) {
                    Write-Success "PostgreSQL tests passed"
                    $postgresCount = (Select-String -Path "test_results_postgres.txt" -Pattern "^--- PASS:").Count
                    Write-Info "PostgreSQL: $postgresCount individual tests passed"
                } else {
                    Write-Error "PostgreSQL tests failed"
                    Write-Info "See test_results_postgres.txt for details"
                }

                # Clean up env var
                Remove-Item Env:DB_TEST_POSTGRES -ErrorAction SilentlyContinue
            } else {
                Write-Error "PostgreSQL failed to become ready within $timeout seconds"
                $postgresResult = 1
            }
        }
    } else {
        Write-Info "Skipping PostgreSQL tests (--SkipDocker specified)"
    }
}

# ============================================
# Summary
# ============================================

Write-Header "Test Results Summary"

$totalTests = 0
$passedTests = 0
$failedTests = 0

if ($null -ne $sqliteResult) {
    $totalTests++
    if ($sqliteResult -eq 0) {
        Write-Success "SQLite: PASSED"
        $passedTests++
    } else {
        Write-Error "SQLite: FAILED"
        $failedTests++
    }
}

if ($null -ne $mysqlResult) {
    $totalTests++
    if ($mysqlResult -eq 0) {
        Write-Success "MySQL: PASSED"
        $passedTests++
    } else {
        Write-Error "MySQL: FAILED"
        $failedTests++
    }
}

if ($null -ne $postgresResult) {
    $totalTests++
    if ($postgresResult -eq 0) {
        Write-Success "PostgreSQL: PASSED"
        $passedTests++
    } else {
        Write-Error "PostgreSQL: FAILED"
        $failedTests++
    }
}

Write-Host ""
Write-Host "Total Databases Tested: $totalTests" -ForegroundColor Cyan
Write-Host "Passed: $passedTests" -ForegroundColor Green
Write-Host "Failed: $failedTests" -ForegroundColor Red

if (-not $SkipDocker) {
    Write-Host ""
    Write-Step "To stop test databases:"
    Write-Info "  docker-compose -f docker-compose.test.yml down"
    Write-Step "To clean up (remove volumes):"
    Write-Info "  docker-compose -f docker-compose.test.yml down -v"
}

Write-Host ""

# Exit with appropriate code
if ($failedTests -gt 0) {
    Write-Header "Some tests failed - review output above"
    exit 1
} else {
    Write-Header "All database tests passed! ✅"
    exit 0
}

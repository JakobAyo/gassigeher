# E2E Test Data Generator
# Wrapper around gentestdata.ps1 for E2E testing
# Usage: .\e2e-tests\gen-e2e-testdata.ps1

param(
    [string]$DatabasePath = ".\e2e-tests\test.db"
)

Write-Host "================================================================" -ForegroundColor Cyan
Write-Host "      Gassigeher E2E Test Data Generator" -ForegroundColor Cyan
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host ""

# Ensure database exists (server should have created it)
if (-not (Test-Path $DatabasePath)) {
    Write-Host "[INFO] Database not found at: $DatabasePath" -ForegroundColor Yellow
    Write-Host "[INFO] Waiting for server to create database..." -ForegroundColor Yellow

    # Wait up to 10 seconds for database to be created
    $waitCount = 0
    while (-not (Test-Path $DatabasePath) -and $waitCount -lt 10) {
        Start-Sleep -Seconds 1
        $waitCount++
    }

    if (-not (Test-Path $DatabasePath)) {
        Write-Host "[ERROR] Database still not found. Please start the server first." -ForegroundColor Red
        exit 1
    }
}

Write-Host "[OK] Found database at: $DatabasePath" -ForegroundColor Green
Write-Host ""

# Call the main gentestdata.ps1 script with E2E database path
Write-Host "[INFO] Calling gentestdata.ps1 with E2E database..." -ForegroundColor Yellow

# Create a temporary env file for E2E testing
$tempEnv = ".\.env.e2e.tmp"
@"
DATABASE_PATH=$DatabasePath
PORT=8080
JWT_SECRET=test-jwt-secret-for-e2e-only
SUPER_ADMIN_EMAIL=admin@test.com
"@ | Out-File -FilePath $tempEnv -Encoding UTF8

try {
    # Call the main script
    & ".\scripts\gentestdata.ps1" -DatabasePath $DatabasePath -EnvFile $tempEnv

    if ($LASTEXITCODE -eq 0) {
        Write-Host ""
        Write-Host "================================================================" -ForegroundColor Green
        Write-Host "      E2E Test Data Generated Successfully!" -ForegroundColor Green
        Write-Host "================================================================" -ForegroundColor Green
        Write-Host ""
        Write-Host "Test Users for E2E Testing:" -ForegroundColor Cyan
        Write-Host "  admin@test.com  (password: test123) - Admin with Orange level" -ForegroundColor White
        Write-Host "  Plus 11 more users with various experience levels" -ForegroundColor White
        Write-Host ""
        Write-Host "Test Dogs:" -ForegroundColor Cyan
        Write-Host "  18 dogs (7 green, 6 blue, 5 orange)" -ForegroundColor White
        Write-Host "  2 are marked as unavailable for testing" -ForegroundColor White
        Write-Host ""
    } else {
        Write-Host "[ERROR] Failed to generate test data" -ForegroundColor Red
        exit 1
    }
} finally {
    # Clean up temp env file
    if (Test-Path $tempEnv) {
        Remove-Item $tempEnv
    }
}

# DONE: E2E test data generation wrapper

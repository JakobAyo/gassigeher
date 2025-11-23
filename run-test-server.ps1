# PowerShell script to run Gassigeher server for E2E testing

$env:DATABASE_PATH = ".\e2e-tests\test.db"
$env:PORT = "8080"
$env:JWT_SECRET = "test-jwt-secret-for-e2e-only"
$env:SUPER_ADMIN_EMAIL = "admin@test.com"
$env:UPLOAD_DIR = ".\e2e-tests\test-uploads"
$env:GMAIL_CLIENT_ID = ""
$env:GMAIL_CLIENT_SECRET = ""
$env:GMAIL_REFRESH_TOKEN = ""
$env:GMAIL_FROM_EMAIL = ""

Write-Host "Starting Gassigeher server for E2E testing..." -ForegroundColor Green
Write-Host "Database: $env:DATABASE_PATH"
Write-Host "Port: $env:PORT"
Write-Host "Super Admin: $env:SUPER_ADMIN_EMAIL"
Write-Host ""
Write-Host "Server will create database automatically on first run"
Write-Host "Press Ctrl+C to stop server"
Write-Host ""

.\gassigeher.exe

# // DONE: PowerShell script to run server with test config

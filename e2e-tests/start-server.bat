@echo off
REM Start Gassigeher server with test environment variables
REM Run this in a separate terminal before running tests

cd ..

set DATABASE_PATH=.\e2e-tests\test.db
set PORT=8080
set JWT_SECRET=test-jwt-secret-for-e2e-only-do-not-use-in-production
set ADMIN_EMAILS=admin@test.com
set UPLOAD_DIR=.\e2e-tests\test-uploads
set GMAIL_CLIENT_ID=
set GMAIL_CLIENT_SECRET=
set GMAIL_REFRESH_TOKEN=
set GMAIL_FROM_EMAIL=

echo Starting Gassigeher server for E2E testing...
echo Database: %DATABASE_PATH%
echo Port: %PORT%
echo Admin: %ADMIN_EMAILS%
echo.
echo Press Ctrl+C to stop server
echo.

gassigeher.exe

REM // DONE: Batch script to start server with test environment

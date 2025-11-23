# How to Fix German Umlauts (ä, ö, ü, ß)

## Problem Summary

German umlauts are displaying as corrupted characters:
- "Müller" appears as "MÃ¼ller"
- "Schäferhund" appears as "SchÃ¤ferhund"
- "gehört" appears as "gehÃ¶rt"

## Root Causes Found

1. **HTTP Response Headers**: Missing `charset=utf-8` declaration
2. **PowerShell Script**: Test data generator had encoding issues
3. **Database**: Contains already-corrupted data that needs regeneration

## Complete Fix (Follow ALL Steps)

### Step 1: Verify Code Fixes (Already Done ✅)

The following files have been fixed:

**File: `internal/handlers/auth_handler.go` (line 441)**
```go
// OLD: w.Header().Set("Content-Type", "application/json")
// NEW:
w.Header().Set("Content-Type", "application/json; charset=utf-8")
```

**File: `scripts/gentestdata.ps1` (lines 10-13)**
```powershell
# Added UTF-8 encoding initialization
$OutputEncoding = [System.Text.Encoding]::UTF8
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$PSDefaultParameterValues['*:Encoding'] = 'utf8'
```

**File: `scripts/gentestdata.ps1` (lines 375-377)**
```powershell
# Changed from Out-File to proper UTF-8 without BOM
$utf8NoBom = New-Object System.Text.UTF8Encoding($false)
$sqlFilePath = Join-Path $PSScriptRoot "..\$sqlFile"
[System.IO.File]::WriteAllText($sqlFilePath, $sql.ToString(), $utf8NoBom)
```

### Step 2: Regenerate Database with Correct Data

**CRITICAL**: The current database has corrupted data. You MUST regenerate it.

```powershell
# Open PowerShell in project root
cd C:\Users\tranm\work\gassigeher

# Run script with UTF-8 encoding
powershell.exe -ExecutionPolicy Bypass -Command "[Console]::OutputEncoding = [System.Text.Encoding]::UTF8; .\scripts\gentestdata.ps1"
```

**Expected Output:**
```
================================================================
         Test Data Generation Complete!
================================================================

Summary:
  Users:                 12 (1 admin, 1 inactive)
  Dogs:                  18 (2 unavailable)
  Bookings:              ~90 (spanning 28 days)
  Blocked Dates:         3
  Experience Requests:   4 (2 pending, 1 approved, 1 denied)
```

### Step 3: Rebuild Application

```bash
# In project root
go build -o gassigeher.exe ./cmd/server
```

### Step 4: Restart Server

```bash
# Stop old server first (Ctrl+C in terminal where it's running)
# Or kill process:
taskkill /F /IM gassigeher.exe

# Start new server
.\gassigeher.exe
```

### Step 5: Clear Browser Cache

**Option A - Hard Refresh (Quick):**
- Press `Ctrl+Shift+R` (Chrome/Edge)
- Or `Ctrl+F5`

**Option B - DevTools (Thorough):**
1. Press `F12` to open DevTools
2. Right-click the **Refresh** button
3. Select **"Empty Cache and Hard Reload"**

**Option C - Incognito/Private (Guaranteed):**
- Open new Incognito/Private window
- Navigate to `http://localhost:8080`

### Step 6: Verify Fix

1. **Login** to application (password: `test123`)
2. **Check admin-users.html** → Should see "Müller" (not "MÃ¼ller")
3. **Check dashboard** → Notes should show "gehört" (not "gehÃ¶rt")
4. **Open DevTools** → Network tab → Check Response Headers
   - Should see: `Content-Type: application/json; charset=utf-8`

## Verification Commands

**Check database directly:**
```bash
cd C:\Users\tranm\work\gassigeher
sqlite3 gassigeher.db "SELECT DISTINCT breed FROM dogs;"
```

**Expected output (correct):**
```
Beagle
Boxer
Schäferhund  ← Should be correct ä, not Ã¤
```

**Check current data (wrong):**
```
SchÃ¤ferhund  ← This means database needs regeneration
```

## Common Mistakes

❌ **Skipping database regeneration** → Umlauts still broken (data is corrupted)
❌ **Not clearing browser cache** → Seeing old cached responses
❌ **Not restarting server** → Old binary still running without charset fix
❌ **Running PowerShell script without UTF-8** → Generates corrupted data again

## Why This Happened

1. **PowerShell Default Encoding**: Windows PowerShell defaults to the console code page (usually Windows-1252 or similar), not UTF-8
2. **SQLite3 CLI**: When executing SQL files, character encoding depends on console settings
3. **HTTP Headers**: Without explicit charset, browsers guess encoding, often incorrectly
4. **Cascading Effect**: Bad data in DB → sent over HTTP without charset → browser misinterprets → user sees garbage

## Files Modified

- ✅ `internal/handlers/auth_handler.go` - Added charset=utf-8 to responses
- ✅ `scripts/gentestdata.ps1` - Fixed PowerShell UTF-8 encoding
- 📋 `docs/BugsFromManualTests.md` - Documented the fix
- 📋 `FIX_UMLAUTS_INSTRUCTIONS.md` - This file (instructions)

## Test Credentials

After regenerating database:
- **Super Admin**: Use email from `.env` file (`SUPER_ADMIN_EMAIL`), password: `test123`
- **Users**: Any generated user email, password: `test123`

## Need Help?

If umlauts still appear broken after following ALL steps:

1. **Check Response Headers**:
   - Open DevTools (F12)
   - Network tab
   - Click on any `/api/users` or `/api/dogs` request
   - Check Response Headers for `Content-Type`
   - Should show: `application/json; charset=utf-8`

2. **Check Database**:
   ```bash
   sqlite3 gassigeher.db "SELECT name FROM users WHERE name LIKE '%ü%' OR name LIKE '%ö%' OR name LIKE '%ä%';"
   ```
   - Should show correct umlauts, not "Ã¼", "Ã¶", "Ã¤"

3. **Check Browser**:
   - Try Incognito/Private window
   - Check browser's Character Encoding setting (should be Auto-detect or UTF-8)

---

**Status**: ✅ Code fixes applied, ready for database regeneration and testing
**Last Updated**: 2025-11-21

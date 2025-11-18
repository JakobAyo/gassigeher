# Gassigeher E2E Tests

End-to-end tests using Playwright for the Gassigeher dog walking booking system.

## Quick Start

```bash
# Install dependencies
cd e2e-tests
npm install

# Install Playwright browsers
npm run install:browsers

# Run all tests (headed mode - see browser)
npm run test:headed

# Run all tests (headless mode)
npm test

# Run specific test file
npx playwright test tests/01-public-pages.spec.js

# Debug a test
npm run test:debug

# View test report
npm run report
```

## Test Structure

```
e2e-tests/
├── tests/                      # Test specifications
│   ├── 01-public-pages.spec.js
│   ├── 02-authentication.spec.js
│   └── 03-user-profile.spec.js
├── pages/                      # Page Object Model
│   ├── BasePage.js
│   ├── LoginPage.js
│   ├── RegisterPage.js
│   └── DashboardPage.js
├── fixtures/                   # Test fixtures
│   ├── database.js
│   └── auth.js
├── utils/                      # Utility helpers
│   ├── db-helpers.js
│   └── german-text.js
├── playwright.config.js        # Playwright configuration
├── global-setup.js             # One-time setup
└── global-teardown.js          # Cleanup
```

## Configuration

Tests are configured to:
- Run against `http://localhost:8080`
- Use separate test database (`test.db`)
- Start Go server automatically
- Run in headed mode for debugging
- Capture screenshots/videos on failure

## Test Data

Test data is automatically seeded in `global-setup.js`:
- **Users**: green@test.com, blue@test.com, orange@test.com, admin@test.com
- **Password**: test123 (for all users)
- **Dogs**: 9 dogs (3 green, 3 blue, 3 orange)

## Running Tests

### All Tests
```bash
npm test                    # Headless, chromium only
npm run test:headed         # Headed (see browser)
npm run test:chrome         # Desktop Chrome only
npm run test:mobile         # Mobile viewports
```

### Specific Tests
```bash
# Run one file
npx playwright test tests/01-public-pages.spec.js

# Run tests matching pattern
npx playwright test -g "login"

# Run specific test
npx playwright test tests/02-authentication.spec.js:42
```

### Debug Mode
```bash
# Interactive UI mode (best for development)
npm run test:ui

# Step-through debug mode
npm run test:debug

# Debug specific test
npx playwright test tests/02-authentication.spec.js --debug
```

## Viewing Results

```bash
# Open HTML report
npm run report

# View trace for failed test
npx playwright show-trace trace.zip
```

## Writing New Tests

1. Create test file in `tests/`
2. Import Page Objects from `pages/`
3. Use `test.describe()` for grouping
4. Use `test()` for individual tests
5. Follow naming: `should [action] [expected result]`

Example:
```javascript
const { test, expect } = require('@playwright/test');
const LoginPage = require('../pages/LoginPage');

test.describe('Feature Name', () => {
  test('should do something successfully', async ({ page }) => {
    const loginPage = new LoginPage(page);
    await loginPage.goto();
    await loginPage.login('test@example.com', 'password');

    expect(page.url()).toContain('dashboard.html');
  });
});
```

## Debugging Tips

1. **Run in headed mode** to see what's happening:
   ```bash
   npm run test:headed
   ```

2. **Use console.log** to inspect values:
   ```javascript
   console.log('Current URL:', page.url());
   ```

3. **Take screenshots**:
   ```javascript
   await page.screenshot({ path: 'debug.png' });
   ```

4. **Use Playwright Inspector**:
   ```bash
   npm run test:debug
   ```

5. **Check test artifacts** in `test-results/` folder

## Common Issues

### Server not starting
- Make sure Go app is built: `go build -o gassigeher.exe ./cmd/server`
- Check port 8080 is not in use

### Database errors
- Delete `test.db` and restart tests
- Check `global-setup.js` runs successfully

### Tests timing out
- Increase timeout in `playwright.config.js`
- Check network tab in headed mode

### Flaky tests
- Add `await page.waitForLoadState('networkidle')`
- Use `await expect(locator).toBeVisible()` instead of `isVisible()`

## Status

**Phase 1 Complete**: ✅ Foundation + First 3 Test Files
- 01-public-pages.spec.js (15 tests)
- 02-authentication.spec.js (20+ tests)
- 03-user-profile.spec.js (15+ tests)

**Total Tests**: 50+ tests covering public pages, authentication, and user profiles

## Next Steps

- Add more test files (dogs, bookings, calendar, admin)
- Run tests to find bugs
- Document bugs found
- Fix bugs and re-run tests

---

See [E2ETestingPlan.md](../E2ETestingPlan.md) for complete testing strategy.

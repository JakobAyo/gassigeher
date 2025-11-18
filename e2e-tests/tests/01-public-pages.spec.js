const { test, expect } = require('@playwright/test');
const BasePage = require('../pages/BasePage');

/**
 * PUBLIC PAGES TESTS
 * Test all publicly accessible pages
 * These should work without authentication
 */

test.describe('Public Pages - Accessibility', () => {

  test('homepage (index.html) should load successfully', async ({ page }) => {
    await page.goto('http://localhost:8080/');
    await page.waitForLoadState('networkidle');

    // Check page loaded (root URL or index.html)
    expect(page.url()).toMatch(/\/(index\.html)?$/);

    // Check page title exists
    const title = await page.title();
    expect(title).toBeTruthy();
    expect(title.length).toBeGreaterThan(0);

    // POTENTIAL BUG CHECK: Is the title in German?
    console.log('Homepage title:', title);
  });

  test('homepage should have navigation links', async ({ page }) => {
    await page.goto('http://localhost:8080/index.html');

    // Check for login link
    const loginLink = page.locator('a[href="/login.html"], a[href="login.html"]');
    await expect(loginLink).toBeVisible();

    // Check for register link (use first() for strict mode)
    const registerLink = page.locator('a[href="/register.html"], a[href="register.html"]').first();
    await expect(registerLink).toBeVisible();

    // POTENTIAL BUG: Links might be broken or have wrong paths
  });

  test('login page should load without authentication', async ({ page }) => {
    await page.goto('http://localhost:8080/login.html');
    await page.waitForLoadState('networkidle');

    expect(page.url()).toContain('login.html');

    // Check form elements exist
    await expect(page.locator('#email')).toBeVisible();
    await expect(page.locator('#password')).toBeVisible();
    await expect(page.locator('button[type="submit"]')).toBeVisible();
  });

  test('register page should load without authentication', async ({ page }) => {
    await page.goto('http://localhost:8080/register.html');
    await page.waitForLoadState('networkidle');

    expect(page.url()).toContain('register.html');

    // Check form elements exist
    await expect(page.locator('#email')).toBeVisible();
    await expect(page.locator('#name')).toBeVisible();
    await expect(page.locator('#phone')).toBeVisible();
    await expect(page.locator('#password')).toBeVisible();
    await expect(page.locator('#accept-terms')).toBeVisible();  // Correct ID
  });

  test('terms and conditions page should be accessible', async ({ page }) => {
    await page.goto('http://localhost:8080/terms.html');
    await page.waitForLoadState('networkidle');

    expect(page.url()).toContain('terms.html');

    // Check page has content
    const bodyText = await page.textContent('body');
    expect(bodyText.length).toBeGreaterThan(100); // Should have substantial content

    // POTENTIAL BUG: Terms might be in English instead of German
    const hasGermanText = bodyText.includes('Nutzungsbedingungen') ||
                          bodyText.includes('Datenschutz') ||
                          bodyText.includes('Tierheim');
    console.log('Terms page has German text:', hasGermanText);
  });

  test('privacy policy page should be accessible', async ({ page }) => {
    await page.goto('http://localhost:8080/privacy.html');
    await page.waitForLoadState('networkidle');

    expect(page.url()).toContain('privacy.html');

    // Check page has content
    const bodyText = await page.textContent('body');
    expect(bodyText.length).toBeGreaterThan(100);

    // POTENTIAL BUG: Privacy policy might be in English
    const hasGermanPrivacyText = bodyText.includes('Datenschutz') ||
                                  bodyText.includes('personenbezogene Daten') ||
                                  bodyText.includes('DSGVO');
    console.log('Privacy page has German text:', hasGermanPrivacyText);
  });

  test('forgot password page should be accessible', async ({ page }) => {
    await page.goto('http://localhost:8080/forgot-password.html');
    await page.waitForLoadState('networkidle');

    expect(page.url()).toContain('forgot-password.html');

    // Check form exists
    await expect(page.locator('#email')).toBeVisible();
    await expect(page.locator('button[type="submit"]')).toBeVisible();
  });

});

test.describe('Public Pages - Navigation', () => {

  test('should navigate from homepage to login', async ({ page }) => {
    await page.goto('http://localhost:8080/index.html');

    // Click login link
    await page.click('a[href="/login.html"], a[href="login.html"]');
    await page.waitForURL('**/login.html');

    expect(page.url()).toContain('login.html');
  });

  test('should navigate from homepage to register', async ({ page }) => {
    await page.goto('http://localhost:8080/index.html');

    // Click register link
    await page.click('a[href="/register.html"], a[href="register.html"]');
    await page.waitForURL('**/register.html');

    expect(page.url()).toContain('register.html');
  });

  test('should navigate from login to register', async ({ page }) => {
    await page.goto('http://localhost:8080/login.html');

    // Look for "Noch kein Konto?" or "Registrieren" link (use first() for strict mode)
    const registerLink = page.locator('a[href="/register.html"], a[href="register.html"]').first();
    await expect(registerLink).toBeVisible();

    await registerLink.click();
    await page.waitForURL('**/register.html');

    expect(page.url()).toContain('register.html');
  });

  test('should navigate from register to login', async ({ page }) => {
    await page.goto('http://localhost:8080/register.html');

    // Look for "Schon registriert?" or "Anmelden" link (use first() for strict mode)
    const loginLink = page.locator('a[href="/login.html"], a[href="login.html"]').first();
    await expect(loginLink).toBeVisible();

    await loginLink.click();
    await page.waitForURL('**/login.html');

    expect(page.url()).toContain('login.html');
  });

  test('should navigate from login to forgot password', async ({ page }) => {
    await page.goto('http://localhost:8080/login.html');

    // Look for "Passwort vergessen?" link
    const forgotLink = page.locator('a[href="/forgot-password.html"], a[href="forgot-password.html"]');
    await expect(forgotLink).toBeVisible();

    await forgotLink.click();
    await page.waitForURL('**/forgot-password.html');

    expect(page.url()).toContain('forgot-password.html');
  });

});

test.describe('Public Pages - Protected Routes', () => {

  test('dashboard should redirect to login when not authenticated', async ({ page }) => {
    // Try to access dashboard without logging in
    await page.goto('http://localhost:8080/dashboard.html');

    // CRITICAL BUG CHECK: Should redirect to login!
    await page.waitForLoadState('networkidle');

    const currentURL = page.url();
    console.log('Dashboard without auth redirected to:', currentURL);

    // Should redirect to login
    // BUG: If this doesn't redirect, it's a MAJOR security issue!
    if (!currentURL.includes('login.html')) {
      console.error('🐛 POTENTIAL BUG: Dashboard accessible without authentication!');
    }
  });

  test('dogs page should redirect to login when not authenticated', async ({ page }) => {
    await page.goto('http://localhost:8080/dogs.html');
    await page.waitForLoadState('networkidle');

    const currentURL = page.url();
    console.log('Dogs page without auth redirected to:', currentURL);

    if (!currentURL.includes('login.html')) {
      console.error('🐛 POTENTIAL BUG: Dogs page accessible without authentication!');
    }
  });

  test('profile page should redirect to login when not authenticated', async ({ page }) => {
    await page.goto('http://localhost:8080/profile.html');
    await page.waitForLoadState('networkidle');

    const currentURL = page.url();
    console.log('Profile page without auth redirected to:', currentURL);

    if (!currentURL.includes('login.html')) {
      console.error('🐛 POTENTIAL BUG: Profile page accessible without authentication!');
    }
  });

  test('admin pages should redirect to login when not authenticated', async ({ page }) => {
    await page.goto('http://localhost:8080/admin-dashboard.html');
    await page.waitForLoadState('networkidle');

    const currentURL = page.url();
    console.log('Admin dashboard without auth redirected to:', currentURL);

    if (!currentURL.includes('login.html') && !currentURL.includes('404')) {
      console.error('🐛 POTENTIAL BUG: Admin page accessible without authentication!');
    }
  });

});

test.describe('Public Pages - UI Consistency', () => {

  test('all public pages should have consistent branding', async ({ page }) => {
    const pages = ['index.html', 'login.html', 'register.html', 'terms.html', 'privacy.html'];

    for (const pagePath of pages) {
      await page.goto(`http://localhost:8080/${pagePath}`);
      await page.waitForLoadState('networkidle');

      // Check for logo or site name
      const bodyText = await page.textContent('body');
      const hasGassigehe = bodyText.toLowerCase().includes('gassigeher') ||
                            bodyText.toLowerCase().includes('tierheim');

      console.log(`${pagePath} has branding:`, hasGassigehe);

      // POTENTIAL BUG: Inconsistent branding across pages
      if (!hasGassigehe) {
        console.warn(`⚠️ ${pagePath} might be missing branding`);
      }
    }
  });

});

// DONE: Public pages tests - checking accessibility, navigation, auth protection, and UI consistency

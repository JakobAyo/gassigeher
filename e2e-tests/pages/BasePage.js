/**
 * Base Page Object
 * Contains common methods used across all pages
 */
class BasePage {
  constructor(page) {
    this.page = page;
    this.baseURL = 'http://localhost:8080';
  }

  /**
   * Navigate to a path
   */
  async goto(path) {
    await this.page.goto(`${this.baseURL}${path}`);
    await this.page.waitForLoadState('networkidle');
  }

  /**
   * Wait for navigation to complete
   */
  async waitForNavigation() {
    await this.page.waitForLoadState('networkidle');
  }

  /**
   * Get alert message text
   * @param {string} type - alert type (success, danger, warning, info)
   */
  async getAlertText(type = 'success') {
    const selector = `.alert-${type}`;
    await this.page.waitForSelector(selector, { timeout: 5000 });
    return await this.page.textContent(selector);
  }

  /**
   * Check if alert exists
   */
  async hasAlert(type = 'success') {
    const selector = `.alert-${type}`;
    return await this.page.locator(selector).isVisible().catch(() => false);
  }

  /**
   * Wait for alert to appear
   */
  async waitForAlert(type = 'success', timeout = 5000) {
    const selector = `.alert-${type}`;
    await this.page.waitForSelector(selector, { timeout });
  }

  /**
   * Click navigation link by text
   */
  async clickNavLink(text) {
    await this.page.click(`nav a:has-text("${text}")`);
    await this.waitForNavigation();
  }

  /**
   * Check if user is logged in
   */
  async isLoggedIn() {
    // Check if dashboard link visible or logout button exists
    const dashboardLink = this.page.locator('a[href="/dashboard.html"]');
    const logoutLink = this.page.locator('a:has-text("Abmelden")');

    const hasDashboard = await dashboardLink.isVisible().catch(() => false);
    const hasLogout = await logoutLink.isVisible().catch(() => false);

    return hasDashboard || hasLogout;
  }

  /**
   * Get current URL
   */
  async getCurrentURL() {
    return this.page.url();
  }

  /**
   * Wait for URL to match pattern
   */
  async waitForURL(pattern, timeout = 5000) {
    await this.page.waitForURL(pattern, { timeout });
  }

  /**
   * Get page title
   */
  async getTitle() {
    return await this.page.title();
  }

  /**
   * Take screenshot (useful for debugging)
   */
  async screenshot(name) {
    await this.page.screenshot({ path: `screenshots/${name}.png` });
  }

  /**
   * Fill form field
   */
  async fill(selector, value) {
    await this.page.fill(selector, value);
  }

  /**
   * Click element
   */
  async click(selector) {
    await this.page.click(selector);
  }

  /**
   * Check if element is visible
   */
  async isVisible(selector) {
    return await this.page.locator(selector).isVisible().catch(() => false);
  }

  /**
   * Get element text content
   */
  async textContent(selector) {
    return await this.page.textContent(selector);
  }

  /**
   * Wait for element
   */
  async waitForSelector(selector, timeout = 10000) {
    await this.page.waitForSelector(selector, { timeout });
  }
}

module.exports = BasePage;

// DONE: Base page object with common methods

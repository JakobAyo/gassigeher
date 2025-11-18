const BasePage = require('./BasePage');

/**
 * Login Page Object
 */
class LoginPage extends BasePage {
  constructor(page) {
    super(page);

    // Selectors
    this.emailInput = '#email';
    this.passwordInput = '#password';
    this.submitButton = 'button[type="submit"]';
    this.errorAlert = '.alert-error';  // Corrected: CSS uses alert-error not alert-danger
    this.successAlert = '.alert-success';
    this.registerLink = 'a[href="/register.html"]';
    this.forgotPasswordLink = 'a[href="/forgot-password.html"]';
  }

  /**
   * Navigate to login page
   */
  async goto() {
    await super.goto('/login.html');
  }

  /**
   * Login with credentials
   */
  async login(email, password) {
    await this.page.fill(this.emailInput, email);
    await this.page.fill(this.passwordInput, password);
    await this.page.click(this.submitButton);
  }

  /**
   * Login and wait for redirect to dashboard
   */
  async loginAndWait(email, password) {
    await this.login(email, password);
    // Login has 1-second delay before redirect, so wait longer
    await this.page.waitForURL('**/dashboard.html', { timeout: 10000 });
  }

  /**
   * Get error message
   */
  async getErrorMessage() {
    await this.page.waitForSelector(this.errorAlert, { timeout: 3000 });
    return await this.page.textContent(this.errorAlert);
  }

  /**
   * Check if error is visible
   */
  async hasError() {
    return await this.page.locator(this.errorAlert).isVisible().catch(() => false);
  }

  /**
   * Click register link
   */
  async goToRegister() {
    await this.page.click(this.registerLink);
    await this.page.waitForURL('**/register.html');
  }

  /**
   * Click forgot password link
   */
  async goToForgotPassword() {
    await this.page.click(this.forgotPasswordLink);
    await this.page.waitForURL('**/forgot-password.html');
  }
}

module.exports = LoginPage;

// DONE: Login page object with login methods

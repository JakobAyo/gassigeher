const BasePage = require('./BasePage');

/**
 * Register Page Object
 */
class RegisterPage extends BasePage {
  constructor(page) {
    super(page);

    // Selectors
    this.emailInput = '#email';
    this.nameInput = '#name';
    this.phoneInput = '#phone';
    this.passwordInput = '#password';
    this.confirmPasswordInput = '#confirm-password';
    this.termsCheckbox = '#accept-terms';  // Correct ID from register.html
    this.submitButton = 'button[type="submit"]';
    this.errorAlert = '.alert-error';  // Corrected: CSS uses alert-error not alert-danger
    this.successAlert = '.alert-success';
    this.loginLink = 'a[href="/login.html"]';
  }

  /**
   * Navigate to register page
   */
  async goto() {
    await super.goto('/register.html');
  }

  /**
   * Fill registration form
   */
  async fillForm({ email, name, phone, password, acceptTerms = true }) {
    if (email) await this.page.fill(this.emailInput, email);
    if (name) await this.page.fill(this.nameInput, name);
    if (phone) await this.page.fill(this.phoneInput, phone);
    if (password) await this.page.fill(this.passwordInput, password);

    if (acceptTerms) {
      await this.page.check(this.termsCheckbox);
    }
  }

  /**
   * Submit registration form
   */
  async submit() {
    await this.page.click(this.submitButton);
  }

  /**
   * Register user with all fields
   */
  async register({ email, name, phone, password, acceptTerms = true }) {
    await this.fillForm({ email, name, phone, password, acceptTerms });
    await this.submit();
  }

  /**
   * Get error message
   */
  async getErrorMessage() {
    await this.page.waitForSelector(this.errorAlert, { timeout: 3000 });
    return await this.page.textContent(this.errorAlert);
  }

  /**
   * Get success message
   */
  async getSuccessMessage() {
    await this.page.waitForSelector(this.successAlert, { timeout: 3000 });
    return await this.page.textContent(this.successAlert);
  }

  /**
   * Check if error is visible
   */
  async hasError() {
    return await this.page.locator(this.errorAlert).isVisible().catch(() => false);
  }

  /**
   * Check if success is visible
   */
  async hasSuccess() {
    return await this.page.locator(this.successAlert).isVisible().catch(() => false);
  }

  /**
   * Go to login page
   */
  async goToLogin() {
    await this.page.click(this.loginLink);
    await this.page.waitForURL('**/login.html');
  }
}

module.exports = RegisterPage;

// DONE: Register page object with form filling and submission

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const { setupAdminAuth } = require('./fixtures/auth');

/**
 * Global setup for E2E tests
 * Runs once before all tests
 * Uses existing gentestdata.ps1 script for realistic test data
 */
module.exports = async (config) => {
  console.log('');
  console.log('═══════════════════════════════════════════════════');
  console.log('🚀 Global Setup: Preparing E2E Test Environment');
  console.log('═══════════════════════════════════════════════════');
  console.log('');

  try {
    const testDbPath = path.resolve(__dirname, 'test.db');

    // Step 1: Delete existing test database
    console.log('📦 Setting up test database...');
    if (fs.existsSync(testDbPath)) {
      fs.unlinkSync(testDbPath);
      console.log('   ✅ Deleted existing test.db');
    }

    // Step 2: Wait for server to create database
    console.log('⏳ Waiting for server to create database...');
    let waitCount = 0;
    while (!fs.existsSync(testDbPath) && waitCount < 15) {
      await new Promise(resolve => setTimeout(resolve, 1000));
      waitCount++;
    }

    if (!fs.existsSync(testDbPath)) {
      throw new Error('Server did not create test database after 15 seconds');
    }
    console.log('   ✅ Database created by server');

    // Step 3: Generate test data using existing PowerShell script
    console.log('🌱 Generating test data using gentestdata.ps1...');
    console.log('');

    const scriptPath = path.resolve(__dirname, '../scripts/gentestdata.ps1');
    const command = `powershell -ExecutionPolicy Bypass -File "${scriptPath}" -DatabasePath "${testDbPath}"`;

    try {
      execSync(command, {
        stdio: 'inherit',
        cwd: path.resolve(__dirname, '..'),
        env: {
          ...process.env,
          ADMIN_EMAILS: 'admin@test.com',
        }
      });
      console.log('');
      console.log('   ✅ Test data generated successfully');
    } catch (error) {
      console.error('   ❌ Failed to generate test data:', error.message);
      throw error;
    }

    // Step 4: Pre-authenticate admin user
    console.log('🔐 Pre-authenticating admin user...');
    await setupAdminAuth();

    console.log('');
    console.log('✅ Global setup complete!');
    console.log('═══════════════════════════════════════════════════');
    console.log('');

  } catch (error) {
    console.error('');
    console.error('❌ Global setup failed:', error.message);
    console.error('═══════════════════════════════════════════════════');
    console.error('');
    throw error;
  }
};

// DONE: Global setup updated to use existing gentestdata.ps1 script


// DONE: Global setup runs once before all tests

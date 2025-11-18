const { cleanupDatabase } = require('./fixtures/database');

/**
 * Global teardown for E2E tests
 * Runs once after all tests complete
 */
module.exports = async (config) => {
  console.log('');
  console.log('═══════════════════════════════════════════════════');
  console.log('🧹 Global Teardown: Cleaning Up Test Environment');
  console.log('═══════════════════════════════════════════════════');
  console.log('');

  try {
    await cleanupDatabase();

    console.log('');
    console.log('✅ Global teardown complete!');
    console.log('═══════════════════════════════════════════════════');
    console.log('');

  } catch (error) {
    console.error('');
    console.error('❌ Global teardown failed:', error);
    console.error('═══════════════════════════════════════════════════');
    console.error('');
  }
};

// DONE: Global teardown runs once after all tests

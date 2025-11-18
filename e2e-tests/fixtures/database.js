const DBHelper = require('../utils/db-helpers');
const path = require('path');
const fs = require('fs');

/**
 * Database fixture for E2E tests
 * Handles setup, seeding, and cleanup
 */

const TEST_DB_PATH = path.resolve(__dirname, '../test.db');

/**
 * Setup test database
 * The database will be created automatically by the Go server on first run
 */
async function setupDatabase() {
  console.log('📦 Setting up test database...');

  // Delete existing test database
  if (fs.existsSync(TEST_DB_PATH)) {
    fs.unlinkSync(TEST_DB_PATH);
    console.log('   Deleted existing test.db');
  }

  // Database will be created by Go server when it starts
  // Migrations run automatically in internal/database/database.go
  console.log('   Database will be created by server on startup');
}

/**
 * Seed initial test data
 */
async function seedInitialData() {
  console.log('🌱 Seeding initial test data...');

  const db = new DBHelper(TEST_DB_PATH);
  await db.connect();

  try {
    // Create test users
    const greenUserId = await db.createUser({
      email: 'green@test.com',
      name: 'Green User',
      experience_level: 'green',
      is_verified: 1,
      is_active: 1,
    });
    console.log('   ✅ Created green@test.com');

    const blueUserId = await db.createUser({
      email: 'blue@test.com',
      name: 'Blue User',
      experience_level: 'blue',
      is_verified: 1,
      is_active: 1,
    });
    console.log('   ✅ Created blue@test.com');

    const orangeUserId = await db.createUser({
      email: 'orange@test.com',
      name: 'Orange User',
      experience_level: 'orange',
      is_verified: 1,
      is_active: 1,
    });
    console.log('   ✅ Created orange@test.com');

    const adminUserId = await db.createUser({
      email: 'admin@test.com',
      name: 'Admin User',
      experience_level: 'orange',
      is_verified: 1,
      is_active: 1,
    });
    console.log('   ✅ Created admin@test.com');

    const unverifiedUserId = await db.createUser({
      email: 'unverified@test.com',
      name: 'Unverified User',
      experience_level: 'green',
      is_verified: 0,
      is_active: 1,
    });
    console.log('   ✅ Created unverified@test.com');

    const inactiveUserId = await db.createUser({
      email: 'inactive@test.com',
      name: 'Inactive User',
      experience_level: 'green',
      is_verified: 1,
      is_active: 0,
    });
    console.log('   ✅ Created inactive@test.com');

    // Create test dogs (3 per category)

    // Green dogs
    const greenDog1 = await db.createDog({
      name: 'Luna',
      breed: 'Golden Retriever',
      category: 'green',
      size: 'large',
      age: 2,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Luna (green)');

    const greenDog2 = await db.createDog({
      name: 'Max',
      breed: 'Labrador',
      category: 'green',
      size: 'large',
      age: 3,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Max (green)');

    const greenDog3 = await db.createDog({
      name: 'Bella',
      breed: 'Beagle',
      category: 'green',
      size: 'medium',
      age: 4,
      is_available: 0,
      unavailable_reason: 'In training',
    });
    console.log('   ✅ Created dog: Bella (green, unavailable)');

    // Blue dogs
    const blueDog1 = await db.createDog({
      name: 'Rocky',
      breed: 'German Shepherd',
      category: 'blue',
      size: 'large',
      age: 4,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Rocky (blue)');

    const blueDog2 = await db.createDog({
      name: 'Daisy',
      breed: 'Border Collie',
      category: 'blue',
      size: 'medium',
      age: 3,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Daisy (blue)');

    const blueDog3 = await db.createDog({
      name: 'Charlie',
      breed: 'Husky',
      category: 'blue',
      size: 'large',
      age: 5,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Charlie (blue)');

    // Orange dogs
    const orangeDog1 = await db.createDog({
      name: 'Rex',
      breed: 'Rottweiler',
      category: 'orange',
      size: 'large',
      age: 6,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Rex (orange)');

    const orangeDog2 = await db.createDog({
      name: 'Zeus',
      breed: 'Doberman',
      category: 'orange',
      size: 'large',
      age: 5,
      is_available: 1,
    });
    console.log('   ✅ Created dog: Zeus (orange)');

    const orangeDog3 = await db.createDog({
      name: 'Thor',
      breed: 'Pitbull',
      category: 'orange',
      size: 'large',
      age: 4,
      is_available: 0,
      unavailable_reason: 'Veterinary care',
    });
    console.log('   ✅ Created dog: Thor (orange, unavailable)');

    console.log('🎉 Test data seeded successfully!');
  } catch (error) {
    console.error('❌ Error seeding data:', error);
    throw error;
  } finally {
    db.close();
  }
}

/**
 * Cleanup database
 */
async function cleanupDatabase() {
  console.log('🧹 Cleaning up test database...');

  if (fs.existsSync(TEST_DB_PATH)) {
    fs.unlinkSync(TEST_DB_PATH);
    console.log('   ✅ Test database deleted');
  }

  // Clean up test uploads directory
  const uploadsPath = path.resolve(__dirname, '../test-uploads');
  if (fs.existsSync(uploadsPath)) {
    fs.rmSync(uploadsPath, { recursive: true, force: true });
    console.log('   ✅ Test uploads deleted');
  }
}

module.exports = {
  setupDatabase,
  seedInitialData,
  cleanupDatabase,
};

// DONE: Database fixture for setup, seeding, and cleanup

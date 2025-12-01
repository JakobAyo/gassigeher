const Database = require('better-sqlite3');
const TEST_PASSWORD_HASH = '$2a$10$LT4jdYaamd5Sxed9IhHTKuedmp/AvzGH27pJwCFzxAqAuO0c6OqfC';

function setupTestData(dbPath) {
    const db = new Database(dbPath);
    const now = new Date().toISOString();       
    const stmt = db.prepare('INSERT OR REPLACE INTO users (email, name, phone, password_hash, experience_level, is_verified, is_active, is_admin, is_super_admin, terms_accepted_at, last_activity_at, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)');
    stmt.run('admin@tierheim-goeppingen.de', 'Admin Name', null, TEST_PASSWORD_HASH, 'orange', 1, 1, 1, 1, now, now, now);
    stmt.run('green@test.com', 'Green User', null, TEST_PASSWORD_HASH, 'green', 1, 1, 0, 0, now, now, now);
    stmt.run('blue@test.com', 'Blue User', null, TEST_PASSWORD_HASH, 'blue', 1, 1, 0, 0, now, now, now);
    stmt.run('delete-me@test.com', 'Delete Me User', null, TEST_PASSWORD_HASH, 'green', 1, 1, 0, 0, now, now, now);
    db.close()
}

module.exports = { setupTestData };
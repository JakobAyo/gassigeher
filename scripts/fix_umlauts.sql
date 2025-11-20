-- Fix corrupted UTF-8 umlauts in database
-- This script fixes double-encoded UTF-8 characters
-- Run with: sqlite3 gassigeher.db < fix_umlauts.sql

-- Backup recommendation: cp gassigeher.db gassigeher.db.backup

BEGIN TRANSACTION;

-- Fix users table
UPDATE users SET name = REPLACE(REPLACE(REPLACE(REPLACE(name, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE name LIKE '%Ã%';
UPDATE users SET email = REPLACE(REPLACE(REPLACE(REPLACE(email, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE email LIKE '%Ã%';

-- Fix dogs table
UPDATE dogs SET name = REPLACE(REPLACE(REPLACE(REPLACE(name, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE name LIKE '%Ã%';
UPDATE dogs SET breed = REPLACE(REPLACE(REPLACE(REPLACE(breed, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE breed LIKE '%Ã%';
UPDATE dogs SET special_needs = REPLACE(REPLACE(REPLACE(REPLACE(special_needs, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE special_needs LIKE '%Ã%';
UPDATE dogs SET pickup_location = REPLACE(REPLACE(REPLACE(REPLACE(pickup_location, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE pickup_location LIKE '%Ã%';
UPDATE dogs SET walk_route = REPLACE(REPLACE(REPLACE(REPLACE(walk_route, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE walk_route LIKE '%Ã%';
UPDATE dogs SET special_instructions = REPLACE(REPLACE(REPLACE(REPLACE(special_instructions, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE special_instructions LIKE '%Ã%';
UPDATE dogs SET unavailable_reason = REPLACE(REPLACE(REPLACE(REPLACE(unavailable_reason, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE unavailable_reason LIKE '%Ã%';

-- Fix bookings table
UPDATE bookings SET user_notes = REPLACE(REPLACE(REPLACE(REPLACE(user_notes, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE user_notes LIKE '%Ã%';
UPDATE bookings SET admin_cancellation_reason = REPLACE(REPLACE(REPLACE(REPLACE(admin_cancellation_reason, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE admin_cancellation_reason LIKE '%Ã%';

-- Fix blocked_dates table
UPDATE blocked_dates SET reason = REPLACE(REPLACE(REPLACE(REPLACE(reason, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE reason LIKE '%Ã%';

-- Fix experience_requests table
UPDATE experience_requests SET admin_message = REPLACE(REPLACE(REPLACE(REPLACE(admin_message, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE admin_message LIKE '%Ã%';

-- Fix reactivation_requests table
UPDATE reactivation_requests SET admin_message = REPLACE(REPLACE(REPLACE(REPLACE(admin_message, 'Ã¤', 'ä'), 'Ã¶', 'ö'), 'Ã¼', 'ü'), 'ÃŸ', 'ß') WHERE admin_message LIKE '%Ã%';

COMMIT;

-- Verify the fixes
SELECT 'Fixed users:' AS table_name, COUNT(*) AS count FROM users WHERE name LIKE '%ü%' OR name LIKE '%ö%' OR name LIKE '%ä%';
SELECT 'Fixed dogs:' AS table_name, COUNT(*) AS count FROM dogs WHERE name LIKE '%ü%' OR name LIKE '%ö%' OR name LIKE '%ä%' OR breed LIKE '%ü%' OR breed LIKE '%ö%' OR breed LIKE '%ä%';
SELECT 'Fixed bookings:' AS table_name, COUNT(*) AS count FROM bookings WHERE user_notes LIKE '%ü%' OR user_notes LIKE '%ö%' OR user_notes LIKE '%ä%';

-- Show sample of fixed data
SELECT 'Sample users:' AS info, id, name, email FROM users WHERE name LIKE '%ü%' OR name LIKE '%ö%' OR name LIKE '%ä%' LIMIT 3;
SELECT 'Sample dogs:' AS info, id, name, breed FROM dogs WHERE name LIKE '%ü%' OR name LIKE '%ö%' OR name LIKE '%ä%' OR breed LIKE '%ü%' OR breed LIKE '%ö%' OR breed LIKE '%ä%' LIMIT 3;

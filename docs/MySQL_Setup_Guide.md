# MySQL Setup Guide for Gassigeher

**Purpose:** Step-by-step guide to configure Gassigeher with MySQL
**Difficulty:** Medium
**Time Required:** 15-30 minutes
**Last Updated:** 2025-01-22

---

## Prerequisites

- MySQL 5.7+ or MySQL 8.0+ (recommended)
- Root or administrative access to MySQL server
- Basic knowledge of MySQL commands

---

## Quick Start (Docker - Recommended for Development)

```bash
# 1. Start MySQL container
docker run --name gassigeher-mysql \
  -e MYSQL_ROOT_PASSWORD=rootpass \
  -e MYSQL_DATABASE=gassigeher \
  -e MYSQL_USER=gassigeher_user \
  -e MYSQL_PASSWORD=gassigeher_pass \
  -p 3306:3306 \
  -d mysql:8.0

# 2. Configure Gassigeher (.env)
DB_TYPE=mysql
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=gassigeher_pass

# 3. Run Gassigeher
go run cmd/server/main.go

# Done! Tables created automatically
```

---

## Production Setup (Step-by-Step)

### Step 1: Install MySQL

#### Ubuntu/Debian:
```bash
sudo apt update
sudo apt install mysql-server
sudo systemctl start mysql
sudo systemctl enable mysql
```

#### CentOS/RHEL:
```bash
sudo yum install mysql-server
sudo systemctl start mysqld
sudo systemctl enable mysqld
```

#### macOS:
```bash
brew install mysql
brew services start mysql
```

#### Windows:
Download installer from: https://dev.mysql.com/downloads/installer/

---

### Step 2: Secure MySQL Installation

```bash
sudo mysql_secure_installation
```

**Recommended answers:**
- Set root password: **Yes**
- Remove anonymous users: **Yes**
- Disallow root login remotely: **Yes**
- Remove test database: **Yes**
- Reload privilege tables: **Yes**

---

### Step 3: Create Database and User

```bash
# Login as root
mysql -u root -p

# Create database
CREATE DATABASE gassigeher CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Create user
CREATE USER 'gassigeher_user'@'localhost' IDENTIFIED BY 'your_secure_password';

# Grant privileges
GRANT ALL PRIVILEGES ON gassigeher.* TO 'gassigeher_user'@'localhost';

# Flush privileges
FLUSH PRIVILEGES;

# Verify
SHOW DATABASES;
SELECT User, Host FROM mysql.user WHERE User = 'gassigeher_user';

# Exit
exit;
```

**Important:** Replace `'your_secure_password'` with a strong password!

---

### Step 4: Test Connection

```bash
# Test connection as gassigeher_user
mysql -u gassigeher_user -p gassigeher

# If successful, you should see:
# mysql>

# Exit
exit;
```

---

### Step 5: Configure Gassigeher

Create or edit `.env` file:

```bash
# Database Configuration
DB_TYPE=mysql
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=your_secure_password

# Connection Pool (optional, these are defaults)
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=5

# ... other configuration ...
```

---

### Step 6: Run Gassigeher

```bash
# Run application
go run cmd/server/main.go

# Expected output:
# 2025/01/22 12:00:00 Using database: mysql
# 2025/01/22 12:00:00 Applying migration: 001_create_users_table
# ... (9 migrations)
# 2025/01/22 12:00:00 Applied 9 migration(s)
# 2025/01/22 12:00:00 Server starting on port 8080...
```

---

### Step 7: Verify Database Setup

```bash
# Connect to MySQL
mysql -u gassigeher_user -p gassigeher

# Check tables created
SHOW TABLES;

# Should show:
# +------------------------+
# | Tables_in_gassigeher   |
# +------------------------+
# | blocked_dates          |
# | bookings               |
# | dogs                   |
# | experience_requests    |
# | reactivation_requests  |
# | schema_migrations      |
# | system_settings        |
# | users                  |
# +------------------------+

# Check migration status
SELECT * FROM schema_migrations ORDER BY applied_at;

# Should show 9 migrations applied

# Exit
exit;
```

---

## Remote MySQL Server

### If MySQL is on Different Server

```bash
# .env configuration
DB_TYPE=mysql
DB_HOST=db.example.com  # Remote MySQL server
DB_PORT=3306
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=your_password
```

**Create user with remote access:**
```sql
-- On MySQL server
CREATE USER 'gassigeher_user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON gassigeher.* TO 'gassigeher_user'@'%';
FLUSH PRIVILEGES;
```

**Note:** `'%'` allows access from any host. For security, replace with specific IP:
```sql
CREATE USER 'gassigeher_user'@'192.168.1.100' IDENTIFIED BY 'password';
```

---

## Cloud MySQL (AWS RDS, Google Cloud SQL, etc.)

### AWS RDS for MySQL

**1. Create RDS Instance:**
- Go to AWS RDS Console
- Create database → MySQL
- Select version 8.0
- Choose instance size (t3.micro for dev, t3.small+ for production)
- Set master username and password
- Configure VPC and security groups
- Create database

**2. Get Connection Details:**
- Endpoint: `gassigeher.xxxxx.us-east-1.rds.amazonaws.com`
- Port: `3306`
- Username: Your master username
- Password: Your master password

**3. Create Gassigeher Database:**
```bash
mysql -h gassigeher.xxxxx.us-east-1.rds.amazonaws.com -u admin -p

CREATE DATABASE gassigeher CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
exit;
```

**4. Configure Gassigeher:**
```bash
DB_TYPE=mysql
DB_HOST=gassigeher.xxxxx.us-east-1.rds.amazonaws.com
DB_PORT=3306
DB_NAME=gassigeher
DB_USER=admin
DB_PASSWORD=your_rds_password
```

---

### DigitalOcean Managed MySQL

**1. Create Managed Database:**
- Go to DigitalOcean → Databases
- Create → MySQL 8
- Choose datacenter and size
- Create cluster

**2. Get Connection Details:**
- Host: `db-mysql-xxx.ondigitalocean.com`
- Port: `25060`
- User: `doadmin`
- Password: (shown in UI)
- Database: `defaultdb`

**3. Create Gassigeher Database:**
```bash
mysql -h db-mysql-xxx.ondigitalocean.com -P 25060 -u doadmin -p

CREATE DATABASE gassigeher CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'gassigeher_user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON gassigeher.* TO 'gassigeher_user'@'%';
exit;
```

**4. Configure Gassigeher:**
```bash
DB_TYPE=mysql
DB_HOST=db-mysql-xxx.ondigitalocean.com
DB_PORT=25060
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=your_password
DB_CONNECTION_STRING=gassigeher_user:password@tcp(db-mysql-xxx.ondigitalocean.com:25060)/gassigeher?parseTime=true&charset=utf8mb4
```

---

## Connection String Format

### Standard Format

```
username:password@tcp(host:port)/database?parameters
```

### Full Example

```
gassigeher_user:mypassword@tcp(localhost:3306)/gassigeher?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci
```

### Required Parameters

- `parseTime=true` - **Required** for Go time.Time field scanning
- `charset=utf8mb4` - **Required** for full Unicode support (emoji, etc.)

### Optional Parameters

- `collation=utf8mb4_unicode_ci` - Unicode collation (recommended)
- `loc=UTC` - Timezone location
- `timeout=10s` - Connection timeout
- `readTimeout=30s` - Read timeout
- `writeTimeout=30s` - Write timeout

---

## Troubleshooting

### "Access denied for user"

**Problem:** Wrong username or password

**Solution:**
```sql
-- Reset password
ALTER USER 'gassigeher_user'@'localhost' IDENTIFIED BY 'new_password';
FLUSH PRIVILEGES;
```

---

### "Unknown database 'gassigeher'"

**Problem:** Database not created

**Solution:**
```sql
CREATE DATABASE gassigeher CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

---

### "Can't connect to MySQL server"

**Problem:** MySQL not running or firewall blocking

**Solutions:**
```bash
# Check if MySQL running
sudo systemctl status mysql

# Start MySQL
sudo systemctl start mysql

# Check firewall
sudo ufw allow 3306/tcp
```

---

### "Client does not support authentication protocol"

**Problem:** Old authentication plugin (MySQL 8.0)

**Solution:**
```sql
ALTER USER 'gassigeher_user'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
FLUSH PRIVILEGES;
```

---

### "Too many connections"

**Problem:** Connection pool exhausted

**Solution in .env:**
```bash
DB_MAX_OPEN_CONNS=50  # Increase from default 25
DB_MAX_IDLE_CONNS=10  # Increase from default 5
```

**Or increase MySQL max_connections:**
```sql
SET GLOBAL max_connections = 200;
```

---

## Performance Tuning

### Recommended MySQL Configuration

Edit `/etc/mysql/mysql.conf.d/mysqld.cnf` (Linux):

```ini
[mysqld]
# Connection settings
max_connections = 100

# InnoDB settings (for Gassigeher tables)
innodb_buffer_pool_size = 256M  # Increase for more RAM
innodb_log_file_size = 64M
innodb_flush_log_at_trx_commit = 2  # Better performance, slight risk

# Character set
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci

# Query cache (MySQL 5.7 only, removed in 8.0)
# query_cache_size = 32M
# query_cache_limit = 2M
```

**Restart MySQL after changes:**
```bash
sudo systemctl restart mysql
```

---

## Backup and Restore

### Backup Database

```bash
# Full backup
mysqldump -u gassigeher_user -p gassigeher > gassigeher_backup_$(date +%Y%m%d).sql

# Compressed backup
mysqldump -u gassigeher_user -p gassigeher | gzip > gassigeher_backup_$(date +%Y%m%d).sql.gz

# Automated daily backup (crontab)
0 3 * * * mysqldump -u gassigeher_user -pPASSWORD gassigeher | gzip > /backups/gassigeher_$(date +\%Y\%m\%d).sql.gz
```

### Restore Database

```bash
# Restore from backup
mysql -u gassigeher_user -p gassigeher < gassigeher_backup_20250122.sql

# Restore from compressed backup
gunzip < gassigeher_backup_20250122.sql.gz | mysql -u gassigeher_user -p gassigeher
```

---

## Migration from SQLite

See [Database_Migration_Guide.md](Database_Migration_Guide.md) for complete migration procedures.

**Quick overview:**
1. Export data from SQLite
2. Transform SQL for MySQL
3. Import into MySQL
4. Update .env to use MySQL
5. Verify application works

---

## Monitoring

### Check Database Size

```sql
SELECT
  table_schema AS 'Database',
  ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS 'Size (MB)'
FROM information_schema.TABLES
WHERE table_schema = 'gassigeher'
GROUP BY table_schema;
```

### Check Table Sizes

```sql
SELECT
  table_name AS 'Table',
  ROUND(((data_length + index_length) / 1024 / 1024), 2) AS 'Size (MB)'
FROM information_schema.TABLES
WHERE table_schema = 'gassigeher'
ORDER BY (data_length + index_length) DESC;
```

### Check Connection Count

```sql
SHOW STATUS LIKE 'Threads_connected';
SHOW STATUS LIKE 'Max_used_connections';
```

---

## Security Recommendations

### 1. Use Strong Passwords

```bash
# Generate secure password
openssl rand -base64 32
```

### 2. Limit Network Access

```sql
-- For local-only access
CREATE USER 'gassigeher_user'@'localhost' IDENTIFIED BY 'password';

-- For specific IP
CREATE USER 'gassigeher_user'@'192.168.1.100' IDENTIFIED BY 'password';

-- Avoid '%' (any host) in production
```

### 3. Use SSL/TLS (Production)

```bash
# In .env, add SSL parameter to connection string
DB_CONNECTION_STRING=gassigeher_user:password@tcp(localhost:3306)/gassigeher?parseTime=true&charset=utf8mb4&tls=true
```

### 4. Regular Backups

```bash
# Automated daily backups
0 3 * * * /usr/local/bin/backup_mysql.sh
```

### 5. Keep MySQL Updated

```bash
sudo apt update
sudo apt upgrade mysql-server
```

---

## Docker Compose (Recommended for Development)

**File:** `docker-compose.yml`

```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    container_name: gassigeher-mysql
    environment:
      MYSQL_ROOT_PASSWORD: rootpass
      MYSQL_DATABASE: gassigeher
      MYSQL_USER: gassigeher_user
      MYSQL_PASSWORD: gassigeher_pass
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
      - ./mysql-init:/docker-entrypoint-initdb.d  # Optional: initialization scripts
    command:
      - --character-set-server=utf8mb4
      - --collation-server=utf8mb4_unicode_ci
      - --default-authentication-plugin=mysql_native_password

  gassigeher:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - mysql
    environment:
      DB_TYPE: mysql
      DB_HOST: mysql
      DB_PORT: 3306
      DB_NAME: gassigeher
      DB_USER: gassigeher_user
      DB_PASSWORD: gassigeher_pass

volumes:
  mysql_data:
```

**Usage:**
```bash
docker-compose up -d
```

---

## Verification Checklist

After setup, verify:

- [ ] MySQL server is running
  ```bash
  sudo systemctl status mysql
  ```

- [ ] Database exists
  ```sql
  SHOW DATABASES LIKE 'gassigeher';
  ```

- [ ] User has correct privileges
  ```sql
  SHOW GRANTS FOR 'gassigeher_user'@'localhost';
  ```

- [ ] Character set is UTF8MB4
  ```sql
  SELECT DEFAULT_CHARACTER_SET_NAME, DEFAULT_COLLATION_NAME
  FROM information_schema.SCHEMATA
  WHERE SCHEMA_NAME = 'gassigeher';
  ```

- [ ] Gassigeher connects successfully
  ```bash
  go run cmd/server/main.go
  # Should see: "Using database: mysql"
  ```

- [ ] All tables created
  ```sql
  USE gassigeher;
  SHOW TABLES;
  # Should show 8 tables
  ```

- [ ] Application accessible
  ```bash
  curl http://localhost:8080/
  # Should return HTML
  ```

---

## Common Issues

### Issue: Character Encoding Problems (ä, ö, ü, ß)

**Solution:** Ensure UTF8MB4
```sql
ALTER DATABASE gassigeher CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
ALTER TABLE users CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Repeat for all tables
```

### Issue: Slow Queries

**Solution:** Add indexes (already created by migrations, but verify):
```sql
SHOW INDEX FROM users;
SHOW INDEX FROM dogs;
SHOW INDEX FROM bookings;
```

### Issue: Connection Pool Exhausted

**Solution:** Increase pool size in .env:
```bash
DB_MAX_OPEN_CONNS=50
DB_MAX_IDLE_CONNS=10
```

---

## Maintenance

### Weekly Tasks

- Check disk space: `df -h`
- Check database size (see Monitoring section)
- Review slow query log

### Monthly Tasks

- Backup database
- Optimize tables: `OPTIMIZE TABLE users, dogs, bookings;`
- Check for updates: `sudo apt update && sudo apt list --upgradable | grep mysql`

### Quarterly Tasks

- Review and adjust configuration
- Analyze query performance
- Plan capacity upgrades if needed

---

## Uninstall / Cleanup

### Remove MySQL (if needed)

```bash
# Stop Gassigeher
pkill gassigeher

# Backup first!
mysqldump -u gassigeher_user -p gassigeher > final_backup.sql

# Drop database
mysql -u root -p
DROP DATABASE gassigeher;
DROP USER 'gassigeher_user'@'localhost';
exit;

# Optionally remove MySQL server
sudo apt remove mysql-server
```

### Remove Docker Container

```bash
docker stop gassigeher-mysql
docker rm gassigeher-mysql
docker volume rm gassigeher_mysql_data  # Removes data permanently!
```

---

## Next Steps

- **Completed MySQL Setup?** See [Database_Migration_Guide.md](Database_Migration_Guide.md) to migrate data from SQLite
- **Need PostgreSQL instead?** See [PostgreSQL_Setup_Guide.md](PostgreSQL_Setup_Guide.md)
- **Performance issues?** See DEPLOYMENT.md for optimization tips
- **Questions?** See [Database_Selection_Guide.md](Database_Selection_Guide.md)

---

**Your MySQL database is ready for Gassigeher!** 🎉

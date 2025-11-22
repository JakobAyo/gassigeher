# Database Selection Guide for Gassigeher

**Purpose:** Help you choose the right database for your deployment
**Last Updated:** 2025-01-22
**Supported Databases:** SQLite, MySQL, PostgreSQL

---

## Quick Decision Matrix

| Your Situation | Recommended Database | Why |
|----------------|---------------------|-----|
| **Development** | SQLite | Zero setup, fast, file-based |
| **Small shelter (<500 users)** | SQLite | Simple, reliable, no server needed |
| **Medium shelter (500-5000 users)** | MySQL | Proven, scalable, easy hosting |
| **Large shelter (5000+ users)** | PostgreSQL | Enterprise-grade, highly concurrent |
| **Multiple shelters** | PostgreSQL | Advanced features, better performance |
| **Shared hosting** | MySQL | Most hosting providers support it |
| **Cloud deployment** | MySQL or PostgreSQL | Both well-supported on cloud platforms |

---

## Database Comparison

### SQLite

**Best For:** Development, small deployments, simple setups

**Advantages:**
- ✅ **Zero configuration** - No server to install or configure
- ✅ **File-based** - Single file database (easy backup)
- ✅ **Fast for small datasets** - Excellent performance up to 1000 users
- ✅ **Portable** - Works on Windows, Linux, Mac
- ✅ **No cost** - Free, no licensing
- ✅ **Easy backup** - Just copy the file

**Limitations:**
- ❌ **Limited concurrency** - One writer at a time
- ❌ **No network access** - Can't connect from remote clients
- ❌ **Not for large scale** - Performance degrades above 1000 users
- ❌ **Single server only** - Can't distribute across servers

**Recommended For:**
- Development and testing
- Single-server deployments
- Shelters with <500 users
- Simple hosting requirements

**Max Recommended Users:** 1,000

---

### MySQL

**Best For:** Web applications, medium to large deployments

**Advantages:**
- ✅ **Proven technology** - Used by millions of websites
- ✅ **Great concurrency** - Handles many simultaneous users
- ✅ **Wide hosting support** - Available on most web hosts
- ✅ **Replication** - Can set up read replicas
- ✅ **Good performance** - Optimized for web workloads
- ✅ **Large community** - Extensive documentation and support

**Limitations:**
- ⚠️ **Requires server** - Need to install and maintain MySQL server
- ⚠️ **More complex** - Configuration and tuning needed
- ⚠️ **Memory usage** - Needs dedicated server resources
- ⚠️ **Cost** - May require paid hosting or server

**Recommended For:**
- Production web applications
- Shelters with 500-50,000 users
- Shared hosting environments
- Deployments requiring high availability

**Max Recommended Users:** 100,000+

---

### PostgreSQL

**Best For:** Enterprise applications, complex queries, high concurrency

**Advantages:**
- ✅ **Advanced features** - JSON, full-text search, geospatial
- ✅ **Excellent concurrency** - Best for many simultaneous writers
- ✅ **ACID compliant** - Strong data integrity guarantees
- ✅ **Extensible** - Can add custom functions and types
- ✅ **Standards compliant** - Follows SQL standards closely
- ✅ **Great for analytics** - Complex queries perform well

**Limitations:**
- ⚠️ **More complex setup** - Requires PostgreSQL server
- ⚠️ **Steeper learning curve** - More configuration options
- ⚠️ **Resource intensive** - Needs more RAM than MySQL
- ⚠️ **Less common on shared hosting** - May need VPS or cloud

**Recommended For:**
- Enterprise deployments
- Multiple shelter network
- Shelters with 10,000+ users
- Applications with complex data requirements
- Cloud deployments (AWS RDS, Google Cloud SQL, etc.)

**Max Recommended Users:** 1,000,000+

---

## Feature Comparison

| Feature | SQLite | MySQL | PostgreSQL |
|---------|--------|-------|------------|
| **Setup Time** | 0 min | 15-30 min | 30-45 min |
| **Maintenance** | None | Low-Medium | Medium |
| **Backup** | File copy | mysqldump | pg_dump |
| **Replication** | No | Yes | Yes |
| **Clustering** | No | Yes | Yes |
| **Full-Text Search** | Yes (FTS5) | Yes | Yes (Better) |
| **JSON Support** | Yes | Yes | Yes (Better) |
| **Concurrent Writes** | Limited | Good | Excellent |
| **Transaction Performance** | Excellent | Good | Excellent |
| **Storage Limit** | 281 TB | Unlimited | Unlimited |

---

## Performance Comparison

### Expected Response Times (Typical Queries)

| Operation | SQLite | MySQL | PostgreSQL |
|-----------|--------|-------|------------|
| **Simple SELECT** | 0.1-0.5ms | 0.5-2ms | 0.5-2ms |
| **Complex JOIN** | 1-5ms | 2-10ms | 2-8ms |
| **INSERT** | 0.5-1ms | 1-3ms | 1-3ms |
| **UPDATE** | 0.5-1ms | 1-3ms | 1-3ms |
| **Transaction** | 1-2ms | 2-5ms | 2-5ms |

**Note:** Network latency adds ~0.5-1ms for MySQL/PostgreSQL on remote servers

### Concurrent User Support

| Database | Concurrent Reads | Concurrent Writes | Max Users* |
|----------|------------------|-------------------|------------|
| **SQLite** | Unlimited | 1 at a time | 1,000 |
| **MySQL** | Excellent | Good | 100,000+ |
| **PostgreSQL** | Excellent | Excellent | 1,000,000+ |

*Max users = realistic limit for Gassigeher use case

---

## Cost Comparison

### SQLite

**Server Cost:** $0 (runs on app server)
**Hosting Cost:** Minimal (any server can run it)
**Maintenance Cost:** $0/month
**Backup Cost:** $0 (file copy)

**Total:** ~$5-10/month (app server only)

---

### MySQL

**Server Cost:**
- Shared hosting: Included
- VPS: $10-50/month
- Cloud (AWS RDS): $15-100/month

**Hosting Cost:** Depends on size
**Maintenance Cost:** Low
**Backup Cost:** Included in hosting

**Total:** ~$20-100/month (varies by hosting)

---

### PostgreSQL

**Server Cost:**
- VPS: $20-100/month
- Cloud (AWS RDS): $25-150/month
- Managed (Digital Ocean): $15-60/month

**Hosting Cost:** Depends on size
**Maintenance Cost:** Low-Medium
**Backup Cost:** Included in hosting

**Total:** ~$30-150/month (varies by hosting)

---

## Migration Paths

### Start with SQLite, Grow as Needed

**Recommended Path:**
```
Development → SQLite
  ↓
Small Deployment (< 500 users) → SQLite
  ↓
Growing (500-5000 users) → Migrate to MySQL
  ↓
Enterprise (5000+ users) → Migrate to PostgreSQL (or stay on MySQL)
```

**Migration Difficulty:**
- SQLite → MySQL: Easy (1-2 hours)
- SQLite → PostgreSQL: Easy (1-2 hours)
- MySQL → PostgreSQL: Medium (2-4 hours)

---

## When to Switch Databases

### Signs You've Outgrown SQLite

- ⚠️ Database file > 1 GB
- ⚠️ Frequent "database locked" errors
- ⚠️ Slow queries (>100ms for simple SELECTs)
- ⚠️ More than 10 concurrent users
- ⚠️ Need remote database access
- ⚠️ Plan to exceed 1000 users

**Solution:** Migrate to MySQL or PostgreSQL

---

### Signs You Need PostgreSQL Over MySQL

- ⚠️ Complex queries (many JOINs, subqueries)
- ⚠️ Need for advanced features (JSON, arrays, custom types)
- ⚠️ High write concurrency (many simultaneous bookings)
- ⚠️ Enterprise compliance requirements
- ⚠️ Multi-region deployment planned

**Solution:** Choose PostgreSQL or migrate from MySQL

---

## Database Selection Flowchart

```
Start Here
   ↓
Do you have < 500 users?
   ├─ Yes → Use SQLite ✅
   │         (Simple, free, fast)
   │
   └─ No → Do you need enterprise features?
            ├─ Yes → Use PostgreSQL ✅
            │         (Advanced, scalable)
            │
            └─ No → Use MySQL ✅
                      (Proven, widely supported)
```

---

## Detailed Comparison

### SQLite Use Cases ✅

**Perfect For:**
- Local development
- Demo/staging environments
- Small animal shelters (1-50 dogs, <500 volunteers)
- Single-server deployments
- Embedded applications

**Example Deployment:**
- Small shelter with 20 dogs
- 100 registered volunteers
- 1-5 concurrent users typically
- ~100 bookings per month
- Single VPS server

**Database Size After 1 Year:**
- Users: 100 × ~1KB = 100KB
- Dogs: 20 × ~2KB = 40KB
- Bookings: 1200 × ~500 bytes = 600KB
- **Total: ~1MB** ✅ SQLite handles this easily

---

### MySQL Use Cases ✅

**Perfect For:**
- Growing shelters (100-500 dogs, 500-10,000 volunteers)
- Multiple concurrent users (10-100)
- Web hosting environments
- Moderate write load
- Standard web applications

**Example Deployment:**
- Medium shelter with 200 dogs
- 2,000 registered volunteers
- 10-30 concurrent users typically
- ~1,000 bookings per month
- Cloud hosting (AWS, DigitalOcean, etc.)

**Database Size After 1 Year:**
- Users: 2,000 × ~1KB = 2MB
- Dogs: 200 × ~2KB = 400KB
- Bookings: 12,000 × ~500 bytes = 6MB
- **Total: ~10MB** ✅ MySQL handles this with ease

---

### PostgreSQL Use Cases ✅

**Perfect For:**
- Large shelters or shelter networks
- 500+ dogs, 10,000+ volunteers
- High concurrent users (50-500)
- Heavy write load
- Complex reporting needs
- Multi-region deployments

**Example Deployment:**
- Shelter network with 1,000+ dogs
- 10,000+ registered volunteers
- 50-200 concurrent users
- ~10,000 bookings per month
- Cloud infrastructure
- Advanced analytics

**Database Size After 1 Year:**
- Users: 10,000 × ~1KB = 10MB
- Dogs: 1,000 × ~2KB = 2MB
- Bookings: 120,000 × ~500 bytes = 60MB
- **Total: ~75MB** ✅ PostgreSQL excels at this scale

---

## Configuration Examples

### SQLite (Default)

**.env:**
```bash
# Minimal configuration (or no .env at all)
DATABASE_PATH=./gassigeher.db
```

**That's it!** No server, no additional configuration.

---

### MySQL

**.env:**
```bash
DB_TYPE=mysql
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=your_secure_password
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
```

**Setup Time:** 15-30 minutes (install MySQL, create database, configure)

---

### PostgreSQL

**.env:**
```bash
DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=your_secure_password
DB_SSLMODE=require
DB_MAX_OPEN_CONNS=25
DB_MAX_IDLE_CONNS=5
```

**Setup Time:** 30-45 minutes (install PostgreSQL, create database, configure)

---

## Hosting Recommendations

### For SQLite

**Best Hosting:**
- Any VPS (DigitalOcean, Linode, Vultr)
- Shared hosting (if Go supported)
- Local server

**Requirements:**
- 512MB RAM minimum
- 10GB disk space
- Linux or Windows server

**Cost:** $5-10/month

---

### For MySQL

**Best Hosting:**
- Managed MySQL (AWS RDS, DigitalOcean Managed Databases)
- VPS with MySQL installed
- Shared hosting with MySQL

**Requirements:**
- 1GB RAM minimum (2GB recommended)
- 20GB disk space
- Reliable network

**Cost:** $15-50/month (managed), $10-30/month (VPS + self-managed)

---

### For PostgreSQL

**Best Hosting:**
- Managed PostgreSQL (AWS RDS, Google Cloud SQL, DigitalOcean)
- VPS with PostgreSQL installed
- Heroku Postgres

**Requirements:**
- 2GB RAM minimum (4GB recommended)
- 20GB disk space
- Reliable network

**Cost:** $25-100/month (managed), $20-50/month (VPS + self-managed)

---

## Migration Timing

### When to Migrate

**SQLite → MySQL:**
- Reaching 500-1000 users
- Need for concurrent write access
- Remote database access needed
- Approaching 1GB database size

**SQLite → PostgreSQL:**
- Enterprise requirements
- Need advanced features
- High concurrency expected
- Complex query requirements

**MySQL → PostgreSQL:**
- Outgrowing MySQL performance
- Need PostgreSQL-specific features
- Higher concurrency requirements
- Complex data relationships

**Timeline:** Plan migration before hitting limits, not after

---

## Decision Factors

### Technical Factors

| Factor | SQLite | MySQL | PostgreSQL |
|--------|--------|-------|------------|
| **Concurrent Writers** | 1 | 100+ | 500+ |
| **Query Complexity** | Simple-Medium | Medium-Complex | Very Complex |
| **Data Size** | <10 GB | <1 TB | Unlimited |
| **Setup Complexity** | ⭐ Easy | ⭐⭐⭐ Medium | ⭐⭐⭐⭐ Medium-Hard |
| **Maintenance** | ⭐ None | ⭐⭐⭐ Regular | ⭐⭐⭐⭐ Regular |

### Business Factors

| Factor | SQLite | MySQL | PostgreSQL |
|--------|--------|-------|------------|
| **Initial Cost** | $0 | $$ | $$$ |
| **Ongoing Cost** | $ | $$ | $$$ |
| **Team Expertise** | Easy | Common | Less Common |
| **Vendor Support** | Limited | Excellent | Good |
| **Cloud Options** | Limited | Excellent | Excellent |

---

## Recommendations by Shelter Size

### Tiny Shelter (1-10 dogs, <50 users)

**Recommended:** SQLite
**Why:** Overkill to use anything else
**Cost:** ~$5-10/month (basic VPS)

---

### Small Shelter (10-50 dogs, 50-500 users)

**Recommended:** SQLite
**Why:** Still within SQLite's sweet spot
**Cost:** ~$10-20/month (VPS)

---

### Medium Shelter (50-200 dogs, 500-5000 users)

**Recommended:** MySQL
**Why:** Better concurrency, scalability, hosting options
**Cost:** ~$30-80/month (managed MySQL or VPS)

**Alternative:** PostgreSQL if you have the expertise

---

### Large Shelter (200+ dogs, 5000+ users)

**Recommended:** PostgreSQL
**Why:** Best performance at scale, advanced features
**Cost:** ~$50-150/month (managed PostgreSQL)

**Alternative:** MySQL if already invested in it

---

### Shelter Network (Multiple Locations)

**Recommended:** PostgreSQL
**Why:** Multi-region support, advanced features, scalability
**Cost:** ~$100-500/month (cloud deployment with replication)

---

## Quick Start Guides

### Use SQLite (No Setup)

```bash
# 1. Clone repository
git clone <repo-url>
cd gassigeher

# 2. Configure (or use defaults)
cp .env.example .env
# DATABASE_PATH=./gassigeher.db (default)

# 3. Run
go run cmd/server/main.go

# Done! Database created automatically
```

---

### Use MySQL (With Docker)

```bash
# 1. Start MySQL
docker run --name gassigeher-mysql \
  -e MYSQL_ROOT_PASSWORD=rootpass \
  -e MYSQL_DATABASE=gassigeher \
  -e MYSQL_USER=gassigeher_user \
  -e MYSQL_PASSWORD=gassigeher_pass \
  -p 3306:3306 -d mysql:8.0

# 2. Configure .env
DB_TYPE=mysql
DB_HOST=localhost
DB_PORT=3306
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=gassigeher_pass

# 3. Run application
go run cmd/server/main.go
# Tables created automatically!
```

---

### Use PostgreSQL (With Docker)

```bash
# 1. Start PostgreSQL
docker run --name gassigeher-postgres \
  -e POSTGRES_DB=gassigeher \
  -e POSTGRES_USER=gassigeher_user \
  -e POSTGRES_PASSWORD=gassigeher_pass \
  -p 5432:5432 -d postgres:15

# 2. Configure .env
DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=gassigeher
DB_USER=gassigeher_user
DB_PASSWORD=gassigeher_pass
DB_SSLMODE=disable

# 3. Run application
go run cmd/server/main.go
# Tables created automatically!
```

---

## FAQ

### Can I switch databases later?

**Yes!** Gassigeher supports all three databases. See the [Migration Guide](Database_Migration_Guide.md) for step-by-step instructions.

---

### What if I'm not sure which to choose?

**Start with SQLite.** You can always migrate later. SQLite is perfect for:
- Getting started quickly
- Development and testing
- Small deployments

When you outgrow SQLite, you'll know (slow queries, database locked errors).

---

### Does the API change with different databases?

**No.** The API is identical regardless of database. All features work the same.

---

### Can I use a different database not listed?

**Not currently.** Gassigeher supports SQLite, MySQL, and PostgreSQL. These cover 99% of use cases. To add another database (e.g., SQL Server, Oracle), you'd need to implement a dialect - see CLAUDE.md for details.

---

### Which database is most reliable?

All three are **extremely reliable** when properly configured. Choose based on your needs, not reliability concerns.

- SQLite: Billions of deployments, rock-solid
- MySQL: Powers most of the web, battle-tested
- PostgreSQL: Bank-grade reliability, ACID compliant

---

### Can I use managed database services?

**Yes!** Gassigeher works great with:

**MySQL:**
- AWS RDS for MySQL
- Google Cloud SQL for MySQL
- Azure Database for MySQL
- DigitalOcean Managed MySQL

**PostgreSQL:**
- AWS RDS for PostgreSQL
- Google Cloud SQL for PostgreSQL
- Azure Database for PostgreSQL
- Heroku Postgres
- DigitalOcean Managed PostgreSQL

Just set `DB_HOST`, `DB_USER`, `DB_PASSWORD` to your managed service credentials.

---

## Summary

**For most users:** Start with **SQLite** (simple, free, fast)

**When growing:** Migrate to **MySQL** (proven, widely supported)

**For enterprise:** Use **PostgreSQL** (advanced, scalable)

**Switching databases:** Easy - just change environment variables and run migrations!

---

**Related Documentation:**
- [MySQL Setup Guide](MySQL_Setup_Guide.md) - Detailed MySQL configuration
- [PostgreSQL Setup Guide](PostgreSQL_Setup_Guide.md) - Detailed PostgreSQL configuration
- [Database Migration Guide](Database_Migration_Guide.md) - How to migrate between databases
- [Multi-Database Testing Guide](MultiDatabase_Testing_Guide.md) - Testing with all databases

---

**Need Help Deciding?** Consider:
1. How many users do you expect?
2. What's your technical expertise?
3. What's your budget?
4. Do you have existing database infrastructure?

Still unsure? **Start with SQLite** - you can always migrate later!

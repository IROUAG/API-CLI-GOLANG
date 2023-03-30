# Use the official PostgreSQL image as the base
FROM postgres

# Copy the SQL script to the /docker-entrypoint-initdb.d/ directory
COPY setup.sql /docker-entrypoint-initdb.d/

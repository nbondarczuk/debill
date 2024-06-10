CREATE SCHEMA debill;
CREATE USER debill;
ALTER ROLE debill SET search_path='debill';
GRANT ALL ON SCHEMA debill TO debill;

CREATE USER debill;
CREATE SCHEMA debill;
ALTER ROLE debill SET search_path='debill';
GRANT ALL ON SCHEMA debill TO debill;

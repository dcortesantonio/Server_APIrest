--Documentation: https://www.cockroachlabs.com/docs/v20.1/build-a-go-app-with-cockroachdb.html#step-1-install-the-go-pq-driver
--DataBase: CockRoach and the Go pq driver. [Local Cluster]

-- Create maxroaxh user.
CREATE USER IF NOT EXISTS maxroach;

-- Create DataBase.
CREATE DATABASE IF NOT EXISTS SERVERSDB;

-- Give to the user the necessary permissions.
GRANT ALL ON DATABASE SERVERSDB TO maxroach;

-- CockroachDB SQL statements to create tables.
CREATE TABLE IF NOT EXISTS SERVERSDB.DOMAIN (
  id SERIAL,
  timeSearched TIMESTAMP,
  domainSearched TEXT NOT NULL,
  PRIMARY KEY (id));

CREATE TABLE IF NOT EXISTS SERVERSDB.SERVER (
  id SERIAL,
  serversChanged BOOLEAN NOT NULL,
  MIN_SSLGrade VARCHAR(3) NOT NULL,
  PREV_SSLGrade VARCHAR(3) NOT NULL,
  logo TEXT NOT NULL,
  title TEXT NOT NULL,
  isDown BOOLEAN NOT NULL,
  domainId INT NOT NULL,
  PRIMARY KEY (id),
  INDEX domain_ID (domainId ASC),
  CONSTRAINT domainId
    FOREIGN KEY (domainId)
    REFERENCES SERVERSDB.domain (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);

CREATE TABLE IF NOT EXISTS SERVERSDB.SERVERITEM (
  id SERIAL,
  address TEXT NOT NULL,
  SSL_Grade VARCHAR(3) NOT NULL,
  country VARCHAR(5) NOT NULL,
  owner TEXT NOT NULL,
  serverId INT NOT NULL,
  PRIMARY KEY (id),
  INDEX server_ID (serverId ASC),
  CONSTRAINT serverId
    FOREIGN KEY (serverId)
    REFERENCES SERVERSDB.server (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE);
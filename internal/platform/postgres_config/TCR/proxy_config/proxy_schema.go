package proxy_config


const TABLE_SCHEMA_FOR_PROXY = `
CREATE TABLE proxies (
    uuid TEXT NOT NULL,
    PRIMARY KEY (uuid)
);
CREATE INDEX proxies_index ON proxies (uuid);
`

const TABLE_NAME_FOR_PROXY = "proxies"

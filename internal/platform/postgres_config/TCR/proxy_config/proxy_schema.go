package proxy_config

const TABLE_SCHEMA_FOR_PROXY = `
CREATE TABLE proxies (
    id BIGSERIAL,
    uuid TEXT NOT NULL,
    PRIMARY KEY (uuid)
);
CREATE INDEX proxies_index ON proxies (id, uuid);
CREATE INDEX proxies_desc_index ON proxies (id DESC NULLS LAST);
`

const TABLE_NAME_FOR_PROXY = "proxies"

CREATE TABLE proxies (
    id BIGSERIAL,
    uuid TEXT NOT NULL,
    PRIMARY KEY (uuid)
);
CREATE INDEX proxies_index ON proxies (id, uuid);
CREATE INDEX proxies_desc_index ON proxies (id DESC NULLS LAST);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON proxies
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

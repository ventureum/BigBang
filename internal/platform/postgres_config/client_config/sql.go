package client_config

const TRIGGER_SET_TIMESTAMP_COMMAND = `
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
`

const REGISTER_TIMESTAMP_TRIGGER_COMMAND = `
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON %s
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
`

const LOAD_UUID_EXTENSION = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
`

const LOAD_VOTE_TYPE_ENUM = `
DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vote_type_enum') THEN
  CREATE TYPE vote_type_enum AS ENUM ('DOWN','UP');
END IF;
END$$;
`

const LOAD_ACTOR_TYPE_ENUM = `
DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'actor_type_enum') THEN
  CREATE TYPE actor_type_enum AS ENUM ('USER','KOL', 'ADMIN', 'PF');
END IF;
END$$;
`

const SET_IDLE_IN_TX_SESSION_TIMEOUT = `
set idle_in_transaction_session_timeout = %d;
`

const LOAD_ACTOR_PROFILE_STATUS_ENUM = `
DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'actor_profile_status_enum') THEN
  CREATE TYPE actor_profile_status_enum AS ENUM ('ACTIVATED','DEACTIVATED');
END IF;
END$$;
`

const LOAD_MILESTONE_STATE_ENUM = `
DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'milestone_state_enum') THEN
  CREATE TYPE milestone_state_enum AS ENUM ('COMPLETE','IN_PROGRESS', 'PENDING');
END IF;
END$$;
`
const DROP_MILESTONE_STATE_ENUM = `
DO $$
BEGIN
IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'milestone_state_enum') THEN
  DROP TYPE milestone_state_enum;
END IF;
END$$;
`

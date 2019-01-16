DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'actor_profile_status_enum') THEN
  CREATE TYPE actor_profile_status_enum AS ENUM ('ACTIVATED','DEACTIVATED');
END IF;
END$$;

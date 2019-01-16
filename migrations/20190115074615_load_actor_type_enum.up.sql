DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'actor_type_enum') THEN
  CREATE TYPE actor_type_enum AS ENUM ('USER','KOL', 'ADMIN', 'PF');
END IF;
END$$;

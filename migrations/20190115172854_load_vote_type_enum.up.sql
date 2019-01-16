DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vote_type_enum') THEN
  CREATE TYPE vote_type_enum AS ENUM ('DOWN','UP');
END IF;
END$$;

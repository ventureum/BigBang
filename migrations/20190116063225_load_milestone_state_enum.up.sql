DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'milestone_state_enum') THEN
  CREATE TYPE milestone_state_enum AS ENUM ('COMPLETE','IN_PROGRESS', 'PENDING');
END IF;
END$$;

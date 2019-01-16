DO $$
BEGIN
IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'milestone_state_enum') THEN
  DROP TYPE milestone_state_enum;
END IF;
END$$;

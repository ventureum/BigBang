DO $$
BEGIN
IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'actor_type_enum') THEN
  DROP TYPE actor_type_enum;
END IF;
END$$;

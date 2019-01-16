DO $$
BEGIN
IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'actor_profile_status_enum') THEN
  DROP TYPE actor_profile_status_enum;
END IF;
END$$;

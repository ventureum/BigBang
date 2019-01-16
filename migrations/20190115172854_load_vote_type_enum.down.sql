DO $$
BEGIN
IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'vote_type_enum') THEN
  DROP TYPE vote_type_enum;
END IF;
END$$;

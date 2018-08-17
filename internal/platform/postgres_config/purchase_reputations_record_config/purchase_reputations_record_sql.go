package purchase_reputations_record_config

const UPSERT_PURCHASE_REPUTATIONS_RECORD_COMMAND = `
INSERT INTO purchase_reputations_records
(
  purchaser,
  vetx,
  reputations
)
VALUES 
(
  :purchaser, 
  :vetx,
  :reputations
);
`

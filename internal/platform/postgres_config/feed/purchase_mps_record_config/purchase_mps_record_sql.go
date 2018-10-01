package purchase_mps_record_config

const UPSERT_PURCHASE_MPS_RECORD_COMMAND = `
INSERT INTO purchase_mps_records
(
  purchaser,
  vetx,
  mps
)
VALUES 
(
  :purchaser, 
  :vetx,
  :mps
);
`

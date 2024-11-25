BEGIN;

INSERT INTO
  authors (
    name,
    description,
    created_by_id,
    updated_by_id,
    created_at,
    updated_at,
    deleted_at
  )
VALUES
  (
    'Karl Marx',
    'Desc Karl Marx',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Friedrich Engels',
    'Desc Friedrich Engels',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Vladimir Lenin',
    'Desc Vladimir Lenin',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Joseph Stalin',
    'Desc Joseph Stalin',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Mao Zedong',
    'Desc Mao Zedong',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  );

-- Reset the sequence to the maximum id value
-- Replace `authors_id_seq` with the actual sequence name if different
SELECT
  setval(
    'authors_id_seq',
    (
      SELECT
        MAX(id)
      FROM
        authors
    )
  );

COMMIT;
BEGIN;

INSERT INTO
  categories (
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
    'Physcology',
    'Desc 1',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Science',
    'Desc 2',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Fiction',
    'Desc 3',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Non-Fiction',
    'Desc 4',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Biography',
    'Desc 5',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'History',
    'Desc 6',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Self-Help',
    'Desc 7',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Cooking',
    'Desc 8',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Travel',
    'Desc 9',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Children',
    'Desc 10',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Religion',
    'Desc 11',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Business',
    'Desc 12',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  ),
  (
    'Health',
    'Desc 13',
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    NULL
  );

-- Reset the sequence to the maximum id value
-- Replace `categories_id_seq` with the actual sequence name if different
SELECT
  setval(
    'categories_id_seq',
    (
      SELECT
        MAX(id)
      FROM
        categories
    )
  );

COMMIT;
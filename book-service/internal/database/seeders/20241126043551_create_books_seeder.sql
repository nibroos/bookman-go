BEGIN;

INSERT INTO
  books (
    name,
    description,
    category_id,
    author_id,
    created_by_id,
    created_at,
    updated_by_id,
    updated_at,
    deleted_at
  )
VALUES
  (
    'War and Peace',
    'Desc 1',
    1,
    1,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    '2024-10-30 04:35:51',
    NULL
  ),
  (
    'The Great Gatsby',
    'Desc 2',
    2,
    2,
    NULL,
    '2024-10-30 04:35:51',
    NULL,
    '2024-10-30 04:35:51',
    NULL
  );

COMMIT;
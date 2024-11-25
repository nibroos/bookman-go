BEGIN;

INSERT INTO
  books (
    name,
    description,
    category_id,
    author_id,
    created_by_id,
    updated_by_id,
    created_at,
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
    NULL
  ),
  -- Reset the sequence to the maximum id value
  -- Replace `books_id_seq` with the actual sequence name if different
SELECT
  setval(
    'books_id_seq',
    (
      SELECT
        MAX(id)
      FROM
        books
    )
  );

COMMIT;
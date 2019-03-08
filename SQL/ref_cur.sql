 -- Procedure that returns multiple result sets (cursors)
   CREATE OR REPLACE FUNCTION show_cur(num INTEGER) RETURNS SETOF refcursor AS $$
    DECLARE
      r_cur refcursor;           -- Declare cursor variables
    BEGIN
      OPEN r_cur FOR SELECT name, p_sex FROM person  where id > num;  -- Open the first cursor
      RETURN NEXT r_cur;                            -- Return the cursor to the caller
    END;
    $$ LANGUAGE plpgsql;

--Improve
    CREATE OR REPLACE FUNCTION show_cur2(r_cur REFCURSOR, edge INTEGER) RETURNS SETOF refcursor AS $$
    DECLARE
      num INTEGER := 0;          -- Declare cursor variables
    BEGIN
      OPEN r_cur FOR SELECT name, p_sex FROM person  where id > edge;  -- Open the first cursor
      RETURN NEXT r_cur;                            -- Return the cursor to the caller
    END;
    $$ LANGUAGE plpgsql;

-- Test
BEGIN;
SELECT show_cur2('cur', 2);
FETCH ALL IN "cur";
COMMIT;
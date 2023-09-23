CREATE OR REPLACE FUNCTION concat_name_and_stack(apelido VARCHAR, nome VARCHAR, stack VARCHAR[])
RETURNS VARCHAR
AS $$
BEGIN
  RETURN apelido || ' ' || nome || ' ' || COALESCE(array_to_string(stack, ', '), '');
END;
$$ LANGUAGE plpgsql IMMUTABLE;

CREATE TABLE pessoas (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  apelido VARCHAR(32) UNIQUE NOT NULL,
  nome VARCHAR(100) NOT NULL,
  nascimento VARCHAR(10) NOT NULL,
  stack VARCHAR(32)[],
  busca VARCHAR(200) GENERATED ALWAYS AS (
    concat_name_and_stack(apelido, nome, stack)
  ) STORED
);

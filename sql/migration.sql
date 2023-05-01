CREATE TABLE IF NOT EXISTS public.album ( 
    id serial4 NOT NULL, 
    title varchar(128) NOT NULL, 
    artist varchar(255) NOT NULL, 
    price numeric(5, 2) NOT NULL, 
    CONSTRAINT album_pkey 
        PRIMARY KEY (id) 
);

INSERT INTO public.album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
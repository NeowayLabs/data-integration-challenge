create extension if not exists "uuid-ossp";

create table if not exists companies
  ( id uuid not null default uuid_generate_v4()
  , name text not null
  , zip char(5) not null
  , website text
  , primary key(id)
  );

create unique index if not exists companies_uk
  on companies(name, zip);

create index if not exists companies_name_idx
  on companies(name);

create index if not exists companies_zip_idx
  on companies(zip);

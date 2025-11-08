create table public.users
(
    id           uuid primary key     default gen_random_uuid(),
    supabase_id  uuid        not null unique,
    email        text,
    display_name text,
    created_at   timestamptz not null default now(),
    updated_at   timestamptz not null default now()
);

-- +goose Up
-- +goose StatementBegin

create table if not exists storage.events
(
    id         text                          not null check ( storage.text_non_empty_trimmed_text(id) ),
    version    int         default 0         not null check ( version >= 0 ),
    name       text                          not null check ( storage.text_non_empty_trimmed_text(name) ),
    content    jsonb                         not null,
    status     text        default 'pending' not null check (
        status in ('pending', 'processing', 'completed', 'failed')
        ),
    retries    int         default 0         not null check ( retries >= 0 ),
    expires_at timestamptz,
    created_at timestamptz default now()     not null,
    constraint events_id_pk primary key (id),
    constraint events_id_version_uq unique (id, version),
    constraint events_name_uq unique (name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists storage.events;

-- +goose StatementEnd

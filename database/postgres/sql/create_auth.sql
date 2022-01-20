CREATE TABLE refresh_sessions
(
    "id"            SERIAL PRIMARY KEY,
    "user_id"       integer REFERENCES users (id) ON DELETE CASCADE,
    "refresh_token" varchar(128)             NOT NULL,
    "user_agent"    varchar(196)             NOT NULL, /* user-agent */
    "fingerprint"   varchar(196)             NOT NULL,
    "ip"            varchar(16)              NOT NULL,
    "expires_in"    bigint                   NOT NULL,
    "created_at"    timestamp with time zone NOT NULL DEFAULT now()
);

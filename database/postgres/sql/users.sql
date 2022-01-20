CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(32)              NOT NULL UNIQUE,
    password   VARCHAR(32)              NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE profiles
(
    id              SERIAL PRIMARY KEY,
    user_id         INT REFERENCES users (id) ON DELETE CASCADE,
    sex             VARCHAR(10) CHECK ( sex IN ('male', 'female', 'other') ),
    sex_orientation VARCHAR(10) CHECK ( sex_orientation IN
                                        ('hetero', 'gay',
                                         'lesbian', 'bisexual',
                                         'asexual', 'pansexual',
                                         'demisexual', 'undecided') ),
    email           VARCHAR(100)             NOT NULL UNIQUE,
    firstname       VARCHAR(32)              NOT NULL,
    lastname        VARCHAR(32),
    bio             varchar(418),
    age             INT2 CHECK ( age >= 18 ) NOT NULL
);

CREATE TABLE interests
(
    id   SERIAL PRIMARY KEY,
    name varchar(24) NOT NULL
);

CREATE TABLE users_interests
(
    user_id     INT REFERENCES users (id) ON DELETE CASCADE,
    interest_id INT REFERENCES interests (id) ON DELETE CASCADE
)

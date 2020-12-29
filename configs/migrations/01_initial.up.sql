-- Table of payments received per project phase
CREATE TABLE payments
(
    id         SERIAL PRIMARY KEY,
    phase_id   INTEGER                  NOT NULL,

    amount     INTEGER                  NOT NULL,
    payed      BOOLEAN DEFAULT FALSE,

    date_due   TIMESTAMP WITH TIME ZONE NOT NULL,
    date_payed TIMESTAMP WITH TIME ZONE
);

-- Table of payments payed to employed people
CREATE TABLE user_payment
(
    id         SERIAL PRIMARY KEY,
    user_id    TEXT                     NOT NULL,

    amount     INTEGER                  NOT NULL,
    date_payed TIMESTAMP WITH TIME ZONE NOT NULL
);

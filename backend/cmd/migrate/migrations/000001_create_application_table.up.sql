CREATE TABLE IF NOT EXISTS application(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    created_by VARCHAR(255) NOT NULL
);
-- insert into application (name, created_at, created_by) values ('Account Service', now(), 'mberwanger@datalift.io');

-- modified_by
-- owned
-- date_created #date() CONSTRAINT ers_nn_chg_part_date_created NOT NULL,
-- date_deleted #date() NULL,
-- deleted_by INTEGER NULL,
-- date_suspended #date() NULL,
-- date_locked #date() NULL,
-- date_modified #date() CONSTRAINT ers_nn_chg_part_date_deleted NOT NULL,
-- created_by INTEGER CONSTRAINT ers_nn_chg_part_created_by NOT NULL,
-- owned_by INTEGER CONSTRAINT ers_nn_chg_part_owned_by NOT NULL,
-- modified_by INTEGER CONSTRAINT ers_nn_chg_part_modified_by NOT NULL,
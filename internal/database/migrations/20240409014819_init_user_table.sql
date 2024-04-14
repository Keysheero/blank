-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE Student (
                         Id UUID PRIMARY KEY,
                         Email VARCHAR(255) UNIQUE,
                         Password VARCHAR(255),
                         Status VARCHAR(50),
                         Attendance INT
);

CREATE TABLE Section (
                         Id UUID PRIMARY KEY,
                         Name VARCHAR(255),
                         Capacity INT,
                         Description VARCHAR(255),
                         Location VARCHAR(255)
);

CREATE TABLE StudentSection (
                                Id UUID PRIMARY KEY,
                                Student_Id UUID REFERENCES Student(Id),
                                Section_Id UUID REFERENCES Section(Id),
                                Registration_Date DATE
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE if exists Student;
DROP TABLE if exists Section;
DROP TABLE if exists StudentSection;

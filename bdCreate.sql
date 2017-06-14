CREATE TABLE generalInfo(
    sampleKey bigserial primary key,
    data jsonb,
    cameToServer timestamp default current_timestamp
);

CREATE TABLE agregatedInfo(
    sampleKey bigserial primary key,
    data jsonb,
    cameToServer timestamp default current_timestamp    
);

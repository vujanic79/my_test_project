INSERT INTO task_status VALUES
    ('ACTIVE'),
    ('PAUSED'),
    ('COMPLETED'),
    ('DELETED');

INSERT INTO app_user VALUES
(
     '32c5982d-b0a7-4756-94f0-11a468ffe05d'::uuid,
     NOW(),
     NOW(),
     'Milan',
     'Vujanic',
     'milanmvujanic@gmail.com'
);

INSERT INTO task VALUES
(
    '012d38ae-f234-46f8-bea8-4a3f43b0abb0'::uuid,
    NOW() AT TIME ZONE 'UTC',
    NOW() AT TIME ZONE 'UTC',
    'Task 001',
    'Task description 001',
    'ACTIVE',
    NOW() AT TIME ZONE 'UTC' + interval '7 day',
    '32c5982d-b0a7-4756-94f0-11a468ffe05d'::uuid
);
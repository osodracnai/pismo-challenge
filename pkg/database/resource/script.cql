CREATE KEYSPACE IF NOT EXISTS pismo WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};


create table pismo.accounts (
    account_id      uuid,
    document_number int,
    primary key (account_id)
)
    with caching = {'keys': 'ALL', 'rows_per_partition': '10'}
     and compaction = {'max_threshold': '32', 'min_threshold': '4', 'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
     and compression = {'class': 'org.apache.cassandra.io.compress.LZ4Compressor', 'chunk_length_in_kb': '64'}
     and dclocal_read_repair_chance = 0.1;

create table pismo.transaction (
        transaction_id     uuid,
        account_id         uuid,
        operation_type_id     int,
        amount bigint,
        event_date timestamp,
        primary key (transaction_id, event_date)
)
    with caching = {'keys': 'ALL', 'rows_per_partition': '10'}
     and compaction = {'max_threshold': '32', 'min_threshold': '4', 'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
     and compression = {'class': 'org.apache.cassandra.io.compress.LZ4Compressor', 'chunk_length_in_kb': '64'}
     and dclocal_read_repair_chance = 0.1;
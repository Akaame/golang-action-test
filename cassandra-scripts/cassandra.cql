CREATE KEYSPACE IF NOT EXISTS messaging
WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 };

# hierarchy of messages what is a better way of achieving the same behavior?
CREATE TABLE hierarchy (
  parent_id bigint,
  child_id bigint,
  child_ordering_key TIMEUUID,
  PRIMARY KEY (parent_id, child_ordering_key)
) WITH CLUSTERING ORDER BY(child_ordering_key ASC);

# messages by snowflake message_id
CREATE TABLE messages (
  message_id bigint,
  sender_id bigint,
  sender_name TEXT,
  message TEXT, # soft deletes would say this to "(Deleted)"
  ts TIMESTAMP,
  PRIMARY KEY (message_id)
);

# message by sender
# displayed when someone opens profile for e.g
CREATE TABLE messages_by_sender (
  sender_id bigint,
  ordering_key TIMEUUID
  sender_name TEXT,
  message TEXT,
  message_id bigint,
  original_message_id bigint,
  ts TIMESTAMP,
  PRIMARY KEY (sender_id, ordering_key)
) WITH CLUSTERING ORDER BY (ordering_key DESC);

# https://ably.com/blog/cassandra-counter-columns-nice-in-theory-hazardous-in-practice
# Use counter column carefully
# will be updated frequently needs special treatment
# what is a better way to handle this within messages table
# given that it will be frequently updated?
CREATE TABLE upvotes (
  message_id bigint,
  upvotes int,
  PRIMARY KEY (message_id)
);

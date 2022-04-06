-- Resource: wafv2.ipsets
DROP TABLE IF EXISTS aws_wafv2_ipsets;

-- Resource: wafv2.regex_pattern_sets
DROP TABLE IF EXISTS aws_wafv2_regex_pattern_sets;

-- Resource: iam.virtual_mfa_devices
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices DROP CONSTRAINT aws_iam_virtual_mfa_devices_pk;
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices ADD CONSTRAINT aws_iam_virtual_mfa_devices_pk PRIMARY KEY (serial_number,enable_date);

-- Resource: ec2.egress_only_internet_gateways
DROP TABLE IF EXISTS aws_ec2_egress_only_internet_gateways;

-- Resource: qldb.ledgers
DROP TABLE IF EXISTS aws_qldb_ledger_journal_kinesis_streams;
DROP TABLE IF EXISTS aws_qldb_ledger_journal_s3_exports;
DROP TABLE IF EXISTS aws_qldb_ledgers;
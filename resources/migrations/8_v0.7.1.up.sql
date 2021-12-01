ALTER TABLE IF EXISTS "aws_config_configuration_recorders" ADD COLUMN status_last_error_code text,
                                                           ADD COLUMN status_last_error_message text,
                                                           ADD COLUMN status_last_start_time timestamp without time zone,
                                                           ADD COLUMN status_last_status text,
                                                           ADD COLUMN status_last_status_change_time timestamp without time zone,
                                                           ADD COLUMN status_last_stop_time timestamp without time zone,
                                                           ADD COLUMN status_recording boolean;


ALTER TABLE IF EXISTS "aws_wafv2_web_acls" ADD COLUMN logging_configuration text[];
ALTER TABLE IF EXISTS "aws_waf_web_acls" ADD COLUMN logging_configuration text[];
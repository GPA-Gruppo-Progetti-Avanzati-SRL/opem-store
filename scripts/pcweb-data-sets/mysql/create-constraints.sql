-- company_grouping
alter table company_grouping add
    CONSTRAINT `FK352C4E85FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

alter table   company_grouping_companies add
    CONSTRAINT `FKA7E88A16E4E522A` FOREIGN KEY (`companies`) REFERENCES `companies` (`id`);

alter table   company_grouping_companies add
    CONSTRAINT `FKA7E88A176800F7` FOREIGN KEY (`company_grouping`) REFERENCES `company_grouping` (`id`);

    -- group_authorities
alter table   group_authorities add
    CONSTRAINT `FK771BA1612D01B654` FOREIGN KEY (`authority`) REFERENCES `authorities` (`role_name`);
alter table   group_authorities add
    CONSTRAINT `group_authorities_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`);

    -- users
alter table   users add
    CONSTRAINT `FK6A68E08FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- group_members
alter table   group_members add
    CONSTRAINT `FK3B9C77595A14A413` FOREIGN KEY (`username`) REFERENCES `users` (`username`);
alter table   group_members add
    CONSTRAINT `group_members_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`);

    -- tb_province
alter table   tb_province add
    CONSTRAINT `FKC647A9E12CB6FEF4` FOREIGN KEY (`ID_NAZIONE`) REFERENCES `tb_nazioni` (`COD_UIC`);

    -- tb_addresses
alter table   tb_addresses add
    CONSTRAINT `FK576CD311D13CF34B` FOREIGN KEY (`district`) REFERENCES `tb_province` (`SIGLA`);
alter table   tb_addresses add
    CONSTRAINT `FK576CD311DD32BE5B` FOREIGN KEY (`nation`) REFERENCES `tb_nazioni` (`COD_UIC`);

    -- tb_branch
alter table   tb_branch add
    CONSTRAINT `FK42549C93D13CF34B` FOREIGN KEY (`district`) REFERENCES `tb_province` (`SIGLA`);
alter table   tb_branch add
    CONSTRAINT `FK42549C93FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);
alter table   tb_branch add
    CONSTRAINT `FK_NATION` FOREIGN KEY (`nation`) REFERENCES `tb_nazioni` (`COD_UIC`);

    -- tb_persons
alter table   tb_persons add
    CONSTRAINT `FKD79936ED5B2C6A6` FOREIGN KEY (`birth_prov`) REFERENCES `tb_province` (`SIGLA`);
alter table   tb_persons add
    CONSTRAINT `FKD79936EDD3DD44B5` FOREIGN KEY (`birth_nat`) REFERENCES `tb_nazioni` (`COD_UIC`);
alter table   tb_persons add
    CONSTRAINT `FKD79936EDFCD0B283` FOREIGN KEY (`address1`) REFERENCES `tb_addresses` (`id`);
alter table   tb_persons add
    CONSTRAINT `FKD79936EDFCD0B284` FOREIGN KEY (`address2`) REFERENCES `tb_addresses` (`id`);
alter table   tb_persons add
    CONSTRAINT `FKD79936EDFCD0B285` FOREIGN KEY (`address3`) REFERENCES `tb_addresses` (`id`);

    -- tb_cards
alter table   tb_cards add
    CONSTRAINT `FKF9E84E52546B4098` FOREIGN KEY (`person`) REFERENCES `tb_persons` (`id`);
alter table   tb_cards add
    CONSTRAINT `FKF9E84E52FCD0B283` FOREIGN KEY (`address1`) REFERENCES `tb_addresses` (`id`);
alter table   tb_cards add
    CONSTRAINT `FKF9E84E52FCD0B284` FOREIGN KEY (`address2`) REFERENCES `tb_addresses` (`id`);
alter table   tb_cards add
    CONSTRAINT `FKF9E84E52FCD0B285` FOREIGN KEY (`address3`) REFERENCES `tb_addresses` (`id`);

    -- tb_comuni
alter table   tb_comuni add
    CONSTRAINT `FK43E4BEC09F64FB6A` FOREIGN KEY (`PROVINCIA`) REFERENCES `tb_province` (`SIGLA`);

    -- tb_dom_param
alter table   tb_dom_param add
    CONSTRAINT `FK4725DEBFFB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_file_header_errors
alter table   tb_file_header_errors add
    CONSTRAINT `FK64E6154B6ACA1660` FOREIGN KEY (`errors`) REFERENCES `tb_file_err` (`id`);

alter table   tb_file_header_errors add
    CONSTRAINT `FK64E6154BA4815876` FOREIGN KEY (`tb_file_header`) REFERENCES `tb_file_header` (`id_file`);

    -- tb_operation_log
alter table   tb_operation_log add
    CONSTRAINT `FK72E5A45BFAB6A38E` FOREIGN KEY (`card`) REFERENCES `tb_cards` (`id`);

    -- tb_file_records
alter table   tb_file_records add
    CONSTRAINT `FK_tb_file_records_tb_operation_log` FOREIGN KEY (`operation_log`) REFERENCES `tb_operation_log` (`id`);

    -- tb_file_header_records
alter table   tb_file_header_records add
    CONSTRAINT `FKD08B34621FD5F63A` FOREIGN KEY (`records`) REFERENCES `tb_file_records` (`id`);
alter table   tb_file_header_records add
    CONSTRAINT `FKD08B3462A4815876` FOREIGN KEY (`tb_file_header`) REFERENCES `tb_file_header` (`id_file`);

    -- tb_prod
alter table   tb_prod add
    CONSTRAINT `FKA4FD2288FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

   --  tb_mag_param
alter table   tb_mag_param add
    CONSTRAINT `FK91D7EE1026025532` FOREIGN KEY (`branch`) REFERENCES `tb_branch` (`id`);
alter table   tb_mag_param add
    CONSTRAINT `FK91D7EE103A1BF45C` FOREIGN KEY (`prodotto`) REFERENCES `tb_prod` (`id`);
alter table   tb_mag_param add
    CONSTRAINT `FK91D7EE10FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_mail_addresses
alter table   tb_mail_addresses add
    CONSTRAINT `FK9E09F90BFB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_mail_logs
alter table   tb_mail_logs add
    CONSTRAINT `FK10E8CAA6FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_mail_logs_sendtos
alter table   tb_mail_logs_sendtos add
    CONSTRAINT `FKCF71D037BACC8E9B` FOREIGN KEY (`mail_log`) REFERENCES `tb_mail_logs` (`id`);

    -- tb_mail_templates
alter table   tb_mail_templates add
    CONSTRAINT `FKE00840A2FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_mailing_events
alter table   tb_mailing_events add
    CONSTRAINT `FKCA4086FEFAA564C1` FOREIGN KEY (`mail_template`) REFERENCES `tb_mail_templates` (`id`);
alter table   tb_mailing_events add
    CONSTRAINT `FKCA4086FEFB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_mailing_events_batch_events
alter table   tb_mailing_events_batch_events add
    CONSTRAINT `FK7E50015FFC573A23` FOREIGN KEY (`mailing_event`) REFERENCES `tb_mailing_events` (`id`);

    -- tb_mailing_events_batch_types
alter table   tb_mailing_events_batch_types add
    CONSTRAINT `FK914B1773FC573A23` FOREIGN KEY (`mailing_event`) REFERENCES `tb_mailing_events` (`id`);

   --  tb_mailing_events_mail_addresses
alter table   tb_mailing_events_mail_addresses add
    CONSTRAINT `FK4DD74E7B437DA029` FOREIGN KEY (`mail_addresses`) REFERENCES `tb_mail_addresses` (`id`);
alter table   tb_mailing_events_mail_addresses add
    CONSTRAINT `FK4DD74E7BCA5DCCFB` FOREIGN KEY (`tb_mailing_events`) REFERENCES `tb_mailing_events` (`id`);

    -- tb_mailing_lists
alter table   tb_mailing_lists add
    CONSTRAINT `FK69FBED10FB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

    -- tb_mailing_events_mailing_lists
alter table   tb_mailing_events_mailing_lists add
    CONSTRAINT `FKA97629A0366949FC` FOREIGN KEY (`mailing_lists`) REFERENCES `tb_mailing_lists` (`id`);
alter table   tb_mailing_events_mailing_lists add
    CONSTRAINT `FKA97629A0CA5DCCFB` FOREIGN KEY (`tb_mailing_events`) REFERENCES `tb_mailing_events` (`id`);

    -- tb_mailing_lists_mail_addresses
alter table   tb_mailing_lists_mail_addresses add
    CONSTRAINT `FKDB64C229437DA029` FOREIGN KEY (`mail_addresses`) REFERENCES `tb_mail_addresses` (`id`);
alter table   tb_mailing_lists_mail_addresses add
    CONSTRAINT `FKDB64C229A3CE51EB` FOREIGN KEY (`tb_mailing_lists`) REFERENCES `tb_mailing_lists` (`id`);

    -- tb_prod_funct
alter table   tb_prod_funct add
    CONSTRAINT `FK895BB5B9762740BC` FOREIGN KEY (`prod`) REFERENCES `tb_prod` (`id`);

    -- tb_sub_companies
alter table   tb_sub_companies add
    CONSTRAINT `FK6908D28B10C6BBA` FOREIGN KEY (`address`) REFERENCES `tb_addresses` (`id`);
alter table   tb_sub_companies add
    CONSTRAINT `FK6908D28BFB2BACCC` FOREIGN KEY (`company`) REFERENCES `companies` (`id`);

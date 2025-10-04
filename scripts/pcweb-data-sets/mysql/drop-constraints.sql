-- company_grouping
alter table company_grouping DROP FOREIGN KEY  `FK352C4E85FB2BACCC`;

alter table   company_grouping_companies DROP FOREIGN KEY 
     `FKA7E88A16E4E522A`;

alter table   company_grouping_companies DROP FOREIGN KEY 
      `FKA7E88A176800F7`;
    

    -- group_authorities
alter table   group_authorities DROP FOREIGN KEY 
      `FK771BA1612D01B654`;
   
alter table   group_authorities DROP FOREIGN KEY 
      `group_authorities_ibfk_1`;
  

    -- users
alter table   users DROP FOREIGN KEY 
      `FK6A68E08FB2BACCC`;
   

    -- group_members
alter table   group_members DROP FOREIGN KEY 
      `FK3B9C77595A14A413`;
    
alter table   group_members DROP FOREIGN KEY 
      `group_members_ibfk_1`;
    
    -- tb_province
alter table   tb_province DROP FOREIGN KEY 
      `FKC647A9E12CB6FEF4`;
    

    -- tb_addresses
alter table   tb_addresses DROP FOREIGN KEY 
      `FK576CD311D13CF34B`;
alter table   tb_addresses DROP FOREIGN KEY 
      `FK576CD311DD32BE5B`;
   

    -- tb_branch
alter table   tb_branch DROP FOREIGN KEY 
      `FK42549C93D13CF34B`;
   
alter table   tb_branch DROP FOREIGN KEY 
      `FK42549C93FB2BACCC`;
 
alter table   tb_branch DROP FOREIGN KEY 
      `FK_NATION`;
   

    -- tb_persons
alter table   tb_persons DROP FOREIGN KEY 
      `FKD79936ED5B2C6A6`;
 
alter table   tb_persons DROP FOREIGN KEY 
      `FKD79936EDD3DD44B5`;

alter table   tb_persons DROP FOREIGN KEY
      `FKD79936EDFCD0B283`;

alter table   tb_persons DROP FOREIGN KEY
      `FKD79936EDFCD0B284`;

alter table   tb_persons DROP FOREIGN KEY 
      `FKD79936EDFCD0B285`;
    

    -- tb_cards
alter table   tb_cards DROP FOREIGN KEY 
      `FKF9E84E52546B4098`;

alter table   tb_cards DROP FOREIGN KEY
      `FKF9E84E52FCD0B283`;

alter table   tb_cards DROP FOREIGN KEY 
      `FKF9E84E52FCD0B284`;

alter table   tb_cards DROP FOREIGN KEY
      `FKF9E84E52FCD0B285`;

    -- tb_comuni
alter table   tb_comuni DROP FOREIGN KEY 
      `FK43E4BEC09F64FB6A`;
  
    -- tb_dom_param
alter table   tb_dom_param DROP FOREIGN KEY 
      `FK4725DEBFFB2BACCC`;
    

    -- tb_file_header_errors
alter table   tb_file_header_errors DROP FOREIGN KEY 
      `FK64E6154B6ACA1660`;
   
alter table   tb_file_header_errors DROP FOREIGN KEY 
      `FK64E6154BA4815876`;
    

    -- tb_operation_log
alter table   tb_operation_log DROP FOREIGN KEY 
      `FK72E5A45BFAB6A38E`;
   

    -- tb_file_records
alter table   tb_file_records DROP FOREIGN KEY 
      `FK_tb_file_records_tb_operation_log`;
    
    -- tb_file_header_records
alter table   tb_file_header_records DROP FOREIGN KEY 
      `FKD08B34621FD5F63A`;
   
alter table   tb_file_header_records DROP FOREIGN KEY 
      `FKD08B3462A4815876`;
  

    -- tb_prod
alter table   tb_prod DROP FOREIGN KEY 
      `FKA4FD2288FB2BACCC`;
   

   --  tb_mag_param
alter table   tb_mag_param DROP FOREIGN KEY 
      `FK91D7EE1026025532`;
    
alter table   tb_mag_param DROP FOREIGN KEY 
      `FK91D7EE103A1BF45C`;
  
alter table   tb_mag_param DROP FOREIGN KEY 
      `FK91D7EE10FB2BACCC`;
    

    -- tb_mail_addresses
alter table   tb_mail_addresses DROP FOREIGN KEY 
      `FK9E09F90BFB2BACCC`;
    
    -- tb_mail_logs
alter table   tb_mail_logs DROP FOREIGN KEY 
      `FK10E8CAA6FB2BACCC`;
   

    -- tb_mail_logs_sendtos
alter table   tb_mail_logs_sendtos DROP FOREIGN KEY 
      `FKCF71D037BACC8E9B`;
  

    -- tb_mail_templates
alter table   tb_mail_templates DROP FOREIGN KEY 
      `FKE00840A2FB2BACCC`;
   

    -- tb_mailing_events
alter table   tb_mailing_events DROP FOREIGN KEY 
      `FKCA4086FEFAA564C1`;
    
alter table   tb_mailing_events DROP FOREIGN KEY 
      `FKCA4086FEFB2BACCC`;
   

    -- tb_mailing_events_batch_events
alter table   tb_mailing_events_batch_events DROP FOREIGN KEY 
      `FK7E50015FFC573A23`;
    

    -- tb_mailing_events_batch_types
alter table   tb_mailing_events_batch_types DROP FOREIGN KEY 
      `FK914B1773FC573A23`;
    
   --  tb_mailing_events_mail_addresses
alter table   tb_mailing_events_mail_addresses DROP FOREIGN KEY 
      `FK4DD74E7B437DA029`;
    
alter table   tb_mailing_events_mail_addresses DROP FOREIGN KEY 
      `FK4DD74E7BCA5DCCFB`;
    

    -- tb_mailing_lists
alter table   tb_mailing_lists DROP FOREIGN KEY 
      `FK69FBED10FB2BACCC`;
    

    -- tb_mailing_events_mailing_lists
alter table   tb_mailing_events_mailing_lists DROP FOREIGN KEY 
      `FKA97629A0366949FC`;
    
alter table   tb_mailing_events_mailing_lists DROP FOREIGN KEY 
      `FKA97629A0CA5DCCFB`;
  

    -- tb_mailing_lists_mail_addresses
alter table   tb_mailing_lists_mail_addresses DROP FOREIGN KEY 
      `FKDB64C229437DA029`;
    
alter table   tb_mailing_lists_mail_addresses DROP FOREIGN KEY 
      `FKDB64C229A3CE51EB`;
    

    -- tb_prod_funct
alter table   tb_prod_funct DROP FOREIGN KEY 
      `FK895BB5B9762740BC`;
    
    -- tb_sub_companies
alter table   tb_sub_companies DROP FOREIGN KEY 
      `FK6908D28B10C6BBA`;
alter table   tb_sub_companies DROP FOREIGN KEY 
      `FK6908D28BFB2BACCC`;
   

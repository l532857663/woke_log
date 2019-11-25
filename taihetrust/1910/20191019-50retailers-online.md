# 20190919 沧州50户零售户上线数据库操作

### 创建店铺

\copy store(id,name,address,mch_id,cz_store_no,province,city,district) FROM '/tmp/store.csv' DELIMITER ',' CSV;

### 创建店铺帐户

insert into account (name,address,status) values ('时利和烟酒店','河北省沧州市新华区建设北街新华区人民法院门市楼一排11号','1');
insert into account (name,address,status) values ('丰裕烟酒经销处','河北省沧州市新华区交通南街交东路3号','1');
insert into account (name,address,status) values ('阁龙轩超市','河北省沧州市运河区小王庄乡胡嘴子村','1');
insert into account (name,address,status) values ('秀萍百货超市','河北省沧州市新华区南环中路立交桥北木材楼3单101室','1');
insert into account (name,address,status) values ('殿林超市','河北省沧州市运河区御河路18-1号','1');
insert into account (name,address,status) values ('洪全副食部','河北省沧州市运河区四合市场三区','1');
insert into account (name,address,status) values ('邻福便利店','河北省沧州市运河区北环西路路南运输二场宿舍9-37','1');
insert into account (name,address,status) values ('健民副食部','河北省沧州市新华区维明路沧县粮局宿舍北','1');
insert into account (name,address,status) values ('大千综合超市','河北省沧州市孟村回族自治县辛店镇辛店村辛店村沧盐路188号','1');
insert into account (name,address,status) values ('惠丰超市','河北省沧州市青县上伍乡周官屯村（17号）','1');
insert into account (name,address,status) values ('考粮酒类经销处','','1');

使用手机号 查找sso的信息


### 创建店铺设备

\copy device(store_id,device_id,remarks,uid,name,device_no,device_model) from '/data/device_insert.csv' DELIMITER ',';

### 初始化店铺库存及价格信息

insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-0565727224',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-6802376175',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-5132278216',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-3614539209',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-9405806556',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-6238034691',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-4513027737',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-7421159479',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-3968919268',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-4194580658',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;
insert into sku_stock(store_id,sku_no,spu_id,product_name,quantity,created_at,updated_at) select 'st-9995503152',spu_no,spu_id,product_name,0,now(),now() from spu where spu_type=1 on conflict do nothing;


insert into sales_sku (select 'st-0565727224', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-6802376175', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-5132278216', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-3614539209', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-9405806556', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-6238034691', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-4513027737', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-7421159479', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-3968919268', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-4194580658', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;
insert into sales_sku (select 'st-9995503152', spu_id, spu_no, null, guide_price, guide_price, created_at, updated_at from spu) on conflict do nothing;

### 更新店铺库存
\copy sku_stock (sku_no,quantity,store_id) from '/data/1003.csv' DELIMITER ',';
update sku_stock set spu_id=spu.spu_id from spu where sku_stock.sku_no = spu.spu_no and sku_stock.store_id='1003';
update sku_stock set product_name=spu.product_name from spu where sku_stock.sku_no = spu.spu_no and sku_stock.store_id='1003';
update sku_stock set guide_price=spu.guide_price from spu where sku_stock.sku_no = spu.spu_no and sku_stock.store_id='1003';
update sku_stock set unit='包' where store_id='1003';

insert into sales_sku (select '1003', spu_id, spu_no, guide_price, guide_price, created_at, updated_at from spu where spu_type=1);
insert into sales_sku (select '1003', spu_id, spu_no, guide_price, guide_price, created_at, updated_at from spu where spu_type=2);

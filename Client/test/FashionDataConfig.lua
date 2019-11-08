local title = {['FashionID']=1,['SetID']=2,['Class']=3,['Price']=4,['Color']=5,['MaleBodyMaterialPath']=6,['MaleHighBodyMaterialPath']=7,['FemaleBodyMaterialPath']=8,['FemaleHighBodyMaterialPath']=9,['MaleWeaponMaterialPath']=10,['FemaleWeaponMaterialPath']=11,['Gift']=12}
local data = {
{ 1,1,1,{{1,2,1,1500},{2,2,1,4500},{3,2,1,9000}},1,'jws/jws01_low/jws01_low_mi.mat','jws/jws01_high/jws01_hi_mi.mat','','','jws/jws01_low/jws01_w_low_mi.mat','',0 },
{ 2,1,1,{{3,1,1014001,6}},4,'jws/jws01_low/jws01_low_c01_mi.mat','jws/jws01_high/jws01_hi_c01_mi.mat','','','jws/jws01_low/jws01_w_low_c01_mi.mat','',0 },
{ 3,1,1,{{3,1,1014001,6}},2,'jws/jws01_low/jws01_low_c02_mi.mat','jws/jws01_high/jws01_hi_c02_mi.mat','','','jws/jws01_low/jws01_w_low_c02_mi.mat','',0 },
{ 4,2,1,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},5,'jws/jws02_low/jws02_low_mi.mat','jws/jws02_high/jws02_hi_mi.mat','','','jws/jws02_low/jws02_w_low_mi.mat','',1 },
{ 5,2,1,{{3,1,1014001,12}},1,'jws/jws02_low/jws02_low_c01_mi.mat','jws/jws02_high/jws02_hi_c01_mi.mat','','','jws/jws02_low/jws02_w_low_c01_mi.mat','',0 },
{ 6,3,1,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},5,'jws/jws03_low/jws03_low_mi.mat','jws/jws03_high/jws03_hi_mi.mat','','','jws/jws03_low/jws03_w_low_mi.mat','',1 },
{ 31,4,1,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'jws/jws04_low/jws04_low_mi.mat','jws/jws04_high/jws04_hi_mi.mat','','','jws/jws04_low/jws04_w_low_mi.mat','',1 },
{ 41,5,1,{{1,2,1,4000},{2,2,1,12000},{3,2,1,24000}},1,'jws/jws05_low/jws05_low_mi.mat','jws/jws05_high/jws05_hi_mi.mat','','','jws/jws05_low/jws05_w_low_mi.mat','',1 },
{ 131,6,1,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'jws/jws02_low/jws02_lunar_low_mi.mat','jws/jws02_high/jws02_lunar_hi_mi.mat','','','jws/jws02_low/jws02_lunar_w_low_mi.mat','',1 },
{ 7,1,2,{{1,2,1,1500},{2,2,1,4500},{3,2,1,9000}},8,'ynz_m/ynz_m01_low/ynz_m01_low_mi.mat','ynz_m/ynz_m01_high/ynz_m01_hi_mi.mat','ynz/ynz01_low/ynz01_low_mi.mat','ynz/ynz01_high/ynz01_hi_hair.mat,ynz/ynz01_high/ynz01_hi_mi.mat','ynz_m/ynz_m01_low/ynz_m01_w_low_mi.mat','ynz/ynz01_low/ynz01_w_low_mi.mat',0 },
{ 701,1,2,{{3,1,1014001,6}},4,'ynz_m/ynz_m01_low/ynz_m01_low_c01_mi.mat','ynz_m/ynz_m01_high/ynz_m01_hi_c01_mi.mat','','','ynz_m/ynz_m01_low/ynz_m01_w_low_c01_mi.mat','',0 },
{ 702,1,2,{{3,1,1014001,6}},3,'ynz_m/ynz_m01_low/ynz_m01_low_c02_mi.mat','ynz_m/ynz_m01_high/ynz_m01_hi_c02_mi.mat','','','ynz_m/ynz_m01_low/ynz_m01_w_low_c02_mi.mat','',0 },
{ 703,1,2,{{3,1,1014001,6}},7,'ynz_m/ynz_m01_low/ynz_m01_low_c03_mi.mat','ynz_m/ynz_m01_high/ynz_m01_hi_c03_mi.mat','','','ynz_m/ynz_m01_low/ynz_m01_w_low_c03_mi.mat','',0 },
{ 704,1,2,{{3,1,1014001,6}},9,'ynz_m/ynz_m01_low/ynz_m01_low_c04_mi.mat','ynz_m/ynz_m01_high/ynz_m01_hi_c04_mi.mat','','','ynz_m/ynz_m01_low/ynz_m01_w_low_c04_mi.mat','',0 },
{ 8,1,2,{{3,1,1014001,6}},1,'','','ynz/ynz01_low/ynz01_low_c01_mi.mat','ynz/ynz01_high/ynz01_hi_hair.mat,ynz/ynz01_high/ynz01_hi_c01_mi.mat','','ynz/ynz01_low/ynz01_w_low_c01_mi.mat',0 },
{ 9,1,2,{{3,1,1014001,6}},2,'','','ynz/ynz01_low/ynz01_low_c02_mi.mat','ynz/ynz01_high/ynz01_hi_hair.mat,ynz/ynz01_high/ynz01_hi_c02_mi.mat','','ynz/ynz01_low/ynz01_w_low_c02_mi.mat',0 },
{ 10,2,2,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},9,'ynz_m/ynz_m02_low/ynz_m02_low_mi.mat','ynz_m/ynz_m02_high/ynz_m02_hi_mi.mat','ynz/ynz02_low/ynz02_low_mi.mat','ynz/ynz02_high/ynz02_hi_mi.mat','ynz_m/ynz_m02_low/ynz_m02_w_low_mi.mat','ynz/ynz02_low/ynz02_w_low_mi.mat',1 },
{ 11,2,2,{{3,1,1014001,12}},6,'ynz_m/ynz_m02_low/ynz_m02_low_c01_mi.mat','ynz_m/ynz_m02_high/ynz_m02_hi_c01_mi.mat','ynz/ynz02_low/ynz02_low_c01_mi.mat','ynz/ynz02_high/ynz02_hi_c01_mi.mat','ynz_m/ynz_m02_low/ynz_m02_w_low_c01_mi.mat','ynz/ynz02_low/ynz02_w_low_c01_mi.mat',0 },
{ 20803,2,2,{{3,1,1014001,12}},10,'ynz_m/ynz_m02_low/ynz_m02_low_c03_mi.mat','ynz_m/ynz_m02_high/ynz_m02_hi_c03_mi.mat','','','ynz_m/ynz_m02_low/ynz_m02_w_low_c03_mi.mat','',0 },
{ 20804,2,2,{{3,1,1014001,12}},9,'ynz_m/ynz_m02_low/ynz_m02_low_c04_mi.mat','ynz_m/ynz_m02_high/ynz_m02_hi_c04_mi.mat','','','ynz_m/ynz_m02_low/ynz_m02_w_low_c04_mi.mat','',0 },
{ 20805,2,2,{{3,1,1014001,12}},4,'ynz_m/ynz_m02_low/ynz_m02_low_c02_mi.mat','ynz_m/ynz_m02_high/ynz_m02_hi_c02_mi.mat','','','ynz_m/ynz_m02_low/ynz_m02_w_low_c02_mi.mat','',0 },
{ 12,3,2,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},4,'','','ynz/ynz03_low/ynz03_low_mi.mat','ynz/ynz03_high/ynz03_hi_mi.mat','','ynz/ynz03_low/ynz03_w_low_mi.mat',1 },
{ 51,4,2,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'','','ynz/ynz04_low/ynz04_low_mi.mat','ynz/ynz04_high/ynz04_hi_mi.mat','','ynz/ynz04_low/ynz04_w_low_mi.mat',1 },
{ 61,5,2,{{1,2,1,4000},{2,2,1,12000},{3,2,1,24000}},1,'','','ynz/ynz05_low/ynz05_low_mi.mat','ynz/ynz05_high/ynz05_hi_mi.mat','','ynz/ynz05_low/ynz05_w_low_mi.mat',1 },
{ 141,6,2,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'','','ynz/ynz02_low/ynz02_lunar_low_mi.mat','ynz/ynz02_high/ynz02_lunar_hi_mi.mat','','ynz/ynz02_low/ynz02_lunar_w_low_mi.mat',1 },
{ 13,1,3,{{1,2,1,1500},{2,2,1,4500},{3,2,1,9000}},4,'qxs/qxs01_low/qxs01_low_mi.mat','qxs/qxs01_high/qxs01_hi_mi.mat','qxs_f/qxs_f01_low/qxs_f01_low_mi.mat','qxs_f/qxs_f01_high/qxs_f01_hi_mi.mat','qxs/qxs01_low/qxs01_w_low_mi.mat','qxs_f/qxs_f01_high/qxs_f01_w_hi_mi.mat',0 },
{ 14,1,3,{{3,1,1014001,6}},1,'qxs/qxs01_low/qxs01_low_c01_mi.mat','qxs/qxs01_high/qxs01_hi_c01_mi.mat','','','qxs/qxs01_low/qxs01_w_low_c01_mi.mat','',0 },
{ 15,1,3,{{3,1,1014001,6}},8,'qxs/qxs01_low/qxs01_low_c02_mi.mat','qxs/qxs01_high/qxs01_hi_c02_mi.mat','','','qxs/qxs01_low/qxs01_w_low_c02_mi.mat','',0 },
{ 1501,1,3,{{3,1,1014001,6}},8,'','','qxs_f/qxs_f01_low/qxs_f01_low_c01_mi.mat','qxs_f/qxs_f01_high/qxs_f01_hi_c01_mi.mat','','qxs_f/qxs_f01_low/qxs_f01_w_low_c01_mi.mat',0 },
{ 1502,1,3,{{3,1,1014001,6}},2,'','','qxs_f/qxs_f01_low/qxs_f01_low_c02_mi.mat','qxs_f/qxs_f01_high/qxs_f01_hi_c02_mi.mat','','qxs_f/qxs_f01_low/qxs_f01_w_low_c02_mi.mat',0 },
{ 1503,1,3,{{3,1,1014001,6}},1,'','','qxs_f/qxs_f01_low/qxs_f01_low_c03_mi.mat','qxs_f/qxs_f01_high/qxs_f01_hi_c03_mi.mat','','qxs_f/qxs_f01_low/qxs_f01_w_low_c03_mi.mat',0 },
{ 16,2,3,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},9,'qxs/qxs02_low/qxs02_low_mi.mat','qxs/qxs02_high/qxs02_hi_mi.mat','qxs_f/qxs_f02_low/qxs_f02_low_c05_mi.mat,qxs_f/qxs_f02_low/qxs_f02_low_bl_c05_mi.mat','qxs_f/qxs_f02_high/qxs_f02_hi_c05_mi.mat,qxs_f/qxs_f02_high/qxs_f02_hi_bl_c05_mi.mat','qxs/qxs02_low/qxs02_w_low_mi.mat','qxs_f/qxs_f02_low/qxs_f02_w_low_c05_mi.mat',1 },
{ 17,2,3,{{3,1,1014001,12}},5,'qxs/qxs02_low/qxs02_low_c01_mi.mat','qxs/qxs02_high/qxs02_hi_c01_mi.mat','qxs_f/qxs_f02_low/qxs_f02_low_c01_mi.mat,qxs_f/qxs_f02_low/qxs_f02_low_bl_c01_mi.mat','qxs_f/qxs_f02_high/qxs_f02_hi_c01_mi.mat,qxs_f/qxs_f02_high/qxs_f02_hi_bl_c01_mi.mat','qxs/qxs02_low/qxs02_w_low_c01_mi.mat','qxs_f/qxs_f02_low/qxs_f02_w_low_c01_mi.mat',0 },
{ 30803,2,3,{{3,1,1014001,12}},9,'','','qxs_f/qxs_f02_low/qxs_f02_low_c02_mi.mat,qxs_f/qxs_f02_low/qxs_f02_low_bl_c02_mi.mat','qxs_f/qxs_f02_high/qxs_f02_hi_c02_mi.mat,qxs_f/qxs_f02_high/qxs_f02_hi_bl_c02_mi.mat','','qxs_f/qxs_f02_low/qxs_f02_w_low_c02_mi.mat',0 },
{ 30804,2,3,{{3,1,1014001,12}},11,'','','qxs_f/qxs_f02_low/qxs_f02_low_mi.mat,qxs_f/qxs_f02_low/qxs_f02_low_bl_mi.mat','qxs_f/qxs_f02_high/qxs_f02_hi_mi.mat,qxs_f/qxs_f02_high/qxs_f02_hi_bl_mi.mat','','qxs_f/qxs_f02_low/qxs_f02_w_low_mi.mat',0 },
{ 30805,2,3,{{3,1,1014001,12}},1,'','','qxs_f/qxs_f02_low/qxs_f02_low_c04_mi.mat,qxs_f/qxs_f02_low/qxs_f02_low_bl_c04_mi.mat','qxs_f/qxs_f02_high/qxs_f02_hi_c04_mi.mat,qxs_f/qxs_f02_high/qxs_f02_hi_bl_c04_mi.mat','','qxs_f/qxs_f02_low/qxs_f02_w_low_c04_mi.mat',0 },
{ 18,3,3,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},4,'qxs/qxs03_low/qxs03_low_mi.mat','qxs/qxs03_high/qxs03_hi_mi.mat','','','qxs/qxs03_low/qxs03_w_low_mi.mat','',1 },
{ 71,4,3,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'qxs/qxs04_low/qxs04_low_mi.mat','qxs/qxs04_high/qxs04_hi_mi.mat','','','qxs/qxs04_low/qxs04_w_low_mi.mat','',1 },
{ 81,5,3,{{1,2,1,4000},{2,2,1,12000},{3,2,1,24000}},1,'qxs/qxs05_low/qxs05_low_mi.mat','qxs/qxs05_high/qxs05_hi_mi.mat','','','qxs/qxs05_low/qxs05_w_low_mi.mat','',1 },
{ 151,6,3,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'qxs/qxs02_low/qxs02_lunar_low_mi.mat','qxs/qxs02_high/qxs02_lunar_hi_mi.mat','','','qxs/qxs02_low/qxs02_lunar_w_low_mi.mat','',1 },
{ 19,1,4,{{1,2,1,1500},{2,2,1,4500},{3,2,1,9000}},5,'gdj/gdj01_low/gdj01_low_mi.mat','gdj/gdj01_high/gdj01_hi_mi.mat','','','gdj/gdj01_low/gdj01_w_low_mi.mat','',0 },
{ 20,1,4,{{3,1,1014001,6}},1,'gdj/gdj01_low/gdj01_low_c01_mi.mat','gdj/gdj01_high/gdj01_hi_c01_mi.mat','','','gdj/gdj01_low/gdj01_w_low_c01_mi.mat','',0 },
{ 21,1,4,{{3,1,1014001,6}},2,'gdj/gdj01_low/gdj01_low_c02_mi.mat','gdj/gdj01_high/gdj01_hi_c02_mi.mat','','','gdj/gdj01_low/gdj01_w_low_c02_mi.mat','',0 },
{ 22,2,4,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},9,'gdj/gdj02_low/gdj02_low_mi.mat','gdj/gdj02_high/gdj02_hi_mi.mat','','','gdj/gdj02_low/gdj02_w_low_mi.mat','',1 },
{ 23,2,4,{{3,1,1014001,12}},4,'gdj/gdj02_low/gdj02_low_c01_mi.mat','gdj/gdj02_high/gdj02_hi_c01_mi.mat','','','gdj/gdj02_low/gdj02_w_low_c01_mi.mat','',0 },
{ 24,3,4,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},4,'gdj/gdj03_low/gdj03_low_mi.mat','gdj/gdj03_high/gdj03_hi_mi.mat','','','gdj/gdj03_low/gdj03_w_low_mi.mat','',1 },
{ 91,4,4,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'gdj/gdj04_low/gdj04_low_mi.mat','gdj/gdj04_high/gdj04_hi_mi.mat','','','gdj/gdj04_low/gdj04_w_low_mi.mat','',1 },
{ 101,5,4,{{1,2,1,4000},{2,2,1,12000},{3,2,1,24000}},1,'gdj/gdj05_low/gdj05_low_mi.mat','gdj/gdj05_high/gdj05_hi_mi.mat','','','gdj/gdj05_low/gdj05_w_low_mi.mat','',1 },
{ 161,6,4,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'gdj/gdj02_low/gdj02_lunar_low_mi.mat','gdj/gdj02_high/gdj02_lunar_hi_mi.mat','','','gdj/gdj02_low/gdj02_lunar_low_mi.mat','',1 },
{ 25,1,5,{{1,2,1,1500},{2,2,1,4500},{3,2,1,9000}},2,'','','zyz/zyz01_low/zyz01_low_mi.mat','zyz/zyz01_high/zyz01_hi_mi.mat','','zyz/zyz01_low/zyz01_w_low_mi.mat',0 },
{ 26,1,5,{{3,1,1014001,6}},4,'','','zyz/zyz01_low/zyz01_low_c01_mi.mat','zyz/zyz01_high/zyz01_hi_c01_mi.mat','','zyz/zyz01_low/zyz01_w_low_c01_mi.mat',0 },
{ 27,1,5,{{3,1,1014001,6}},1,'','','zyz/zyz01_low/zyz01_low_c02_mi.mat','zyz/zyz01_high/zyz01_hi_c02_mi.mat','','zyz/zyz01_low/zyz01_w_low_c02_mi.mat',0 },
{ 28,2,5,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},4,'','','zyz/zyz02_low/zyz02_low_mi.mat','zyz/zyz02_high/zyz02_hi_mi.mat','','zyz/zyz02_low/zyz02_w_low_mi.mat',1 },
{ 29,2,5,{{3,1,1014001,12}},8,'','','zyz/zyz02_low/zyz02_low_c01_mi.mat','zyz/zyz02_high/zyz02_hi_c01_mi.mat','','zyz/zyz02_low/zyz02_w_low_c01_mi.mat',0 },
{ 30,3,5,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},1,'','','zyz/zyz03_low/zyz03_low_mi.mat','zyz/zyz03_high/zyz03_hi_mi.mat','','zyz/zyz03_low/zyz03_w_low_mi.mat',1 },
{ 111,4,5,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'','','zyz/zyz04_low/zyz04_low_mi.mat','zyz/zyz04_high/zyz04_hi_mi.mat','','zyz/zyz04_low/zyz04_w_low_mi.mat',1 },
{ 121,5,5,{{1,2,1,4000},{2,2,1,12000},{3,2,1,24000}},1,'','','zyz/zyz05_low/zyz05_low_mi.mat','zyz/zyz05_high/zyz05_hi_mi.mat','','zyz/zyz05_low/zyz05_w_low_mi.mat,zyz/zyz05_low/zyz05_w_glow_low_mi.mat',1 },
{ 171,6,5,{{1,2,1,3000},{2,2,1,9000},{3,2,1,18000}},1,'','','zyz/zyz02_low/zyz02_lunar_low_mi.mat','zyz/zyz02_high/zyz02_lunar_hi_mi.mat','','zyz/zyz02_low/zyz02_lunar_w_low_mi.mat',1 },
{ 10101,7,1,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},9,'jws/jws06_low/jws06_low_mi.mat','jws/jws06_high/jws06_hi_mi.mat','','','jws/jws06_low/jws06_w_low_mi.mat','',1 },
{ 10201,7,1,{{3,1,1014001,36}},10,'jws/jws06_low/jws06_low_c01_mi.mat','jws/jws06_high/jws06_hi_c01_mi.mat','','','jws/jws06_low/jws06_w_low_c01_mi.mat','',0 },
{ 10401,7,1,{{3,1,1014001,36}},8,'jws/jws06_low/jws06_low_c03_mi.mat','jws/jws06_high/jws06_hi_c03_mi.mat','','','jws/jws06_low/jws06_w_low_c03_mi.mat','',0 },
{ 20101,7,2,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},9,'ynz_m/ynz_m06_low/ynz_m06_low_mi.mat','ynz_m/ynz_m06_high/ynz_m06_hi_mi.mat','ynz/ynz06_low/ynz06_low_mi.mat','ynz/ynz06_high/ynz06_hi_mi.mat','ynz_m/ynz_m06_low/ynz_m06_w_low_mi.mat','ynz/ynz06_low/ynz06_w_low_mi.mat',1 },
{ 20201,7,2,{{3,1,1014001,36}},10,'','','ynz/ynz06_low/ynz06_low_c01_mi.mat','ynz/ynz06_high/ynz06_hi_c01_mi.mat','','ynz/ynz06_low/ynz06_w_low_c01_mi.mat',0 },
{ 20301,7,2,{{3,1,1014001,36}},1,'','','ynz/ynz06_low/ynz06_low_c02_mi.mat','ynz/ynz06_high/ynz06_hi_c02_mi.mat','','ynz/ynz06_low/ynz06_w_low_c02_mi.mat',0 },
{ 20601,7,2,{{3,1,1014001,36}},1,'ynz_m/ynz_m06_low/ynz_m06_low_c02_mi.mat','ynz_m/ynz_m06_high/ynz_m06_hi_c02_mi.mat','','','ynz_m/ynz_m06_low/ynz_m06_w_low_c02_mi.mat','',0 },
{ 20701,7,2,{{3,1,1014001,36}},4,'ynz_m/ynz_m06_low/ynz_m06_low_c03_mi.mat','ynz_m/ynz_m06_high/ynz_m06_hi_c03_mi.mat','','','ynz_m/ynz_m06_low/ynz_m06_w_low_c03_mi.mat','',0 },
{ 30101,7,3,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},9,'qxs/qxs06_low/qxs06_low_mi.mat','qxs/qxs06_high/qxs06_hi_mi.mat','qxs_f/qxs_f06_low/qxs_f06_low_mi.mat','qxs_f/qxs_f06_high/qxs_f06_hi_mi.mat','qxs/qxs06_low/qxs06_w_low_mi.mat','qxs_f/qxs_f06_low/qxs_f06_w_low_mi.mat',1 },
{ 30201,7,3,{{3,1,1014001,36}},1,'','','qxs_f/qxs_f06_low/qxs_f06_low_c02_mi.mat','qxs_f/qxs_f06_high/qxs_f06_hi_c02_mi.mat','','qxs_f/qxs_f06_low/qxs_f06_w_low_c02_mi.mat',0 },
{ 30301,7,3,{{3,1,1014001,36}},4,'','','qxs_f/qxs_f06_low/qxs_f06_low_c03_mi.mat','qxs_f/qxs_f06_high/qxs_f06_hi_c03_mi.mat','','qxs_f/qxs_f06_low/qxs_f06_w_low_c03_mi.mat',0 },
{ 30601,7,3,{{3,1,1014001,36}},10,'qxs/qxs06_low/qxs06_low_c01_mi.mat','qxs/qxs06_high/qxs06_hi_c01_mi.mat','','','qxs/qxs06_low/qxs06_w_low_c01_mi.mat','',0 },
{ 30701,7,3,{{3,1,1014001,36}},4,'qxs/qxs06_low/qxs06_low_c02_mi.mat','qxs/qxs06_high/qxs06_hi_c02_mi.mat','','','qxs/qxs06_low/qxs06_w_low_c02_mi.mat','',0 },
{ 40101,7,4,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},9,'gdj/gdj06_low/gdj06_low_mi.mat','gdj/gdj06_high/gdj06_hi_mi.mat','','','gdj/gdj06_low/gdj06_w_low_mi.mat','',1 },
{ 40201,7,4,{{3,1,1014001,36}},10,'gdj/gdj06_low/gdj06_low_c01_mi.mat','gdj/gdj06_high/gdj06_hi_c01_mi.mat','','','gdj/gdj06_low/gdj06_w_low_c01_mi.mat','',0 },
{ 40401,7,4,{{3,1,1014001,36}},1,'gdj/gdj06_low/gdj06_low_c03_mi.mat','gdj/gdj06_high/gdj06_hi_c03_mi.mat','','','gdj/gdj06_low/gdj06_w_low_c03_mi.mat','',0 },
{ 50101,7,5,{{1,2,1,6000},{2,2,1,18000},{3,2,1,28000}},9,'','','zyz/zyz06_low/zyz06_low_mi.mat','zyz/zyz06_high/zyz06_hi_mi.mat','','zyz/zyz06_low/zyz06_w_low_mi.mat',1 },
{ 50201,7,5,{{3,1,1014001,36}},10,'','','zyz/zyz06_low/zyz06_low_c01_mi.mat','zyz/zyz06_high/zyz06_hi_c01_mi.mat','','zyz/zyz06_low/zyz06_w_low_c01_mi.mat',0 },
{ 50301,7,5,{{3,1,1014001,36}},1,'','','zyz/zyz06_low/zyz06_low_c02_mi.mat','zyz/zyz06_high/zyz06_hi_c02_mi.mat','','zyz/zyz06_low/zyz06_w_low_c02_mi.mat',0 },
}
return {data, title}
local title = {['ID']=1,['Name']=2,['TextName']=3,['IconID']=4,['Rate']=5,['ItemID']=6,['BuffID']=7,['SkillID']=8,['AttackParam']=9,['Vector3Test']=10,['Vector2Test']=11,['AwardList']=12,['Type']=13,['Pos']=14,['PosSingle']=15,['SkillCache']=16,['ArrayEx']=17,['ArrayExFloat']=18,['IsRobot']=19}
local data = {
[1]={ 100,'黑猫警长','黑猫警长，',0,0.1,100,{10},{{x=4,y=6},{x=8,y=9}},1,{{x=3,y=4,z=5}},{x=1,y=2},{10,3,4},0,{{X=100,Y=89},{X=20,Y=30}},{X=100,Y=89},{0,0.333},{{0,1},{2,3}},{{0.1,0.2,0.3},{0.4,0.5,0.6}},true },
[2]={ 101,'葫芦\n娃','',0,0.5,100,{3,1},nil,0,nil,{x=0,y=0},{1,3,5},2,{{X=100,Y=89},{X=20,Y=39}},{X=100,Y=20},{0,0.467},nil,nil,false },
[3]={ 102,'舒"克"','',0,0.3,100,nil,nil,0,nil,{x=0,y=0},nil,3,{{}},{X=100,Y=89},{0,0.533},nil,nil,false },
['one']={ 103,'贝\n塔','',0,0.6,100,nil,nil,0,nil,{x=0,y=0},nil,1,{{}},{X=0,Y=0},{0,0.533},nil,nil,false },
[5]={ 104,'邋遢大王','',0,0.7,100,nil,nil,0,nil,{x=0,y=0},nil,2,{{}},{X=0,Y=0},nil,nil,nil,false },
}
return {data, title}

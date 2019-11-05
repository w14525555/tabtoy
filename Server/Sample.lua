local title = {['ID']=1,['Name']=2,['TextName']=3,['IconID']=4,['Rate']=5,['ItemID']=6,['BuffID']=7,['SkillID']=8,['AttackParam']=9,['Vector3Test']=10,['Vector2Test']=11,['key']=12,['AwardList']=13,['Type']=14,['Pos']=15,['PosSingle']=16,['SkillCache']=17,['ArrayEx']=18,['ArrayExFloat']=19,['IsRobot']=20}
local data = {
{ 100,'黑猫警长','黑猫警长，',0,0.1,100,{10},{{x=4,y=6},{x=8,y=9}},1,{{x=3,y=4,z=5}},{x=1,y=2},1,{10,3,4},0,{{X=100,Y=89},{X=20,Y=30}},{X=100,Y=89},{0,0.333},{{0,1},{2,3}},{{0.1,0.2,0.3},{0.4,0.5,0.6}},true },
{ 101,'葫芦\n娃','',0,0.5,100,{3,1},{},0,{},{x=0,y=0},2,{1,3,5},2,{{X=100,Y=89},{X=20,Y=39}},{X=100,Y=20},{0,0.467},{},{},false },
{ 102,'舒\"克\"','',0,0.3,100,{},{},0,{},{x=0,y=0},3,{},3,{{}},{X=100,Y=89},{0,0.533},{},{},false },
{ 103,'贝\n塔','',0,0.6,100,{},{},0,{},{x=0,y=0},'one',{},1,{{}},{X=0,Y=0},{0,0.533},{},{},false },
{ 104,'邋遢大王','',0,0.7,100,{},{},0,{},{x=0,y=0},5,{},2,{{}},{X=0,Y=0},{},{},{},false },
}
local mt = {__index = function (table,key)
    local temp = title[key]
    if temp then
        return table[temp]
    end
    return nil
end}
for k, v in pairs(data) do
    setmetatable(v, mt)
end
return {data=data, title=title}

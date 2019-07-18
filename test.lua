-- Generated by github.com/davyxu/tabtoy
-- Version: 2.9.1

local tab = {
	Sample = {
		{ ID = 100, Name = "黑猫警长", EmptyName = "", IconID = 0, NumericalRate = 0.6, ItemID = 100, BuffID = { 10 }, Pos = { X= 100, Y= 89 }, Type = "Leader", SkillID = { 4, 6, 7 }, AttackParam = { Value= 1 }, SingleStruct = { HP= 100, AttackRate= 1.2 }, StrStruct = { { HP= 3, ExType= "Leader" }, { HP= 10, ExType= "Monkey" } } 	},
		{ ID = 101, Name = "葫芦\n娃", EmptyName = "", IconID = 0, NumericalRate = 0.8, ItemID = 100, BuffID = { 3, 1 }, Pos = {  }, Type = "Pig", SkillID = { 1 }, AttackParam = {  }, SingleStruct = { HP= 10, AttackRate= 0, ExType= "Leader" }, StrStruct = { {  } } 	},
		{ ID = 102, Name = "舒\"克\"", EmptyName = "", IconID = 0, NumericalRate = 0.7, ItemID = 100, BuffID = {  }, Pos = {  }, Type = "Hammer", SkillID = {  }, AttackParam = {  }, SingleStruct = { HP= 10, AttackRate= 0, ExType= "Leader" }, StrStruct = { {  } } 	},
		{ ID = 103, Name = "贝\n塔", EmptyName = "", IconID = 0, NumericalRate = 0, ItemID = 100, BuffID = {  }, Pos = {  }, Type = "Monkey", SkillID = {  }, AttackParam = {  }, SingleStruct = { HP= 10, AttackRate= 0, ExType= "Leader" }, StrStruct = { {  } } 	},
		{ ID = 104, Name = "邋遢大王", EmptyName = "", IconID = 0, NumericalRate = 1, ItemID = 100, BuffID = {  }, Pos = {  }, Type = "Pig", SkillID = {  }, AttackParam = {  }, SingleStruct = { HP= 10, AttackRate= 0, ExType= "Leader" }, StrStruct = { {  } } 	}
	}

}


-- ID
tab.SampleByID = {}
for _, rec in pairs(tab.Sample) do
	tab.SampleByID[rec.ID] = rec
end

-- Name
tab.SampleByName = {}
for _, rec in pairs(tab.Sample) do
	tab.SampleByName[rec.Name] = rec
end

tab.Enum = {
	ActorType = {
		[0] = "Leader",
		[2] = "Pig",
		Monkey = 1,
		Hammer = 3,
	},
}

return tab
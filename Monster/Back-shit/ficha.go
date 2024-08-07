package main

type monster struct {
	Nome        string   `json:"Nome" bson:"Nome"`
	Vida        int64    `json:"Vida" bson:"Vida"`
	VD          int64    `json:"VD" bson:"VD"`
	DanoSan     string   `json:"DanoSan" bson:"DanoSan"`
	Armadura    int64    `json:"Armadura" bson:"Armadura"`
	Resistencia string   `json:"Resistencia" bson:"Resistencia"`
	Atributos   atributo `json:"Atributos" bson:"Atributos"`
	Elemento    string   `json:"Elemento" bson:"Elemento"`
	Passiva     string   `json:"Passiva" bson:"Passiva"`
	Acoes       string   `json:"Acoes" bson:"Acoes"`
}

type atributo struct {
	FOR int `json:"FOR" bson:"FOR"`
	AGI int `json:"AGI" bson:"AGI"`
	VIG int `json:"VIG" bson:"VIG"`
	PRE int `json:"PRE" bson:"PRE"`
	INT int `json:"INT" bson:"INT"`
}

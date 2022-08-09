package stringconcater

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ConcatChar string `md:"concatChar"`
}

type Input struct {
	FirstString  string `md:"firstString,required"`
	SecondString string `md:"secondString,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	s1, _ := coerce.ToString(values["firstString"])
	r.FirstString = s1
	s2, _ := coerce.ToString(values["secondString"])
	r.SecondString = s2
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"firstString":  r.FirstString,
		"secondString": r.SecondString,
	}
}

type Output struct {
	ConcatedString string `md:"concatedString"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["concatedString"])
	o.ConcatedString = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"concatedString": o.ConcatedString,
	}
}

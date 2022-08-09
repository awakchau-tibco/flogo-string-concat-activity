package stringconcater

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ConcatChar string `md:"concatChar"`
}

func (s *Settings) FromMap(values map[string]interface{}) error {
	c, err := coerce.ToString(values["concatChar"])
	if err != nil {
		return err
	}
	s.ConcatChar = c
	return nil
}

func (s *Settings) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"concatChar": s.ConcatChar,
	}
}

type Input struct {
	FirstString  string `md:"firstString,required"`
	SecondString string `md:"secondString,required"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	s1, err := coerce.ToString(values["firstString"])
	if err != nil {
		return err
	}
	i.FirstString = s1
	s2, err := coerce.ToString(values["secondString"])
	if err != nil {
		return err
	}
	i.SecondString = s2
	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"firstString":  i.FirstString,
		"secondString": i.SecondString,
	}
}

type Output struct {
	ConcatedString string `md:"concatedString"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, err := coerce.ToString(values["concatedString"])
	if err != nil {
		return err
	}
	o.ConcatedString = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"concatedString": o.ConcatedString,
	}
}

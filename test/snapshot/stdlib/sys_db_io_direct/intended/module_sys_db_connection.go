package main

type sys__db__Connection interface {
	request(s *string) sys__db__ResultSet
}

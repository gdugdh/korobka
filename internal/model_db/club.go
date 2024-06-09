package model_db

import "time"

type Club struct {
	Id          int64  `db:"id"`
	IdCreator   int64  `db:"id_creator"`
	Name        string `db:"name"`
	Image       string `db:"image"`
	Description string `db:"description"`
}

type Discipline struct {
	Id   int64  `db:"id_discipline"`
	Name string `db:"name_discipline"`
}

type Location struct {
	Id        int64   `db:"id_location"`
	Name      string  `db:"name_location"`
	Address   string  `db:"address_location"`
	Latitude  float64 `db:"latitude_location"`
	Longitude float64 `db:"longitude_location"`
}

type Match struct {
	Id            int64       `db:"id"`
	Location      *Location   `db:""`
	Discipline    *Discipline `db:""`
	Description   string      `db:"description"`
	DatetimeStart time.Time   `db:"datetime_start"`
	DatetimeEnd   time.Time   `db:"datetime_end"`
	Status        int64       `db:"status"`
}

type User struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	FullName string `db:"full_name"`
	Email    string `db:"email"`
	Avatar   string `db:"avatar"`
}

type Team struct {
	Id             int64   `db:"id_team"`
	IdMatch        int64   `db:"id_match_team"`
	Club           *Club   `db:""`
	Status         int64   `db:"status_team"`
	CountPlayer    int64   `db:"count_player_team"`
	CurCountPlayer int64   `db:"cur_count_player_team"`
	Players        []*User `db:""`
}

/*
club v, discipline v, location v, matchv, user v, Team v

message Club {
  int64 id = 1;
  int64 idCreator = 2;
  string name = 3;
  string image = 4;
  string description = 5;
}

message Discipline {
  int64 id = 1;
  string name = 2;
}

message Location {
  int64 id = 1;
  string name = 2;
  string address = 3;
  double latitude = 4;
  double longitude = 5;
}

message Match {
  int64 id = 1;
  Location location = 2;
  Discipline discipline = 3;
  string description = 4;
  google.protobuf.Timestamp datetime_start = 5;
  google.protobuf.Timestamp datetime_end = 6;
  int64 status = 7;
}

message User {
  int64 id = 1;
  string username = 2;
  string full_name = 3;
  string email = 4;
  string avatar = 5;
}

message Team {
  int64 id = 1;
  int64 id_match = 2;
  int64 id_club = 3;
  int64 status = 4;
  repeated User players = 5;
}

message MatchDetail {
  int64 id = 1;
  Location location = 2;
  Discipline discipline = 3;
  string description = 4;
  google.protobuf.Timestamp datetime_start = 5;
  google.protobuf.Timestamp datetime_end = 6;
  int64 status = 7;
  Team team1 = 8;
  Team team2 = 9;
}
*/

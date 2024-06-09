package converter

import (
	"golang-boilerplate/internal/model_db"
	desc "golang-boilerplate/pkg/club_v1"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ClubM2D(club *model_db.Club) *desc.Club {
	return &desc.Club{
		Id:          wrapperspb.Int64(club.Id),
		IdCreator:   wrapperspb.Int64(club.IdCreator),
		Name:        wrapperspb.String(club.Name),
		Image:       wrapperspb.String(club.Image),
		Description: wrapperspb.String(club.Description),
	}
}

func ClubD2M(club *desc.Club) *model_db.Club {
	log.Println(club)
	return &model_db.Club{
		Id:          club.Id.GetValue(),
		IdCreator:   club.IdCreator.GetValue(),
		Name:        club.Name.GetValue(),
		Image:       club.Image.GetValue(),
		Description: club.Description.GetValue(),
	}
}

func DisciplineM2D(discipline *model_db.Discipline) *desc.Discipline {
	return &desc.Discipline{
		Id:   discipline.Id,
		Name: discipline.Name,
	}
}

func LocationM2D(location *model_db.Location) *desc.Location {
	return &desc.Location{
		Id:        location.Id,
		Name:      location.Name,
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}
}

func MatchM2D(match *model_db.Match, teams []*model_db.Team) *desc.Match {
	desc_matches := &desc.Match{
		Id:            match.Id,
		Location:      LocationM2D(match.Location),
		Discipline:    DisciplineM2D(match.Discipline),
		Description:   match.Description,
		DatetimeStart: timestamppb.New(match.DatetimeStart),
		DatetimeEnd:   timestamppb.New(match.DatetimeEnd),
		Status:        match.Status,
		Teams:         make([]*desc.Team, 0),
	}

	for _, team := range teams {
		desc_matches.Teams = append(desc_matches.Teams, TeamM2D(team))
	}

	return desc_matches
}

func UserM2D(user *model_db.User) *desc.User {
	return &desc.User{
		Id:       user.Id,
		Username: user.Username,
		FullName: user.FullName,
		Avatar:   user.Avatar,
		Email:    user.Email,
	}
}

func TeamM2D(team *model_db.Team) *desc.Team {
	desc_team := &desc.Team{
		Id:             team.Id,
		IdMatch:        team.IdMatch,
		Club:           ClubM2D(team.Club),
		CountPlayer:    team.CountPlayer,
		CurCountPlayer: team.CurCountPlayer,
		Status:         team.Status,
	}

	for _, player := range team.Players {
		desc_team.Players = append(desc_team.Players, UserM2D(player))
	}

	return desc_team
}

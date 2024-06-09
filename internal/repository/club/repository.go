package club

import (
	"context"
	"log"

	"golang-boilerplate/internal/client/db"
	"golang-boilerplate/internal/converter"
	"golang-boilerplate/internal/model_db"
	"golang-boilerplate/internal/repository"
	desc "golang-boilerplate/pkg/club_v1"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ClubRepository {
	return &repo{db: db}
}

func (r *repo) FinishedMatchList(ctx context.Context, id_discipline int64) ([]*desc.Match, error) {
	q := db.Query{
		Name: "club_repository.ActiveMatchList SELECT matches",
		QueryRaw: `SELECT 
		m.id, 
		
		l.id as "id_location",
		l.name as "name_location",
		l.address as "address_location",
		l.latitude as "latitude_location",
		l.longitude as "longitude_location",
		
		d.id as "id_discipline",
		d.name as "name_discipline",
		
		m.description, 
		m.datetime_start, 
		m.datetime_end, 
		m.status
		
		FROM public."Match" m
		INNER JOIN public."Location" l ON m.id_location = l.id
		INNER JOIN public."Discipline" d ON m.id_discipline = d.id
		WHERE status != 0 AND d.id=$1;`}

	var matches []*model_db.Match
	if err := r.db.DB().ScanAllContext(ctx, &matches, q, id_discipline); err != nil {
		return nil, err
	}

	return r.MatchesM2D(ctx, matches)
}

func (r *repo) ActiveMatchList(ctx context.Context, id_discipline int64) ([]*desc.Match, error) {
	q := db.Query{
		Name: "club_repository.ActiveMatchList SELECT matches",
		QueryRaw: `SELECT 
		m.id, 
		
		l.id as "id_location",
		l.name as "name_location",
		l.address as "address_location",
		l.latitude as "latitude_location",
		l.longitude as "longitude_location",
		
		d.id as "id_discipline",
		d.name as "name_discipline",
		
		m.description, 
		m.datetime_start, 
		m.datetime_end, 
		m.status
		
		FROM public."Match" m
		INNER JOIN public."Location" l ON m.id_location = l.id
		INNER JOIN public."Discipline" d ON m.id_discipline = d.id
		WHERE status = 0 AND d.id=$1;`}

	var matches []*model_db.Match
	if err := r.db.DB().ScanAllContext(ctx, &matches, q, id_discipline); err != nil {
		return nil, err
	}

	return r.MatchesM2D(ctx, matches)
}

func (r *repo) MatchesM2D(ctx context.Context, matches []*model_db.Match) ([]*desc.Match, error) {
	var desc_matches []*desc.Match

	for _, match := range matches {
		var teams []*model_db.Team
		team_q := db.Query{
			Name: "club_repository.ActiveMatchList SELECT teams",
			QueryRaw: `SELECT 
			t.id as "id_team", 
			t.id_match as "id_match_team",
		
			t.status as "status_team", 
			t.count_player as "count_player_team", 
			t.cur_count_player as "cur_count_player_team"
		FROM public."Team" t
		WHERE t.id_match = $1;`}
		err := r.db.DB().ScanAllContext(ctx, &teams, team_q, match.Id)
		if err != nil {
			return nil, err
		}

		// teams

		// club_q := db.Query{
		// 	Name: "club_repository.ActiveMatchList SELECT teams",
		// 	QueryRaw: `SELECT
		// 		c.id,
		// 		c.id_creator,
		// 		c.name,
		// 		c.image,
		// 		c.description,
		// 		FROM public."Club" c
		// 		WHERE c.id = $1`}

		for _, team := range teams {
			club := new(model_db.Club)
			club_q := db.Query{
				Name:     "club_repository.ActiveMatchList SELECT club",
				QueryRaw: `SELECT id, id_creator, name, image, description FROM public."Club" WHERE id IN (SELECT id_club FROM public."Team" WHERE id=$1 )`,
			}

			if err := r.db.DB().QueryRowContext(ctx, club_q, team.Id).Scan(&club.Id, &club.IdCreator, &club.Name, &club.Image, &club.Description); err != nil {
				log.Println(err)
				team.Club = &model_db.Club{}
			} else {
				team.Club = club
			}

			var players []*model_db.User
			players_q := db.Query{
				Name: "club_repository.ActiveMatchList SELECT players",
				QueryRaw: `SELECT id, username, email, full_name, avatar
				FROM public."User"
			WHERE id IN (SELECT id_user FROM public."PlayerInTeam" WHERE id_team = $1);`}
			if err := r.db.DB().ScanAllContext(ctx, &players, players_q, team.Id); err != nil {
				return nil, err
			}
			team.Players = players
		}
		desc_matches = append(desc_matches, converter.MatchM2D(match, teams))
	}

	return desc_matches, nil
}

func (r *repo) FriendList(ctx context.Context, id_user int64) ([]*desc.User, error) {
	q := db.Query{
		Name: "club_repository.ActiveMatchList SELECT matches",
		QueryRaw: `SELECT id, username, email, full_name, avatar
		FROM public."User"
		WHERE id IN (SELECT id_user1 FROM public."Friendship" WHERE id_user2 = $1) OR 
			  id IN (SELECT id_user2 FROM public."Friendship" WHERE id_user1 = $1)`}

	var friends []*model_db.User
	if err := r.db.DB().ScanAllContext(ctx, &friends, q, id_user); err != nil {
		return nil, err
	}

	var desc_users []*desc.User

	for _, friend := range friends {
		desc_users = append(desc_users, converter.UserM2D(friend))
	}

	return desc_users, nil
}

// func (i *Implementation) FinishedMatchList(ctx context.Context, req *desc.FinishedMatchListRequest) (*desc.FinishedMatchListResponse, error) {
// 	return i.clubRepository.FinishedMatchList(ctx, req.GetId())
// }

// func (r *repo) Create(ctx context.Context, club *model_db.Club) (int64, error) {
// 	q := db.Query{
// 		Name:     "club_repository.Create",
// 		QueryRaw: "INSERT INTO public.\"Club\" (id_creator, name, image, description) VALUES ($1, $2, $3,  $4) RETURNING id;",
// 	}

// 	var id int64
// 	err := r.db.DB().QueryRowContext(ctx, q, club.IdCreator, club.Name, club.Image, club.Description).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (r *repo) Get(ctx context.Context, id int64) (*model_db.Club, error) {
// 	q := db.Query{
// 		Name:     "club_repository.Get",
// 		QueryRaw: "SELECT * FROM public.\"Club\" WHERE id = $1;",
// 	}

// 	var club model_db.Club
// 	err := r.db.DB().ScanOneContext(ctx, &club, q, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &club, nil
// }

// func (r *repo) Update(ctx context.Context, updateInfo *model_db.ClubUpdate) error {
// 	q := db.Query{
// 		Name:     "club_repository.Update",
// 		QueryRaw: "UPDATE public.\"Club\" SET name = $1, image = $2, description = $3 WHERE id = $4;",
// 	}

// 	_, err := r.db.DB().ExecContext(ctx, q, updateInfo.Name, updateInfo.Image, updateInfo.Description, updateInfo.Id)

// 	return err
// }

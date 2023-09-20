package gapi

import (
	"context"
	"strings"
	"time"

	models "github.com/scm-dev1dev5/mtm-community-forum/user_rpc/models"
	"github.com/scm-dev1dev5/mtm-community-forum/user_rpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (teamServer *TeamServer) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.TeamResponse, error) {
	team := &models.CreateTeamRequest{
		Name:        req.GetName(),
		DeparmentId: req.GetDepartmentId(),
	}

	newTeam, err := teamServer.teamService.CreateTeam(team)
	if err != nil {
		if strings.Contains(err.Error(), "name already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.TeamResponse{

		XId:          newTeam.Id.Hex(),
		Name:         newTeam.Name,
		DepartmentId: newTeam.DeparmentId,
		CreatedAt:    timestamppb.New(newTeam.CreateAt),
		UpdatedAt:    timestamppb.New(newTeam.UpdatedAt),
	}
	return res, nil
}

func (teamServer *TeamServer) GetTeam(ctx context.Context, req *pb.TeamRequest) (*pb.TeamResponse, error) {
	teamId := req.GetXId()

	team, err := teamServer.teamService.GetTeam(teamId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.TeamResponse{

		XId:          team.Id.Hex(),
		Name:         team.Name,
		DepartmentId: team.DeparmentId,
		CreatedAt:    timestamppb.New(team.CreateAt),
		UpdatedAt:    timestamppb.New(team.UpdatedAt),
	}
	return res, nil
}

func (teamServer *TeamServer) GetTeams(req *pb.GetTeamsRequest, stream pb.TeamService_GetTeamsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	teams := teamServer.teamService.GetTeamList(int(page), int(limit))

	for _, team := range teams {
		stream.Send(&pb.Team{
			XId:          team.Id.Hex(),
			Name:         team.Name,
			DepartmentId: team.DeparmentId,
			CreatedAt:    timestamppb.New(team.CreateAt),
			UpdatedAt:    timestamppb.New(team.UpdatedAt),
		})
	}

	return nil
}

func (teamServer *TeamServer) UpdateTeam(ctx context.Context, req *pb.TeamUpdateRequest) (*pb.TeamResponse, error) {
	teamId := req.GetXId()

	team := &models.UpdateTeam{
		Name:        req.GetName(),
		DeparmentId: req.GetDepartmentId(),
		UpdatedAt:   time.Now(),
	}

	updatedTeam, err := teamServer.teamService.UpdateTeam(teamId, team)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.TeamResponse{

		XId:          updatedTeam.Id.Hex(),
		Name:         updatedTeam.Name,
		DepartmentId: updatedTeam.DeparmentId,
		CreatedAt:    timestamppb.New(updatedTeam.CreateAt),
		UpdatedAt:    timestamppb.New(updatedTeam.UpdatedAt),
	}
	return res, nil
}

func (teamServer *TeamServer) DeleteTeam(ctx context.Context, req *pb.TeamRequest) (*pb.DeleteTeamResponse, error) {
	teamId := req.GetXId()

	err := teamServer.teamService.DeleteTeam(teamId)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res := &pb.DeleteTeamResponse{
		Success: true,
	}
	return res, nil
}

func (teamServer *TeamServer) GetTeamByDeparmentId(req *pb.TeamRequestByDeparmentId, stream pb.TeamService_GetTeamByDeparmentIdServer) error {
	deparmentId := req.GetDepartmentId()

	deparments, err := teamServer.teamService.GetTeamByDeparmentId(deparmentId)

	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, deparment := range deparments {
		stream.Send(&pb.Team{
			XId:          deparment.Id.Hex(),
			Name:         deparment.Name,
			DepartmentId: deparment.DeparmentId,
			CreatedAt:    timestamppb.New(deparment.CreateAt),
			UpdatedAt:    timestamppb.New(deparment.UpdatedAt),
		})
	}

	return nil
}

func (teamServer TeamServer) GetTeamCount(context context.Context, req *pb.GetTeamsRequest) (*pb.TeamCountResponse, error) {
	var page = req.GetPage()
	var limit = req.GetLimit()
	teams := teamServer.teamService.GetTeamCount(int(page), int(limit))
	res := &pb.TeamCountResponse{
		Count: int64(teams),
	}
	return res, nil
}

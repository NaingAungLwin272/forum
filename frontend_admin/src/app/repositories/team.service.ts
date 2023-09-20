import { Injectable } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Team } from '../interfaces/team';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TeamService extends ApiService {
  getTeamList(pageIndex?: number, pageSize?: number): Observable<Team[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/teams?page=${pageIndex}&limit=${pageSize}`,
      {}
    );
  }

  getTeamCount(): Observable<{ count: number }> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/teams/count`,
      {}
    )
  }

  getTeamByDepartmentId(id: string): Observable<Team[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/departments/${id}/teams`,
      {}
    );
  }

  createTeam(data: Team): Observable<Team> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/team`,
      { ...data }
    )
  }

  updateTeam(id: string, data: Team): Observable<Team> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/teams/${id}`,
      { ...data }
    )
  }

  deleteTeam(id: string): Observable<Team> {
    return this.apiConnecter(
      'DELETE',
      `${this.apiEndpoint}/teams/${id}`,
      {}
    )
  }
}

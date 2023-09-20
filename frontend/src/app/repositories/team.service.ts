import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';
import { Team } from '../interfaces/team';

@Injectable({
  providedIn: 'root'
})
export class TeamService extends ApiService {

  getTeam(id: string): Observable<Team> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/teams/${id}`,
      {}
    );
  }
}

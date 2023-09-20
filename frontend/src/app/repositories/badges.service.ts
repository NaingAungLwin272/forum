import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserPoint } from '../interfaces/badge';
import { ApiService } from '../services/api.service';

@Injectable({
  providedIn: 'root'
})
export class BadgesService extends ApiService {
  updateUserPoint(id: string, isSolved: boolean): Observable<UserPoint> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/user_point/${id}?is_solved=${isSolved}`,
      {}
    )
  }
}

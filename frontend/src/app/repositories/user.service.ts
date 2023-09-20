import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';
import { User, UserSummary } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class UserService extends ApiService {

  getUserList(): Observable<User[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users`,
      {}
    )
  }

  getUser(id: string): Observable<User> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}`,
      {}
    )
  }


  getUserProfile(id: string): Observable<User> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}
      `,
      {}
    );
  }

  getUserSummary(userId: string): Observable<UserSummary> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${userId}/summary`,
      {}
    );
  }

  editProfile(id: string | undefined, formData: any): Observable<User> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/users/${id}`,
      formData
    )
  }

  getUserByDisplayName(name: string): Observable<User> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/uname/${name}`,
      {}
    )
  }
}
